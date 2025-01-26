package migrations

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type MigrationTable struct {
	gorm.Model
	Name         string `gorm:"unique"`
	SchemaFields string `gorm:"type:json"`
}

// Parse metadata from given model
func GetMetaData(db *gorm.DB, model interface{}) (string, string, *schema.Schema, error) {
	// Prepare statements
	stmt := &gorm.Statement{DB: db}

	// Parse model schema
	if err := stmt.Parse(model); err != nil {
		return "", "", nil, fmt.Errorf("error parsing model: %w", err)
	}

	// Serialize model fields
	var fields []string
	for _, field := range stmt.Schema.Fields {
		fields = append(fields, field.DBName)
	}

	// Convert model fields to string
	if fieldsString, err := json.Marshal(fields); err != nil {
		return "", "", nil, fmt.Errorf("error: %w", err)
	} else {
		return string(fieldsString), stmt.Schema.Table, stmt.Schema, nil
	}
}

// check migration and migrate, migration table
func CheckMigrateMigrationTable(db *gorm.DB) error {
	// Migrate `MigrationTable` return error if not success
	if err := db.Migrator().AutoMigrate(&MigrationTable{}); err != nil {
		return fmt.Errorf("error migrating %T: %w", &MigrationTable{}, err)
	}
	// return nil if everything is ok
	return nil
}

// Migrate applies all pending migrations
func Migrate(db *gorm.DB) error {
	// Check migration table is ready if not throw an error
	if err := CheckMigrateMigrationTable(db); err != nil {
		return err
	}

	fmt.Println("Starting migration process...")

	// Loop through each model in modelLists and apply migrations
	for _, model := range modelLists {
		// Parse metadata from model
		modelSchemaField, modelName, _, err := GetMetaData(db, model)
		if err != nil {
			return fmt.Errorf("error migrating %T: %w", model, err)
		}

		// Table doesn't exist, creating the table
		fmt.Printf("Table for model %T does not exist. Creating table...\n", model)
		if err := db.Migrator().CreateTable(model); err != nil {
			return fmt.Errorf("failed to create table for model %T: %w", model, err)
		}

		// Record the migration to migration table
		db.Create(&MigrationTable{Name: modelName, SchemaFields: modelSchemaField})
		fmt.Printf("Successfully created table for model %T\n", model)
	}

	fmt.Println("Migration process completed successfully.")
	return nil
}

// Migrate applies all pending migrations
func MigrateWithCleanUp(db *gorm.DB) error {
	// Check migration table is ready if not throw an error
	if err := CheckMigrateMigrationTable(db); err != nil {
		return err
	}

	fmt.Println("Starting migration process...")

	// Loop through each model in modelLists and apply migrations
	for _, model := range modelLists {
		// Parse metadata from model
		modelSchemaField, modelName, modelSchema, err := GetMetaData(db, model)
		if err != nil {
			return fmt.Errorf("error migrating %T: %w", model, err)
		}

		// Check if the table exists
		if !db.Migrator().HasTable(model) {
			// Table doesn't exist, creating the table
			fmt.Printf("Table for model %T does not exist. Creating table...\n", model)
			if err := db.Migrator().CreateTable(model); err != nil {
				return fmt.Errorf("failed to create table for model %T: %w", model, err)
			}

			// Record the migration to migration table
			db.Create(&MigrationTable{Name: modelName, SchemaFields: modelSchemaField})
			fmt.Printf("Successfully created table for model %T\n", model)

		} else {

			// If Table Already exist check field changes
			// get previous schema fields
			// if schema field are not same then migrate again
			var m MigrationTable
			db.Where("name = ?", modelName).First(&m)
			var oldFiedls []string
			if err := json.Unmarshal([]byte(m.SchemaFields), &oldFiedls); err != nil {
				return fmt.Errorf("error migrating %T: %w", model, err)
			}
			fiedlsMap := make(map[string]struct{})
			for _, item := range modelSchema.Fields {
				fiedlsMap[item.DBName] = struct{}{}
			}
			fmt.Printf("Table for model %T already exists. Running migration...\n", model)
			// Auto migrate the model
			db.AutoMigrate(model)

			// Remove auto unused field
			for _, field := range oldFiedls {
				// check new field exist of not in current db schema
				if _, found := fiedlsMap[field]; !found {
					// not exist delete the column
					db.Migrator().DropColumn(model, field)
				}
			}

			db.Model(&MigrationTable{}).Where("name = ?", modelName).Update("SchemaFields", modelSchemaField)
		}
	}

	fmt.Println("Migration process completed successfully.")
	return nil
}

// Rollback the last applied migration
func Rollback(db *gorm.DB) error {
	// Loop through each model in modelLists
	for _, m := range modelLists {
		// Check if the table exists before trying to drop it
		if db.Migrator().HasTable(m) {
			// Attempt to drop the table
			if err := db.Migrator().DropTable(m); err != nil {
				return fmt.Errorf("failed to drop table for model %T: %w", m, err)
			}
			fmt.Printf("Successfully dropped table for model: %T\n", m)
		} else {
			fmt.Printf("Table for model %T does not exist, skipping drop.\n", m)
		}
	}

	// Delete all migration history or Drop table
	db.Migrator().DropTable(&MigrationTable{})

	return nil
}

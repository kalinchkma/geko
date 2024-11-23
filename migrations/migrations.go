package migrations

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Migration struct {
	ID      string
	Up      func(*gorm.DB) error
	Down    func(*gorm.DB) error
	Applied bool
}

type MigrationTable struct {
	gorm.Model
	ID   string `gorm:"primaryKey"`
	Name string
}

// Migration list
// Here you will add models to migrate
var migrations = []Migration{
	{
		ID:   "user_",
		Up:   Up_user,
		Down: Down_user,
	},
	{
		ID:   "guild_",
		Up:   Up_guild,
		Down: Down_guild,
	},
}

type MigrationLog struct {
	ID        string `gorm:"primaryKey"`
	TrackIdS  []int
	AppliedAt time.Time
}

// Create model on database
func Up(db *gorm.DB, m *gorm.Model) error {
	return db.AutoMigrate(m)
}

// Drop model from database
func Down(db *gorm.DB, m *gorm.Model) error {
	return db.Migrator().DropTable(m)
}

// Migrate applies all pending migration
func Migrate(db *gorm.DB) error {
	db.AutoMigrate(&MigrationLog{})
	var appliedMigrations []MigrationLog
	db.Find(&appliedMigrations)

	appliedIDs := map[string]bool{}

	for _, m := range appliedMigrations {
		appliedIDs[m.ID] = true
	}

	for _, m := range migrations {
		if !appliedIDs[m.ID] {
			fmt.Printf("Applying migrationL: %s\n", m.ID)
			if err := m.Up(db); err != nil {
				return fmt.Errorf("failed to apply migration %s: %w", m.ID, err)
			}
			db.Create(&MigrationLog{ID: m.ID, AppliedAt: time.Now()})
		}
	}
	return nil
}

// Rollback the last applied migration
func Rollback(db *gorm.DB) error {
	// Ensure the MigrationLog table exists
	db.AutoMigrate(&MigrationLog{})

	// Retrieve the most recently migration
	var lastMigration MigrationLog
	r := db.Order("applied_at desc").First(&lastMigration)

	// check migation applied or not
	// if no migration applied, no roll back
	if r.RowsAffected == 0 {
		fmt.Println("No migrations to roll back.")
		return nil
	}

	// find the corresponding migration in the list
	var migration *Migration
	for _, m := range migrations {
		if m.ID == lastMigration.ID {
			migration = &m
			break
		}
	}

	if migration == nil {
		return fmt.Errorf("migration with ID %s not found in migrations list", lastMigration.ID)
	}

	// Roll back the migration
	fmt.Printf("Rolling back migration: %s\n", migration.ID)
	if err := migration.Down(db); err != nil {
		return fmt.Errorf("failed to roll back migration %s: %w", migration.ID, err)
	}

	// Remove the migration log entry
	if err := db.Delete(&lastMigration).Error; err != nil {
		return fmt.Errorf("failed to delete migration log for %s: %w", migration.ID, err)
	}

	fmt.Printf("Successfully rolled back migrations: %s\n", migration.ID)
	return nil
}

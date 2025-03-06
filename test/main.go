package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// ValidateJSONPayload validates a JSON payload against a set of rules.
func ValidateJSONPayload[T any](payload []byte) error {
	var data map[string]any

	// Unmarshal the JSON payload into a map
	if err := json.Unmarshal(payload, &data); err != nil {
		return fmt.Errorf("invalid JSON: %v", err)
	}

	var body T

	v := reflect.ValueOf(body)
	t := reflect.TypeOf(body)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	fmt.Println("validating")
	for i := range v.NumField() {
		field := t.Field(i)
		// value := v.Field(i)
		tag := field.Tag.Get("json")
		fmt.Println("Field type", field.Type)
		// fmt.Println("field", field)
		// fmt.Println("value", value)
		// fmt.Println("tag", tag)
		fmt.Println("dynamic fatch data", data[tag])
	}

	// // Example validation rules
	// if _, ok := data["name"]; !ok {
	// 	return errors.New("field 'name' is required")
	// }

	// if name, ok := data["name"].(string); !ok || name == "" {
	// 	return errors.New("field 'name' must be a non-empty string")
	// }

	// if age, ok := data["age"].(float64); ok {
	// 	if age < 0 {
	// 		return errors.New("field 'age' must be a non-negative number")
	// 	}
	// } else if _, ok := data["age"]; ok {
	// 	return errors.New("field 'age' must be a number")
	// }

	// Add more validation rules as needed

	return nil
}

type User struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Designation string `json:"designation"`
}

func main() {
	// Example JSON payload
	payload := []byte(`{
		"name": "John Doe",
		"age": "30",
		"designation": "developer"
	}`)

	// Validate the JSON payload
	if err := ValidateJSONPayload[User](payload); err != nil {
		fmt.Println("Validation failed:", err)
	} else {
		fmt.Println("Validation succeeded: JSON payload is valid.")
	}
}

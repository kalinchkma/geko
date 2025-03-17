package lib

import "testing"

func TestCamelToSnakeCase(t *testing.T) {
	input := "testHudai"
	result := "test_hudai"
	tv, err := CamelToSnake(input)
	if err != nil {
		t.Fatal("\nFalied camel to snakecase", err)
	} else if tv == result {
		t.Log("Pass", tv)
	} else {
		t.Fatalf("\nFalied expected: %v\n got: %v", result, tv)
	}

	input = "testHudaiTestRamTam"
	result = "test_hudai_test_ram_tam"
	tv, err = CamelToSnake(input)
	if err != nil {
		t.Fatal("\nFalied camel to snakecase on case 2", err)
	} else if tv == result {
		t.Log("Pass", tv)
	} else {
		t.Fatalf("\nFailed expected: %v\n got: %v", result, tv)
	}

}

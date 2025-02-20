package utils

import "testing"

func TestCamelToSnake(t *testing.T) {
	testData := []string{
		"UserID",
		"testData",
		"TestValue",
		"Test-Value",
		"Test_value",
	}

	for _, v := range testData {
		t.Log(CamelToSnake(v))
	}
}

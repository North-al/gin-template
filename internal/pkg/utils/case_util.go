package utils

import (
	"regexp"
	"strings"
)

// CamelToSnake 将驼峰命名转换为蛇形命名
// example: CamelToSnake("userId") -> "user_id"
// example: CamelToSnake("testData") -> "test_data"
// example: CamelToSnake("TestValue") -> "test_value"
func CamelToSnake(s string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

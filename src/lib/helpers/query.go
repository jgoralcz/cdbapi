package helpers

import (
	"errors"
	"strconv"
	"strings"
)

var ErrorLessThanZero = errors.New("number is less than 0")
var ErrorGreaterThanMax = errors.New("number is greater than max")

// MaxLimit receives a user input string and decides whether
// it is an integer and between the default value and max value.
// If it is not an integer it will default to the default value.
// If it is above the max limit, the provided maxLimit will be used.
func MaxLimit(userLimit string, defaultLimit int, maxLimit int) int {
	if len(userLimit) > 5 {
		return defaultLimit
	}

	limit, err := NumberOverMax(userLimit)

	if err != nil || limit < defaultLimit {
		return defaultLimit
	}

	if limit > maxLimit {
		return maxLimit
	}

	return limit
}

// DefaultBoolean receives a user input string and decides whether
// it is a valid boolean (true or false). If it is not a boolean then
// an empty string is used
func DefaultBoolean(userBool string) string {
	if len(userBool) > 5 {
		return ""
	}

	lowerStr := strings.ToLower(userBool)

	if lowerStr != "true" && lowerStr != "false" {
		return ""
	}
	return lowerStr
}

func NumberOverMax(value string) (int, error) {
	num, err := strconv.Atoi(value)

	if err != nil {
		return -1, err
	}

	if num < 0 {
		return -1, ErrorLessThanZero
	}

	if num >= 2_147_483_647 {
		return -1, ErrorGreaterThanMax
	}

	return num, nil
}

// DefaultNumber receives a value and defaultValue. When the value is an empty string,
// the default int is used. When the value cannot be converted, -1 is returned.
// Otherwise the parsed string is converted to an int and returned.
func DefaultNumber(value string, defaultValue int) int {
	if len(value) > 5 {
		return -1
	}

	if value == "" {
		return defaultValue
	}

	valueInt, err := NumberOverMax(value)
	if err != nil {
		return -1
	}

	return valueInt
}

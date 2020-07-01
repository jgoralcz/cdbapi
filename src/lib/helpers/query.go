package helpers

import "strconv"

// MaxLimit receives a user input string and decides whether
// it is an integer and between the default value and max value.
// If it is not an integer it will default to the default value.
// If it is above the max limit, the provided maxLimit will be used.
func MaxLimit(userLimit string, defaultLimit int, maxLimit int) int {
	limit, err := strconv.Atoi(userLimit)

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
	if userBool != "true" && userBool != "false" {
		return ""
	}
	return userBool
}

// DefaultNumber receives a value and defaultValue. When the value is an empty string,
// the default int is used. When the value cannot be converted, -1 is returned.
// Otherwise the parsed string is converted to an int and returned.
func DefaultNumber(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return -1
	}

	return valueInt
}

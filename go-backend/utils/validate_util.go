package utils

func ValidateStringLength(s string, min int, max int) bool {
	return min <= len(s) && len(s) <= max
}

package hw09structvalidator

import "regexp"

// a < b.
func IntIsLess(a int, b int) bool {
	return a < b
}

// a > b.
func IntIsGreater(a int, b int) bool {
	return a > b
}

func StringIsGreater(s string, length int) bool {
	return len(s) > length
}

func IsIn[T interface{ int | string }](v T, list []T) bool {
	for _, item := range list {
		if v == item {
			return true
		}
	}
	return false
}

func IsPatternMatch(s string, regexp regexp.Regexp) bool {
	return regexp.MatchString(s)
}

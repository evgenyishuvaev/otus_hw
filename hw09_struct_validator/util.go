package hw09structvalidator

import (
	"strings"
)

type Constraint struct {
	Name  string
	Value string
}

func GetConstraints(s string) []Constraint {
	if s == "" {
		return []Constraint{}
	}

	res := []Constraint{}
	constraints := strings.Split(s, "|")
	for _, constraint := range constraints {
		constraintAttrs := strings.Split(constraint, ":")
		res = append(res, Constraint{Name: constraintAttrs[0], Value: constraintAttrs[1]})
	}
	return res
}

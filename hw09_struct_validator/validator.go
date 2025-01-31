package hw09structvalidator

import (
	// "fmt"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	res := ""
	for _, validationErr := range v {
		res += fmt.Sprintf("%v: %s\n", validationErr.Field, validationErr.Err)

	}
	return res
}

// Валидатор структур, валидрует согласно прописанным тегам у полей.
func Validate(v interface{}) error {
	vr := reflect.ValueOf(v)
	res := ValidationErrors{}

	if vr.Kind() != reflect.Struct {
		return ErrInvalidTypeProvided
	}
	for i := range vr.NumField() {
		structField := vr.Type().Field(i)
		if !structField.IsExported() {
			continue
		}
		fieldValue := vr.FieldByName(structField.Name)

		fieldErrors := ValidateField(structField, fieldValue)
		if !errors.As(fieldErrors, ValidationErrors{}) {
			return fieldErrors
		}
		res = append(res, fieldErrors.(ValidationErrors)...)
	}

	fmt.Println(res)
	return res
}

func ValidateField(structField reflect.StructField, fieldValue reflect.Value) error {
	res := ValidationErrors{}
	fieldKind := structField.Type.Kind()
	fieldConstraints := structField.Tag.Get("validate")
	constraints := GetConstraints(fieldConstraints)

	for _, constraint := range constraints {
		switch constraint.Name {
		case "len":
			requiredLen, _ := strconv.Atoi(constraint.Value)
			if fieldKind == reflect.String {
				value := fieldValue.Interface().(string)
				if StringIsGreater(value, requiredLen) {
					res = append(res, ValidationError{structField.Name, ErrStringLenGreater})
				}
			}

		case "min":
			minValue, _ := strconv.Atoi(constraint.Value)
			if fieldKind != reflect.Int {
				return ErrBadValidateTagForType
			}
			value := fieldValue.Interface().(int)
			if IntIsLess(value, minValue) {
				res = append(res, ValidationError{structField.Name, ErrIntegerValueLess})
			}

		case "max":
			maxValue, _ := strconv.Atoi(constraint.Value)
			if fieldKind == reflect.Int {
				value := fieldValue.Interface().(int)
				if IntIsGreater(value, maxValue) {
					res = append(res, ValidationError{structField.Name, ErrIntegerValueGreater})
				}
			}
		case "regexp":
			pattern, err := regexp.Compile(constraint.Value)
			if err != nil {

			}

			if fieldKind == reflect.String {
				value := fieldValue.Interface().(int)
				if IntIsGreater(value, maxValue) {
					res = append(res, ValidationError{structField.Name, ErrIntegerValueGreater})
				}
			}

		}
	}
	return res
}

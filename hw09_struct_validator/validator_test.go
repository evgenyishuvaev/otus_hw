package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	// слайс структур для валидации
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			User{
				ID:     "111111111111111111111111111111111111111111111",
				Name:   "Igor",
				Age:    100,
				Email:  "igor@example.ru",
				Role:   "not in",
				Phones: []string{"89997776655"},
				meta:   nil,
			},
			ValidationErrors{
				ValidationError{
					"ID",
					ErrStringLenGreater,
				},
				ValidationError{
					"Age",
					ErrIntegerValueLess,
				},
			},
		},
		{
			"asdasd",
			nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			Validate(tt.in)
		})
	}
}

// func TestParseConstraints(t *testing.T) {
// t.Run("regexp with '|' char", func(t *testing.T) {})
// t.Run("regexp with ':' char", func(t *testing.T) {})
//
// }

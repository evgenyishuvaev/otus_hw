package hw09structvalidator

import "errors"

var (
	ErrStringLenGreater          = errors.New("длина строки не равна заданной")
	ErrStringDoesNotMatchPattern = errors.New("строка не соответстует шаблону")
	ErrStringNotIncludedInSet    = errors.New("строка не входит в заданное множество")

	ErrIntegerValueLess       = errors.New("число меньше заданного")
	ErrIntegerValueGreater    = errors.New("число больше заданного")
	ErrIntegerNotIncludeInSet = errors.New("число не входит в множество")
	ErrInvalidTypeProvided    = errors.New("передан недопустимый тип")

	ErrBadRegexpPattern      = errors.New("ошибка компиляции шаблона")
	ErrBadValidateTagForType = errors.New("тег недопустим для указанного типа")
)

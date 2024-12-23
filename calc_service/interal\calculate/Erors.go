package calc

import "errors"

var (
	ErrDivByZero       = errors.New("Деление на ноль!")
	ErrInvalidBracket  = errors.New("Ошибка со скобками")
	ErrInvalidOperands = errors.New("Проверьте количество операндов(+,-,/,*)")
	ErrInvalidJson     = errors.New("Проверьте коректность json")
	ErrEmptyJson       = errors.New("Пустой json!")
	ErrEmptyExpression = errors.New("Пустое выражение!")
)

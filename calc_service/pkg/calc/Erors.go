package calc

import "errors"

var (
	ErrDivByZero       = errors.New("Деление на ноль")
	ErrInvalidBracket  = errors.New("Проверьте скобки")
	ErrInvalidOperands = errors.New("Проверьте количество операндов(+,-,/,*)")
	ErrInvalidJson     = errors.New("Проверьте коректность написания")
	ErrEmptyJson       = errors.New("Пустой json")
	ErrEmptyExpression = errors.New("Пустое выражение")
)

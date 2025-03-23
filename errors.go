package goRule

import (
	"errors"
	"fmt"
)

var ErrorUnknownDecimalOperator error = errors.New("unknown decimal operator")
var ErrorTypeMismatch error = errors.New("compare type mismatch")
var ErrorUnknownStringOperator error = errors.New("unknown string operator")
var ErrorNoExpression error = errors.New("there is no expression")
var ErrorInvalidValue error = errors.New("invalid value")
var ErrorUnknownType error = errors.New("invalid value")
var ErrorInvalidFunctionCall error = errors.New("invalid function call")
var ErrorInvalidExpression error = errors.New("invalid expression")
var ErrorInvalidOperator error = errors.New("invalid operator")
var ErrorSyntaxError error = errors.New("syntax error")

func NewSyntaxError(v string) error {
	return fmt.Errorf("%w at line %v", ErrorSyntaxError, v)
}

func NewErrorInvalidOperator(op string, t string) error {
	return fmt.Errorf("%w: %s on %s", ErrorInvalidOperator, op, t)
}

func NewErrorTypeMismatch(v string, v2 string) error {
	return fmt.Errorf("%w: %s and %s", ErrorTypeMismatch, v, v2)
}

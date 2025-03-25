package rule

import (
	"errors"
	"fmt"
)

// Error variables used throughout the package:
var (
	// ErrorUnknownDecimalOperator is returned when a decimal comparison operator is unrecognized.
	ErrorUnknownDecimalOperator = errors.New("unknown decimal operator")

	// ErrorTypeMismatch is returned when two values fail a strict type check.
	ErrorTypeMismatch = errors.New("compare type mismatch")

	// ErrorUnknownStringOperator is returned when a string-based operator (co, sw, ew, in, pr) is unknown.
	ErrorUnknownStringOperator = errors.New("unknown string operator")

	// ErrorNoExpression is returned when the expression tree is nil or missing.
	ErrorNoExpression = errors.New("there is no expression")

	// ErrorInvalidValue is returned for invalid numeric/string parse attempts.
	ErrorInvalidValue = errors.New("invalid value")

	// ErrorUnknownType is returned for unknown or unsupported type annotations.
	ErrorUnknownType = errors.New("invalid value")

	// ErrorInvalidFunctionCall is returned when a function call parse is incomplete or malformed.
	ErrorInvalidFunctionCall = errors.New("invalid function call")

	// ErrorInvalidExpression is returned when the parse tree visitor hits an unexpected node.
	ErrorInvalidExpression = errors.New("invalid expression")

	// ErrorInvalidOperator is returned for an unknown or unsupported operator in a comparison.
	ErrorInvalidOperator = errors.New("invalid operator")

	// ErrorSyntaxError is used for general syntax errors in the input query.
	ErrorSyntaxError = errors.New("syntax error")
)

// newSyntaxError wraps the standard ErrorSyntaxError with line/column info.
func newSyntaxError(v string) error {
	return fmt.Errorf("%w at line %v", ErrorSyntaxError, v)
}

// newErrorInvalidOperator constructs an error indicating the given operator is invalid for a particular type.
func newErrorInvalidOperator(op string, t string) error {
	return fmt.Errorf("%w: %s on %s", ErrorInvalidOperator, op, t)
}

// newErrorTypeMismatch constructs an error indicating the given types do not match during strict comparison.
func newErrorTypeMismatch(v string, v2 string) error {
	return fmt.Errorf("%w: %s and %s", ErrorTypeMismatch, v, v2)
}

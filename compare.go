package rule

import (
	"fmt"
	"github.com/shopspring/decimal"
	"golang.org/x/exp/constraints"
	"reflect"
	"strconv"
	"strings"
)

// compareDecimal compares two decimal.Decimal values using the provided operator string.
// It returns a boolean indicating whether the comparison passed and an error if the operator is unknown.
//
// Supported operators: eq, ne, gt, lt, ge, le
func compareDecimal(leftVal decimal.Decimal, operator string, rightVal decimal.Decimal) (bool, error) {
	switch operator {
	case "eq", "==":
		return leftVal.Equal(rightVal), nil
	case "ne", "!=":
		return !leftVal.Equal(rightVal), nil
	case "gt", ">":
		return leftVal.GreaterThan(rightVal), nil
	case "lt", "<":
		return leftVal.LessThan(rightVal), nil
	case "ge", ">=":
		return leftVal.GreaterThanOrEqual(rightVal), nil
	case "le", "<=":
		return leftVal.LessThanOrEqual(rightVal), nil
	}
	return false, ErrorUnknownDecimalOperator
}

// compareOperator is the main comparison function used by the evaluation engine.
// It accepts leftVal, an operator string, rightVal, and a strictTypeCheck boolean.
//
// If strictTypeCheck == true, then reflect.TypeOf(leftVal) must match reflect.TypeOf(rightVal);
// if not, a type mismatch error is returned.
//
// Non-decimal numeric comparisons are delegated to compareOrdered() for standard ordering (>, <, etc.).
// String operations (co, sw, ew, in, pr) go to compareStringOps().
// Decimal comparisons go to compareDecimal().
func compareOperator(leftVal any, operator string, rightVal any, strictTypeCheck bool) (bool, error) {
	// Enforce strict type check if requested
	if strictTypeCheck {
		if reflect.TypeOf(leftVal) != reflect.TypeOf(rightVal) {
			return false, newErrorTypeMismatch(reflect.TypeOf(rightVal).String(), reflect.TypeOf(leftVal).String())
		}
	}

	// Handle special string-based operators
	switch operator {
	case "co", "sw", "ew", "in", "pr":
		return compareStringOps(leftVal, operator, rightVal)
	}

	// Handle decimals
	if ld, ok := leftVal.(decimal.Decimal); ok {
		rd := rightVal.(decimal.Decimal)
		return compareDecimal(ld, operator, rd)
	}

	// Otherwise handle numeric or boolean
	switch l := leftVal.(type) {

	case int:
		if strictTypeCheck {
			return compareOrdered(int64(l), operator, int64(rightVal.(int)))
		}
		i64, err := strconv.ParseInt(fmt.Sprintf("%v", rightVal), 10, 64)
		if err != nil {
			return false, ErrorInvalidValue
		}
		return compareOrdered(int64(l), operator, i64)

	case int32:
		if strictTypeCheck {
			return compareOrdered(l, operator, rightVal.(int32))
		}
		i32, err := strconv.ParseInt(fmt.Sprintf("%v", rightVal), 10, 32)
		if err != nil {
			return false, ErrorInvalidValue
		}
		return compareOrdered(int64(l), operator, i32)

	case int64:
		if strictTypeCheck {
			return compareOrdered(l, operator, rightVal.(int64))
		}
		i64, err := strconv.ParseInt(fmt.Sprintf("%v", rightVal), 10, 64)
		if err != nil {
			return false, ErrorInvalidValue
		}
		return compareOrdered(l, operator, i64)

	case uint:
		if strictTypeCheck {
			return compareOrdered(uint64(l), operator, uint64(rightVal.(uint)))
		}
		u64, err := strconv.ParseUint(fmt.Sprintf("%v", rightVal), 10, 64)
		if err != nil {
			return false, ErrorInvalidValue
		}
		return compareOrdered(l, operator, uint(u64))

	case uint32:
		if strictTypeCheck {
			return compareOrdered(uint64(l), operator, uint64(rightVal.(uint32)))
		}
		u32, err := strconv.ParseUint(fmt.Sprintf("%v", rightVal), 10, 32)
		if err != nil {
			return false, ErrorInvalidValue
		}
		return compareOrdered(l, operator, uint32(u32))

	case uint64:
		if strictTypeCheck {
			return compareOrdered(l, operator, rightVal.(uint64))
		}
		u64, err := strconv.ParseUint(fmt.Sprintf("%v", rightVal), 10, 64)
		if err != nil {
			return false, ErrorInvalidValue
		}
		return compareOrdered(l, operator, u64)

	case float32:
		if strictTypeCheck {
			return compareOrdered(float64(l), operator, float64(rightVal.(float32)))
		}
		f32, err := strconv.ParseFloat(fmt.Sprintf("%v", rightVal), 32)
		if err != nil {
			return false, ErrorInvalidValue
		}
		return compareOrdered(l, operator, float32(f32))

	case float64:
		if strictTypeCheck {
			return compareOrdered(l, operator, rightVal.(float64))
		}
		f64, err := strconv.ParseFloat(fmt.Sprintf("%v", rightVal), 64)
		if err != nil {
			return false, ErrorInvalidValue
		}
		return compareOrdered(l, operator, f64)

	case string:
		if strictTypeCheck {
			return compareOrdered(l, operator, rightVal.(string))
		}
		// Fallback: treat both as strings
		return compareOrdered(l, operator, fmt.Sprintf("%v", rightVal))

	case bool:
		if operator == "eq" {
			return l == rightVal, nil
		} else if operator == "ne" {
			return l != rightVal, nil
		} else {
			return false, newErrorInvalidOperator(operator, reflect.TypeOf(leftVal).String())
		}

	default:
		return false, newErrorTypeMismatch(reflect.TypeOf(rightVal).String(), reflect.TypeOf(leftVal).String())
	}
}

// compareOrdered is a generic helper for numeric or string ordering comparisons
// (i.e., eq, ne, gt, lt, ge, le). It uses Go generics constraints.Ordered.
//
// The left and right arguments must be of the same constraints.Ordered type T.
func compareOrdered[T constraints.Ordered](left T, operator string, right T) (bool, error) {
	switch operator {
	case "eq":
		return left == right, nil
	case "ne":
		return left != right, nil
	case "gt":
		return left > right, nil
	case "lt":
		return left < right, nil
	case "ge":
		return left >= right, nil
	case "le":
		return left <= right, nil
	default:
		return false, newErrorInvalidOperator(operator, reflect.TypeOf(left).String())
	}
}

// compareStringOps handles string-based comparisons: co, sw, ew, in, pr.
//
// - co => "contains": strings.Contains(leftVal, rightVal)
// - sw => "starts with": strings.HasPrefix(leftVal, rightVal)
// - ew => "ends with": strings.HasSuffix(leftVal, rightVal)
// - in => substring check (implemented as strings.Contains(rightVal, leftVal))
// - pr => "present" => leftVal != nil
func compareStringOps(leftVal any, operator string, rightVal any) (bool, error) {
	// "pr" => param is present (not nil). We only check for nil in leftVal.
	if operator == "pr" {
		return leftVal != nil, nil
	}

	// Convert both sides to string
	l := fmt.Sprint(leftVal)
	r := fmt.Sprint(rightVal)

	switch operator {
	case "co":
		// "contains"
		return strings.Contains(l, r), nil
	case "sw":
		// "starts with"
		return strings.HasPrefix(l, r), nil
	case "ew":
		// "ends with"
		return strings.HasSuffix(l, r), nil
	case "in":
		// "in" => check if l is in r
		return strings.Contains(r, l), nil
	}

	return false, ErrorUnknownStringOperator
}

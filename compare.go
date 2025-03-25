package rule

import (
	"fmt"
	"github.com/shopspring/decimal"
	"golang.org/x/exp/constraints"
	"reflect"
	"strconv"
	"strings"
)

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

func compareOperator(leftVal any, operator string, rightVal any, strictTypeCheck bool) (bool, error) {
	if strictTypeCheck {
		if reflect.TypeOf(leftVal) != reflect.TypeOf(rightVal) {
			return false, NewErrorTypeMismatch(reflect.TypeOf(rightVal).String(), reflect.TypeOf(leftVal).String())
		}
	}

	switch operator {
	case "co", "sw", "ew", "in", "pr":
		return compareStringOps(leftVal, operator, rightVal)
	}

	if ld, ok := leftVal.(decimal.Decimal); ok {
		rd := rightVal.(decimal.Decimal)
		return compareDecimal(ld, operator, rd)
	}

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
		return compareOrdered(l, operator, fmt.Sprintf("%v", rightVal))
	case bool:
		if operator == "eq" {
			return l == rightVal, nil
		} else if operator == "ne" {
			return l != rightVal, nil
		} else {
			return false, NewErrorInvalidOperator(operator, reflect.TypeOf(leftVal).String())
		}
	default:
		return false, NewErrorTypeMismatch(reflect.TypeOf(rightVal).String(), reflect.TypeOf(leftVal).String())
	}
}

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
		return false, NewErrorInvalidOperator(operator, reflect.TypeOf(left).String())
	}
}

func compareStringOps(leftVal any, operator string, rightVal any) (bool, error) {
	// We can handle "pr" quickly if that means "leftVal is not nil"
	if operator == "pr" {
		return leftVal != nil, nil
	}

	// Convert both sides to string for "contains", "starts with", "ends with", "in", etc.
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
		// "in" = "leftVal in rightVal" or vice versa
		// Implementation differs, but here's an example:
		return strings.Contains(r, l), nil
	}
	return false, ErrorUnknownStringOperator
}
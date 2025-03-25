package rule

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/shopspring/decimal"
	parser "github.com/sky93/go-rule/internal/antlr4"
	"strconv"
	"strings"
)

// SyntaxError captures ANTLR syntax errors, storing them in errorListener.
func (l *errorListener) SyntaxError(_ antlr.Recognizer, _ any, line, column int, msg string, _ antlr.RecognitionException) {
	l.hasErrors = true
	l.errMsg = NewSyntaxError(fmt.Sprintf("%d:%d: %s", line, column, msg))
}

// evaluate traverses the exprTree to evaluate the final boolean outcome, given a map
// of param.id -> actual value. (This is used by Rule.Evaluate.)
func (e *exprTree) evaluate(values map[int]any, debugMode bool) (bool, error) {
	if e == nil {
		return false, ErrorNoExpression
	}

	// Handle logical operator nodes
	switch e.op {
	case "and":
		lRes, lErr := e.left.evaluate(values, debugMode)
		if lErr != nil {
			return false, lErr
		}
		rRes, rErr := e.right.evaluate(values, debugMode)
		if rErr != nil {
			return false, rErr
		}
		res := lRes && rRes
		if e.not {
			return !res, nil
		}
		return res, nil

	case "or":
		lRes, lErr := e.left.evaluate(values, debugMode)
		if lErr != nil {
			return false, lErr
		}
		rRes, rErr := e.right.evaluate(values, debugMode)
		if rErr != nil {
			return false, rErr
		}
		res := lRes || rRes
		if e.not {
			return !res, nil
		}
		return res, nil
	}

	// Leaf node => do param-based comparison
	p := e.param
	val, ok := values[p.id]
	if !ok {
		// If operator is "pr", presence is false => no value
		if p.operator == "pr" {
			return e.not == true, nil
		}
		// Otherwise treat missing param as false
		if e.not {
			return true, nil
		}
		return false, nil
	}

	out, err := compareOperator(val, p.operator, p.compareValue, p.strictTypeCheck)
	if debugMode {
		fmt.Printf(
			"Name: %s, left Value: %v<%T>, Operator:%s, right Value: %v<%T>, Strict Type Check: %t, Result: %t\n",
			p.Name, val, val, p.operator, p.compareValue, p.compareValue, p.strictTypeCheck, out,
		)
	}
	if err != nil {
		return false, err
	}
	if e.not {
		return !out, nil
	}
	return out, nil
}

// parseFunctionCall extracts the function name and arguments from the parse context.
func (v *queryVisitor) parseFunctionCall(ctx parser.IFunctionCallContext) (string, []FunctionArgument, error) {
	if ctx == nil {
		return "", nil, ErrorInvalidFunctionCall
	}
	name := ctx.ATTRNAME().GetText()
	var args []FunctionArgument
	argList := ctx.ArgList()
	if argList != nil {
		vals := argList.AllValue()
		for _, vc := range vals {
			val, t, _, err := v.parseValue(vc)
			if err != nil {
				return "", nil, err
			}
			args = append(args, FunctionArgument{
				ArgumentType: t,
				Value:        val,
			})
		}
	}
	return name, args, nil
}

// getAttrName returns the full text of an attrPath node (e.g. "user.age" or "account.details").
func (v *queryVisitor) getAttrName(ctx parser.IAttrPathContext) string {
	return ctx.GetText()
}

// parseTypedValue inspects the typedValue parse tree, applying user type annotations if present.
func (v *queryVisitor) parseTypedValue(tv parser.ITypedValueContext) (any, ArgumentType, bool, error) {
	switch typedNode := tv.(type) {
	case *parser.TypedStringContext:
		userType := "[s]"
		strictTypeCheck := false
		if ann := typedNode.TypeAnnotation(); ann != nil {
			userType = ann.GetText()
			strictTypeCheck = true
		}
		value, argType, err := v.applyUserType(unquoteString(typedNode.STRING().GetText()), userType)
		return value, argType, strictTypeCheck, err

	case *parser.TypedDoubleContext:
		userType := "[f64]"
		strictTypeCheck := false
		if ann := typedNode.TypeAnnotation(); ann != nil {
			userType = ann.GetText()
			strictTypeCheck = true
		}

		value, argType, err := v.applyUserType(typedNode.DOUBLE().GetText(), userType)
		return value, argType, strictTypeCheck, err

	case *parser.TypedIntegerContext:
		userType := "[i64]"
		strictTypeCheck := false
		if ann := typedNode.TypeAnnotation(); ann != nil {
			userType = ann.GetText()
			strictTypeCheck = true
		}

		value, argType, err := v.applyUserType(typedNode.GetText(), userType)
		return value, argType, strictTypeCheck, err

	default:
		return nil, ArgTypeUnknown, false, ErrorUnknownDecimalOperator
	}
}

// applyUserType parses rawVal (string) according to the user-specified annotation in userType
// (e.g. "f64", "i32", etc.) then returns the typed value plus its ArgumentType.
func (v *queryVisitor) applyUserType(rawVal string, userType string) (any, ArgumentType, error) {
	if len(userType) >= 2 && strings.HasPrefix(userType, "[") && strings.HasSuffix(userType, "]") {
		userType = userType[1 : len(userType)-1] // e.g. "f64" or "i64"
	}

	switch userType {
	case "f64":
		fl, err := strconv.ParseFloat(rawVal, 64)
		if err != nil {
			return nil, ArgTypeUnknown, ErrorInvalidValue
		}
		return fl, ArgTypeFloat64, nil

	case "d":
		dec, _ := decimal.NewFromString(fmt.Sprint(rawVal))
		return dec, ArgTypeDecimal, nil

	case "i":
		i, err := strconv.Atoi(rawVal)
		if err != nil {
			return nil, ArgTypeUnknown, ErrorInvalidValue
		}
		return i, ArgTypeInteger, nil

	case "ui":
		u64, err := strconv.ParseUint(rawVal, 10, 64)
		if err != nil {
			return nil, ArgTypeUnknown, ErrorInvalidValue
		}
		return uint(u64), ArgTypeUnsignedInteger, nil

	case "i64":
		i64, err := strconv.ParseInt(rawVal, 10, 64)
		if err != nil {
			return nil, ArgTypeUnknown, ErrorInvalidValue
		}
		return i64, ArgTypeInteger64, nil

	case "ui64":
		u64, err := strconv.ParseUint(rawVal, 10, 64)
		if err != nil {
			return nil, ArgTypeUnknown, ErrorInvalidValue
		}
		return u64, ArgTypeUnsignedInteger64, nil

	case "ui32":
		u32, err := strconv.ParseUint(rawVal, 10, 32)
		if err != nil {
			return nil, ArgTypeUnknown, ErrorInvalidValue
		}
		return uint32(u32), ArgTypeUnsignedInteger32, nil

	case "i32":
		i32, err := strconv.ParseInt(rawVal, 10, 32)
		if err != nil {
			return nil, ArgTypeUnknown, ErrorInvalidValue
		}
		return int32(i32), ArgTypeInteger32, nil

	case "s":
		return rawVal, ArgTypeString, nil

	case "f32":
		fl, err := strconv.ParseFloat(rawVal, 32)
		if err != nil {
			return nil, ArgTypeUnknown, ErrorInvalidValue
		}
		return float32(fl), ArgTypeFloat32, nil

	default:
		return nil, ArgTypeUnknown, ErrorUnknownType
	}
}

// parseValue handles the different parse-tree "Value" possibilities:
// typedVal, boolean, null, listOfInts, listOfDoubles, listOfStrings.
//
// If typedVal has a user annotation (e.g. [f64]"123.45"), it enforces strictTypeCheck for Evaluate.
func (v *queryVisitor) parseValue(valCtx parser.IValueContext) (any, ArgumentType, bool, error) {
	switch node := valCtx.(type) {

	case *parser.TypedValContext:
		return v.parseTypedValue(node.TypedValue())

	case *parser.BooleanContext:
		txt := strings.ToLower(valCtx.GetText())
		if txt == "true" {
			return true, ArgTypeBoolean, true, nil
		}
		return false, ArgTypeBoolean, true, nil

	case *parser.NullContext:
		return nil, ArgTypeNull, false, nil

	case *parser.ListOfIntsContext:
		return valCtx.GetText(), ArgTypeList, false, nil

	case *parser.ListOfDoublesContext:
		return valCtx.GetText(), ArgTypeList, false, nil

	case *parser.ListOfStringsContext:
		return valCtx.GetText(), ArgTypeList, false, nil

	default:
		return "", ArgTypeUnknown, false, ErrorInvalidValue
	}
}

// unquoteString removes the surrounding quotes of a string literal, plus minimal un-escaping.
func unquoteString(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		inner := s[1 : len(s)-1]
		inner = strings.ReplaceAll(inner, `\"`, `"`)
		inner = strings.ReplaceAll(inner, `\\`, `\`)
		return inner
	}
	return s
}

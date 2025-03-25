package rule

import (
	"fmt"
	"github.com/shopspring/decimal"
	parser "github.com/sky93/go-rule/internal/antlr4"
	"golang.org/x/exp/constraints"
	"reflect"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func (l *ErrorListener) SyntaxError(_ antlr.Recognizer, _ any, line, column int, msg string, _ antlr.RecognitionException) {
	l.hasErrors = true
	l.errMsg = NewSyntaxError(fmt.Sprintf("%d:%d: %s", line, column, msg))
}

func (g *Rule) Evaluate(values []Evaluation) (bool, error) {
	valuesMap := make(map[int]any)
	for _, value := range values {
		valuesMap[value.Param.id] = value.Result
	}
	return g.exprTree.evaluate(valuesMap, g.debugMode)
}

func (e *exprTree) evaluate(values map[int]any, debugMode bool) (bool, error) {
	if e == nil {
		return false, ErrorNoExpression
	}

	// If it's a logical node ("and"/"or")
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

	// Otherwise it's a leaf node with param
	p := e.param

	// If the user didn't provide a Value for p.Name, check if operator is "pr" (presence)
	val, ok := values[p.id]
	if !ok {
		if p.operator == "pr" {
			// 'pr' means "present". So if it's not present, that's false, or invert if not
			return e.not == true, nil
		}
		// Otherwise we can't evaluate. We'll treat it as false
		if e.not {
			return true, nil
		}
		return false, nil
	}

	// Compare
	out, err := compareOperator(val, p.operator, p.compareValue, p.strictTypeCheck)
	if debugMode {
		fmt.Printf("Name: %s, left Value: %v<%T>, Operator:%s, right Value: %v<%T>, Strict Type Check: %t, Result: %t\n", p.Name, val, val, p.operator, p.compareValue, p.compareValue, p.strictTypeCheck, out)
	}
	if err != nil {
		return false, err
	}
	if e.not {
		return !out, nil
	}
	return out, nil
}

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

// compareOperator is a helper that compares leftVal operator rightVal as strings or numeric.
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

// -------------------------------------------------------------------
// 4) Our custom visitor to build (exprTree + []Parameter) from the parse tree
// -------------------------------------------------------------------

type queryVisitor struct {
	antlr.ParseTreeVisitor // embed base
	parameters             []Parameter
	parser.BaseJsonQueryVisitor

	// We'll return the top-level exprTree
}

func (v *queryVisitor) VisitRoot(ctx *parser.RootContext) (any, error) {
	return v.Visit(ctx.Query())
}

func (v *queryVisitor) Visit(ctx parser.IQueryContext) (any, error) {
	switch actual := ctx.(type) {
	case *parser.ParenExpContext:
		return v.visitParenExp(actual)
	case *parser.LogicalExpContext:
		return v.visitLogicalExp(actual)
	case *parser.PresentExpContext:
		return v.visitPresentExp(actual)
	case *parser.CompareExpContext:
		return v.visitCompareExp(actual)
	default:
		return nil, ErrorInvalidExpression
	}
}

// parenExp => (NOT? SP?)? LPAREN (SP?)? query (SP?)? RPAREN
func (v *queryVisitor) visitParenExp(ctx *parser.ParenExpContext) (*exprTree, error) {
	sub, err := v.Visit(ctx.Query())
	if err != nil {
		return nil, err
	}
	subExp := sub.(*exprTree)
	notToken := ctx.NOT()
	if notToken != nil {
		subExp.not = !subExp.not
	}
	return subExp, nil
}

// presentExp => attrPath SP 'pr'
func (v *queryVisitor) visitPresentExp(ctx *parser.PresentExpContext) (*exprTree, error) {
	name := v.getAttrName(ctx.AttrPath())
	// We'll store "pr" as an operator meaning "present".
	p := Parameter{
		id:        len(v.parameters),
		Name:      name,
		InputType: Expression,
		operator:  "pr",
	}
	v.parameters = append(v.parameters, p)
	return &exprTree{param: &p}, nil
}

// logicalExp => query SP LOGICAL_OPERATOR SP query
func (v *queryVisitor) visitLogicalExp(ctx *parser.LogicalExpContext) (*exprTree, error) {
	leftAny, err := v.Visit(ctx.Query(0))
	if err != nil {
		return nil, err
	}
	rightAny, err := v.Visit(ctx.Query(1))
	if err != nil {
		return nil, err
	}
	leftNode, _ := leftAny.(*exprTree)
	rightNode, _ := rightAny.(*exprTree)

	op := strings.ToLower(ctx.LOGICAL_OPERATOR().GetText()) // "and" or "or"

	return &exprTree{
		op:    op,
		left:  leftNode,
		right: rightNode,
	}, nil
}

// compareExp => (attrPath | functionCall) SP operator SP Value
func (v *queryVisitor) visitCompareExp(ctx *parser.CompareExpContext) (*exprTree, error) {
	isFunc := ctx.AttrPath().FunctionCall() != nil
	var name string
	var err error
	var funcArgs []FunctionArgument

	if isFunc {
		name, funcArgs, err = v.parseFunctionCall(ctx.AttrPath().FunctionCall())
		if err != nil {
			return nil, err
		}
	} else {
		name = v.getAttrName(ctx.AttrPath())
	}
	opText := strings.ToLower(ctx.GetOp().GetText())
	valCtx := ctx.Value()
	val, valType, strict, err := v.parseValue(valCtx)
	if err != nil {
		return nil, err
	}

	p := Parameter{
		id:              len(v.parameters),
		Name:            name,
		operator:        opText,
		compareValue:    val,
		InputType:       Expression,
		Expression:      valType,
		strictTypeCheck: strict,
	}
	if isFunc {
		p.InputType = FunctionCall
		p.FunctionArguments = funcArgs
	}
	v.parameters = append(v.parameters, p)
	return &exprTree{param: &p}, nil
}

// parseFunctionCall extracts function Name + arguments
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

// getAttrName descends into attrPath -> subAttr -> ...
// For a simple parse, we can just grab ctx.GetText() or build it carefully.
func (v *queryVisitor) getAttrName(ctx parser.IAttrPathContext) string {
	return ctx.GetText()
}

func (v *queryVisitor) parseTypedValue(tv parser.ITypedValueContext) (any, ArgumentType, bool, error) {
	switch typedNode := tv.(type) {
	case *parser.TypedStringContext:
		userType := "[s]"
		strictTypeCheck := false
		if ann := typedNode.TypeAnnotation(); ann != nil {
			userType = ann.GetText() // e.g. "[f64]", "[ui32]", etc.
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

// applyUserType might do something like convert an int64 to float64 if user typed [f64], or
// parse as decimal if [d], etc.
func (v *queryVisitor) applyUserType(rawVal string, userType string) (any, ArgumentType, error) {
	if len(userType) >= 2 && strings.HasPrefix(userType, "[") && strings.HasSuffix(userType, "]") {
		userType = userType[1 : len(userType)-1] // e.g. "f64" or "i64", etc.
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

// parseValue interprets the parse tree as a go literal + ArgumentType.
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

func unquoteString(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		inner := s[1 : len(s)-1]
		// We can parse escapes if needed. For brevity, do a partial:
		inner = strings.ReplaceAll(inner, `\"`, `"`)
		inner = strings.ReplaceAll(inner, `\\`, `\`)
		return inner
	}
	return s
}

// -------------------------------------------------------------------
// 5) Public ParseQuery function
// -------------------------------------------------------------------

// ParseQuery returns the expression tree, the slice of parameters, and an error if any.
func ParseQuery(input string, config *Config) (Rule, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			// Convert the panic to an error
			switch v := r.(type) {
			case error:
				err = v // If it's already an error, use it directly
			case string:
				err = fmt.Errorf("panic occurred: %s", v) // If it's a string, wrap it in an error
			default:
				err = fmt.Errorf("unexpected panic: %v", v) // Handle other types
			}
		}
	}()

	debugMode := false
	if config != nil {
		if config.DebugMode {
			debugMode = true
		}
	}

	is := antlr.NewInputStream(input)
	lexer := parser.NewJsonQueryLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewJsonQueryParser(stream)
	// Attach custom error listener to catch syntax errors:
	errListener := &ErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)

	tree := p.Root() // parse

	if errListener.hasErrors {
		return Rule{}, errListener.errMsg
	}

	vis := &queryVisitor{}
	exprAny, err := vis.VisitRoot(tree.(*parser.RootContext))
	if err != nil {
		return Rule{}, err
	}
	expr, _ := exprAny.(*exprTree)

	return Rule{
		exprTree:  *expr,
		Params:    vis.parameters,
		debugMode: debugMode,
	}, err
}

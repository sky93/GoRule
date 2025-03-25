package rule

import (
	"github.com/antlr4-go/antlr/v4"
	parser "github.com/sky93/go-rule/internal/antlr4"
	"strings"
)

// queryVisitor is the main AST visitor used by ParseQuery to build an exprTree of logical expressions.
//
// See parse.go for queryVisitor's parseValue / parseTypedValue and other helper methods.
type queryVisitor struct {
	antlr.ParseTreeVisitor
	parameters []Parameter
	parser.BaseSCIMQueryVisitor
}

// visitRoot is an entry point for the parse tree's RootContext, returning an exprTree or error.
func (v *queryVisitor) visitRoot(ctx *parser.RootContext) (any, error) {
	return v.visit(ctx.Query())
}

// visit is a dispatcher that calls one of the typed visit methods below, based on the parse node type.
func (v *queryVisitor) visit(ctx parser.IQueryContext) (any, error) {
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

// visitParenExp handles parentheses, e.g. (someExp). A NOT token may prefix the group.
func (v *queryVisitor) visitParenExp(ctx *parser.ParenExpContext) (*exprTree, error) {
	sub, err := v.visit(ctx.Query())
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

// visitPresentExp handles a "pr" operator, e.g. "attribute pr" meaning "attribute is present".
func (v *queryVisitor) visitPresentExp(ctx *parser.PresentExpContext) (*exprTree, error) {
	name := v.getAttrName(ctx.AttrPath())
	p := Parameter{
		id:        len(v.parameters),
		Name:      name,
		InputType: Expression,
		operator:  "pr",
	}
	v.parameters = append(v.parameters, p)
	return &exprTree{param: &p}, nil
}

// visitLogicalExp handles expressions joined by "and" / "or", e.g. "query and query".
func (v *queryVisitor) visitLogicalExp(ctx *parser.LogicalExpContext) (*exprTree, error) {
	leftAny, err := v.visit(ctx.Query(0))
	if err != nil {
		return nil, err
	}
	rightAny, err := v.visit(ctx.Query(1))
	if err != nil {
		return nil, err
	}
	leftNode := leftAny.(*exprTree)
	rightNode := rightAny.(*exprTree)

	op := strings.ToLower(ctx.LOGICAL_OPERATOR().GetText())
	return &exprTree{
		op:    op,
		left:  leftNode,
		right: rightNode,
	}, nil
}

// visitCompareExp handles a single comparison: (attrPath|functionCall) operator value
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

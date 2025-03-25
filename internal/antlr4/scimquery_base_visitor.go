// Code generated from internal/antlr4/SCIMQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SCIMQuery

import "github.com/antlr4-go/antlr/v4"

type BaseSCIMQueryVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSCIMQueryVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitCompareExp(ctx *CompareExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitParenExp(ctx *ParenExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitPresentExp(ctx *PresentExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitLogicalExp(ctx *LogicalExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitAttrPath(ctx *AttrPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitArgList(ctx *ArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitSubAttr(ctx *SubAttrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitTypedString(ctx *TypedStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitTypedDouble(ctx *TypedDoubleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitTypedInteger(ctx *TypedIntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitTypedVal(ctx *TypedValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitNull(ctx *NullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitListOfInts(ctx *ListOfIntsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitListOfDoubles(ctx *ListOfDoublesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitListOfStrings(ctx *ListOfStringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitListStrings(ctx *ListStringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitSubListOfStrings(ctx *SubListOfStringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitListDoubles(ctx *ListDoublesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitSubListOfDoubles(ctx *SubListOfDoublesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitListInts(ctx *ListIntsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSCIMQueryVisitor) VisitSubListOfInts(ctx *SubListOfIntsContext) interface{} {
	return v.VisitChildren(ctx)
}

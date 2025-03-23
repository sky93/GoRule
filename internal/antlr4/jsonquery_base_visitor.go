// Code generated from internal/antlr4/JsonQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // JsonQuery

import "github.com/antlr4-go/antlr/v4"

type BaseJsonQueryVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJsonQueryVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitCompareExp(ctx *CompareExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitParenExp(ctx *ParenExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitPresentExp(ctx *PresentExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitLogicalExp(ctx *LogicalExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitAttrPath(ctx *AttrPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitArgList(ctx *ArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitSubAttr(ctx *SubAttrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitTypedString(ctx *TypedStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitTypedDouble(ctx *TypedDoubleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitTypedInteger(ctx *TypedIntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitTypedVal(ctx *TypedValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitNull(ctx *NullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitVersion(ctx *VersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListOfInts(ctx *ListOfIntsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListOfDoubles(ctx *ListOfDoublesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListOfStrings(ctx *ListOfStringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListStrings(ctx *ListStringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitSubListOfStrings(ctx *SubListOfStringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListDoubles(ctx *ListDoublesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitSubListOfDoubles(ctx *SubListOfDoublesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListInts(ctx *ListIntsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitSubListOfInts(ctx *SubListOfIntsContext) interface{} {
	return v.VisitChildren(ctx)
}

// Code generated from internal/antlr4/SCIMQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SCIMQuery

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SCIMQueryParser.
type SCIMQueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SCIMQueryParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#compareExp.
	VisitCompareExp(ctx *CompareExpContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#parenExp.
	VisitParenExp(ctx *ParenExpContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#presentExp.
	VisitPresentExp(ctx *PresentExpContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#logicalExp.
	VisitLogicalExp(ctx *LogicalExpContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#attrPath.
	VisitAttrPath(ctx *AttrPathContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#typeAnnotation.
	VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#argList.
	VisitArgList(ctx *ArgListContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#subAttr.
	VisitSubAttr(ctx *SubAttrContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#typedString.
	VisitTypedString(ctx *TypedStringContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#typedDouble.
	VisitTypedDouble(ctx *TypedDoubleContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#typedInteger.
	VisitTypedInteger(ctx *TypedIntegerContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#typedVal.
	VisitTypedVal(ctx *TypedValContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#boolean.
	VisitBoolean(ctx *BooleanContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#null.
	VisitNull(ctx *NullContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#listOfInts.
	VisitListOfInts(ctx *ListOfIntsContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#listOfDoubles.
	VisitListOfDoubles(ctx *ListOfDoublesContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#listOfStrings.
	VisitListOfStrings(ctx *ListOfStringsContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#listStrings.
	VisitListStrings(ctx *ListStringsContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#subListOfStrings.
	VisitSubListOfStrings(ctx *SubListOfStringsContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#listDoubles.
	VisitListDoubles(ctx *ListDoublesContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#subListOfDoubles.
	VisitSubListOfDoubles(ctx *SubListOfDoublesContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#listInts.
	VisitListInts(ctx *ListIntsContext) interface{}

	// Visit a parse tree produced by SCIMQueryParser#subListOfInts.
	VisitSubListOfInts(ctx *SubListOfIntsContext) interface{}
}

// Code generated from internal/antlr4/JsonQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // JsonQuery

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by JsonQueryParser.
type JsonQueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JsonQueryParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#compareExp.
	VisitCompareExp(ctx *CompareExpContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#parenExp.
	VisitParenExp(ctx *ParenExpContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#presentExp.
	VisitPresentExp(ctx *PresentExpContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#logicalExp.
	VisitLogicalExp(ctx *LogicalExpContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#attrPath.
	VisitAttrPath(ctx *AttrPathContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#typeAnnotation.
	VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#argList.
	VisitArgList(ctx *ArgListContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#subAttr.
	VisitSubAttr(ctx *SubAttrContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#typedString.
	VisitTypedString(ctx *TypedStringContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#typedDouble.
	VisitTypedDouble(ctx *TypedDoubleContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#typedInteger.
	VisitTypedInteger(ctx *TypedIntegerContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#typedVal.
	VisitTypedVal(ctx *TypedValContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#boolean.
	VisitBoolean(ctx *BooleanContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#null.
	VisitNull(ctx *NullContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#version.
	VisitVersion(ctx *VersionContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#listOfInts.
	VisitListOfInts(ctx *ListOfIntsContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#listOfDoubles.
	VisitListOfDoubles(ctx *ListOfDoublesContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#listOfStrings.
	VisitListOfStrings(ctx *ListOfStringsContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#listStrings.
	VisitListStrings(ctx *ListStringsContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#subListOfStrings.
	VisitSubListOfStrings(ctx *SubListOfStringsContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#listDoubles.
	VisitListDoubles(ctx *ListDoublesContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#subListOfDoubles.
	VisitSubListOfDoubles(ctx *SubListOfDoublesContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#listInts.
	VisitListInts(ctx *ListIntsContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#subListOfInts.
	VisitSubListOfInts(ctx *SubListOfIntsContext) interface{}
}

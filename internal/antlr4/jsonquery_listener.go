// Code generated from internal/antlr4/JsonQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // JsonQuery

import "github.com/antlr4-go/antlr/v4"

// JsonQueryListener is a complete listener for a parse tree produced by JsonQueryParser.
type JsonQueryListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterCompareExp is called when entering the compareExp production.
	EnterCompareExp(c *CompareExpContext)

	// EnterParenExp is called when entering the parenExp production.
	EnterParenExp(c *ParenExpContext)

	// EnterPresentExp is called when entering the presentExp production.
	EnterPresentExp(c *PresentExpContext)

	// EnterLogicalExp is called when entering the logicalExp production.
	EnterLogicalExp(c *LogicalExpContext)

	// EnterAttrPath is called when entering the attrPath production.
	EnterAttrPath(c *AttrPathContext)

	// EnterTypeAnnotation is called when entering the typeAnnotation production.
	EnterTypeAnnotation(c *TypeAnnotationContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterArgList is called when entering the argList production.
	EnterArgList(c *ArgListContext)

	// EnterSubAttr is called when entering the subAttr production.
	EnterSubAttr(c *SubAttrContext)

	// EnterTypedString is called when entering the typedString production.
	EnterTypedString(c *TypedStringContext)

	// EnterTypedDouble is called when entering the typedDouble production.
	EnterTypedDouble(c *TypedDoubleContext)

	// EnterTypedInteger is called when entering the typedInteger production.
	EnterTypedInteger(c *TypedIntegerContext)

	// EnterTypedVal is called when entering the typedVal production.
	EnterTypedVal(c *TypedValContext)

	// EnterBoolean is called when entering the boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterNull is called when entering the null production.
	EnterNull(c *NullContext)

	// EnterVersion is called when entering the version production.
	EnterVersion(c *VersionContext)

	// EnterListOfInts is called when entering the listOfInts production.
	EnterListOfInts(c *ListOfIntsContext)

	// EnterListOfDoubles is called when entering the listOfDoubles production.
	EnterListOfDoubles(c *ListOfDoublesContext)

	// EnterListOfStrings is called when entering the listOfStrings production.
	EnterListOfStrings(c *ListOfStringsContext)

	// EnterListStrings is called when entering the listStrings production.
	EnterListStrings(c *ListStringsContext)

	// EnterSubListOfStrings is called when entering the subListOfStrings production.
	EnterSubListOfStrings(c *SubListOfStringsContext)

	// EnterListDoubles is called when entering the listDoubles production.
	EnterListDoubles(c *ListDoublesContext)

	// EnterSubListOfDoubles is called when entering the subListOfDoubles production.
	EnterSubListOfDoubles(c *SubListOfDoublesContext)

	// EnterListInts is called when entering the listInts production.
	EnterListInts(c *ListIntsContext)

	// EnterSubListOfInts is called when entering the subListOfInts production.
	EnterSubListOfInts(c *SubListOfIntsContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitCompareExp is called when exiting the compareExp production.
	ExitCompareExp(c *CompareExpContext)

	// ExitParenExp is called when exiting the parenExp production.
	ExitParenExp(c *ParenExpContext)

	// ExitPresentExp is called when exiting the presentExp production.
	ExitPresentExp(c *PresentExpContext)

	// ExitLogicalExp is called when exiting the logicalExp production.
	ExitLogicalExp(c *LogicalExpContext)

	// ExitAttrPath is called when exiting the attrPath production.
	ExitAttrPath(c *AttrPathContext)

	// ExitTypeAnnotation is called when exiting the typeAnnotation production.
	ExitTypeAnnotation(c *TypeAnnotationContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitArgList is called when exiting the argList production.
	ExitArgList(c *ArgListContext)

	// ExitSubAttr is called when exiting the subAttr production.
	ExitSubAttr(c *SubAttrContext)

	// ExitTypedString is called when exiting the typedString production.
	ExitTypedString(c *TypedStringContext)

	// ExitTypedDouble is called when exiting the typedDouble production.
	ExitTypedDouble(c *TypedDoubleContext)

	// ExitTypedInteger is called when exiting the typedInteger production.
	ExitTypedInteger(c *TypedIntegerContext)

	// ExitTypedVal is called when exiting the typedVal production.
	ExitTypedVal(c *TypedValContext)

	// ExitBoolean is called when exiting the boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitNull is called when exiting the null production.
	ExitNull(c *NullContext)

	// ExitVersion is called when exiting the version production.
	ExitVersion(c *VersionContext)

	// ExitListOfInts is called when exiting the listOfInts production.
	ExitListOfInts(c *ListOfIntsContext)

	// ExitListOfDoubles is called when exiting the listOfDoubles production.
	ExitListOfDoubles(c *ListOfDoublesContext)

	// ExitListOfStrings is called when exiting the listOfStrings production.
	ExitListOfStrings(c *ListOfStringsContext)

	// ExitListStrings is called when exiting the listStrings production.
	ExitListStrings(c *ListStringsContext)

	// ExitSubListOfStrings is called when exiting the subListOfStrings production.
	ExitSubListOfStrings(c *SubListOfStringsContext)

	// ExitListDoubles is called when exiting the listDoubles production.
	ExitListDoubles(c *ListDoublesContext)

	// ExitSubListOfDoubles is called when exiting the subListOfDoubles production.
	ExitSubListOfDoubles(c *SubListOfDoublesContext)

	// ExitListInts is called when exiting the listInts production.
	ExitListInts(c *ListIntsContext)

	// ExitSubListOfInts is called when exiting the subListOfInts production.
	ExitSubListOfInts(c *SubListOfIntsContext)
}

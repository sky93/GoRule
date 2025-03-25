// Code generated from internal/antlr4/SCIMQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SCIMQuery

import "github.com/antlr4-go/antlr/v4"

// BaseSCIMQueryListener is a complete listener for a parse tree produced by SCIMQueryParser.
type BaseSCIMQueryListener struct{}

var _ SCIMQueryListener = &BaseSCIMQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSCIMQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSCIMQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSCIMQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSCIMQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseSCIMQueryListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseSCIMQueryListener) ExitRoot(ctx *RootContext) {}

// EnterCompareExp is called when production compareExp is entered.
func (s *BaseSCIMQueryListener) EnterCompareExp(ctx *CompareExpContext) {}

// ExitCompareExp is called when production compareExp is exited.
func (s *BaseSCIMQueryListener) ExitCompareExp(ctx *CompareExpContext) {}

// EnterParenExp is called when production parenExp is entered.
func (s *BaseSCIMQueryListener) EnterParenExp(ctx *ParenExpContext) {}

// ExitParenExp is called when production parenExp is exited.
func (s *BaseSCIMQueryListener) ExitParenExp(ctx *ParenExpContext) {}

// EnterPresentExp is called when production presentExp is entered.
func (s *BaseSCIMQueryListener) EnterPresentExp(ctx *PresentExpContext) {}

// ExitPresentExp is called when production presentExp is exited.
func (s *BaseSCIMQueryListener) ExitPresentExp(ctx *PresentExpContext) {}

// EnterLogicalExp is called when production logicalExp is entered.
func (s *BaseSCIMQueryListener) EnterLogicalExp(ctx *LogicalExpContext) {}

// ExitLogicalExp is called when production logicalExp is exited.
func (s *BaseSCIMQueryListener) ExitLogicalExp(ctx *LogicalExpContext) {}

// EnterAttrPath is called when production attrPath is entered.
func (s *BaseSCIMQueryListener) EnterAttrPath(ctx *AttrPathContext) {}

// ExitAttrPath is called when production attrPath is exited.
func (s *BaseSCIMQueryListener) ExitAttrPath(ctx *AttrPathContext) {}

// EnterTypeAnnotation is called when production typeAnnotation is entered.
func (s *BaseSCIMQueryListener) EnterTypeAnnotation(ctx *TypeAnnotationContext) {}

// ExitTypeAnnotation is called when production typeAnnotation is exited.
func (s *BaseSCIMQueryListener) ExitTypeAnnotation(ctx *TypeAnnotationContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseSCIMQueryListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseSCIMQueryListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterArgList is called when production argList is entered.
func (s *BaseSCIMQueryListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BaseSCIMQueryListener) ExitArgList(ctx *ArgListContext) {}

// EnterSubAttr is called when production subAttr is entered.
func (s *BaseSCIMQueryListener) EnterSubAttr(ctx *SubAttrContext) {}

// ExitSubAttr is called when production subAttr is exited.
func (s *BaseSCIMQueryListener) ExitSubAttr(ctx *SubAttrContext) {}

// EnterTypedString is called when production typedString is entered.
func (s *BaseSCIMQueryListener) EnterTypedString(ctx *TypedStringContext) {}

// ExitTypedString is called when production typedString is exited.
func (s *BaseSCIMQueryListener) ExitTypedString(ctx *TypedStringContext) {}

// EnterTypedDouble is called when production typedDouble is entered.
func (s *BaseSCIMQueryListener) EnterTypedDouble(ctx *TypedDoubleContext) {}

// ExitTypedDouble is called when production typedDouble is exited.
func (s *BaseSCIMQueryListener) ExitTypedDouble(ctx *TypedDoubleContext) {}

// EnterTypedInteger is called when production typedInteger is entered.
func (s *BaseSCIMQueryListener) EnterTypedInteger(ctx *TypedIntegerContext) {}

// ExitTypedInteger is called when production typedInteger is exited.
func (s *BaseSCIMQueryListener) ExitTypedInteger(ctx *TypedIntegerContext) {}

// EnterTypedVal is called when production typedVal is entered.
func (s *BaseSCIMQueryListener) EnterTypedVal(ctx *TypedValContext) {}

// ExitTypedVal is called when production typedVal is exited.
func (s *BaseSCIMQueryListener) ExitTypedVal(ctx *TypedValContext) {}

// EnterBoolean is called when production boolean is entered.
func (s *BaseSCIMQueryListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production boolean is exited.
func (s *BaseSCIMQueryListener) ExitBoolean(ctx *BooleanContext) {}

// EnterNull is called when production null is entered.
func (s *BaseSCIMQueryListener) EnterNull(ctx *NullContext) {}

// ExitNull is called when production null is exited.
func (s *BaseSCIMQueryListener) ExitNull(ctx *NullContext) {}

// EnterListOfInts is called when production listOfInts is entered.
func (s *BaseSCIMQueryListener) EnterListOfInts(ctx *ListOfIntsContext) {}

// ExitListOfInts is called when production listOfInts is exited.
func (s *BaseSCIMQueryListener) ExitListOfInts(ctx *ListOfIntsContext) {}

// EnterListOfDoubles is called when production listOfDoubles is entered.
func (s *BaseSCIMQueryListener) EnterListOfDoubles(ctx *ListOfDoublesContext) {}

// ExitListOfDoubles is called when production listOfDoubles is exited.
func (s *BaseSCIMQueryListener) ExitListOfDoubles(ctx *ListOfDoublesContext) {}

// EnterListOfStrings is called when production listOfStrings is entered.
func (s *BaseSCIMQueryListener) EnterListOfStrings(ctx *ListOfStringsContext) {}

// ExitListOfStrings is called when production listOfStrings is exited.
func (s *BaseSCIMQueryListener) ExitListOfStrings(ctx *ListOfStringsContext) {}

// EnterListStrings is called when production listStrings is entered.
func (s *BaseSCIMQueryListener) EnterListStrings(ctx *ListStringsContext) {}

// ExitListStrings is called when production listStrings is exited.
func (s *BaseSCIMQueryListener) ExitListStrings(ctx *ListStringsContext) {}

// EnterSubListOfStrings is called when production subListOfStrings is entered.
func (s *BaseSCIMQueryListener) EnterSubListOfStrings(ctx *SubListOfStringsContext) {}

// ExitSubListOfStrings is called when production subListOfStrings is exited.
func (s *BaseSCIMQueryListener) ExitSubListOfStrings(ctx *SubListOfStringsContext) {}

// EnterListDoubles is called when production listDoubles is entered.
func (s *BaseSCIMQueryListener) EnterListDoubles(ctx *ListDoublesContext) {}

// ExitListDoubles is called when production listDoubles is exited.
func (s *BaseSCIMQueryListener) ExitListDoubles(ctx *ListDoublesContext) {}

// EnterSubListOfDoubles is called when production subListOfDoubles is entered.
func (s *BaseSCIMQueryListener) EnterSubListOfDoubles(ctx *SubListOfDoublesContext) {}

// ExitSubListOfDoubles is called when production subListOfDoubles is exited.
func (s *BaseSCIMQueryListener) ExitSubListOfDoubles(ctx *SubListOfDoublesContext) {}

// EnterListInts is called when production listInts is entered.
func (s *BaseSCIMQueryListener) EnterListInts(ctx *ListIntsContext) {}

// ExitListInts is called when production listInts is exited.
func (s *BaseSCIMQueryListener) ExitListInts(ctx *ListIntsContext) {}

// EnterSubListOfInts is called when production subListOfInts is entered.
func (s *BaseSCIMQueryListener) EnterSubListOfInts(ctx *SubListOfIntsContext) {}

// ExitSubListOfInts is called when production subListOfInts is exited.
func (s *BaseSCIMQueryListener) ExitSubListOfInts(ctx *SubListOfIntsContext) {}

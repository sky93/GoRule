// Code generated from internal/antlr4/JsonQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // JsonQuery

import "github.com/antlr4-go/antlr/v4"

// BaseJsonQueryListener is a complete listener for a parse tree produced by JsonQueryParser.
type BaseJsonQueryListener struct{}

var _ JsonQueryListener = &BaseJsonQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJsonQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJsonQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJsonQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJsonQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseJsonQueryListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseJsonQueryListener) ExitRoot(ctx *RootContext) {}

// EnterCompareExp is called when production compareExp is entered.
func (s *BaseJsonQueryListener) EnterCompareExp(ctx *CompareExpContext) {}

// ExitCompareExp is called when production compareExp is exited.
func (s *BaseJsonQueryListener) ExitCompareExp(ctx *CompareExpContext) {}

// EnterParenExp is called when production parenExp is entered.
func (s *BaseJsonQueryListener) EnterParenExp(ctx *ParenExpContext) {}

// ExitParenExp is called when production parenExp is exited.
func (s *BaseJsonQueryListener) ExitParenExp(ctx *ParenExpContext) {}

// EnterPresentExp is called when production presentExp is entered.
func (s *BaseJsonQueryListener) EnterPresentExp(ctx *PresentExpContext) {}

// ExitPresentExp is called when production presentExp is exited.
func (s *BaseJsonQueryListener) ExitPresentExp(ctx *PresentExpContext) {}

// EnterLogicalExp is called when production logicalExp is entered.
func (s *BaseJsonQueryListener) EnterLogicalExp(ctx *LogicalExpContext) {}

// ExitLogicalExp is called when production logicalExp is exited.
func (s *BaseJsonQueryListener) ExitLogicalExp(ctx *LogicalExpContext) {}

// EnterAttrPath is called when production attrPath is entered.
func (s *BaseJsonQueryListener) EnterAttrPath(ctx *AttrPathContext) {}

// ExitAttrPath is called when production attrPath is exited.
func (s *BaseJsonQueryListener) ExitAttrPath(ctx *AttrPathContext) {}

// EnterTypeAnnotation is called when production typeAnnotation is entered.
func (s *BaseJsonQueryListener) EnterTypeAnnotation(ctx *TypeAnnotationContext) {}

// ExitTypeAnnotation is called when production typeAnnotation is exited.
func (s *BaseJsonQueryListener) ExitTypeAnnotation(ctx *TypeAnnotationContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseJsonQueryListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseJsonQueryListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterArgList is called when production argList is entered.
func (s *BaseJsonQueryListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BaseJsonQueryListener) ExitArgList(ctx *ArgListContext) {}

// EnterSubAttr is called when production subAttr is entered.
func (s *BaseJsonQueryListener) EnterSubAttr(ctx *SubAttrContext) {}

// ExitSubAttr is called when production subAttr is exited.
func (s *BaseJsonQueryListener) ExitSubAttr(ctx *SubAttrContext) {}

// EnterTypedString is called when production typedString is entered.
func (s *BaseJsonQueryListener) EnterTypedString(ctx *TypedStringContext) {}

// ExitTypedString is called when production typedString is exited.
func (s *BaseJsonQueryListener) ExitTypedString(ctx *TypedStringContext) {}

// EnterTypedDouble is called when production typedDouble is entered.
func (s *BaseJsonQueryListener) EnterTypedDouble(ctx *TypedDoubleContext) {}

// ExitTypedDouble is called when production typedDouble is exited.
func (s *BaseJsonQueryListener) ExitTypedDouble(ctx *TypedDoubleContext) {}

// EnterTypedInteger is called when production typedInteger is entered.
func (s *BaseJsonQueryListener) EnterTypedInteger(ctx *TypedIntegerContext) {}

// ExitTypedInteger is called when production typedInteger is exited.
func (s *BaseJsonQueryListener) ExitTypedInteger(ctx *TypedIntegerContext) {}

// EnterTypedVal is called when production typedVal is entered.
func (s *BaseJsonQueryListener) EnterTypedVal(ctx *TypedValContext) {}

// ExitTypedVal is called when production typedVal is exited.
func (s *BaseJsonQueryListener) ExitTypedVal(ctx *TypedValContext) {}

// EnterBoolean is called when production boolean is entered.
func (s *BaseJsonQueryListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production boolean is exited.
func (s *BaseJsonQueryListener) ExitBoolean(ctx *BooleanContext) {}

// EnterNull is called when production null is entered.
func (s *BaseJsonQueryListener) EnterNull(ctx *NullContext) {}

// ExitNull is called when production null is exited.
func (s *BaseJsonQueryListener) ExitNull(ctx *NullContext) {}

// EnterVersion is called when production version is entered.
func (s *BaseJsonQueryListener) EnterVersion(ctx *VersionContext) {}

// ExitVersion is called when production version is exited.
func (s *BaseJsonQueryListener) ExitVersion(ctx *VersionContext) {}

// EnterListOfInts is called when production listOfInts is entered.
func (s *BaseJsonQueryListener) EnterListOfInts(ctx *ListOfIntsContext) {}

// ExitListOfInts is called when production listOfInts is exited.
func (s *BaseJsonQueryListener) ExitListOfInts(ctx *ListOfIntsContext) {}

// EnterListOfDoubles is called when production listOfDoubles is entered.
func (s *BaseJsonQueryListener) EnterListOfDoubles(ctx *ListOfDoublesContext) {}

// ExitListOfDoubles is called when production listOfDoubles is exited.
func (s *BaseJsonQueryListener) ExitListOfDoubles(ctx *ListOfDoublesContext) {}

// EnterListOfStrings is called when production listOfStrings is entered.
func (s *BaseJsonQueryListener) EnterListOfStrings(ctx *ListOfStringsContext) {}

// ExitListOfStrings is called when production listOfStrings is exited.
func (s *BaseJsonQueryListener) ExitListOfStrings(ctx *ListOfStringsContext) {}

// EnterListStrings is called when production listStrings is entered.
func (s *BaseJsonQueryListener) EnterListStrings(ctx *ListStringsContext) {}

// ExitListStrings is called when production listStrings is exited.
func (s *BaseJsonQueryListener) ExitListStrings(ctx *ListStringsContext) {}

// EnterSubListOfStrings is called when production subListOfStrings is entered.
func (s *BaseJsonQueryListener) EnterSubListOfStrings(ctx *SubListOfStringsContext) {}

// ExitSubListOfStrings is called when production subListOfStrings is exited.
func (s *BaseJsonQueryListener) ExitSubListOfStrings(ctx *SubListOfStringsContext) {}

// EnterListDoubles is called when production listDoubles is entered.
func (s *BaseJsonQueryListener) EnterListDoubles(ctx *ListDoublesContext) {}

// ExitListDoubles is called when production listDoubles is exited.
func (s *BaseJsonQueryListener) ExitListDoubles(ctx *ListDoublesContext) {}

// EnterSubListOfDoubles is called when production subListOfDoubles is entered.
func (s *BaseJsonQueryListener) EnterSubListOfDoubles(ctx *SubListOfDoublesContext) {}

// ExitSubListOfDoubles is called when production subListOfDoubles is exited.
func (s *BaseJsonQueryListener) ExitSubListOfDoubles(ctx *SubListOfDoublesContext) {}

// EnterListInts is called when production listInts is entered.
func (s *BaseJsonQueryListener) EnterListInts(ctx *ListIntsContext) {}

// ExitListInts is called when production listInts is exited.
func (s *BaseJsonQueryListener) ExitListInts(ctx *ListIntsContext) {}

// EnterSubListOfInts is called when production subListOfInts is entered.
func (s *BaseJsonQueryListener) EnterSubListOfInts(ctx *SubListOfIntsContext) {}

// ExitSubListOfInts is called when production subListOfInts is exited.
func (s *BaseJsonQueryListener) ExitSubListOfInts(ctx *SubListOfIntsContext) {}

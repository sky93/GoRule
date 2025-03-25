package rule

import "github.com/antlr4-go/antlr/v4"

// InputType indicates if a Parameter is for a function call (FunctionCall) or a direct attribute expression (Expression).
type InputType int

const (
	// FunctionCall indicates the Parameter references a user-defined function plus arguments, e.g. get_user("abc").
	FunctionCall InputType = iota

	// Expression indicates a standard attribute-based comparison, e.g. age gt 30.
	Expression
)

// String returns the string representation of the InputType, either "FunctionCall" or "Expression".
func (it InputType) String() string {
	if it == FunctionCall {
		return "FunctionCall"
	}
	return "Expression"
}

// ArgumentType indicates how to interpret a Parameter or FunctionArgument's value (string, int, decimal, etc.).
type ArgumentType int

const (
	// ArgTypeUnknown means the type could not be determined or does not match known annotations.
	ArgTypeUnknown ArgumentType = iota

	// ArgTypeString indicates a string type.
	ArgTypeString

	// ArgTypeInteger indicates a standard int type (no specific bit width).
	ArgTypeInteger

	// ArgTypeUnsignedInteger indicates a standard uint type (no specific bit width).
	ArgTypeUnsignedInteger

	// ArgTypeFloat64 indicates a float64 type.
	ArgTypeFloat64

	// ArgTypeBoolean indicates a bool type.
	ArgTypeBoolean

	// ArgTypeNull indicates a nil or null value.
	ArgTypeNull

	// ArgTypeList indicates some bracketed list syntax, e.g. [1,2,3].
	ArgTypeList

	// ArgTypeFloat32 indicates a float32 type.
	ArgTypeFloat32

	// ArgTypeInteger32 indicates a 32-bit integer type (int32).
	ArgTypeInteger32

	// ArgTypeInteger64 indicates a 64-bit integer type (int64).
	ArgTypeInteger64

	// ArgTypeUnsignedInteger64 indicates a 64-bit unsigned integer type (uint64).
	ArgTypeUnsignedInteger64

	// ArgTypeUnsignedInteger32 indicates a 32-bit unsigned integer type (uint32).
	ArgTypeUnsignedInteger32

	// ArgTypeDecimal indicates a decimal.Decimal type (from shopspring/decimal).
	ArgTypeDecimal
)

// argToString maps ArgumentType to a short descriptor for debugging or logging.
var argToString = map[ArgumentType]string{
	ArgTypeUnknown:           "unknown",
	ArgTypeString:            "string",
	ArgTypeInteger:           "int",
	ArgTypeUnsignedInteger:   "uint",
	ArgTypeFloat64:           "float64",
	ArgTypeBoolean:           "bool",
	ArgTypeNull:              "nil",
	ArgTypeList:              "list",
	ArgTypeFloat32:           "float32",
	ArgTypeInteger32:         "int32",
	ArgTypeInteger64:         "int64",
	ArgTypeUnsignedInteger64: "uint64",
	ArgTypeUnsignedInteger32: "uint32",
	ArgTypeDecimal:           "decimal",
}

// String returns the string representation of the ArgumentType (for debugging).
func (at ArgumentType) String() string {
	if s, ok := argToString[at]; ok {
		return s
	}
	return "unknown"
}

// FunctionArgument represents a single argument in a function call, e.g. get_author("abc", 123).
type FunctionArgument struct {
	ArgumentType ArgumentType // The determined type (ArgTypeString, ArgTypeInteger, etc.)
	Value        any          // The actual parsed value
}

// Parameter holds all information needed to represent a single condition or function invocation
// within a query. For example, "age gt 18" or "get_author("Book") eq "Walt Whitman".
//
// Fields:
//   - id: unique ID for internal evaluation references
//   - Name: the attribute name or function name
//   - InputType: either Expression or FunctionCall
//   - FunctionArguments: slice of arguments if it's a function call
//   - strictTypeCheck: indicates if type annotations must be strictly enforced
//   - Expression: if InputType == Expression, describes the type annotation discovered
//   - operator: the SCIM-like operator (eq, gt, etc.)
//   - compareValue: the RHS value for the comparison (number, string, decimal, etc.)
type Parameter struct {
	id                int
	Name              string
	InputType         InputType
	FunctionArguments []FunctionArgument
	strictTypeCheck   bool
	Expression        ArgumentType
	operator          string
	compareValue      any
}

// exprTree is an internal node in the expression tree built during parsing.
//
// Fields:
//   - not: indicates a NOT operation on this node
//   - op: the logical operator ("and"/"or") or "" for leaves
//   - left, right: subtrees if op is non-empty
//   - param: if this is a leaf node, param references a Parameter to evaluate
type exprTree struct {
	not   bool
	op    string
	left  *exprTree
	right *exprTree
	param *Parameter
}

// errorListener is an ANTLR error listener capturing syntax issues.
type errorListener struct {
	*antlr.DefaultErrorListener
	hasErrors bool
	errMsg    error
}

// Rule represents the final compiled query. After parsing, a Rule contains:
//   - Params: all discovered parameters
//   - exprTree: the root of the expression tree for logical ops
//   - debugMode: flag enabling debug prints during evaluation
type Rule struct {
	Params    []Parameter
	exprTree  exprTree
	debugMode bool
}

// Config controls optional ParseQuery() behaviors.
//
// Fields:
//   - DebugMode: if true, evaluation debug lines are printed to stdout
type Config struct {
	DebugMode bool
}

// Evaluation couples a parsed Parameter with an actual value for runtime evaluation.
// The Evaluate() method will iterate over these pairs to resolve the final query result.
type Evaluation struct {
	Param  Parameter
	Result any
}

package rule

import "github.com/antlr4-go/antlr/v4"

// InputType indicates if a Parameter is for a function call or a simple attribute expression.
type InputType int

const (
	FunctionCall InputType = iota
	Expression
)

func (it InputType) String() string  {
	if it == FunctionCall {
		return "FunctionCall"
	}
	return "Expression"
}

// ArgumentType indicates how to interpret the argument's Value.
type ArgumentType int

const (
	ArgTypeUnknown ArgumentType = iota
	ArgTypeString  ArgumentType = iota
	ArgTypeInteger
	ArgTypeUnsignedInteger
	ArgTypeFloat64
	ArgTypeBoolean
	ArgTypeNull
	ArgTypeList
	ArgTypeFloat32
	ArgTypeInteger32
	ArgTypeInteger64
	ArgTypeUnsignedInteger64
	ArgTypeUnsignedInteger32
	ArgTypeDecimal
)

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

func (at ArgumentType) String() string {
	if s, ok := argToString[at]; ok {
		return s
	}
	return "unknown"
}

// FunctionArgument represents a single argument in a function call, e.g. (30, "margin").
type FunctionArgument struct {
	ArgumentType ArgumentType
	Value        any
}

// Parameter represents either a function call or an attribute expression.
//
// For example, "tradeVolDay(30, "margin") dgt "100"" is a function call parameter.
// Meanwhile, "usr_id eq 17" is an attribute expression parameter.
//
// The fields Name, operator, compareValue, etc. let the user see what was parsed.
//
//   - If InputType == FunctionCall, then functionArguments may be set
//     (like the 30, "margin" in tradeVolDay(30, "margin")).
//   - If InputType == Expression, then operator & compareValue represent a direct comparison
//     (like eq "someValue" or gt 100).
type Parameter struct {
	id                int
	Name              string
	InputType         InputType
	FunctionArguments []FunctionArgument
	strictTypeCheck   bool

	// If InputType == Expression, this indicates the type of compareValue (string, int, etc.).
	Expression ArgumentType

	// Operator is one of eq, ne, gt, lt, ge, le, co, sw, ew, in, dgt, etc.
	operator     string
	compareValue any
}

type exprTree struct {
	Not   bool      // if true, invert the result of this node
	Op    string    // "", "and", or "or"
	Left  *exprTree // used if Op is logical (and/or)
	Right *exprTree // used if Op is logical (and/or)
	Param *Parameter
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	hasErrors bool
	errMsg    error
}

type GoRule struct {
	exprTree  exprTree
	Params    []Parameter
	debugMode bool
}

type Config struct {
	DebugMode bool
}

type Evaluation struct {
	Param  Parameter
	Result any
}

package rule

import "github.com/antlr4-go/antlr/v4"

// InputType indicates if a Parameter is for a function call or a simple attribute expression.
type InputType int

const (
	FunctionCall InputType = iota
	Expression
)

// String returns the string representation of the InputType, either "FunctionCall" or "Expression", based on its value.
func (it InputType) String() string {
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

// String returns the string representation of the ArgumentType based on a predefined map or "unknown" if not found.
func (at ArgumentType) String() string {
	if s, ok := argToString[at]; ok {
		return s
	}
	return "unknown"
}

// FunctionArgument represents a single argument in a function call, e.g. (1, "Whitman").
type FunctionArgument struct {
	ArgumentType ArgumentType
	Value        any
}

// Parameter represents either a function call or an attribute expression.
//
// For example, "user_age(12345) gt 18" is a function call parameter.
// Meanwhile, "user_age eq 18" is an attribute expression parameter.
//
// The fields Name lets the user see what was parsed.
//
//   - If InputType == FunctionCall, then functionArguments may be set
//     (like the 1, "Whitman" in user_age(1, "Whitman")).
//   - If InputType == Expression, then operator & compareValue represent a direct comparison
//     (like eq "someValue" or gt 18).
type Parameter struct {
	id                int
	Name              string
	InputType         InputType
	FunctionArguments []FunctionArgument
	strictTypeCheck   bool
	// If InputType == Expression, this indicates the type of compareValue (string, int, etc.).
	Expression   ArgumentType
	operator     string
	compareValue any
}

type exprTree struct {
	not  bool
	op    string
	left  *exprTree
	right *exprTree
	param *Parameter
}

// errorListener is a custom error listener that captures syntax errors during parsing and stores error information.
type errorListener struct {
	*antlr.DefaultErrorListener
	hasErrors bool
	errMsg    error
}

// Rule represents a rule containing a slice of Parameter for evaluation and processing.
type Rule struct {
	Params    []Parameter
	exprTree  exprTree
	debugMode bool
}

// Config represents a configuration with options to control behavior such as enabling debug mode.
type Config struct {
	DebugMode bool
}

// Evaluation represents a parameter-result pair used in the evaluation process.
// The Param field holds the input Parameter being evaluated.
// The Result field holds the evaluation output for the given Parameter.
type Evaluation struct {
	Param  Parameter
	Result any
}

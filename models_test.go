package rule

import (
	"errors"
	"testing"
)

// TestInputTypeString verifies the String() method on InputType
func TestInputTypeString(t *testing.T) {
	cases := []struct {
		it       InputType
		expected string
	}{
		{FunctionCall, "FunctionCall"},
		{Expression, "Expression"},
	}

	for i, c := range cases {
		got := c.it.String()
		if got != c.expected {
			t.Errorf("[%d] InputType(%v).String() => %q, want %q",
				i, c.it, got, c.expected)
		}
	}
}

// TestArgumentTypeString verifies the String() method on ArgumentType
func TestArgumentTypeString(t *testing.T) {
	cases := []struct {
		at       ArgumentType
		expected string
	}{
		{ArgTypeUnknown, "unknown"},
		{ArgTypeString, "string"},
		{ArgTypeInteger, "int"},
		{ArgTypeUnsignedInteger, "uint"},
		{ArgTypeFloat64, "float64"},
		{ArgTypeBoolean, "bool"},
		{ArgTypeNull, "nil"},
		{ArgTypeList, "list"},
		{ArgTypeFloat32, "float32"},
		{ArgTypeInteger32, "int32"},
		{ArgTypeInteger64, "int64"},
		{ArgTypeUnsignedInteger64, "uint64"},
		{ArgTypeUnsignedInteger32, "uint32"},
		{ArgTypeDecimal, "decimal"},
	}

	for i, c := range cases {
		got := c.at.String()
		if got != c.expected {
			t.Errorf("[%d] ArgumentType(%v).String() => %q, want %q",
				i, c.at, got, c.expected)
		}
	}
}

// TestFunctionArgumentAndParameter just instantiates these structs.
func TestFunctionArgumentAndParameter(t *testing.T) {
	// Create a function argument
	arg := FunctionArgument{
		ArgumentType: ArgTypeString,
		Value:        "test",
	}
	if arg.ArgumentType != ArgTypeString {
		t.Errorf("expected ArgTypeString, got %v", arg.ArgumentType)
	}
	if arg.Value != "test" {
		t.Errorf("expected Value='test', got %v", arg.Value)
	}

	// Create a Parameter
	param := Parameter{
		id:                1,
		Name:              "testParam",
		InputType:         FunctionCall,
		FunctionArguments: []FunctionArgument{arg},
		strictTypeCheck:   true,
		Expression:        ArgTypeBoolean,
		operator:          "eq",
		compareValue:      true,
	}
	if param.id != 1 {
		t.Errorf("expected id=1, got %d", param.id)
	}
	if param.Name != "testParam" {
		t.Errorf("expected Name='testParam', got %q", param.Name)
	}
	if param.InputType != FunctionCall {
		t.Errorf("expected InputType=FunctionCall, got %v", param.InputType)
	}
	if len(param.FunctionArguments) != 1 {
		t.Errorf("expected 1 FunctionArgument, got %d", len(param.FunctionArguments))
	}
	if !param.strictTypeCheck {
		t.Errorf("expected strictTypeCheck=true")
	}
	if param.Expression != ArgTypeBoolean {
		t.Errorf("expected Expression=ArgTypeBoolean, got %v", param.Expression)
	}
	if param.operator != "eq" {
		t.Errorf("expected operator='eq', got %q", param.operator)
	}
	if val, ok := param.compareValue.(bool); !ok || val != true {
		t.Errorf("expected compareValue=true (bool), got %v", param.compareValue)
	}
}

// TestExprTree instantiates a small exprTree node.
func TestExprTree(t *testing.T) {
	p := &Parameter{Name: "dummy"}
	tree := exprTree{
		not:   true,
		op:    "and",
		left:  nil,
		right: nil,
		param: p,
	}
	if !tree.not {
		t.Errorf("expected not=true")
	}
	if tree.op != "and" {
		t.Errorf("expected op='and', got %q", tree.op)
	}
	if tree.param != p {
		t.Errorf("exprTree param mismatch")
	}
}

// TestErrorListener ensures we can set the hasErrors / errMsg properly
func TestErrorListener(t *testing.T) {
	el := &errorListener{}
	el.SyntaxError(nil, nil, 1, 2, "test error", nil)

	if !el.hasErrors {
		t.Errorf("expected hasErrors=true after SyntaxError()")
	}
	if el.errMsg == nil {
		t.Errorf("expected errMsg to be set")
	} else if el.errMsg.Error() == "" {
		t.Errorf("expected non-empty error message")
	}
}

// TestRuleAndConfig does minimal usage coverage for these structs
func TestRuleAndConfig(t *testing.T) {
	r := Rule{
		Params:    []Parameter{},
		exprTree:  exprTree{},
		debugMode: true,
	}
	if !r.debugMode {
		t.Error("expected debugMode=true")
	}

	conf := Config{DebugMode: false}
	if conf.DebugMode {
		t.Error("expected DebugMode=false")
	}
}

// TestEvaluation ensures referencing the struct covers basic lines
func TestEvaluation(t *testing.T) {
	p := Parameter{Name: "evaluationParam"}
	e := Evaluation{
		Param:  p,
		Result: 123,
	}
	if e.Param.Name != "evaluationParam" {
		t.Errorf("got unexpected param name %q", e.Param.Name)
	}
	if e.Result != 123 {
		t.Errorf("expected result=123, got %v", e.Result)
	}
}

// TestNonExistentArgumentType ensures a fallback if we do something invalid
func TestNonExistentArgumentType(t *testing.T) {
	var weirdArgumentType ArgumentType = 9999 // not in map
	if weirdArgumentType.String() != "unknown" {
		t.Errorf("expected 'unknown' for out-of-range ArgumentType, got %q",
			weirdArgumentType.String())
	}
}

// TestNonExistentInputType ensures coverage for an out-of-range InputType
func TestNonExistentInputType(t *testing.T) {
	var weirdInputType InputType = 9999
	if got := weirdInputType.String(); got != "Expression" {
		t.Errorf("expected fallback to 'Expression' for unknown InputType, got %q", got)
	}
}

// TestCustomANTLRListener just a quick check that errorListener can handle other methods (no-op)
func TestCustomANTLRListener(t *testing.T) {
	el := &errorListener{}
	el.ReportAmbiguity(nil, nil, 0, 0, false, nil, nil)
	el.ReportAttemptingFullContext(nil, nil, 0, 0, nil, nil)
	el.ReportContextSensitivity(nil, nil, 0, 0, 0, nil)
	// no panics => pass
}

// TestErrorListenerCustomError ensures we can set our own error to simulate a parse fail
func TestErrorListenerCustomError(t *testing.T) {
	el := &errorListener{}
	el.hasErrors = true
	el.errMsg = errors.New("custom error for coverage")
	if el.errMsg.Error() != "custom error for coverage" {
		t.Errorf("errorListener errMsg not matching expected message")
	}
}

package goRule

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
	"testing"
)

// TestParseQueryAndEvaluate is a higher-level “integration” style test
// that covers both the parsing phase and the evaluation phase.
func TestParseQueryAndEvaluate(t *testing.T) {
	tests := []struct {
		name        string
		query       string
		evals       []Evaluation
		wantResult  bool
		wantErr     bool
		errContains string
	}{
		{
			name:  "Basic eq success",
			query: `usr_id eq 17`,
			evals: []Evaluation{
				{
					Param: Parameter{
						id:        0,
						Name:      "usr_id",
						InputType: Expression,
					},
					Result: 17,
				},
			},
			wantResult: true,
		},
		{
			name:  "Basic eq fail",
			query: `usr_id eq 17`,
			evals: []Evaluation{
				{
					Param: Parameter{
						id:        0,
						Name:      "usr_id",
						InputType: Expression,
					},
					Result: 18,
				},
			},
			wantResult: false,
		},
		{
			name:  "Absent param => false",
			query: `someParam eq 123`,
			// no evals => param missing
			evals:      nil,
			wantResult: false,
		},
		{
			name:  "Present operator => true if param present",
			query: `someParam pr`,
			evals: []Evaluation{
				{
					Param: Parameter{
						id:       0,
						Name:     "someParam",
						operator: "pr", // the code sets this
					},
					Result: 99,
				},
			},
			wantResult: true,
		},
		{
			name:       "Present operator => false if param missing",
			query:      `someParam pr`,
			evals:      []Evaluation{},
			wantResult: false,
		},
		{
			name:  "String operators - co",
			query: `usr_name co "Doe"`,
			evals: []Evaluation{
				{
					Param: Parameter{
						id:   0,
						Name: "usr_name",
					},
					Result: "John Doe",
				},
			},
			wantResult: true,
		},
		{
			name:  "String operators - sw, false scenario",
			query: `country sw "Br"`,
			evals: []Evaluation{
				{
					Param: Parameter{
						id:   0,
						Name: "country",
					},
					Result: "Argentina",
				},
			},
			wantResult: false,
		},
		{
			name:  "Boolean eq true",
			query: `is_ready eq true`,
			evals: []Evaluation{
				{
					Param: Parameter{
						id:   0,
						Name: "is_ready",
					},
					Result: true,
				},
			},
			wantResult: true,
		},
		{
			name:  "Complex logic with parentheses",
			query: `(usr_id eq 17 and age gt 20) or usr_id eq 99`,
			evals: []Evaluation{
				{Param: Parameter{id: 0, Name: "usr_id"}, Result: 17},
				{Param: Parameter{id: 1, Name: "age"}, Result: 21},
			},
			wantResult: true,
		},
		{
			name:  "Decimal typed with annotation => eq",
			query: `amount eq [d]"12.34"`,
			evals: []Evaluation{
				{
					Param: Parameter{
						id:   0,
						Name: "amount",
					},
					Result: decimal.NewFromFloat(12.34),
				},
			},
			wantResult: true,
		},
		{
			name:  "Decimal typed mismatch => error in strict mode",
			query: `amount eq [d]"12.34"`,
			evals: []Evaluation{
				{
					Param: Parameter{
						id:   0,
						Name: "amount",
						// Suppose user gave an int, so mismatch
					},
					Result: 12,
				},
			},
			wantErr:     true,
			errContains: "compare type mismatch",
		},
		{
			name:  "Arithmetic operators test => le",
			query: `score le 100`,
			evals: []Evaluation{
				{
					Param: Parameter{
						id:   0,
						Name: "score",
					},
					Result: 99,
				},
			},
			wantResult: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 1) Parse
			rule, err := ParseQuery(tc.query, nil)
			if err != nil {
				if !tc.wantErr {
					t.Fatalf("unexpected parse error: %v", err)
				}
				// If we DO expect an error, check the message
				if tc.errContains != "" && !strings.Contains(err.Error(), tc.errContains) {
					t.Fatalf("error mismatch, want substring %q, got: %v", tc.errContains, err)
				}
				return
			}

			// 2) Evaluate
			got, err := rule.Evaluate(tc.evals)
			if err != nil {
				if !tc.wantErr {
					t.Fatalf("unexpected eval error: %v", err)
				}
				if tc.errContains != "" && !strings.Contains(err.Error(), tc.errContains) {
					t.Fatalf("error mismatch, want substring %q, got: %v", tc.errContains, err)
				}
				return
			}
			if tc.wantErr {
				t.Fatalf("expected error but got none")
			}
			if got != tc.wantResult {
				t.Errorf("Evaluate() got %v, want %v", got, tc.wantResult)
			}
		})
	}
}

// TestStrictTypeCheckSomeCasts checks a few “annotation vs actual” combos.
func TestStrictTypeCheckSomeCasts(t *testing.T) {
	// eq, with annotation
	query := `val eq [f64]"123.456"`
	rule, err := ParseQuery(query, nil)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	// Provide float64 => should pass
	evals := []Evaluation{
		{
			Param:  rule.Params[0],
			Result: 123.456,
		},
	}
	ok, err := rule.Evaluate(evals)
	if err != nil {
		t.Errorf("evaluate error: %v", err)
	}
	if !ok {
		t.Errorf("expected true with matching float64")
	}

	// Provide int => mismatch error
	evals = []Evaluation{
		{
			Param:  rule.Params[0],
			Result: 123,
		},
	}
	_, err = rule.Evaluate(evals)
	if err == nil {
		t.Errorf("expected type mismatch error but got nil")
	}
}

// TestFunctionCall covers usage of InputType=FunctionCall with arguments
func TestFunctionCall(t *testing.T) {
	query := `testFunc(15, "abc") eq 1`
	rule, err := ParseQuery(query, nil)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if len(rule.Params) != 1 {
		t.Fatalf("expected 1 param, got %d", len(rule.Params))
	}
	param := rule.Params[0]
	if param.InputType != FunctionCall {
		t.Fatalf("expected function call param, got: %v", param.InputType)
	}
	if param.Name != "testFunc" {
		t.Errorf("expected name=testFunc, got: %s", param.Name)
	}
	if len(param.FunctionArguments) != 2 {
		t.Errorf("expected 2 function arguments, got %d", len(param.FunctionArguments))
	}
	// Evaluate => we must supply param.value (the function result)
	// plus we *optionally* confirm we got the same function arguments:
	if param.FunctionArguments[0].Value != int64(15) {
		t.Errorf("expected first arg=15, got: %+v", param.FunctionArguments[0].Value)
	}
	if param.FunctionArguments[1].Value != "abc" {
		t.Errorf("expected second arg='abc', got: %+v", param.FunctionArguments[1].Value)
	}

	// Provide the final function result => 1
	evals := []Evaluation{
		{Param: param, Result: 1},
	}
	got, err := rule.Evaluate(evals)
	if err != nil {
		t.Fatalf("evaluation error: %v", err)
	}
	if !got {
		t.Fatalf("expected true but got false")
	}
}

// TestBadQueries checks parse-time errors
func TestBadQueries(t *testing.T) {
	bad := []string{
		``,                      // empty
		`(`,                     // incomplete
		`usr_id eq`,             // missing compare value
		`usr_id eq "abc" extra`, // trailing
		`((usr_id eq 1)`,        // unbalanced parens
		`x in`,                  // missing RHS
		`functionCall(,5) eq 1`, // weird comma
	}
	for _, b := range bad {
		_, err := ParseQuery(b, nil)
		if err == nil {
			t.Errorf("expected parse error for %q, got nil", b)
		}
	}
}

// TestUnknownOperator ensures we handle unknown operators at evaluation time
func TestUnknownOperator(t *testing.T) {
	// We can forcibly place an unknown operator in the Parameter, or parse something that
	// your grammar doesn't (but let's do it forcibly).
	rule := GoRule{
		exprTree: exprTree{
			Param: &Parameter{
				id:         0,
				Name:       "dummy",
				operator:   "xxx", // not recognized
				InputType:  Expression,
				Expression: ArgTypeString,
			},
		},
		Params: []Parameter{
			{
				id:         0,
				Name:       "dummy",
				operator:   "xxx",
				InputType:  Expression,
				Expression: ArgTypeString,
			},
		},
	}
	// Provide "some string"
	evals := []Evaluation{
		{
			Param:  rule.Params[0],
			Result: "hello",
		},
	}
	_, err := rule.Evaluate(evals)
	if err == nil {
		t.Error("expected error for unknown operator, got nil")
	}
	if !strings.Contains(err.Error(), "invalid operator") {
		t.Errorf("error should mention invalid operator, got: %v", err)
	}
}

// Example usage test (just to show quick usage)
func ExampleParseQuery() {
	rule, err := ParseQuery(`(active eq true) and (score gt 50)`, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Supply actual param values:
	evals := []Evaluation{
		{
			Param:  rule.Params[0], // active
			Result: true,
		},
		{
			Param:  rule.Params[1], // score
			Result: 75,
		},
	}
	ok, err := rule.Evaluate(evals)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(ok)
	// Output: true
}

func BenchmarkParseQuery(b *testing.B) {
	query := `(usr_id eq 15 or usr_id eq 99) and amount gt [d]"12345.67"`
	for i := 0; i < b.N; i++ {
		_, err := ParseQuery(query, nil)
		if err != nil {
			b.Fatalf("ParseQuery error: %v", err)
		}
	}
}

func BenchmarkEvaluate(b *testing.B) {
	query := `(usr_id eq 17 or usr_id eq 99) and amount gt [d]"12345.67"`
	rule, err := ParseQuery(query, nil)
	if err != nil {
		b.Fatalf("ParseQuery error: %v", err)
	}

	// example data for evaluation
	evals := []Evaluation{
		{Param: rule.Params[0], Result: int64(17)},
		{Param: rule.Params[1], Result: int64(99999)}, // amount -> decimal
	}

	b.ResetTimer() // ignore parse time in benchmark

	for i := 0; i < b.N; i++ {
		_, err := rule.Evaluate(evals)
		if err != nil {
			b.Fatalf("Evaluate error: %v", err)
		}
	}
}

package rule

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestParseTypedValueAnnotations(t *testing.T) {
	table := []struct {
		queryPart string
		wantType  ArgumentType
		wantVal   any
	}{
		{`[i32]"123"`, ArgTypeInteger32, int32(123)},
		{`[i64]"456"`, ArgTypeInteger64, int64(456)},
		{`[f64]"3.14"`, ArgTypeFloat64, 3.14},
		{`[f32]"2.5"`, ArgTypeFloat32, float32(2.5)},
		{`[ui64]"42"`, ArgTypeUnsignedInteger64, uint64(42)},
		{`[ui]"42"`, ArgTypeUnsignedInteger, uint(42)},
		{`[d]"12.34"`, ArgTypeDecimal, decimal.NewFromFloat(12.34)},
		{`[s]"Hello"`, ArgTypeString, "Hello"},
	}

	for _, tc := range table {
		// We'll embed this typedValue into a simple eq comparison for full parse:
		query := `myParam eq ` + tc.queryPart
		r, err := ParseQuery(query, nil)
		if err != nil {
			t.Errorf("ParseQuery(%q) => unexpected error: %v", query, err)
			continue
		}
		if len(r.Params) != 1 {
			t.Errorf("ParseQuery(%q) => expected 1 param, got %d", query, len(r.Params))
			continue
		}
		gotParam := r.Params[0]
		gotVal := gotParam.compareValue
		gotType := gotParam.Expression

		if gotType != tc.wantType {
			t.Errorf("expected ArgType=%v, got %v", tc.wantType, gotType)
		}
		// Compare actual value
		// For decimals or floats, do approximate checks or use .Equal():
		switch v := tc.wantVal.(type) {
		case decimal.Decimal:
			decVal, ok := gotVal.(decimal.Decimal)
			if !ok || !decVal.Equal(v) {
				t.Errorf("expected decimal %v, got %v", v, gotVal)
			}
		default:
			if gotVal != v {
				t.Errorf("expected value %v, got %v", v, gotVal)
			}
		}
	}
}

func TestParseFunctionCall(t *testing.T) {
	// Test parse of a function call with multiple arguments (some typed).
	query := `myFunc([i64]"42", "text", [f64]"3.14") eq 1`
	r, err := ParseQuery(query, nil)
	if err != nil {
		t.Errorf("ParseQuery error: %v", err)
		return
	}
	if len(r.Params) != 1 {
		t.Errorf("expected 1 param, got %d", len(r.Params))
		return
	}
	p := r.Params[0]
	if p.Name != "myFunc" {
		t.Errorf("expected function name 'myFunc', got %q", p.Name)
	}
	if p.InputType != FunctionCall {
		t.Errorf("expected FunctionCall, got %s", p.InputType.String())
	}
	if len(p.FunctionArguments) != 3 {
		t.Errorf("expected 3 arguments, got %d", len(p.FunctionArguments))
	}
}

func TestParseNestedParentheses(t *testing.T) {
	query := `((age gt 18) and (score lt 100)) or (status eq "active")`
	r, err := ParseQuery(query, nil)
	if err != nil {
		t.Errorf("ParseQuery error: %v", err)
	}
	// We won't evaluate; just confirm it parsed 3 parameters.
	if len(r.Params) != 3 {
		t.Errorf("expected 3 params, got %d", len(r.Params))
	}
}

func TestParseInvalidQueries(t *testing.T) {
	queries := []string{
		``,                        // empty
		`(age gt 18`,              // missing closing paren
		`someParam eq`,            // missing value
		`myFunc(,) eq 2`,          // invalid comma usage
		`unknownOp(12) abc "xyz"`, // unknown op "abc" => parse error
	}

	for _, q := range queries {
		_, err := ParseQuery(q, nil)
		if err == nil {
			t.Errorf("expected parse error for %q but got none", q)
		}
	}
}

package rule

import (
	"github.com/shopspring/decimal"
	"strings"
	"testing"
)

func TestCompareDecimal(t *testing.T) {
	left := decimal.NewFromFloat(10.5)
	right := decimal.NewFromFloat(10.5)

	cases := []struct {
		op   string
		want bool
	}{
		{"eq", true},
		{"==", true},
		{"ne", false},
		{"!=", false},
		{"gt", false},
		{">", false},
		{"lt", false},
		{"<", false},
		{"ge", true},
		{">=", true},
		{"le", true},
		{"<=", true},
	}

	for _, c := range cases {
		got, err := compareDecimal(left, c.op, right)
		if err != nil {
			t.Errorf("compareDecimal(%v, %s, %v) => unexpected error: %v", left, c.op, right, err)
			continue
		}
		if got != c.want {
			t.Errorf("compareDecimal(%v, %s, %v) => %v, want %v", left, c.op, right, got, c.want)
		}
	}

	// Try an unknown operator
	_, err := compareDecimal(left, "??", right)
	if err == nil {
		t.Error("expected error for unknown operator, got nil")
	}
}

func TestCompareOperator_Basics(t *testing.T) {
	// This covers compareOperator() with numeric, string, boolean, etc.
	// We can call compareOperator() directly to test all branches.
	type testCase struct {
		left          any
		operator      string
		right         any
		strict        bool
		wantBool      bool
		wantErrSubstr string // optional substring check in error
	}

	tests := []testCase{
		// Integers
		{int64(10), "eq", int64(10), true, true, ""},
		{int64(10), "eq", 10, false, true, ""}, // non-strict, will cast
		{int64(10), "ne", int64(5), true, true, ""},
		{int64(10), "gt", int64(5), true, true, ""},
		{int64(5), "lt", int64(10), true, true, ""},
		{int64(10), "ge", int64(10), true, true, ""},
		{int64(9), "le", int64(9), true, true, ""},
		// Strict mismatch
		{int64(10), "eq", 10, true, false, "mismatch"}, // strict => mismatch

		// Floats
		{float64(3.14), "eq", float64(3.14), true, true, ""},
		{float64(3.14), "eq", float32(3.14), false, true, ""}, // non-strict okay
		{float64(2.1), "gt", float64(2.0), true, true, ""},

		// Booleans
		{true, "eq", true, false, true, ""},
		{false, "ne", true, false, true, ""},
		{false, "gt", false, false, false, "invalid operator"}, // invalid

		// Strings
		{"hello", "eq", "hello", true, true, ""},
		{"hello", "eq", "world", true, false, ""},
		{"test", "co", "es", false, true, ""},         // "contains"
		{"start", "sw", "st", false, true, ""},        // "starts with"
		{"finish", "ew", "sh", false, true, ""},       // "ends with"
		{"abc", "in", "zzzzabczzzz", false, true, ""}, // "in" => check if left is in right
		{"abc", "in", "abd", false, false, ""},
		{"abc", "in", "abc", false, true, ""},
		{"anything", "pr", nil, false, true, ""}, // "present" => leftVal != nil => true, and leftVal is "anything" => tricky, must pass
	}

	for i, tc := range tests {
		got, err := compareOperator(tc.left, tc.operator, tc.right, tc.strict)
		if tc.wantErrSubstr != "" {
			if err == nil || (err != nil && !contains(err.Error(), tc.wantErrSubstr)) {
				t.Errorf("[%d] compareOperator(%v, %s, %v, strict=%t) => error=%v, want substring %q",
					i, tc.left, tc.operator, tc.right, tc.strict, err, tc.wantErrSubstr)
			}
			continue
		}
		if err != nil {
			t.Errorf("[%d] unexpected error: %v", i, err)
			continue
		}
		if got != tc.wantBool {
			t.Errorf("[%d] compareOperator(%v, %s, %v, strict=%t) got=%v, want=%v", i, tc.left, tc.operator, tc.right, tc.strict, got, tc.wantBool)
		}
	}
}

func TestCompareStringOps(t *testing.T) {
	// Directly test compareStringOps(...) for coverage
	type testCase struct {
		left, right string
		op          string
		want        bool
	}

	data := []testCase{
		{"Hello World", "World", "co", true}, // contains
		{"Hello", "He", "sw", true},          // starts with
		{"Goodbye", "bye", "ew", true},       // ends with
		{"abc", "zzzabczzz", "in", true},     // "abc" in "zzzabczzz" => true
		{"missing", "", "co", false},
		{"abc", "ABC", "co", false}, // case-sensitive
		{"One", "", "sw", false},
		{"One", "", "ew", false},
		// "pr" => just checks leftVal != nil
		{"anything", "", "pr", true},
	}

	for _, d := range data {
		got, err := compareStringOps(d.left, d.op, d.right)
		if err != nil {
			t.Errorf("compareStringOps(%q, %s, %q) => err=%v", d.left, d.op, d.right, err)
			continue
		}
		if got != d.want {
			t.Errorf("compareStringOps(%q, %s, %q) => %v, want %v", d.left, d.op, d.right, got, d.want)
		}
	}

	// Unknown operator => must fail
	_, err := compareStringOps("abc", "unknownOp", "def")
	if err == nil {
		t.Error("expected error for unknown string operator, got nil")
	}
}

// Helper for substring check:
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

package rule

import (
	"github.com/antlr4-go/antlr/v4"
	parser "github.com/sky93/go-rule/internal/antlr4"
)

// Evaluate applies the stored exprTree logic to a slice of Evaluation structs.
// Each Evaluation links a Parameter in g.Params to a real runtime value. The result is a bool.
func (g *Rule) Evaluate(values []Evaluation) (bool, error) {
	valuesMap := make(map[int]any)
	for _, value := range values {
		valuesMap[value.Param.id] = value.Result
	}
	return g.exprTree.evaluate(valuesMap, g.debugMode)
}

// ParseQuery takes a SCIM-like query (e.g. `age gt 30 and (lang eq "en" or lang eq "fr")`) and
// compiles it into a Rule object. Optional config can enable DebugMode.
//
// If parsing fails due to syntax errors or other issues, an error is returned.
// Otherwise, the returned Rule can be used for Evaluate().
func ParseQuery(input string, config *Config) (Rule, error) {
	debugMode := false
	if config != nil && config.DebugMode {
		debugMode = true
	}

	is := antlr.NewInputStream(input)
	lexer := parser.NewSCIMQueryLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSCIMQueryParser(stream)

	// Attach a custom error listener to catch syntax errors
	errListener := &errorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)

	// Parse the input
	tree := p.Root()
	if errListener.hasErrors {
		return Rule{}, errListener.errMsg
	}

	// Build internal expression tree
	vis := &queryVisitor{}
	exprAny, err := vis.visitRoot(tree.(*parser.RootContext))
	if err != nil {
		return Rule{}, err
	}

	expr, _ := exprAny.(*exprTree)
	return Rule{
		exprTree:  *expr,
		Params:    vis.parameters,
		debugMode: debugMode,
	}, err
}

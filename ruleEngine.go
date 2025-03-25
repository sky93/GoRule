package rule

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	parser "github.com/sky93/go-rule/internal/antlr4"
)

func (g *Rule) Evaluate(values []Evaluation) (bool, error) {
	valuesMap := make(map[int]any)
	for _, value := range values {
		valuesMap[value.Param.id] = value.Result
	}
	return g.exprTree.evaluate(valuesMap, g.debugMode)
}

func ParseQuery(input string, config *Config) (Rule, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			// Convert the panic to an error
			switch v := r.(type) {
			case error:
				err = v // If it's already an error, use it directly
			case string:
				err = fmt.Errorf("panic occurred: %s", v) // If it's a string, wrap it in an error
			default:
				err = fmt.Errorf("unexpected panic: %v", v) // Handle other types
			}
		}
	}()

	debugMode := false
	if config != nil {
		if config.DebugMode {
			debugMode = true
		}
	}

	is := antlr.NewInputStream(input)
	lexer := parser.NewSCIMQueryLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSCIMQueryParser(stream)
	// Attach custom error listener to catch syntax errors:
	errListener := &errorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)

	tree := p.Root() // parse

	if errListener.hasErrors {
		return Rule{}, errListener.errMsg
	}

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

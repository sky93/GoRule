package main

import (
	"fmt"
	"github.com/sky93/go-rule"
	"log"
)

func main() {
	input := `book_pages gt 100 and (language eq "en" or language eq "fr") and price pr and in_stock eq true`

	exp, err := rule.ParseQuery(input, nil)
	if err != nil {
		log.Fatalf("Parse error: %v", err)
	}

	fmt.Println("Discovered parameters:")
	values := make([]rule.Evaluation, 0)
	for i, p := range exp.Params {
		if p.Name == "book_pages" {
			values = append(values, rule.Evaluation{
				Param:  p,
				Result: 150,
			})
		}
		if p.Name == "language" {
			values = append(values, rule.Evaluation{
				Param:  p,
				Result: "en",
			})
		}
		if p.Name == "price" {
			values = append(values, rule.Evaluation{
				Param:  p,
				Result: 100,
			})
		}
		if p.Name == "in_stock" {
			values = append(values, rule.Evaluation{
				Param:  p,
				Result: true,
			})
		}

		if p.InputType == rule.FunctionCall {
			fmt.Printf("\t%d)  Name=%q\n\t\tType: %s\n\t\tFunction Args: %v\n\t\tExpected Type: %v\n\n", i+1, p.Name, p.InputType.String(), p.FunctionArguments, p.Expression.String())
		} else {
			fmt.Printf("\t%d)  Name=%q\n\t\tType: %s\n\t\tExpected Type: %v\n\n", i+1, p.Name, p.InputType.String(), p.Expression.String())
		}
	}

	res, err := exp.Evaluate(values)
	if err != nil {
		log.Fatalf("Evaluation error: %v", err)
	}
	fmt.Printf("Evaluation => %v\n", res)
}

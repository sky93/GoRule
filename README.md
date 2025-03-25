![go-rule logo](/logo.png)


# go-rule

[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/sky93/go-rule)](https://goreportcard.com/report/github.com/sky93/go-rule)
![Go Version](https://img.shields.io/badge/go_version-1.23+-green)
[![Coverage Status](https://coveralls.io/repos/github/sky93/go-rule/badge.svg)](https://coveralls.io/github/sky93/go-rule)
[![Go Reference](https://pkg.go.dev/badge/github.com/sky93/go-rule.svg)](https://pkg.go.dev/github.com/sky93/go-rule)

A **lightweight rule parsing and evaluation library** for Go. Define human-readable queries (e.g. `score gt 100 and active eq true`) and evaluate them against real-world data. Supports **parentheses**, **logical operators** (`and`, `or`, `not`), **type annotations** (`[i64]`, `[f64]`, `[d]`, etc.), **function calls**, and more.

> **Key features**
> - Simple query language (operators like `eq`, `gt`, `lt`, `pr`, `in`, `co`, `sw`, `ew`, etc.).
> - Built on [ANTLR4](https://github.com/antlr4-go/antlr) grammar.
> - Support for typed values (`[i64]"123"`, `[f64]"123.45"`, `[d]"12.34"`) and typed comparisons (strict type checks).
> - Handle decimals via [shopspring/decimal](https://github.com/shopspring/decimal).
> - Evaluate queries with dynamic data (like JSON objects, struct fields, or custom function results).
> - Easily embed in your own projects and run rule-based filtering or validations.

---

## Table of Contents

1. [Installation](#installation)
2. [Quick Start Example](#quick-start-example)
3. [Usage](#usage)
    - [Parsing Queries](#parsing-queries)
    - [Evaluating a Parsed Rule](#evaluating-a-parsed-rule)
    - [Working with Typed Values](#working-with-typed-values)
    - [Function Calls](#function-calls)
    - [Supported Operators](#supported-operators)
    - [Debug/Logging](#debuglogging)
4. [Advanced Examples](#advanced-examples)
5. [Testing](#testing)
6. [Full API Documentation](#full-api-documentation)
7. [License](#license)

---

## Installation

```bash
go get github.com/sky93/go-rule
```

---

## Quick Start Example

There's a self-contained example in [`_example/simple/main.go`](./_example/simple/main.go).

Below is a short version:

```go
package main

import (
    "fmt"
    "log"

    "github.com/sky93/go-rule"
)

func main() {
    // A simple query
    input := `book_pages gt 100 and (language eq "en" or language eq "fr") and price pr and in_stock eq true`

    // Parse the query
    parsedRule, err := rule.ParseQuery(input, nil)
    if err != nil {
        log.Fatalf("Parse error: %v", err)
    }

    fmt.Println("Discovered parameters:")
    // Prepare evaluations with actual data
    values := []rule.Evaluation{
        {
            Param:  parsedRule.Params[0], // e.g. "book_pages"
            Result: 150,
        },
        {
            Param:  parsedRule.Params[1], // e.g. "language"
            Result: "en",
        },
        {
            Param:  parsedRule.Params[2], // e.g. "price"
            Result: 10.0,
        },
        {
            Param:  parsedRule.Params[3], // e.g. "in_stock"
            Result: true,
        },
    }

    // Evaluate
    ok, err := parsedRule.Evaluate(values)
    if err != nil {
        log.Fatalf("Evaluation error: %v", err)
    }
    fmt.Printf("Evaluation => %v\n", ok)
}
```

**Output**:
```
Discovered parameters:
 ... (various parameter info) ...
Evaluation => true
```

---

## Usage

### Parsing Queries

Use `rule.ParseQuery(queryString, config)` to parse a textual rule:

```go
import "github.com/sky93/go-rule"

ruleSet, err := rule.ParseQuery(
    `(usr_id eq 100 or usr_id eq 101) and amount gt [d]"12.34"`, 
    nil,
)
// ruleSet is a `Rule` struct with an internal expression tree + discovered Params.
```

- **`queryString`** can contain:
    - **Comparison**: `attrName operator value`
    - **Logical**: `( ... ) and/or ( ... )`, plus `not`
    - **Type annotation**: `[i64]"123"`, `[f64]"123.45"`, `[d]"12.34"`, `[ui]` (and more)
    - **Presence**: `someField pr` (true if field is present)
    - **String operators**: `co` (contains), `sw` (starts with), `ew` (ends with), `in`

**Parsing Errors**  
If the syntax is invalid, `ParseQuery` returns an error. For example, unbalanced parentheses or unknown tokens.

### Evaluating a Parsed Rule

Once parsed, you get a `rule.Rule` that contains:

- `Params`: The discovered parameters (name, operator, typed expression, etc.).

To evaluate:

1. **Identify each parameter** in `ruleSet.Params`.
2. Construct a slice of `rule.Evaluation` items, each linking a **Param** from `ruleSet.Params` to an actual **Result** from your data.
3. Call `ruleSet.Evaluate(evals []Evaluation)` => returns `(bool, error)`.

```go
ok, err := ruleSet.Evaluate([]rule.Evaluation{
    {
        Param:  ruleSet.Params[0], // param with Name="age"
        Result: 20,
    },
    {
        Param:  ruleSet.Params[1], // param with Name="can_drive"
        Result: true,
    },
})
fmt.Println(ok)   // => true
fmt.Println(err)  // => nil
```

### Working with Typed Values

You can specify a type annotation in the query, for example `[i64]"123"`, `[f64]"123.45"`, `[d]"12.34"`.

**Strict Type Checking**:  
If a query param is `[f64]"123.45"`, the library enforces that your provided `Result` is `float64`, otherwise you'll get a type mismatch error.

**Supported type annotations**:

| Annotation | Meaning / Go Type                                          |
|------------|------------------------------------------------------------|
| `[i64]`    | `int64`                                                    |
| `[ui64]`   | `uint64`                                                   |
| `[i]`      | `int`                                                      |
| `[ui]`     | `uint`                                                     |
| `[i32]`    | `int32`                                                    |
| `[ui32]`   | `uint32`                                                   |
| `[f64]`    | `float64`                                                  |
| `[f32]`    | `float32`                                                  |
| `[d]`      | [`decimal.Decimal`](https://github.com/shopspring/decimal) |
| `[s]`      | `string`                                                   |

### Function Calls

You can have queries like:
```
get_author("Song of Myself") eq "Walt Whitman"
```
Here:
- `get_author` is the function name.
- `"Song of Myself"` is function argument.
- `eq` is equal operator.
- `"Walt Whitman"` is the compare value.

`ParseQuery` sets `Parameter.InputType = FunctionCall` with `Parameter.FunctionArguments`.  
For final evaluation, you must supply a single numeric/string/boolean (etc.) `Result` for `Param`:

```go
// If we have: get_author("Song of Myself") eq "Walt Whitman"
param := ruleSet.Params[0]
// param.FunctionArguments => e.g. [ {ArgTypeString, Value: "Song of Myself"} ]

ok, err := ruleSet.Evaluate([]rule.Evaluation{
  {
    Param:  param,
    Result: "Walt Whitman", // the real function result
  },
})
// => true
```

### Supported Operators

| Operator | Meaning                  |
|----------|--------------------------|
| `eq`     | equals                   |
| `ne`     | not equals               |
| `gt`     | greater than             |
| `lt`     | less than                |
| `ge`     | greater or equal         |
| `le`     | less or equal            |
| `co`     | contains (substring)     |
| `sw`     | starts with              |
| `ew`     | ends with                |
| `in`     | "in" check (substring)   |
| `pr`     | present (non-nil check)  |

**Logical**: `and`, `or`, plus optional `not` prefix.  
**Parentheses**: `( expr )`

### Debug/Logging

You can pass a config with `DebugMode=true`:

```go
ruleSet, err := rule.ParseQuery(`age gt 18`, &rule.Config{DebugMode: true})
if err != nil {
    log.Fatal(err)
}
// Evaluate => will print debug messages to stdout
```

---

## Advanced Examples

See [`ruleEngine_test.go`](./ruleEngine_test.go) for in-depth test scenarios:
- Decimal usage
- Complex boolean logic
- Strict type checking
- Absent parameters vs. `pr`
- Nested parentheses
- Function calls with arguments
- Error handling

Also, `_example/simple/main.go` shows a small command-line usage.

---

## Testing

This repository includes unit tests in `ruleEngine_test.go`. Run them with:

```bash
go test ./...
```

You’ll see coverage for query parsing, expression evaluation, typed operators, error handling, etc.

---

## Full API Documentation

Browse all public functions, types, and methods on **[pkg.go.dev](https://pkg.go.dev/github.com/sky93/go-rule)** or read the doc comments in the source code.

---

## License

This project is licensed under the [MIT License](LICENSE).  
Copyright &copy; 2025 [Sepehr Mohaghegh](https://github.com/sky93).

> Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files ... [full MIT license text](./LICENSE).

**Logo Attribution & Dependencies**
- The “goopher” logo is copied from [avivcarmi.com](https://avivcarmi.com/we-need-to-talk-about-the-bad-sides-of-go/).
- This library uses [shopspring/decimal](https://github.com/shopspring/decimal) and [ANTLR4 for Go](https://github.com/antlr4-go/antlr).
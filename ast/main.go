package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"io/ioutil"
)

func main() {
	r, err := ioutil.ReadFile("./.vimrc.bundle")
	if err != nil {
		panic(err)
	}

	print(string(r))

	expr, err := parser.ParseExpr(string(r))
	if err != nil {
		panic(err)
	}

	ast.Inspect(expr, func(n ast.Node) bool {
		if n != nil {
			fmt.Printf("%[1]v(%[1]T)\n", n)
		}
		return true
	})
}

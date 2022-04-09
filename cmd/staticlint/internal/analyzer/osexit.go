package analyzer

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

// OsExitCheck - the variable of type &analysis.Analyzer for creating custom
// analyzer.
var OsExitCheck = &analysis.Analyzer{
	Name: "exitcheck",
	Doc:  "interdict os.Exit",
	Run:  run,
}

// run contains the main code of checking. Accepts the pointer to struct
// analysis.Pass in parameter.
// The function tracks usage os.Exit.
func run(pass *analysis.Pass) (interface{}, error) {
	expr := func(x *ast.ExprStmt) {
		if call, ok := x.X.(*ast.CallExpr); ok {
			if selector, ok := call.Fun.(*ast.SelectorExpr); ok {
				if name, ok := selector.X.(*ast.Ident); ok {
					if name.Name == "os" {
						if fnname := selector.Sel.Name; fnname == "Exit" {
							pass.Reportf(x.Pos(), "error, os.Exit using")
						}
					}
				}
			}
		}
	}
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			if x, ok := node.(*ast.ExprStmt); ok {
				expr(x)
			}
			return true
		})
	}
	return nil, nil
}

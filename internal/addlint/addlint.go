package addlint

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

// OsExitCheck - переменная типа &analysis.Analyzer для создания пользовательского
// анализатора.
var OsExitCheck = &analysis.Analyzer{
	Name: "exitcheck",
	Doc:  "interdict os.Exit",
	Run:  run,
}

// run содержит основной код проверки. Принимает в параметре
// указатель на структуру analysis.Pass.
// Функция отслеживает использование os.Exit.
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

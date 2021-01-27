package bund

import (
	// "github.com/pieterclaerhout/go-log"

	"github.com/vulogov/Bund/internal/conf"
	"github.com/vulogov/Bund/internal/parse"
)

func Eval() {
	Init()
	parse.Parse(*conf.Expr)
}

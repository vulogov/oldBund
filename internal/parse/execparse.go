package parse

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/vulogov/Bund/internal/parser"
	"github.com/vulogov/Bund/internal/vm"
)

type bundExecListener struct {
	*parser.BaseBundListener
	vm      *vm.VM
	inBlock bool
	inPair  bool
}

func ParserExec(code string) {
	_input := antlr.NewInputStream(code)
	lexer := parser.NewBundLexer(_input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewBundParser(stream)
	listener := new(bundExecListener)
	listener.vm = vm.New("inline")
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Expressions())
}

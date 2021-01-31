package parse

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/vulogov/Bund/internal/parser"
	"github.com/vulogov/Bund/internal/vm"
)

type bundListener struct {
	*parser.BaseBundListener
	vm *vm.VM
}

func ParserPrint(code string) {
	_input := antlr.NewInputStream(code)
	lexer := parser.NewBundLexer(_input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewBundParser(stream)
	listener := new(bundListener)
	listener.vm = new(vm.VM)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Expressions())
}

func (l *bundListener) ExitPair(c *parser.PairContext) {
	fmt.Printf("pair: %s %+v \n", c.GetText(), c.GetHead())
}

func (l *bundListener) ExitDirective(c *parser.DirectiveContext) {
	fmt.Printf("directive: %s \n", c.GetText())
}

func (l *bundListener) ExitCmd(c *parser.CmdContext) {
	fmt.Printf("command: %s \n", c.GetText())
}

func (l *bundListener) ExitLambda(c *parser.LambdaContext) {
	fmt.Printf("lambda: %s \n", c.GetText())
}

func (l *bundListener) ExitTrue_term(c *parser.True_termContext) {
	fmt.Printf("TRUE: %v \n", c.GetValue().GetTokenType())
}

func (l *bundListener) ExitInteger(c *parser.IntegerContext) {
	fmt.Printf("INT: %v %v \n", c.GetValue().GetTokenType(), c.GetValue().GetText())
}

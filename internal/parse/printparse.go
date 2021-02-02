package parse

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/vulogov/Bund/internal/parser"
	"github.com/vulogov/Bund/internal/vm"
)

type bundListener struct {
	*parser.BaseBundListener
	vm      *vm.VM
	inBlock bool
	inPair  bool
}

func ParserPrint(code string) {
	_input := antlr.NewInputStream(code)
	lexer := parser.NewBundLexer(_input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewBundParser(stream)
	listener := new(bundListener)
	listener.vm = vm.New("printer")
	listener.inBlock = false
	listener.inPair = false
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Expressions())
}

func (l *bundListener) EnterBlock(c *parser.BlockContext) {
	fmt.Printf("open block\n")
}
func (l *bundListener) ExitBlock(c *parser.BlockContext) {
	fmt.Printf("close block: %v\n", c.GetValues())
}
func (l *bundListener) EnterPair(c *parser.PairContext) {
	fmt.Printf("open pair\n")
}
func (l *bundListener) ExitPair(c *parser.PairContext) {
	fmt.Printf("close pair: %s %v %v\n", c.GetText(), c.GetHead(), c.GetTail().GetText())
}
func (l *bundListener) ExitNth(c *parser.NthContext) {
	fmt.Printf("NTH: %v\n", c.GetValue())
}

func (l *bundListener) ExitDirective(c *parser.DirectiveContext) {
	fmt.Printf("directive: %s \n", c.GetText())
}

func (l *bundListener) ExitCmd(c *parser.CmdContext) {
	fmt.Printf("command: %s \n", c.GetText())
}

func (l *bundListener) ExitName_term(c *parser.Name_termContext) {
	fmt.Printf("name: %s \n", c.GetText())
}

func (l *bundListener) ExitLambda(c *parser.LambdaContext) {
	fmt.Printf("lambda: %s \n", c.GetText())
}

func (l *bundListener) ExitInteger(c *parser.IntegerContext) {
	fmt.Printf("INT: %v \n", c.GetValue().GetText())
}
func (l *bundListener) ExitFloat_term(c *parser.Float_termContext) {
	fmt.Printf("FLOAT: %v \n", c.GetValue().GetText())
}
func (l *bundListener) ExitTrue_term(c *parser.True_termContext) {
	fmt.Printf("TRUE is so TRUE\n")
}
func (l *bundListener) ExitFalse_term(c *parser.False_termContext) {
	fmt.Printf("FALSE is not TRUE\n")
}
func (l *bundListener) ExitString_term(c *parser.String_termContext) {
	fmt.Printf("STRING: %v\n", c.GetValue())
}

func (l *bundListener) ExitBegin(c *parser.BeginContext) {
	fmt.Printf("WILL PUSH TO BEGIN OF STACK\n")
}

func (l *bundListener) ExitEnd(c *parser.EndContext) {
	fmt.Printf("WILL PUSH TO END OF STACK\n")
}
func (l *bundListener) ExitDrop_term(c *parser.Drop_termContext) {
	fmt.Printf("DROP LAST VALE FROM STACK\n")
}
func (l *bundListener) ExitDuplicate_term(c *parser.Duplicate_termContext) {
	fmt.Printf("DUPLICATE LAST VALE FROM STACK\n")
}

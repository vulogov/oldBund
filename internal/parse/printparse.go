package parse

import (
	"fmt"
	"strconv"

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
	fmt.Printf("directive: %v %v \n", c.GetOp().GetText(), c.GetName().GetText())
}
func (l *bundListener) ExitSys_directive(c *parser.Sys_directiveContext) {
	fmt.Printf("SYS directive: %v %v %v \n", c.GetSys().GetText(), c.GetOp().GetText(), c.GetName().GetText())
}

func (l *bundListener) ExitCmd(c *parser.CmdContext) {
	fmt.Printf("command: %s \n", c.GetText())
}
func (l *bundListener) ExitSys_cmd(c *parser.Sys_cmdContext) {
	fmt.Printf("SYS command: %v = %v \n", c.GetSys().GetText(), c.GetCommand().GetText())
}

func (l *bundListener) ExitName_term(c *parser.Name_termContext) {
	fmt.Printf("name: %s \n", c.GetText())
}
func (l *bundListener) ExitSys_name_term(c *parser.Sys_name_termContext) {
	fmt.Printf("SYS name: %v %v \n", c.GetSys().GetText(), c.GetValue().GetText())
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
func (l *bundListener) EnterBin_integer(c *parser.Bin_integerContext) {
	v, err := strconv.ParseInt(c.GetValue().GetText(), 0, 64)
	if err == nil {
		fmt.Printf("FROMBIN: %v\n", v)
	} else {
		fmt.Printf("FROMBIN Error: %v %v\n", c.GetValue().GetText(), err)
	}
}
func (l *bundListener) EnterHex_integer(c *parser.Hex_integerContext) {
	v, err := strconv.ParseInt(c.GetValue().GetText(), 0, 64)
	if err == nil {
		fmt.Printf("FROMHEX: %v\n", v)
	} else {
		fmt.Printf("FROMHEX Error: %v %v\n", c.GetValue().GetText(), err)
	}
}
func (l *bundListener) EnterOct_integer(c *parser.Oct_integerContext) {
	v, err := strconv.ParseInt(c.GetValue().GetText(), 0, 64)
	if err == nil {
		fmt.Printf("FROMOCT: %v\n", v)
	} else {
		fmt.Printf("FROMOCT Error: %v %v\n", c.GetValue().GetText(), err)
	}
}

func (l *bundListener) ExitBegin(c *parser.BeginContext) {
	fmt.Printf("WILL PUSH TO BEGIN OF STACK\n")
}

func (l *bundListener) ExitEnd(c *parser.EndContext) {
	fmt.Printf("WILL PUSH TO END OF STACK\n")
}
func (l *bundListener) ExitDrop_term(c *parser.Drop_termContext) {
	fmt.Printf("DROP LAST VALUE FROM STACK\n")
}
func (l *bundListener) ExitDuplicate_term(c *parser.Duplicate_termContext) {
	fmt.Printf("DUPLICATE LAST VALUE FROM STACK\n")
}

func (l *bundListener) EnterExec(c *parser.ExecContext) {
	fmt.Printf("EXEC: %v\n", c.GetValue().GetText())
}

func (l *bundListener) ExitChannel_out(c *parser.Channel_outContext) {
	fmt.Printf("CHANNEL OUT: %v\n", c.GetName().GetText())
}

func (l *bundListener) ExitChannel_in(c *parser.Channel_inContext) {
	fmt.Printf("CHANNEL IN: %v\n", c.GetName().GetText())
}

func (l *bundListener) ExitSys_channel_out(c *parser.Sys_channel_outContext) {
	fmt.Printf("SYS CHANNEL OUT: %v %v\n", c.GetSys().GetText(), c.GetName().GetText())
}

func (l *bundListener) ExitSys_channel_in(c *parser.Sys_channel_inContext) {
	fmt.Printf("SYS CHANNEL IN: %v %v\n", c.GetSys().GetText(), c.GetName().GetText())
}

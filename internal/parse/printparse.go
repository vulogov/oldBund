package parse

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/Bund/internal/parser"
	"github.com/vulogov/Bund/internal/vm"
)

type bundListener struct {
	*parser.BaseBundListener
	vm      *vm.VM
	inBlock bool
	inPair  bool
	errors  int
}

type bundErrorListener struct {
	antlr.ErrorListener
	errors int
}

func ParserPrint(code string) {
	errorListener := new(bundErrorListener)
	_input := antlr.NewInputStream(code)
	lexer := parser.NewBundLexer(_input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewBundParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)
	listener := new(bundListener)
	listener.vm = vm.New("printer")
	listener.inBlock = false
	listener.inPair = false
	if errorListener.errors > 0 {
		log.Errorf("%v lexer errors detected.", errorListener.errors)
		return
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Expressions())
	if errorListener.errors > 0 {
		log.Errorf("Errors detected: %v", errorListener.errors)
	} else {
		log.Debug("No errors detected")
	}
}

func (l *bundErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	log.Errorf("Syntax error line=%v, column=%v : %v", line, column, msg)
	l.errors += 1
}
func (l *bundErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	l.errors += 1
}
func (l *bundErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	l.errors += 1
}
func (l *bundErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	l.errors += 1
}

func (l *bundListener) EnterBlock(c *parser.BlockContext) {
	log.Info("OPEN BLOCK")
}
func (l *bundListener) ExitBlock(c *parser.BlockContext) {
	log.Infof("CLOSE BLOCK")
}
func (l *bundListener) EnterPair(c *parser.PairContext) {
	log.Info("OPEN PAIR")
}
func (l *bundListener) ExitPair(c *parser.PairContext) {
	log.Infof("CLOSE PAIR: %s %v %v\n", c.GetText(), c.GetHead(), c.GetTail().GetText())
}
func (l *bundListener) ExitNth(c *parser.NthContext) {
	log.Infof("NTH: %v", c.GetValue())
}

func (l *bundListener) ExitDirective(c *parser.DirectiveContext) {
	log.Infof("DIRECTIVE: %v %v \n", c.GetOp().GetText(), c.GetName().GetText())
}
func (l *bundListener) ExitSys_directive(c *parser.Sys_directiveContext) {
	log.Infof("SYS DIRECTIVE: %v %v %v \n", c.GetSys().GetText(), c.GetOp().GetText(), c.GetName().GetText())
}

func (l *bundListener) ExitCmd(c *parser.CmdContext) {
	log.Infof("COMMAND: %s \n", c.GetText())
}
func (l *bundListener) ExitSys_cmd(c *parser.Sys_cmdContext) {
	log.Infof("SYS COMMAND: %v = %v \n", c.GetSys().GetText(), c.GetCommand().GetText())
}

func (l *bundListener) ExitName_term(c *parser.Name_termContext) {
	log.Infof("NAME: %s \n", c.GetText())
}
func (l *bundListener) ExitSys_name_term(c *parser.Sys_name_termContext) {
	log.Infof("SYS name: %v %v \n", c.GetSys().GetText(), c.GetValue().GetText())
}

func (l *bundListener) EnterLambda(c *parser.LambdaContext) {
	log.Infof("define lambda: %v \n", c.GetName().GetText())
}
func (l *bundListener) ExitLambda(c *parser.LambdaContext) {
	log.Infof("lambda: %s \n", c.GetText())
}

func (l *bundListener) EnterLambda_cmd(c *parser.Lambda_cmdContext) {
	if c == nil {
		log.Error("Empty lambda CMD")
		return
	}
	n := c.GetName()
	if n == nil {
		log.Errorf("Empty lambda name is not allowed")
		return
	}
	log.Infof("DEFINE LAMBDA COMMAND: %v \n", c.GetName().GetText())
}
func (l *bundListener) ExitLambda_cmd(c *parser.Lambda_cmdContext) {
	log.Infof("COMMAND: %s \n", c.GetText())
}

func (l *bundListener) ExitInteger(c *parser.IntegerContext) {
	log.Infof("INT: %v \n", c.GetValue().GetText())
}
func (l *bundListener) ExitFloat_term(c *parser.Float_termContext) {
	log.Infof("FLOAT: %v \n", c.GetValue().GetText())
}
func (l *bundListener) ExitTrue_term(c *parser.True_termContext) {
	log.Infof("TRUE is so TRUE\n")
}
func (l *bundListener) ExitFalse_term(c *parser.False_termContext) {
	log.Infof("FALSE is not TRUE\n")
}
func (l *bundListener) ExitString_term(c *parser.String_termContext) {
	log.Infof("STRING: %v\n", c.GetValue())
}
func (l *bundListener) EnterBin_integer(c *parser.Bin_integerContext) {
	v, err := strconv.ParseInt(c.GetValue().GetText(), 0, 64)
	if err == nil {
		log.Infof("FROMBIN: %v\n", v)
	} else {
		log.Errorf("FROMBIN Error: %v %v\n", c.GetValue().GetText(), err)
	}
}
func (l *bundListener) EnterHex_integer(c *parser.Hex_integerContext) {
	v, err := strconv.ParseInt(c.GetValue().GetText(), 0, 64)
	if err == nil {
		log.Infof("FROMHEX: %v\n", v)
	} else {
		log.Infof("FROMHEX Error: %v %v\n", c.GetValue().GetText(), err)
	}
}
func (l *bundListener) EnterOct_integer(c *parser.Oct_integerContext) {
	v, err := strconv.ParseInt(c.GetValue().GetText(), 0, 64)
	if err == nil {
		log.Infof("FROMOCT: %v\n", v)
	} else {
		log.Infof("FROMOCT Error: %v %v\n", c.GetValue().GetText(), err)
		l.errors += 1
	}
}

func (l *bundListener) ExitBegin(c *parser.BeginContext) {
	log.Warnf("WILL PUSH TO BEGIN OF STACK\n")
}

func (l *bundListener) ExitEnd(c *parser.EndContext) {
	log.Warnf("WILL PUSH TO END OF STACK\n")
}
func (l *bundListener) ExitDrop_term(c *parser.Drop_termContext) {
	log.Warnf("DROP LAST VALUE FROM STACK\n")
}
func (l *bundListener) ExitDuplicate_term(c *parser.Duplicate_termContext) {
	log.Warnf("DUPLICATE LAST VALUE FROM STACK\n")
}

func (l *bundListener) EnterExec(c *parser.ExecContext) {
	log.Warnf("EXEC: %v\n", c.GetValue().GetText())
}

func (l *bundListener) ExitChannel_out(c *parser.Channel_outContext) {
	log.Infof("CHANNEL OUT: %v\n", c.GetName().GetText())
}

func (l *bundListener) ExitChannel_in(c *parser.Channel_inContext) {
	log.Infof("CHANNEL IN: %v\n", c.GetName().GetText())
}

func (l *bundListener) ExitSys_channel_out(c *parser.Sys_channel_outContext) {
	log.Infof("SYS CHANNEL OUT: %v %v\n", c.GetSys().GetText(), c.GetName().GetText())
}

func (l *bundListener) ExitSys_channel_in(c *parser.Sys_channel_inContext) {
	log.Infof("SYS CHANNEL IN: %v %v\n", c.GetSys().GetText(), c.GetName().GetText())
}

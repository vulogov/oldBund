// Code generated from Bund.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Bund

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 26, 80, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 26,
	10, 2, 12, 2, 14, 2, 29, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 5, 3, 41, 10, 3, 3, 4, 3, 4, 7, 4, 45, 10, 4, 12,
	4, 14, 4, 48, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 6,
	6, 58, 10, 6, 13, 6, 14, 6, 59, 3, 6, 3, 6, 3, 7, 6, 7, 65, 10, 7, 13,
	7, 14, 7, 66, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 73, 10, 8, 12, 8, 14, 8, 76,
	11, 8, 3, 8, 3, 8, 3, 8, 2, 2, 9, 2, 4, 6, 8, 10, 12, 14, 2, 2, 2, 94,
	2, 27, 3, 2, 2, 2, 4, 40, 3, 2, 2, 2, 6, 42, 3, 2, 2, 2, 8, 51, 3, 2, 2,
	2, 10, 57, 3, 2, 2, 2, 12, 64, 3, 2, 2, 2, 14, 68, 3, 2, 2, 2, 16, 26,
	5, 4, 3, 2, 17, 26, 5, 6, 4, 2, 18, 26, 5, 10, 6, 2, 19, 26, 5, 12, 7,
	2, 20, 26, 5, 14, 8, 2, 21, 26, 7, 24, 2, 2, 22, 26, 7, 25, 2, 2, 23, 26,
	7, 19, 2, 2, 24, 26, 7, 20, 2, 2, 25, 16, 3, 2, 2, 2, 25, 17, 3, 2, 2,
	2, 25, 18, 3, 2, 2, 2, 25, 19, 3, 2, 2, 2, 25, 20, 3, 2, 2, 2, 25, 21,
	3, 2, 2, 2, 25, 22, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 24, 3, 2, 2, 2,
	26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 3, 3, 2,
	2, 2, 29, 27, 3, 2, 2, 2, 30, 41, 7, 9, 2, 2, 31, 41, 7, 11, 2, 2, 32,
	41, 7, 12, 2, 2, 33, 41, 7, 13, 2, 2, 34, 41, 7, 14, 2, 2, 35, 41, 7, 16,
	2, 2, 36, 41, 7, 17, 2, 2, 37, 41, 7, 15, 2, 2, 38, 41, 7, 18, 2, 2, 39,
	41, 5, 8, 5, 2, 40, 30, 3, 2, 2, 2, 40, 31, 3, 2, 2, 2, 40, 32, 3, 2, 2,
	2, 40, 33, 3, 2, 2, 2, 40, 34, 3, 2, 2, 2, 40, 35, 3, 2, 2, 2, 40, 36,
	3, 2, 2, 2, 40, 37, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 39, 3, 2, 2, 2,
	41, 5, 3, 2, 2, 2, 42, 46, 7, 3, 2, 2, 43, 45, 5, 4, 3, 2, 44, 43, 3, 2,
	2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 49,
	3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 50, 7, 3, 2, 2, 50, 7, 3, 2, 2, 2,
	51, 52, 7, 4, 2, 2, 52, 53, 5, 4, 3, 2, 53, 54, 5, 4, 3, 2, 54, 55, 7,
	5, 2, 2, 55, 9, 3, 2, 2, 2, 56, 58, 7, 22, 2, 2, 57, 56, 3, 2, 2, 2, 58,
	59, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 61, 3, 2, 2,
	2, 61, 62, 7, 18, 2, 2, 62, 11, 3, 2, 2, 2, 63, 65, 7, 23, 2, 2, 64, 63,
	3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2,
	67, 13, 3, 2, 2, 2, 68, 69, 7, 6, 2, 2, 69, 70, 7, 18, 2, 2, 70, 74, 7,
	7, 2, 2, 71, 73, 5, 4, 3, 2, 72, 71, 3, 2, 2, 2, 73, 76, 3, 2, 2, 2, 74,
	72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 77, 3, 2, 2, 2, 76, 74, 3, 2, 2,
	2, 77, 78, 7, 8, 2, 2, 78, 15, 3, 2, 2, 2, 9, 25, 27, 40, 46, 59, 66, 74,
}
var literalNames = []string{
	"", "'|'", "'('", "')'", "'['", "']'", "'.'", "", "", "", "", "", "", "",
	"", "", "", "':'", "';'", "'/'", "", "", "'^'", "'_'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "INTEGER", "DECIMAL_INTEGER", "OCT_INTEGER",
	"HEX_INTEGER", "BIN_INTEGER", "FLOAT_NUMBER", "STRING", "TRUE", "FALSE",
	"NAME", "TOBEGIN", "TOEND", "SLASH", "DIR", "CMD", "DUPLICATE", "DROP",
	"SKIP_",
}

var ruleNames = []string{
	"expressions", "term", "block", "pair", "directive", "cmd", "lambda",
}

type BundParser struct {
	*antlr.BaseParser
}

// NewBundParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *BundParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewBundParser(input antlr.TokenStream) *BundParser {
	this := new(BundParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Bund.g4"

	return this
}

// BundParser tokens.
const (
	BundParserEOF             = antlr.TokenEOF
	BundParserT__0            = 1
	BundParserT__1            = 2
	BundParserT__2            = 3
	BundParserT__3            = 4
	BundParserT__4            = 5
	BundParserT__5            = 6
	BundParserINTEGER         = 7
	BundParserDECIMAL_INTEGER = 8
	BundParserOCT_INTEGER     = 9
	BundParserHEX_INTEGER     = 10
	BundParserBIN_INTEGER     = 11
	BundParserFLOAT_NUMBER    = 12
	BundParserSTRING          = 13
	BundParserTRUE            = 14
	BundParserFALSE           = 15
	BundParserNAME            = 16
	BundParserTOBEGIN         = 17
	BundParserTOEND           = 18
	BundParserSLASH           = 19
	BundParserDIR             = 20
	BundParserCMD             = 21
	BundParserDUPLICATE       = 22
	BundParserDROP            = 23
	BundParserSKIP_           = 24
)

// BundParser rules.
const (
	BundParserRULE_expressions = 0
	BundParserRULE_term        = 1
	BundParserRULE_block       = 2
	BundParserRULE_pair        = 3
	BundParserRULE_directive   = 4
	BundParserRULE_cmd         = 5
	BundParserRULE_lambda      = 6
)

// IExpressionsContext is an interface to support dynamic dispatch.
type IExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionsContext differentiates from other interfaces.
	IsExpressionsContext()
}

type ExpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionsContext() *ExpressionsContext {
	var p = new(ExpressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_expressions
	return p
}

func (*ExpressionsContext) IsExpressionsContext() {}

func NewExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionsContext {
	var p = new(ExpressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_expressions

	return p
}

func (s *ExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExpressionsContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ExpressionsContext) AllDirective() []IDirectiveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDirectiveContext)(nil)).Elem())
	var tst = make([]IDirectiveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDirectiveContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Directive(i int) IDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *ExpressionsContext) AllCmd() []ICmdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICmdContext)(nil)).Elem())
	var tst = make([]ICmdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICmdContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Cmd(i int) ICmdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICmdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *ExpressionsContext) AllLambda() []ILambdaContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILambdaContext)(nil)).Elem())
	var tst = make([]ILambdaContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILambdaContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Lambda(i int) ILambdaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILambdaContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILambdaContext)
}

func (s *ExpressionsContext) AllDUPLICATE() []antlr.TerminalNode {
	return s.GetTokens(BundParserDUPLICATE)
}

func (s *ExpressionsContext) DUPLICATE(i int) antlr.TerminalNode {
	return s.GetToken(BundParserDUPLICATE, i)
}

func (s *ExpressionsContext) AllDROP() []antlr.TerminalNode {
	return s.GetTokens(BundParserDROP)
}

func (s *ExpressionsContext) DROP(i int) antlr.TerminalNode {
	return s.GetToken(BundParserDROP, i)
}

func (s *ExpressionsContext) AllTOBEGIN() []antlr.TerminalNode {
	return s.GetTokens(BundParserTOBEGIN)
}

func (s *ExpressionsContext) TOBEGIN(i int) antlr.TerminalNode {
	return s.GetToken(BundParserTOBEGIN, i)
}

func (s *ExpressionsContext) AllTOEND() []antlr.TerminalNode {
	return s.GetTokens(BundParserTOEND)
}

func (s *ExpressionsContext) TOEND(i int) antlr.TerminalNode {
	return s.GetToken(BundParserTOEND, i)
}

func (s *ExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterExpressions(s)
	}
}

func (s *ExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitExpressions(s)
	}
}

func (p *BundParser) Expressions() (localctx IExpressionsContext) {
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BundParserRULE_expressions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__1)|(1<<BundParserT__3)|(1<<BundParserINTEGER)|(1<<BundParserOCT_INTEGER)|(1<<BundParserHEX_INTEGER)|(1<<BundParserBIN_INTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserNAME)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND)|(1<<BundParserDIR)|(1<<BundParserCMD)|(1<<BundParserDUPLICATE)|(1<<BundParserDROP))) != 0 {
		p.SetState(23)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case BundParserT__1, BundParserINTEGER, BundParserOCT_INTEGER, BundParserHEX_INTEGER, BundParserBIN_INTEGER, BundParserFLOAT_NUMBER, BundParserSTRING, BundParserTRUE, BundParserFALSE, BundParserNAME:
			{
				p.SetState(14)
				p.Term()
			}

		case BundParserT__0:
			{
				p.SetState(15)
				p.Block()
			}

		case BundParserDIR:
			{
				p.SetState(16)
				p.Directive()
			}

		case BundParserCMD:
			{
				p.SetState(17)
				p.Cmd()
			}

		case BundParserT__3:
			{
				p.SetState(18)
				p.Lambda()
			}

		case BundParserDUPLICATE:
			{
				p.SetState(19)
				p.Match(BundParserDUPLICATE)
			}

		case BundParserDROP:
			{
				p.SetState(20)
				p.Match(BundParserDROP)
			}

		case BundParserTOBEGIN:
			{
				p.SetState(21)
				p.Match(BundParserTOBEGIN)
			}

		case BundParserTOEND:
			{
				p.SetState(22)
				p.Match(BundParserTOEND)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(BundParserINTEGER, 0)
}

func (s *TermContext) OCT_INTEGER() antlr.TerminalNode {
	return s.GetToken(BundParserOCT_INTEGER, 0)
}

func (s *TermContext) HEX_INTEGER() antlr.TerminalNode {
	return s.GetToken(BundParserHEX_INTEGER, 0)
}

func (s *TermContext) BIN_INTEGER() antlr.TerminalNode {
	return s.GetToken(BundParserBIN_INTEGER, 0)
}

func (s *TermContext) FLOAT_NUMBER() antlr.TerminalNode {
	return s.GetToken(BundParserFLOAT_NUMBER, 0)
}

func (s *TermContext) TRUE() antlr.TerminalNode {
	return s.GetToken(BundParserTRUE, 0)
}

func (s *TermContext) FALSE() antlr.TerminalNode {
	return s.GetToken(BundParserFALSE, 0)
}

func (s *TermContext) STRING() antlr.TerminalNode {
	return s.GetToken(BundParserSTRING, 0)
}

func (s *TermContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *TermContext) Pair() IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *BundParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BundParserRULE_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(38)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Match(BundParserINTEGER)
		}

	case BundParserOCT_INTEGER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(29)
			p.Match(BundParserOCT_INTEGER)
		}

	case BundParserHEX_INTEGER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(30)
			p.Match(BundParserHEX_INTEGER)
		}

	case BundParserBIN_INTEGER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(31)
			p.Match(BundParserBIN_INTEGER)
		}

	case BundParserFLOAT_NUMBER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(32)
			p.Match(BundParserFLOAT_NUMBER)
		}

	case BundParserTRUE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(33)
			p.Match(BundParserTRUE)
		}

	case BundParserFALSE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(34)
			p.Match(BundParserFALSE)
		}

	case BundParserSTRING:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(35)
			p.Match(BundParserSTRING)
		}

	case BundParserNAME:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(36)
			p.Match(BundParserNAME)
		}

	case BundParserT__1:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(37)
			p.Pair()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetValues returns the values rule context list.
	GetValues() []ITermContext

	// SetValues sets the values rule context list.
	SetValues([]ITermContext)

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_term  ITermContext
	values []ITermContext
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Get_term() ITermContext { return s._term }

func (s *BlockContext) Set_term(v ITermContext) { s._term = v }

func (s *BlockContext) GetValues() []ITermContext { return s.values }

func (s *BlockContext) SetValues(v []ITermContext) { s.values = v }

func (s *BlockContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *BlockContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *BundParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BundParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(BundParserT__0)
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__1)|(1<<BundParserINTEGER)|(1<<BundParserOCT_INTEGER)|(1<<BundParserHEX_INTEGER)|(1<<BundParserBIN_INTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserNAME))) != 0 {
		{
			p.SetState(41)

			var _x = p.Term()

			localctx.(*BlockContext)._term = _x
		}
		localctx.(*BlockContext).values = append(localctx.(*BlockContext).values, localctx.(*BlockContext)._term)

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(47)
		p.Match(BundParserT__0)
	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHead returns the head rule contexts.
	GetHead() ITermContext

	// GetTail returns the tail rule contexts.
	GetTail() ITermContext

	// SetHead sets the head rule contexts.
	SetHead(ITermContext)

	// SetTail sets the tail rule contexts.
	SetTail(ITermContext)

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	head   ITermContext
	tail   ITermContext
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) GetHead() ITermContext { return s.head }

func (s *PairContext) GetTail() ITermContext { return s.tail }

func (s *PairContext) SetHead(v ITermContext) { s.head = v }

func (s *PairContext) SetTail(v ITermContext) { s.tail = v }

func (s *PairContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *PairContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *BundParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BundParserRULE_pair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(BundParserT__1)
	}
	{
		p.SetState(50)

		var _x = p.Term()

		localctx.(*PairContext).head = _x
	}
	{
		p.SetState(51)

		var _x = p.Term()

		localctx.(*PairContext).tail = _x
	}
	{
		p.SetState(52)
		p.Match(BundParserT__2)
	}

	return localctx
}

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_DIR returns the _DIR token.
	Get_DIR() antlr.Token

	// GetName returns the name token.
	GetName() antlr.Token

	// Set_DIR sets the _DIR token.
	Set_DIR(antlr.Token)

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetOp returns the op token list.
	GetOp() []antlr.Token

	// SetOp sets the op token list.
	SetOp([]antlr.Token)

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_DIR   antlr.Token
	op     []antlr.Token
	name   antlr.Token
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_directive
	return p
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) Get_DIR() antlr.Token { return s._DIR }

func (s *DirectiveContext) GetName() antlr.Token { return s.name }

func (s *DirectiveContext) Set_DIR(v antlr.Token) { s._DIR = v }

func (s *DirectiveContext) SetName(v antlr.Token) { s.name = v }

func (s *DirectiveContext) GetOp() []antlr.Token { return s.op }

func (s *DirectiveContext) SetOp(v []antlr.Token) { s.op = v }

func (s *DirectiveContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *DirectiveContext) AllDIR() []antlr.TerminalNode {
	return s.GetTokens(BundParserDIR)
}

func (s *DirectiveContext) DIR(i int) antlr.TerminalNode {
	return s.GetToken(BundParserDIR, i)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterDirective(s)
	}
}

func (s *DirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitDirective(s)
	}
}

func (p *BundParser) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BundParserRULE_directive)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == BundParserDIR {
		{
			p.SetState(54)

			var _m = p.Match(BundParserDIR)

			localctx.(*DirectiveContext)._DIR = _m
		}
		localctx.(*DirectiveContext).op = append(localctx.(*DirectiveContext).op, localctx.(*DirectiveContext)._DIR)

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(59)

		var _m = p.Match(BundParserNAME)

		localctx.(*DirectiveContext).name = _m
	}

	return localctx
}

// ICmdContext is an interface to support dynamic dispatch.
type ICmdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CMD returns the _CMD token.
	Get_CMD() antlr.Token

	// Set_CMD sets the _CMD token.
	Set_CMD(antlr.Token)

	// GetCommand returns the command token list.
	GetCommand() []antlr.Token

	// SetCommand sets the command token list.
	SetCommand([]antlr.Token)

	// IsCmdContext differentiates from other interfaces.
	IsCmdContext()
}

type CmdContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_CMD    antlr.Token
	command []antlr.Token
}

func NewEmptyCmdContext() *CmdContext {
	var p = new(CmdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_cmd
	return p
}

func (*CmdContext) IsCmdContext() {}

func NewCmdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdContext {
	var p = new(CmdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_cmd

	return p
}

func (s *CmdContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdContext) Get_CMD() antlr.Token { return s._CMD }

func (s *CmdContext) Set_CMD(v antlr.Token) { s._CMD = v }

func (s *CmdContext) GetCommand() []antlr.Token { return s.command }

func (s *CmdContext) SetCommand(v []antlr.Token) { s.command = v }

func (s *CmdContext) AllCMD() []antlr.TerminalNode {
	return s.GetTokens(BundParserCMD)
}

func (s *CmdContext) CMD(i int) antlr.TerminalNode {
	return s.GetToken(BundParserCMD, i)
}

func (s *CmdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterCmd(s)
	}
}

func (s *CmdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitCmd(s)
	}
}

func (p *BundParser) Cmd() (localctx ICmdContext) {
	localctx = NewCmdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BundParserRULE_cmd)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(61)

				var _m = p.Match(BundParserCMD)

				localctx.(*CmdContext)._CMD = _m
			}
			localctx.(*CmdContext).command = append(localctx.(*CmdContext).command, localctx.(*CmdContext)._CMD)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// ILambdaContext is an interface to support dynamic dispatch.
type ILambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetBody returns the body rule context list.
	GetBody() []ITermContext

	// SetBody sets the body rule context list.
	SetBody([]ITermContext)

	// IsLambdaContext differentiates from other interfaces.
	IsLambdaContext()
}

type LambdaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyLambdaContext() *LambdaContext {
	var p = new(LambdaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_lambda
	return p
}

func (*LambdaContext) IsLambdaContext() {}

func NewLambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaContext {
	var p = new(LambdaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_lambda

	return p
}

func (s *LambdaContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaContext) GetName() antlr.Token { return s.name }

func (s *LambdaContext) SetName(v antlr.Token) { s.name = v }

func (s *LambdaContext) Get_term() ITermContext { return s._term }

func (s *LambdaContext) Set_term(v ITermContext) { s._term = v }

func (s *LambdaContext) GetBody() []ITermContext { return s.body }

func (s *LambdaContext) SetBody(v []ITermContext) { s.body = v }

func (s *LambdaContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *LambdaContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *LambdaContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *LambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterLambda(s)
	}
}

func (s *LambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitLambda(s)
	}
}

func (p *BundParser) Lambda() (localctx ILambdaContext) {
	localctx = NewLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BundParserRULE_lambda)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(BundParserT__3)
	}
	{
		p.SetState(67)

		var _m = p.Match(BundParserNAME)

		localctx.(*LambdaContext).name = _m
	}
	{
		p.SetState(68)
		p.Match(BundParserT__4)
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__1)|(1<<BundParserINTEGER)|(1<<BundParserOCT_INTEGER)|(1<<BundParserHEX_INTEGER)|(1<<BundParserBIN_INTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserNAME))) != 0 {
		{
			p.SetState(69)

			var _x = p.Term()

			localctx.(*LambdaContext)._term = _x
		}
		localctx.(*LambdaContext).body = append(localctx.(*LambdaContext).body, localctx.(*LambdaContext)._term)

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(75)
		p.Match(BundParserT__5)
	}

	return localctx
}

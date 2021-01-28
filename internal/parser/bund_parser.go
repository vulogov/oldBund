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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 15, 51, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 7, 2, 12, 10, 2, 12,
	2, 14, 2, 15, 11, 2, 3, 2, 7, 2, 18, 10, 2, 12, 2, 14, 2, 21, 11, 2, 5,
	2, 23, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	3, 34, 10, 3, 3, 4, 3, 4, 3, 4, 7, 4, 39, 10, 4, 12, 4, 14, 4, 42, 11,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 2, 2, 6, 2, 4, 6, 8,
	2, 2, 2, 58, 2, 22, 3, 2, 2, 2, 4, 33, 3, 2, 2, 2, 6, 35, 3, 2, 2, 2, 8,
	45, 3, 2, 2, 2, 10, 12, 5, 4, 3, 2, 11, 10, 3, 2, 2, 2, 12, 15, 3, 2, 2,
	2, 13, 11, 3, 2, 2, 2, 13, 14, 3, 2, 2, 2, 14, 23, 3, 2, 2, 2, 15, 13,
	3, 2, 2, 2, 16, 18, 5, 6, 4, 2, 17, 16, 3, 2, 2, 2, 18, 21, 3, 2, 2, 2,
	19, 17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 23, 3, 2, 2, 2, 21, 19, 3,
	2, 2, 2, 22, 13, 3, 2, 2, 2, 22, 19, 3, 2, 2, 2, 23, 3, 3, 2, 2, 2, 24,
	34, 7, 6, 2, 2, 25, 34, 7, 8, 2, 2, 26, 34, 7, 9, 2, 2, 27, 34, 7, 10,
	2, 2, 28, 34, 7, 11, 2, 2, 29, 34, 7, 13, 2, 2, 30, 34, 7, 14, 2, 2, 31,
	34, 7, 12, 2, 2, 32, 34, 5, 8, 5, 2, 33, 24, 3, 2, 2, 2, 33, 25, 3, 2,
	2, 2, 33, 26, 3, 2, 2, 2, 33, 27, 3, 2, 2, 2, 33, 28, 3, 2, 2, 2, 33, 29,
	3, 2, 2, 2, 33, 30, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 32, 3, 2, 2, 2,
	34, 5, 3, 2, 2, 2, 35, 36, 7, 3, 2, 2, 36, 40, 5, 4, 3, 2, 37, 39, 5, 4,
	3, 2, 38, 37, 3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41,
	3, 2, 2, 2, 41, 43, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 43, 44, 7, 3, 2, 2,
	44, 7, 3, 2, 2, 2, 45, 46, 7, 4, 2, 2, 46, 47, 5, 4, 3, 2, 47, 48, 5, 4,
	3, 2, 48, 49, 7, 5, 2, 2, 49, 9, 3, 2, 2, 2, 7, 13, 19, 22, 33, 40,
}
var literalNames = []string{
	"", "'|'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "INTEGER", "DECIMAL_INTEGER", "OCT_INTEGER", "HEX_INTEGER",
	"BIN_INTEGER", "FLOAT_NUMBER", "STRING", "TRUE", "FALSE", "SKIP_",
}

var ruleNames = []string{
	"expressions", "term", "block", "pair",
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
	BundParserINTEGER         = 4
	BundParserDECIMAL_INTEGER = 5
	BundParserOCT_INTEGER     = 6
	BundParserHEX_INTEGER     = 7
	BundParserBIN_INTEGER     = 8
	BundParserFLOAT_NUMBER    = 9
	BundParserSTRING          = 10
	BundParserTRUE            = 11
	BundParserFALSE           = 12
	BundParserSKIP_           = 13
)

// BundParser rules.
const (
	BundParserRULE_expressions = 0
	BundParserRULE_term        = 1
	BundParserRULE_block       = 2
	BundParserRULE_pair        = 3
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

	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(11)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__1)|(1<<BundParserINTEGER)|(1<<BundParserOCT_INTEGER)|(1<<BundParserHEX_INTEGER)|(1<<BundParserBIN_INTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserTRUE)|(1<<BundParserFALSE))) != 0 {
			{
				p.SetState(8)
				p.Term()
			}

			p.SetState(13)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == BundParserT__0 {
			{
				p.SetState(14)
				p.Block()
			}

			p.SetState(19)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

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

	p.SetState(31)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Match(BundParserINTEGER)
		}

	case BundParserOCT_INTEGER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.Match(BundParserOCT_INTEGER)
		}

	case BundParserHEX_INTEGER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(24)
			p.Match(BundParserHEX_INTEGER)
		}

	case BundParserBIN_INTEGER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(25)
			p.Match(BundParserBIN_INTEGER)
		}

	case BundParserFLOAT_NUMBER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(26)
			p.Match(BundParserFLOAT_NUMBER)
		}

	case BundParserTRUE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(27)
			p.Match(BundParserTRUE)
		}

	case BundParserFALSE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(28)
			p.Match(BundParserFALSE)
		}

	case BundParserSTRING:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(29)
			p.Match(BundParserSTRING)
		}

	case BundParserT__1:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(30)
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

	// GetHead returns the head rule contexts.
	GetHead() ITermContext

	// SetHead sets the head rule contexts.
	SetHead(ITermContext)

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	head   ITermContext
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

func (s *BlockContext) GetHead() ITermContext { return s.head }

func (s *BlockContext) SetHead(v ITermContext) { s.head = v }

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
		p.SetState(33)
		p.Match(BundParserT__0)
	}
	{
		p.SetState(34)

		var _x = p.Term()

		localctx.(*BlockContext).head = _x
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__1)|(1<<BundParserINTEGER)|(1<<BundParserOCT_INTEGER)|(1<<BundParserHEX_INTEGER)|(1<<BundParserBIN_INTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserTRUE)|(1<<BundParserFALSE))) != 0 {
		{
			p.SetState(35)
			p.Term()
		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(41)
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
		p.SetState(43)
		p.Match(BundParserT__1)
	}
	{
		p.SetState(44)

		var _x = p.Term()

		localctx.(*PairContext).head = _x
	}
	{
		p.SetState(45)

		var _x = p.Term()

		localctx.(*PairContext).tail = _x
	}
	{
		p.SetState(46)
		p.Match(BundParserT__2)
	}

	return localctx
}

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 28, 182,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 75, 10, 2, 12, 2, 14, 2, 78, 11, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 98, 10, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 112, 10, 4,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20,
	7, 20, 147, 10, 20, 12, 20, 14, 20, 150, 11, 20, 3, 20, 3, 20, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 7, 26, 175,
	10, 26, 12, 26, 14, 26, 178, 11, 26, 3, 26, 3, 26, 3, 26, 2, 2, 27, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 50, 2, 2, 2, 208, 2, 76, 3, 2, 2, 2, 4, 97, 3, 2, 2, 2,
	6, 111, 3, 2, 2, 2, 8, 113, 3, 2, 2, 2, 10, 115, 3, 2, 2, 2, 12, 117, 3,
	2, 2, 2, 14, 119, 3, 2, 2, 2, 16, 121, 3, 2, 2, 2, 18, 123, 3, 2, 2, 2,
	20, 125, 3, 2, 2, 2, 22, 127, 3, 2, 2, 2, 24, 129, 3, 2, 2, 2, 26, 131,
	3, 2, 2, 2, 28, 133, 3, 2, 2, 2, 30, 136, 3, 2, 2, 2, 32, 138, 3, 2, 2,
	2, 34, 140, 3, 2, 2, 2, 36, 142, 3, 2, 2, 2, 38, 144, 3, 2, 2, 2, 40, 153,
	3, 2, 2, 2, 42, 158, 3, 2, 2, 2, 44, 161, 3, 2, 2, 2, 46, 165, 3, 2, 2,
	2, 48, 167, 3, 2, 2, 2, 50, 170, 3, 2, 2, 2, 52, 75, 5, 22, 12, 2, 53,
	75, 5, 24, 13, 2, 54, 75, 5, 20, 11, 2, 55, 75, 5, 18, 10, 2, 56, 75, 5,
	16, 9, 2, 57, 75, 5, 14, 8, 2, 58, 75, 5, 12, 7, 2, 59, 75, 5, 10, 6, 2,
	60, 75, 5, 8, 5, 2, 61, 75, 5, 26, 14, 2, 62, 75, 5, 28, 15, 2, 63, 75,
	5, 30, 16, 2, 64, 75, 5, 32, 17, 2, 65, 75, 5, 34, 18, 2, 66, 75, 5, 36,
	19, 2, 67, 75, 5, 40, 21, 2, 68, 75, 5, 38, 20, 2, 69, 75, 5, 42, 22, 2,
	70, 75, 5, 44, 23, 2, 71, 75, 5, 46, 24, 2, 72, 75, 5, 48, 25, 2, 73, 75,
	5, 50, 26, 2, 74, 52, 3, 2, 2, 2, 74, 53, 3, 2, 2, 2, 74, 54, 3, 2, 2,
	2, 74, 55, 3, 2, 2, 2, 74, 56, 3, 2, 2, 2, 74, 57, 3, 2, 2, 2, 74, 58,
	3, 2, 2, 2, 74, 59, 3, 2, 2, 2, 74, 60, 3, 2, 2, 2, 74, 61, 3, 2, 2, 2,
	74, 62, 3, 2, 2, 2, 74, 63, 3, 2, 2, 2, 74, 64, 3, 2, 2, 2, 74, 65, 3,
	2, 2, 2, 74, 66, 3, 2, 2, 2, 74, 67, 3, 2, 2, 2, 74, 68, 3, 2, 2, 2, 74,
	69, 3, 2, 2, 2, 74, 70, 3, 2, 2, 2, 74, 71, 3, 2, 2, 2, 74, 72, 3, 2, 2,
	2, 74, 73, 3, 2, 2, 2, 75, 78, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77,
	3, 2, 2, 2, 77, 3, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 79, 98, 5, 24, 13, 2,
	80, 98, 5, 22, 12, 2, 81, 98, 5, 20, 11, 2, 82, 98, 5, 18, 10, 2, 83, 98,
	5, 16, 9, 2, 84, 98, 5, 14, 8, 2, 85, 98, 5, 12, 7, 2, 86, 98, 5, 10, 6,
	2, 87, 98, 5, 8, 5, 2, 88, 98, 5, 26, 14, 2, 89, 98, 5, 30, 16, 2, 90,
	98, 5, 32, 17, 2, 91, 98, 5, 34, 18, 2, 92, 98, 5, 36, 19, 2, 93, 98, 5,
	40, 21, 2, 94, 98, 5, 38, 20, 2, 95, 98, 5, 42, 22, 2, 96, 98, 5, 46, 24,
	2, 97, 79, 3, 2, 2, 2, 97, 80, 3, 2, 2, 2, 97, 81, 3, 2, 2, 2, 97, 82,
	3, 2, 2, 2, 97, 83, 3, 2, 2, 2, 97, 84, 3, 2, 2, 2, 97, 85, 3, 2, 2, 2,
	97, 86, 3, 2, 2, 2, 97, 87, 3, 2, 2, 2, 97, 88, 3, 2, 2, 2, 97, 89, 3,
	2, 2, 2, 97, 90, 3, 2, 2, 2, 97, 91, 3, 2, 2, 2, 97, 92, 3, 2, 2, 2, 97,
	93, 3, 2, 2, 2, 97, 94, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 96, 3, 2, 2,
	2, 98, 5, 3, 2, 2, 2, 99, 112, 5, 22, 12, 2, 100, 112, 5, 24, 13, 2, 101,
	112, 5, 20, 11, 2, 102, 112, 5, 18, 10, 2, 103, 112, 5, 16, 9, 2, 104,
	112, 5, 14, 8, 2, 105, 112, 5, 12, 7, 2, 106, 112, 5, 10, 6, 2, 107, 112,
	5, 8, 5, 2, 108, 112, 5, 26, 14, 2, 109, 112, 5, 42, 22, 2, 110, 112, 5,
	46, 24, 2, 111, 99, 3, 2, 2, 2, 111, 100, 3, 2, 2, 2, 111, 101, 3, 2, 2,
	2, 111, 102, 3, 2, 2, 2, 111, 103, 3, 2, 2, 2, 111, 104, 3, 2, 2, 2, 111,
	105, 3, 2, 2, 2, 111, 106, 3, 2, 2, 2, 111, 107, 3, 2, 2, 2, 111, 108,
	3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 111, 110, 3, 2, 2, 2, 112, 7, 3, 2, 2,
	2, 113, 114, 7, 17, 2, 2, 114, 9, 3, 2, 2, 2, 115, 116, 7, 18, 2, 2, 116,
	11, 3, 2, 2, 2, 117, 118, 7, 16, 2, 2, 118, 13, 3, 2, 2, 2, 119, 120, 7,
	15, 2, 2, 120, 15, 3, 2, 2, 2, 121, 122, 7, 13, 2, 2, 122, 17, 3, 2, 2,
	2, 123, 124, 7, 12, 2, 2, 124, 19, 3, 2, 2, 2, 125, 126, 7, 11, 2, 2, 126,
	21, 3, 2, 2, 2, 127, 128, 7, 9, 2, 2, 128, 23, 3, 2, 2, 2, 129, 130, 7,
	14, 2, 2, 130, 25, 3, 2, 2, 2, 131, 132, 7, 19, 2, 2, 132, 27, 3, 2, 2,
	2, 133, 134, 7, 23, 2, 2, 134, 135, 7, 19, 2, 2, 135, 29, 3, 2, 2, 2, 136,
	137, 7, 26, 2, 2, 137, 31, 3, 2, 2, 2, 138, 139, 7, 27, 2, 2, 139, 33,
	3, 2, 2, 2, 140, 141, 7, 20, 2, 2, 141, 35, 3, 2, 2, 2, 142, 143, 7, 21,
	2, 2, 143, 37, 3, 2, 2, 2, 144, 148, 7, 3, 2, 2, 145, 147, 5, 4, 3, 2,
	146, 145, 3, 2, 2, 2, 147, 150, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 148,
	149, 3, 2, 2, 2, 149, 151, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 151, 152,
	7, 3, 2, 2, 152, 39, 3, 2, 2, 2, 153, 154, 7, 4, 2, 2, 154, 155, 5, 26,
	14, 2, 155, 156, 5, 6, 4, 2, 156, 157, 7, 5, 2, 2, 157, 41, 3, 2, 2, 2,
	158, 159, 7, 24, 2, 2, 159, 160, 7, 19, 2, 2, 160, 43, 3, 2, 2, 2, 161,
	162, 7, 23, 2, 2, 162, 163, 7, 24, 2, 2, 163, 164, 7, 19, 2, 2, 164, 45,
	3, 2, 2, 2, 165, 166, 7, 25, 2, 2, 166, 47, 3, 2, 2, 2, 167, 168, 7, 23,
	2, 2, 168, 169, 7, 25, 2, 2, 169, 49, 3, 2, 2, 2, 170, 171, 7, 6, 2, 2,
	171, 172, 7, 19, 2, 2, 172, 176, 7, 7, 2, 2, 173, 175, 5, 4, 3, 2, 174,
	173, 3, 2, 2, 2, 175, 178, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 176, 177,
	3, 2, 2, 2, 177, 179, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 179, 180, 7, 8,
	2, 2, 180, 51, 3, 2, 2, 2, 8, 74, 76, 97, 111, 148, 176,
}
var literalNames = []string{
	"", "'|'", "'{'", "'}'", "'['", "']'", "'.'", "", "", "", "", "", "", "",
	"", "", "", "", "':'", "';'", "'/'", "", "", "", "'^'", "'_'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "INTEGER", "DECIMAL_INTEGER", "OCT_INTEGER",
	"HEX_INTEGER", "BIN_INTEGER", "NTH", "FLOAT_NUMBER", "STRING", "TRUE",
	"FALSE", "NAME", "TOBEGIN", "TOEND", "SLASH", "SYS", "DIR", "CMD", "DUPLICATE",
	"DROP", "SKIP_",
}

var ruleNames = []string{
	"expressions", "term", "pvalue", "true_term", "false_term", "string_term",
	"float_term", "bin_integer", "hex_integer", "oct_integer", "integer", "nth",
	"name_term", "sys_name_term", "duplicate_term", "drop_term", "begin", "end",
	"block", "pair", "directive", "sys_directive", "cmd", "sys_cmd", "lambda",
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
	BundParserNTH             = 12
	BundParserFLOAT_NUMBER    = 13
	BundParserSTRING          = 14
	BundParserTRUE            = 15
	BundParserFALSE           = 16
	BundParserNAME            = 17
	BundParserTOBEGIN         = 18
	BundParserTOEND           = 19
	BundParserSLASH           = 20
	BundParserSYS             = 21
	BundParserDIR             = 22
	BundParserCMD             = 23
	BundParserDUPLICATE       = 24
	BundParserDROP            = 25
	BundParserSKIP_           = 26
)

// BundParser rules.
const (
	BundParserRULE_expressions    = 0
	BundParserRULE_term           = 1
	BundParserRULE_pvalue         = 2
	BundParserRULE_true_term      = 3
	BundParserRULE_false_term     = 4
	BundParserRULE_string_term    = 5
	BundParserRULE_float_term     = 6
	BundParserRULE_bin_integer    = 7
	BundParserRULE_hex_integer    = 8
	BundParserRULE_oct_integer    = 9
	BundParserRULE_integer        = 10
	BundParserRULE_nth            = 11
	BundParserRULE_name_term      = 12
	BundParserRULE_sys_name_term  = 13
	BundParserRULE_duplicate_term = 14
	BundParserRULE_drop_term      = 15
	BundParserRULE_begin          = 16
	BundParserRULE_end            = 17
	BundParserRULE_block          = 18
	BundParserRULE_pair           = 19
	BundParserRULE_directive      = 20
	BundParserRULE_sys_directive  = 21
	BundParserRULE_cmd            = 22
	BundParserRULE_sys_cmd        = 23
	BundParserRULE_lambda         = 24
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

func (s *ExpressionsContext) AllInteger() []IIntegerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIntegerContext)(nil)).Elem())
	var tst = make([]IIntegerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIntegerContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Integer(i int) IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *ExpressionsContext) AllNth() []INthContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INthContext)(nil)).Elem())
	var tst = make([]INthContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INthContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Nth(i int) INthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INthContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INthContext)
}

func (s *ExpressionsContext) AllOct_integer() []IOct_integerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOct_integerContext)(nil)).Elem())
	var tst = make([]IOct_integerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOct_integerContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Oct_integer(i int) IOct_integerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOct_integerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOct_integerContext)
}

func (s *ExpressionsContext) AllHex_integer() []IHex_integerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHex_integerContext)(nil)).Elem())
	var tst = make([]IHex_integerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHex_integerContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Hex_integer(i int) IHex_integerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHex_integerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHex_integerContext)
}

func (s *ExpressionsContext) AllBin_integer() []IBin_integerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBin_integerContext)(nil)).Elem())
	var tst = make([]IBin_integerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBin_integerContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Bin_integer(i int) IBin_integerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBin_integerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBin_integerContext)
}

func (s *ExpressionsContext) AllFloat_term() []IFloat_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFloat_termContext)(nil)).Elem())
	var tst = make([]IFloat_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFloat_termContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Float_term(i int) IFloat_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloat_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFloat_termContext)
}

func (s *ExpressionsContext) AllString_term() []IString_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IString_termContext)(nil)).Elem())
	var tst = make([]IString_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IString_termContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) String_term(i int) IString_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IString_termContext)
}

func (s *ExpressionsContext) AllFalse_term() []IFalse_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFalse_termContext)(nil)).Elem())
	var tst = make([]IFalse_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFalse_termContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) False_term(i int) IFalse_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFalse_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFalse_termContext)
}

func (s *ExpressionsContext) AllTrue_term() []ITrue_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITrue_termContext)(nil)).Elem())
	var tst = make([]ITrue_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITrue_termContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) True_term(i int) ITrue_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrue_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITrue_termContext)
}

func (s *ExpressionsContext) AllName_term() []IName_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IName_termContext)(nil)).Elem())
	var tst = make([]IName_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IName_termContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Name_term(i int) IName_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IName_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IName_termContext)
}

func (s *ExpressionsContext) AllSys_name_term() []ISys_name_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISys_name_termContext)(nil)).Elem())
	var tst = make([]ISys_name_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISys_name_termContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Sys_name_term(i int) ISys_name_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISys_name_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISys_name_termContext)
}

func (s *ExpressionsContext) AllDuplicate_term() []IDuplicate_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDuplicate_termContext)(nil)).Elem())
	var tst = make([]IDuplicate_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDuplicate_termContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Duplicate_term(i int) IDuplicate_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDuplicate_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDuplicate_termContext)
}

func (s *ExpressionsContext) AllDrop_term() []IDrop_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDrop_termContext)(nil)).Elem())
	var tst = make([]IDrop_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDrop_termContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Drop_term(i int) IDrop_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDrop_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDrop_termContext)
}

func (s *ExpressionsContext) AllBegin() []IBeginContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBeginContext)(nil)).Elem())
	var tst = make([]IBeginContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBeginContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Begin(i int) IBeginContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBeginContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBeginContext)
}

func (s *ExpressionsContext) AllEnd() []IEndContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEndContext)(nil)).Elem())
	var tst = make([]IEndContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEndContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) End(i int) IEndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEndContext)
}

func (s *ExpressionsContext) AllPair() []IPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairContext)(nil)).Elem())
	var tst = make([]IPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Pair(i int) IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
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

func (s *ExpressionsContext) AllSys_directive() []ISys_directiveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISys_directiveContext)(nil)).Elem())
	var tst = make([]ISys_directiveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISys_directiveContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Sys_directive(i int) ISys_directiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISys_directiveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISys_directiveContext)
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

func (s *ExpressionsContext) AllSys_cmd() []ISys_cmdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISys_cmdContext)(nil)).Elem())
	var tst = make([]ISys_cmdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISys_cmdContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Sys_cmd(i int) ISys_cmdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISys_cmdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISys_cmdContext)
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
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__1)|(1<<BundParserT__3)|(1<<BundParserINTEGER)|(1<<BundParserOCT_INTEGER)|(1<<BundParserHEX_INTEGER)|(1<<BundParserBIN_INTEGER)|(1<<BundParserNTH)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserNAME)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND)|(1<<BundParserSYS)|(1<<BundParserDIR)|(1<<BundParserCMD)|(1<<BundParserDUPLICATE)|(1<<BundParserDROP))) != 0 {
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(50)
				p.Integer()
			}

		case 2:
			{
				p.SetState(51)
				p.Nth()
			}

		case 3:
			{
				p.SetState(52)
				p.Oct_integer()
			}

		case 4:
			{
				p.SetState(53)
				p.Hex_integer()
			}

		case 5:
			{
				p.SetState(54)
				p.Bin_integer()
			}

		case 6:
			{
				p.SetState(55)
				p.Float_term()
			}

		case 7:
			{
				p.SetState(56)
				p.String_term()
			}

		case 8:
			{
				p.SetState(57)
				p.False_term()
			}

		case 9:
			{
				p.SetState(58)
				p.True_term()
			}

		case 10:
			{
				p.SetState(59)
				p.Name_term()
			}

		case 11:
			{
				p.SetState(60)
				p.Sys_name_term()
			}

		case 12:
			{
				p.SetState(61)
				p.Duplicate_term()
			}

		case 13:
			{
				p.SetState(62)
				p.Drop_term()
			}

		case 14:
			{
				p.SetState(63)
				p.Begin()
			}

		case 15:
			{
				p.SetState(64)
				p.End()
			}

		case 16:
			{
				p.SetState(65)
				p.Pair()
			}

		case 17:
			{
				p.SetState(66)
				p.Block()
			}

		case 18:
			{
				p.SetState(67)
				p.Directive()
			}

		case 19:
			{
				p.SetState(68)
				p.Sys_directive()
			}

		case 20:
			{
				p.SetState(69)
				p.Cmd()
			}

		case 21:
			{
				p.SetState(70)
				p.Sys_cmd()
			}

		case 22:
			{
				p.SetState(71)
				p.Lambda()
			}

		}

		p.SetState(76)
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

func (s *TermContext) Nth() INthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INthContext)
}

func (s *TermContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *TermContext) Oct_integer() IOct_integerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOct_integerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOct_integerContext)
}

func (s *TermContext) Hex_integer() IHex_integerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHex_integerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHex_integerContext)
}

func (s *TermContext) Bin_integer() IBin_integerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBin_integerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBin_integerContext)
}

func (s *TermContext) Float_term() IFloat_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloat_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloat_termContext)
}

func (s *TermContext) String_term() IString_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_termContext)
}

func (s *TermContext) False_term() IFalse_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFalse_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFalse_termContext)
}

func (s *TermContext) True_term() ITrue_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrue_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrue_termContext)
}

func (s *TermContext) Name_term() IName_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IName_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IName_termContext)
}

func (s *TermContext) Duplicate_term() IDuplicate_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDuplicate_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDuplicate_termContext)
}

func (s *TermContext) Drop_term() IDrop_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDrop_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDrop_termContext)
}

func (s *TermContext) Begin() IBeginContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBeginContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBeginContext)
}

func (s *TermContext) End() IEndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndContext)
}

func (s *TermContext) Pair() IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *TermContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *TermContext) Directive() IDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *TermContext) Cmd() ICmdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICmdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(95)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserNTH:
		{
			p.SetState(77)
			p.Nth()
		}

	case BundParserINTEGER:
		{
			p.SetState(78)
			p.Integer()
		}

	case BundParserOCT_INTEGER:
		{
			p.SetState(79)
			p.Oct_integer()
		}

	case BundParserHEX_INTEGER:
		{
			p.SetState(80)
			p.Hex_integer()
		}

	case BundParserBIN_INTEGER:
		{
			p.SetState(81)
			p.Bin_integer()
		}

	case BundParserFLOAT_NUMBER:
		{
			p.SetState(82)
			p.Float_term()
		}

	case BundParserSTRING:
		{
			p.SetState(83)
			p.String_term()
		}

	case BundParserFALSE:
		{
			p.SetState(84)
			p.False_term()
		}

	case BundParserTRUE:
		{
			p.SetState(85)
			p.True_term()
		}

	case BundParserNAME:
		{
			p.SetState(86)
			p.Name_term()
		}

	case BundParserDUPLICATE:
		{
			p.SetState(87)
			p.Duplicate_term()
		}

	case BundParserDROP:
		{
			p.SetState(88)
			p.Drop_term()
		}

	case BundParserTOBEGIN:
		{
			p.SetState(89)
			p.Begin()
		}

	case BundParserTOEND:
		{
			p.SetState(90)
			p.End()
		}

	case BundParserT__1:
		{
			p.SetState(91)
			p.Pair()
		}

	case BundParserT__0:
		{
			p.SetState(92)
			p.Block()
		}

	case BundParserDIR:
		{
			p.SetState(93)
			p.Directive()
		}

	case BundParserCMD:
		{
			p.SetState(94)
			p.Cmd()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPvalueContext is an interface to support dynamic dispatch.
type IPvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPvalueContext differentiates from other interfaces.
	IsPvalueContext()
}

type PvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPvalueContext() *PvalueContext {
	var p = new(PvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_pvalue
	return p
}

func (*PvalueContext) IsPvalueContext() {}

func NewPvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PvalueContext {
	var p = new(PvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_pvalue

	return p
}

func (s *PvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *PvalueContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *PvalueContext) Nth() INthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INthContext)
}

func (s *PvalueContext) Oct_integer() IOct_integerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOct_integerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOct_integerContext)
}

func (s *PvalueContext) Hex_integer() IHex_integerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHex_integerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHex_integerContext)
}

func (s *PvalueContext) Bin_integer() IBin_integerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBin_integerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBin_integerContext)
}

func (s *PvalueContext) Float_term() IFloat_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloat_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloat_termContext)
}

func (s *PvalueContext) String_term() IString_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_termContext)
}

func (s *PvalueContext) False_term() IFalse_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFalse_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFalse_termContext)
}

func (s *PvalueContext) True_term() ITrue_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrue_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrue_termContext)
}

func (s *PvalueContext) Name_term() IName_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IName_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IName_termContext)
}

func (s *PvalueContext) Directive() IDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *PvalueContext) Cmd() ICmdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICmdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICmdContext)
}

func (s *PvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterPvalue(s)
	}
}

func (s *PvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitPvalue(s)
	}
}

func (p *BundParser) Pvalue() (localctx IPvalueContext) {
	localctx = NewPvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BundParserRULE_pvalue)

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
	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserINTEGER:
		{
			p.SetState(97)
			p.Integer()
		}

	case BundParserNTH:
		{
			p.SetState(98)
			p.Nth()
		}

	case BundParserOCT_INTEGER:
		{
			p.SetState(99)
			p.Oct_integer()
		}

	case BundParserHEX_INTEGER:
		{
			p.SetState(100)
			p.Hex_integer()
		}

	case BundParserBIN_INTEGER:
		{
			p.SetState(101)
			p.Bin_integer()
		}

	case BundParserFLOAT_NUMBER:
		{
			p.SetState(102)
			p.Float_term()
		}

	case BundParserSTRING:
		{
			p.SetState(103)
			p.String_term()
		}

	case BundParserFALSE:
		{
			p.SetState(104)
			p.False_term()
		}

	case BundParserTRUE:
		{
			p.SetState(105)
			p.True_term()
		}

	case BundParserNAME:
		{
			p.SetState(106)
			p.Name_term()
		}

	case BundParserDIR:
		{
			p.SetState(107)
			p.Directive()
		}

	case BundParserCMD:
		{
			p.SetState(108)
			p.Cmd()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITrue_termContext is an interface to support dynamic dispatch.
type ITrue_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsTrue_termContext differentiates from other interfaces.
	IsTrue_termContext()
}

type True_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyTrue_termContext() *True_termContext {
	var p = new(True_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_true_term
	return p
}

func (*True_termContext) IsTrue_termContext() {}

func NewTrue_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *True_termContext {
	var p = new(True_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_true_term

	return p
}

func (s *True_termContext) GetParser() antlr.Parser { return s.parser }

func (s *True_termContext) GetValue() antlr.Token { return s.value }

func (s *True_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *True_termContext) TRUE() antlr.TerminalNode {
	return s.GetToken(BundParserTRUE, 0)
}

func (s *True_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *True_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *True_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterTrue_term(s)
	}
}

func (s *True_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitTrue_term(s)
	}
}

func (p *BundParser) True_term() (localctx ITrue_termContext) {
	localctx = NewTrue_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BundParserRULE_true_term)

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
		p.SetState(111)

		var _m = p.Match(BundParserTRUE)

		localctx.(*True_termContext).value = _m
	}

	return localctx
}

// IFalse_termContext is an interface to support dynamic dispatch.
type IFalse_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsFalse_termContext differentiates from other interfaces.
	IsFalse_termContext()
}

type False_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyFalse_termContext() *False_termContext {
	var p = new(False_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_false_term
	return p
}

func (*False_termContext) IsFalse_termContext() {}

func NewFalse_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *False_termContext {
	var p = new(False_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_false_term

	return p
}

func (s *False_termContext) GetParser() antlr.Parser { return s.parser }

func (s *False_termContext) GetValue() antlr.Token { return s.value }

func (s *False_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *False_termContext) FALSE() antlr.TerminalNode {
	return s.GetToken(BundParserFALSE, 0)
}

func (s *False_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *False_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *False_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterFalse_term(s)
	}
}

func (s *False_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitFalse_term(s)
	}
}

func (p *BundParser) False_term() (localctx IFalse_termContext) {
	localctx = NewFalse_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BundParserRULE_false_term)

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
		p.SetState(113)

		var _m = p.Match(BundParserFALSE)

		localctx.(*False_termContext).value = _m
	}

	return localctx
}

// IString_termContext is an interface to support dynamic dispatch.
type IString_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsString_termContext differentiates from other interfaces.
	IsString_termContext()
}

type String_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyString_termContext() *String_termContext {
	var p = new(String_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_string_term
	return p
}

func (*String_termContext) IsString_termContext() {}

func NewString_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_termContext {
	var p = new(String_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_string_term

	return p
}

func (s *String_termContext) GetParser() antlr.Parser { return s.parser }

func (s *String_termContext) GetValue() antlr.Token { return s.value }

func (s *String_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *String_termContext) STRING() antlr.TerminalNode {
	return s.GetToken(BundParserSTRING, 0)
}

func (s *String_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterString_term(s)
	}
}

func (s *String_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitString_term(s)
	}
}

func (p *BundParser) String_term() (localctx IString_termContext) {
	localctx = NewString_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BundParserRULE_string_term)

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
		p.SetState(115)

		var _m = p.Match(BundParserSTRING)

		localctx.(*String_termContext).value = _m
	}

	return localctx
}

// IFloat_termContext is an interface to support dynamic dispatch.
type IFloat_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsFloat_termContext differentiates from other interfaces.
	IsFloat_termContext()
}

type Float_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyFloat_termContext() *Float_termContext {
	var p = new(Float_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_float_term
	return p
}

func (*Float_termContext) IsFloat_termContext() {}

func NewFloat_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Float_termContext {
	var p = new(Float_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_float_term

	return p
}

func (s *Float_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Float_termContext) GetValue() antlr.Token { return s.value }

func (s *Float_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *Float_termContext) FLOAT_NUMBER() antlr.TerminalNode {
	return s.GetToken(BundParserFLOAT_NUMBER, 0)
}

func (s *Float_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Float_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Float_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterFloat_term(s)
	}
}

func (s *Float_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitFloat_term(s)
	}
}

func (p *BundParser) Float_term() (localctx IFloat_termContext) {
	localctx = NewFloat_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BundParserRULE_float_term)

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
		p.SetState(117)

		var _m = p.Match(BundParserFLOAT_NUMBER)

		localctx.(*Float_termContext).value = _m
	}

	return localctx
}

// IBin_integerContext is an interface to support dynamic dispatch.
type IBin_integerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsBin_integerContext differentiates from other interfaces.
	IsBin_integerContext()
}

type Bin_integerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyBin_integerContext() *Bin_integerContext {
	var p = new(Bin_integerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_bin_integer
	return p
}

func (*Bin_integerContext) IsBin_integerContext() {}

func NewBin_integerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bin_integerContext {
	var p = new(Bin_integerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_bin_integer

	return p
}

func (s *Bin_integerContext) GetParser() antlr.Parser { return s.parser }

func (s *Bin_integerContext) GetValue() antlr.Token { return s.value }

func (s *Bin_integerContext) SetValue(v antlr.Token) { s.value = v }

func (s *Bin_integerContext) BIN_INTEGER() antlr.TerminalNode {
	return s.GetToken(BundParserBIN_INTEGER, 0)
}

func (s *Bin_integerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bin_integerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bin_integerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterBin_integer(s)
	}
}

func (s *Bin_integerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitBin_integer(s)
	}
}

func (p *BundParser) Bin_integer() (localctx IBin_integerContext) {
	localctx = NewBin_integerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BundParserRULE_bin_integer)

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
		p.SetState(119)

		var _m = p.Match(BundParserBIN_INTEGER)

		localctx.(*Bin_integerContext).value = _m
	}

	return localctx
}

// IHex_integerContext is an interface to support dynamic dispatch.
type IHex_integerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsHex_integerContext differentiates from other interfaces.
	IsHex_integerContext()
}

type Hex_integerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyHex_integerContext() *Hex_integerContext {
	var p = new(Hex_integerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_hex_integer
	return p
}

func (*Hex_integerContext) IsHex_integerContext() {}

func NewHex_integerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Hex_integerContext {
	var p = new(Hex_integerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_hex_integer

	return p
}

func (s *Hex_integerContext) GetParser() antlr.Parser { return s.parser }

func (s *Hex_integerContext) GetValue() antlr.Token { return s.value }

func (s *Hex_integerContext) SetValue(v antlr.Token) { s.value = v }

func (s *Hex_integerContext) HEX_INTEGER() antlr.TerminalNode {
	return s.GetToken(BundParserHEX_INTEGER, 0)
}

func (s *Hex_integerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Hex_integerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Hex_integerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterHex_integer(s)
	}
}

func (s *Hex_integerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitHex_integer(s)
	}
}

func (p *BundParser) Hex_integer() (localctx IHex_integerContext) {
	localctx = NewHex_integerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BundParserRULE_hex_integer)

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
		p.SetState(121)

		var _m = p.Match(BundParserHEX_INTEGER)

		localctx.(*Hex_integerContext).value = _m
	}

	return localctx
}

// IOct_integerContext is an interface to support dynamic dispatch.
type IOct_integerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsOct_integerContext differentiates from other interfaces.
	IsOct_integerContext()
}

type Oct_integerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyOct_integerContext() *Oct_integerContext {
	var p = new(Oct_integerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_oct_integer
	return p
}

func (*Oct_integerContext) IsOct_integerContext() {}

func NewOct_integerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Oct_integerContext {
	var p = new(Oct_integerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_oct_integer

	return p
}

func (s *Oct_integerContext) GetParser() antlr.Parser { return s.parser }

func (s *Oct_integerContext) GetValue() antlr.Token { return s.value }

func (s *Oct_integerContext) SetValue(v antlr.Token) { s.value = v }

func (s *Oct_integerContext) OCT_INTEGER() antlr.TerminalNode {
	return s.GetToken(BundParserOCT_INTEGER, 0)
}

func (s *Oct_integerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Oct_integerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Oct_integerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterOct_integer(s)
	}
}

func (s *Oct_integerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitOct_integer(s)
	}
}

func (p *BundParser) Oct_integer() (localctx IOct_integerContext) {
	localctx = NewOct_integerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BundParserRULE_oct_integer)

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
		p.SetState(123)

		var _m = p.Match(BundParserOCT_INTEGER)

		localctx.(*Oct_integerContext).value = _m
	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) GetValue() antlr.Token { return s.value }

func (s *IntegerContext) SetValue(v antlr.Token) { s.value = v }

func (s *IntegerContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(BundParserINTEGER, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (p *BundParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BundParserRULE_integer)

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
		p.SetState(125)

		var _m = p.Match(BundParserINTEGER)

		localctx.(*IntegerContext).value = _m
	}

	return localctx
}

// INthContext is an interface to support dynamic dispatch.
type INthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsNthContext differentiates from other interfaces.
	IsNthContext()
}

type NthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyNthContext() *NthContext {
	var p = new(NthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_nth
	return p
}

func (*NthContext) IsNthContext() {}

func NewNthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NthContext {
	var p = new(NthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_nth

	return p
}

func (s *NthContext) GetParser() antlr.Parser { return s.parser }

func (s *NthContext) GetValue() antlr.Token { return s.value }

func (s *NthContext) SetValue(v antlr.Token) { s.value = v }

func (s *NthContext) NTH() antlr.TerminalNode {
	return s.GetToken(BundParserNTH, 0)
}

func (s *NthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterNth(s)
	}
}

func (s *NthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitNth(s)
	}
}

func (p *BundParser) Nth() (localctx INthContext) {
	localctx = NewNthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BundParserRULE_nth)

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
		p.SetState(127)

		var _m = p.Match(BundParserNTH)

		localctx.(*NthContext).value = _m
	}

	return localctx
}

// IName_termContext is an interface to support dynamic dispatch.
type IName_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsName_termContext differentiates from other interfaces.
	IsName_termContext()
}

type Name_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyName_termContext() *Name_termContext {
	var p = new(Name_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_name_term
	return p
}

func (*Name_termContext) IsName_termContext() {}

func NewName_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Name_termContext {
	var p = new(Name_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_name_term

	return p
}

func (s *Name_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Name_termContext) GetValue() antlr.Token { return s.value }

func (s *Name_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *Name_termContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Name_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Name_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Name_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterName_term(s)
	}
}

func (s *Name_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitName_term(s)
	}
}

func (p *BundParser) Name_term() (localctx IName_termContext) {
	localctx = NewName_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BundParserRULE_name_term)

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
		p.SetState(129)

		var _m = p.Match(BundParserNAME)

		localctx.(*Name_termContext).value = _m
	}

	return localctx
}

// ISys_name_termContext is an interface to support dynamic dispatch.
type ISys_name_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSys returns the sys token.
	GetSys() antlr.Token

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetSys sets the sys token.
	SetSys(antlr.Token)

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsSys_name_termContext differentiates from other interfaces.
	IsSys_name_termContext()
}

type Sys_name_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	sys    antlr.Token
	value  antlr.Token
}

func NewEmptySys_name_termContext() *Sys_name_termContext {
	var p = new(Sys_name_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_sys_name_term
	return p
}

func (*Sys_name_termContext) IsSys_name_termContext() {}

func NewSys_name_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sys_name_termContext {
	var p = new(Sys_name_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_sys_name_term

	return p
}

func (s *Sys_name_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Sys_name_termContext) GetSys() antlr.Token { return s.sys }

func (s *Sys_name_termContext) GetValue() antlr.Token { return s.value }

func (s *Sys_name_termContext) SetSys(v antlr.Token) { s.sys = v }

func (s *Sys_name_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *Sys_name_termContext) SYS() antlr.TerminalNode {
	return s.GetToken(BundParserSYS, 0)
}

func (s *Sys_name_termContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Sys_name_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sys_name_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sys_name_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterSys_name_term(s)
	}
}

func (s *Sys_name_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitSys_name_term(s)
	}
}

func (p *BundParser) Sys_name_term() (localctx ISys_name_termContext) {
	localctx = NewSys_name_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BundParserRULE_sys_name_term)

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
		p.SetState(131)

		var _m = p.Match(BundParserSYS)

		localctx.(*Sys_name_termContext).sys = _m
	}
	{
		p.SetState(132)

		var _m = p.Match(BundParserNAME)

		localctx.(*Sys_name_termContext).value = _m
	}

	return localctx
}

// IDuplicate_termContext is an interface to support dynamic dispatch.
type IDuplicate_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsDuplicate_termContext differentiates from other interfaces.
	IsDuplicate_termContext()
}

type Duplicate_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyDuplicate_termContext() *Duplicate_termContext {
	var p = new(Duplicate_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_duplicate_term
	return p
}

func (*Duplicate_termContext) IsDuplicate_termContext() {}

func NewDuplicate_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Duplicate_termContext {
	var p = new(Duplicate_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_duplicate_term

	return p
}

func (s *Duplicate_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Duplicate_termContext) GetValue() antlr.Token { return s.value }

func (s *Duplicate_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *Duplicate_termContext) DUPLICATE() antlr.TerminalNode {
	return s.GetToken(BundParserDUPLICATE, 0)
}

func (s *Duplicate_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Duplicate_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Duplicate_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterDuplicate_term(s)
	}
}

func (s *Duplicate_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitDuplicate_term(s)
	}
}

func (p *BundParser) Duplicate_term() (localctx IDuplicate_termContext) {
	localctx = NewDuplicate_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BundParserRULE_duplicate_term)

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
		p.SetState(134)

		var _m = p.Match(BundParserDUPLICATE)

		localctx.(*Duplicate_termContext).value = _m
	}

	return localctx
}

// IDrop_termContext is an interface to support dynamic dispatch.
type IDrop_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsDrop_termContext differentiates from other interfaces.
	IsDrop_termContext()
}

type Drop_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyDrop_termContext() *Drop_termContext {
	var p = new(Drop_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_drop_term
	return p
}

func (*Drop_termContext) IsDrop_termContext() {}

func NewDrop_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Drop_termContext {
	var p = new(Drop_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_drop_term

	return p
}

func (s *Drop_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Drop_termContext) GetValue() antlr.Token { return s.value }

func (s *Drop_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *Drop_termContext) DROP() antlr.TerminalNode {
	return s.GetToken(BundParserDROP, 0)
}

func (s *Drop_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Drop_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Drop_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterDrop_term(s)
	}
}

func (s *Drop_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitDrop_term(s)
	}
}

func (p *BundParser) Drop_term() (localctx IDrop_termContext) {
	localctx = NewDrop_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BundParserRULE_drop_term)

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
		p.SetState(136)

		var _m = p.Match(BundParserDROP)

		localctx.(*Drop_termContext).value = _m
	}

	return localctx
}

// IBeginContext is an interface to support dynamic dispatch.
type IBeginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsBeginContext differentiates from other interfaces.
	IsBeginContext()
}

type BeginContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyBeginContext() *BeginContext {
	var p = new(BeginContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_begin
	return p
}

func (*BeginContext) IsBeginContext() {}

func NewBeginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BeginContext {
	var p = new(BeginContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_begin

	return p
}

func (s *BeginContext) GetParser() antlr.Parser { return s.parser }

func (s *BeginContext) GetValue() antlr.Token { return s.value }

func (s *BeginContext) SetValue(v antlr.Token) { s.value = v }

func (s *BeginContext) TOBEGIN() antlr.TerminalNode {
	return s.GetToken(BundParserTOBEGIN, 0)
}

func (s *BeginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BeginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BeginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterBegin(s)
	}
}

func (s *BeginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitBegin(s)
	}
}

func (p *BundParser) Begin() (localctx IBeginContext) {
	localctx = NewBeginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BundParserRULE_begin)

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
		p.SetState(138)

		var _m = p.Match(BundParserTOBEGIN)

		localctx.(*BeginContext).value = _m
	}

	return localctx
}

// IEndContext is an interface to support dynamic dispatch.
type IEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsEndContext differentiates from other interfaces.
	IsEndContext()
}

type EndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyEndContext() *EndContext {
	var p = new(EndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_end
	return p
}

func (*EndContext) IsEndContext() {}

func NewEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndContext {
	var p = new(EndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_end

	return p
}

func (s *EndContext) GetParser() antlr.Parser { return s.parser }

func (s *EndContext) GetValue() antlr.Token { return s.value }

func (s *EndContext) SetValue(v antlr.Token) { s.value = v }

func (s *EndContext) TOEND() antlr.TerminalNode {
	return s.GetToken(BundParserTOEND, 0)
}

func (s *EndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterEnd(s)
	}
}

func (s *EndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitEnd(s)
	}
}

func (p *BundParser) End() (localctx IEndContext) {
	localctx = NewEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BundParserRULE_end)

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
		p.SetState(140)

		var _m = p.Match(BundParserTOEND)

		localctx.(*EndContext).value = _m
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
	p.EnterRule(localctx, 36, BundParserRULE_block)

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
	{
		p.SetState(142)
		p.Match(BundParserT__0)
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(143)

				var _x = p.Term()

				localctx.(*BlockContext)._term = _x
			}
			localctx.(*BlockContext).values = append(localctx.(*BlockContext).values, localctx.(*BlockContext)._term)

		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	{
		p.SetState(149)
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
	GetHead() IName_termContext

	// GetTail returns the tail rule contexts.
	GetTail() IPvalueContext

	// SetHead sets the head rule contexts.
	SetHead(IName_termContext)

	// SetTail sets the tail rule contexts.
	SetTail(IPvalueContext)

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	head   IName_termContext
	tail   IPvalueContext
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

func (s *PairContext) GetHead() IName_termContext { return s.head }

func (s *PairContext) GetTail() IPvalueContext { return s.tail }

func (s *PairContext) SetHead(v IName_termContext) { s.head = v }

func (s *PairContext) SetTail(v IPvalueContext) { s.tail = v }

func (s *PairContext) Name_term() IName_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IName_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IName_termContext)
}

func (s *PairContext) Pvalue() IPvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPvalueContext)
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
	p.EnterRule(localctx, 38, BundParserRULE_pair)

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
		p.SetState(151)
		p.Match(BundParserT__1)
	}
	{
		p.SetState(152)

		var _x = p.Name_term()

		localctx.(*PairContext).head = _x
	}
	{
		p.SetState(153)

		var _x = p.Pvalue()

		localctx.(*PairContext).tail = _x
	}
	{
		p.SetState(154)
		p.Match(BundParserT__2)
	}

	return localctx
}

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// GetName returns the name token.
	GetName() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
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

func (s *DirectiveContext) GetOp() antlr.Token { return s.op }

func (s *DirectiveContext) GetName() antlr.Token { return s.name }

func (s *DirectiveContext) SetOp(v antlr.Token) { s.op = v }

func (s *DirectiveContext) SetName(v antlr.Token) { s.name = v }

func (s *DirectiveContext) DIR() antlr.TerminalNode {
	return s.GetToken(BundParserDIR, 0)
}

func (s *DirectiveContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
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
	p.EnterRule(localctx, 40, BundParserRULE_directive)

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
		p.SetState(156)

		var _m = p.Match(BundParserDIR)

		localctx.(*DirectiveContext).op = _m
	}
	{
		p.SetState(157)

		var _m = p.Match(BundParserNAME)

		localctx.(*DirectiveContext).name = _m
	}

	return localctx
}

// ISys_directiveContext is an interface to support dynamic dispatch.
type ISys_directiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSys returns the sys token.
	GetSys() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// GetName returns the name token.
	GetName() antlr.Token

	// SetSys sets the sys token.
	SetSys(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsSys_directiveContext differentiates from other interfaces.
	IsSys_directiveContext()
}

type Sys_directiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	sys    antlr.Token
	op     antlr.Token
	name   antlr.Token
}

func NewEmptySys_directiveContext() *Sys_directiveContext {
	var p = new(Sys_directiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_sys_directive
	return p
}

func (*Sys_directiveContext) IsSys_directiveContext() {}

func NewSys_directiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sys_directiveContext {
	var p = new(Sys_directiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_sys_directive

	return p
}

func (s *Sys_directiveContext) GetParser() antlr.Parser { return s.parser }

func (s *Sys_directiveContext) GetSys() antlr.Token { return s.sys }

func (s *Sys_directiveContext) GetOp() antlr.Token { return s.op }

func (s *Sys_directiveContext) GetName() antlr.Token { return s.name }

func (s *Sys_directiveContext) SetSys(v antlr.Token) { s.sys = v }

func (s *Sys_directiveContext) SetOp(v antlr.Token) { s.op = v }

func (s *Sys_directiveContext) SetName(v antlr.Token) { s.name = v }

func (s *Sys_directiveContext) SYS() antlr.TerminalNode {
	return s.GetToken(BundParserSYS, 0)
}

func (s *Sys_directiveContext) DIR() antlr.TerminalNode {
	return s.GetToken(BundParserDIR, 0)
}

func (s *Sys_directiveContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Sys_directiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sys_directiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sys_directiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterSys_directive(s)
	}
}

func (s *Sys_directiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitSys_directive(s)
	}
}

func (p *BundParser) Sys_directive() (localctx ISys_directiveContext) {
	localctx = NewSys_directiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BundParserRULE_sys_directive)

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
		p.SetState(159)

		var _m = p.Match(BundParserSYS)

		localctx.(*Sys_directiveContext).sys = _m
	}
	{
		p.SetState(160)

		var _m = p.Match(BundParserDIR)

		localctx.(*Sys_directiveContext).op = _m
	}
	{
		p.SetState(161)

		var _m = p.Match(BundParserNAME)

		localctx.(*Sys_directiveContext).name = _m
	}

	return localctx
}

// ICmdContext is an interface to support dynamic dispatch.
type ICmdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCommand returns the command token.
	GetCommand() antlr.Token

	// SetCommand sets the command token.
	SetCommand(antlr.Token)

	// IsCmdContext differentiates from other interfaces.
	IsCmdContext()
}

type CmdContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	command antlr.Token
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

func (s *CmdContext) GetCommand() antlr.Token { return s.command }

func (s *CmdContext) SetCommand(v antlr.Token) { s.command = v }

func (s *CmdContext) CMD() antlr.TerminalNode {
	return s.GetToken(BundParserCMD, 0)
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
	p.EnterRule(localctx, 44, BundParserRULE_cmd)

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
		p.SetState(163)

		var _m = p.Match(BundParserCMD)

		localctx.(*CmdContext).command = _m
	}

	return localctx
}

// ISys_cmdContext is an interface to support dynamic dispatch.
type ISys_cmdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSys returns the sys token.
	GetSys() antlr.Token

	// GetCommand returns the command token.
	GetCommand() antlr.Token

	// SetSys sets the sys token.
	SetSys(antlr.Token)

	// SetCommand sets the command token.
	SetCommand(antlr.Token)

	// IsSys_cmdContext differentiates from other interfaces.
	IsSys_cmdContext()
}

type Sys_cmdContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	sys     antlr.Token
	command antlr.Token
}

func NewEmptySys_cmdContext() *Sys_cmdContext {
	var p = new(Sys_cmdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_sys_cmd
	return p
}

func (*Sys_cmdContext) IsSys_cmdContext() {}

func NewSys_cmdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sys_cmdContext {
	var p = new(Sys_cmdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_sys_cmd

	return p
}

func (s *Sys_cmdContext) GetParser() antlr.Parser { return s.parser }

func (s *Sys_cmdContext) GetSys() antlr.Token { return s.sys }

func (s *Sys_cmdContext) GetCommand() antlr.Token { return s.command }

func (s *Sys_cmdContext) SetSys(v antlr.Token) { s.sys = v }

func (s *Sys_cmdContext) SetCommand(v antlr.Token) { s.command = v }

func (s *Sys_cmdContext) SYS() antlr.TerminalNode {
	return s.GetToken(BundParserSYS, 0)
}

func (s *Sys_cmdContext) CMD() antlr.TerminalNode {
	return s.GetToken(BundParserCMD, 0)
}

func (s *Sys_cmdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sys_cmdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sys_cmdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterSys_cmd(s)
	}
}

func (s *Sys_cmdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitSys_cmd(s)
	}
}

func (p *BundParser) Sys_cmd() (localctx ISys_cmdContext) {
	localctx = NewSys_cmdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BundParserRULE_sys_cmd)

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
		p.SetState(165)

		var _m = p.Match(BundParserSYS)

		localctx.(*Sys_cmdContext).sys = _m
	}
	{
		p.SetState(166)

		var _m = p.Match(BundParserCMD)

		localctx.(*Sys_cmdContext).command = _m
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
	p.EnterRule(localctx, 48, BundParserRULE_lambda)
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
		p.SetState(168)
		p.Match(BundParserT__3)
	}
	{
		p.SetState(169)

		var _m = p.Match(BundParserNAME)

		localctx.(*LambdaContext).name = _m
	}
	{
		p.SetState(170)
		p.Match(BundParserT__4)
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__1)|(1<<BundParserINTEGER)|(1<<BundParserOCT_INTEGER)|(1<<BundParserHEX_INTEGER)|(1<<BundParserBIN_INTEGER)|(1<<BundParserNTH)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserNAME)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND)|(1<<BundParserDIR)|(1<<BundParserCMD)|(1<<BundParserDUPLICATE)|(1<<BundParserDROP))) != 0 {
		{
			p.SetState(171)

			var _x = p.Term()

			localctx.(*LambdaContext)._term = _x
		}
		localctx.(*LambdaContext).body = append(localctx.(*LambdaContext).body, localctx.(*LambdaContext)._term)

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(177)
		p.Match(BundParserT__5)
	}

	return localctx
}

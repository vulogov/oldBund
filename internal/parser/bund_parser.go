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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 26, 141,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 7, 2, 61, 10, 2, 12, 2, 14, 2, 64, 11, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 5, 3, 83, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 7, 17,
	113, 10, 17, 12, 17, 14, 17, 116, 11, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21,
	3, 21, 7, 21, 134, 10, 21, 12, 21, 14, 21, 137, 11, 21, 3, 21, 3, 21, 3,
	21, 2, 2, 22, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 2, 2, 2, 156, 2, 62, 3, 2, 2, 2, 4, 82, 3, 2, 2, 2, 6,
	84, 3, 2, 2, 2, 8, 86, 3, 2, 2, 2, 10, 88, 3, 2, 2, 2, 12, 90, 3, 2, 2,
	2, 14, 92, 3, 2, 2, 2, 16, 94, 3, 2, 2, 2, 18, 96, 3, 2, 2, 2, 20, 98,
	3, 2, 2, 2, 22, 100, 3, 2, 2, 2, 24, 102, 3, 2, 2, 2, 26, 104, 3, 2, 2,
	2, 28, 106, 3, 2, 2, 2, 30, 108, 3, 2, 2, 2, 32, 110, 3, 2, 2, 2, 34, 119,
	3, 2, 2, 2, 36, 124, 3, 2, 2, 2, 38, 127, 3, 2, 2, 2, 40, 129, 3, 2, 2,
	2, 42, 61, 5, 20, 11, 2, 43, 61, 5, 18, 10, 2, 44, 61, 5, 16, 9, 2, 45,
	61, 5, 14, 8, 2, 46, 61, 5, 12, 7, 2, 47, 61, 5, 10, 6, 2, 48, 61, 5, 8,
	5, 2, 49, 61, 5, 6, 4, 2, 50, 61, 5, 22, 12, 2, 51, 61, 5, 24, 13, 2, 52,
	61, 5, 26, 14, 2, 53, 61, 5, 28, 15, 2, 54, 61, 5, 30, 16, 2, 55, 61, 5,
	34, 18, 2, 56, 61, 5, 32, 17, 2, 57, 61, 5, 36, 19, 2, 58, 61, 5, 38, 20,
	2, 59, 61, 5, 40, 21, 2, 60, 42, 3, 2, 2, 2, 60, 43, 3, 2, 2, 2, 60, 44,
	3, 2, 2, 2, 60, 45, 3, 2, 2, 2, 60, 46, 3, 2, 2, 2, 60, 47, 3, 2, 2, 2,
	60, 48, 3, 2, 2, 2, 60, 49, 3, 2, 2, 2, 60, 50, 3, 2, 2, 2, 60, 51, 3,
	2, 2, 2, 60, 52, 3, 2, 2, 2, 60, 53, 3, 2, 2, 2, 60, 54, 3, 2, 2, 2, 60,
	55, 3, 2, 2, 2, 60, 56, 3, 2, 2, 2, 60, 57, 3, 2, 2, 2, 60, 58, 3, 2, 2,
	2, 60, 59, 3, 2, 2, 2, 61, 64, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63,
	3, 2, 2, 2, 63, 3, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 65, 83, 5, 20, 11, 2,
	66, 83, 5, 18, 10, 2, 67, 83, 5, 16, 9, 2, 68, 83, 5, 14, 8, 2, 69, 83,
	5, 12, 7, 2, 70, 83, 5, 10, 6, 2, 71, 83, 5, 8, 5, 2, 72, 83, 5, 6, 4,
	2, 73, 83, 5, 22, 12, 2, 74, 83, 5, 24, 13, 2, 75, 83, 5, 26, 14, 2, 76,
	83, 5, 28, 15, 2, 77, 83, 5, 30, 16, 2, 78, 83, 5, 34, 18, 2, 79, 83, 5,
	32, 17, 2, 80, 83, 5, 36, 19, 2, 81, 83, 5, 38, 20, 2, 82, 65, 3, 2, 2,
	2, 82, 66, 3, 2, 2, 2, 82, 67, 3, 2, 2, 2, 82, 68, 3, 2, 2, 2, 82, 69,
	3, 2, 2, 2, 82, 70, 3, 2, 2, 2, 82, 71, 3, 2, 2, 2, 82, 72, 3, 2, 2, 2,
	82, 73, 3, 2, 2, 2, 82, 74, 3, 2, 2, 2, 82, 75, 3, 2, 2, 2, 82, 76, 3,
	2, 2, 2, 82, 77, 3, 2, 2, 2, 82, 78, 3, 2, 2, 2, 82, 79, 3, 2, 2, 2, 82,
	80, 3, 2, 2, 2, 82, 81, 3, 2, 2, 2, 83, 5, 3, 2, 2, 2, 84, 85, 7, 16, 2,
	2, 85, 7, 3, 2, 2, 2, 86, 87, 7, 17, 2, 2, 87, 9, 3, 2, 2, 2, 88, 89, 7,
	15, 2, 2, 89, 11, 3, 2, 2, 2, 90, 91, 7, 14, 2, 2, 91, 13, 3, 2, 2, 2,
	92, 93, 7, 13, 2, 2, 93, 15, 3, 2, 2, 2, 94, 95, 7, 12, 2, 2, 95, 17, 3,
	2, 2, 2, 96, 97, 7, 11, 2, 2, 97, 19, 3, 2, 2, 2, 98, 99, 7, 9, 2, 2, 99,
	21, 3, 2, 2, 2, 100, 101, 7, 18, 2, 2, 101, 23, 3, 2, 2, 2, 102, 103, 7,
	24, 2, 2, 103, 25, 3, 2, 2, 2, 104, 105, 7, 25, 2, 2, 105, 27, 3, 2, 2,
	2, 106, 107, 7, 19, 2, 2, 107, 29, 3, 2, 2, 2, 108, 109, 7, 20, 2, 2, 109,
	31, 3, 2, 2, 2, 110, 114, 7, 3, 2, 2, 111, 113, 5, 4, 3, 2, 112, 111, 3,
	2, 2, 2, 113, 116, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2,
	2, 115, 117, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 117, 118, 7, 3, 2, 2, 118,
	33, 3, 2, 2, 2, 119, 120, 7, 4, 2, 2, 120, 121, 7, 18, 2, 2, 121, 122,
	5, 4, 3, 2, 122, 123, 7, 5, 2, 2, 123, 35, 3, 2, 2, 2, 124, 125, 7, 22,
	2, 2, 125, 126, 7, 18, 2, 2, 126, 37, 3, 2, 2, 2, 127, 128, 7, 23, 2, 2,
	128, 39, 3, 2, 2, 2, 129, 130, 7, 6, 2, 2, 130, 131, 7, 18, 2, 2, 131,
	135, 7, 7, 2, 2, 132, 134, 5, 4, 3, 2, 133, 132, 3, 2, 2, 2, 134, 137,
	3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 138, 3, 2,
	2, 2, 137, 135, 3, 2, 2, 2, 138, 139, 7, 8, 2, 2, 139, 41, 3, 2, 2, 2,
	7, 60, 62, 82, 114, 135,
}
var literalNames = []string{
	"", "'|'", "'('", "')'", "'['", "']'", "'.'", "", "", "", "", "", "", "",
	"", "", "", "':'", "';'", "'/'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "INTEGER", "DECIMAL_INTEGER", "OCT_INTEGER",
	"HEX_INTEGER", "BIN_INTEGER", "FLOAT_NUMBER", "STRING", "TRUE", "FALSE",
	"NAME", "TOBEGIN", "TOEND", "SLASH", "DIR", "CMD", "DUPLICATE", "DROP",
	"SKIP_",
}

var ruleNames = []string{
	"expressions", "term", "true_term", "false_term", "string_term", "float_term",
	"bin_integer", "hex_integer", "oct_integer", "integer", "name_term", "duplicate_term",
	"drop_term", "begin", "end", "block", "pair", "directive", "cmd", "lambda",
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
	BundParserRULE_expressions    = 0
	BundParserRULE_term           = 1
	BundParserRULE_true_term      = 2
	BundParserRULE_false_term     = 3
	BundParserRULE_string_term    = 4
	BundParserRULE_float_term     = 5
	BundParserRULE_bin_integer    = 6
	BundParserRULE_hex_integer    = 7
	BundParserRULE_oct_integer    = 8
	BundParserRULE_integer        = 9
	BundParserRULE_name_term      = 10
	BundParserRULE_duplicate_term = 11
	BundParserRULE_drop_term      = 12
	BundParserRULE_begin          = 13
	BundParserRULE_end            = 14
	BundParserRULE_block          = 15
	BundParserRULE_pair           = 16
	BundParserRULE_directive      = 17
	BundParserRULE_cmd            = 18
	BundParserRULE_lambda         = 19
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
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__1)|(1<<BundParserT__3)|(1<<BundParserINTEGER)|(1<<BundParserOCT_INTEGER)|(1<<BundParserHEX_INTEGER)|(1<<BundParserBIN_INTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserNAME)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND)|(1<<BundParserDIR)|(1<<BundParserCMD)|(1<<BundParserDUPLICATE)|(1<<BundParserDROP))) != 0 {
		p.SetState(58)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case BundParserINTEGER:
			{
				p.SetState(40)
				p.Integer()
			}

		case BundParserOCT_INTEGER:
			{
				p.SetState(41)
				p.Oct_integer()
			}

		case BundParserHEX_INTEGER:
			{
				p.SetState(42)
				p.Hex_integer()
			}

		case BundParserBIN_INTEGER:
			{
				p.SetState(43)
				p.Bin_integer()
			}

		case BundParserFLOAT_NUMBER:
			{
				p.SetState(44)
				p.Float_term()
			}

		case BundParserSTRING:
			{
				p.SetState(45)
				p.String_term()
			}

		case BundParserFALSE:
			{
				p.SetState(46)
				p.False_term()
			}

		case BundParserTRUE:
			{
				p.SetState(47)
				p.True_term()
			}

		case BundParserNAME:
			{
				p.SetState(48)
				p.Name_term()
			}

		case BundParserDUPLICATE:
			{
				p.SetState(49)
				p.Duplicate_term()
			}

		case BundParserDROP:
			{
				p.SetState(50)
				p.Drop_term()
			}

		case BundParserTOBEGIN:
			{
				p.SetState(51)
				p.Begin()
			}

		case BundParserTOEND:
			{
				p.SetState(52)
				p.End()
			}

		case BundParserT__1:
			{
				p.SetState(53)
				p.Pair()
			}

		case BundParserT__0:
			{
				p.SetState(54)
				p.Block()
			}

		case BundParserDIR:
			{
				p.SetState(55)
				p.Directive()
			}

		case BundParserCMD:
			{
				p.SetState(56)
				p.Cmd()
			}

		case BundParserT__3:
			{
				p.SetState(57)
				p.Lambda()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(62)
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
	p.SetState(80)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserINTEGER:
		{
			p.SetState(63)
			p.Integer()
		}

	case BundParserOCT_INTEGER:
		{
			p.SetState(64)
			p.Oct_integer()
		}

	case BundParserHEX_INTEGER:
		{
			p.SetState(65)
			p.Hex_integer()
		}

	case BundParserBIN_INTEGER:
		{
			p.SetState(66)
			p.Bin_integer()
		}

	case BundParserFLOAT_NUMBER:
		{
			p.SetState(67)
			p.Float_term()
		}

	case BundParserSTRING:
		{
			p.SetState(68)
			p.String_term()
		}

	case BundParserFALSE:
		{
			p.SetState(69)
			p.False_term()
		}

	case BundParserTRUE:
		{
			p.SetState(70)
			p.True_term()
		}

	case BundParserNAME:
		{
			p.SetState(71)
			p.Name_term()
		}

	case BundParserDUPLICATE:
		{
			p.SetState(72)
			p.Duplicate_term()
		}

	case BundParserDROP:
		{
			p.SetState(73)
			p.Drop_term()
		}

	case BundParserTOBEGIN:
		{
			p.SetState(74)
			p.Begin()
		}

	case BundParserTOEND:
		{
			p.SetState(75)
			p.End()
		}

	case BundParserT__1:
		{
			p.SetState(76)
			p.Pair()
		}

	case BundParserT__0:
		{
			p.SetState(77)
			p.Block()
		}

	case BundParserDIR:
		{
			p.SetState(78)
			p.Directive()
		}

	case BundParserCMD:
		{
			p.SetState(79)
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
	p.EnterRule(localctx, 4, BundParserRULE_true_term)

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
		p.SetState(82)

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
	p.EnterRule(localctx, 6, BundParserRULE_false_term)

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
		p.SetState(84)

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
	p.EnterRule(localctx, 8, BundParserRULE_string_term)

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
		p.SetState(86)

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
	p.EnterRule(localctx, 10, BundParserRULE_float_term)

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
		p.SetState(88)

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
	p.EnterRule(localctx, 12, BundParserRULE_bin_integer)

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
		p.SetState(90)

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
	p.EnterRule(localctx, 14, BundParserRULE_hex_integer)

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
		p.SetState(92)

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
	p.EnterRule(localctx, 16, BundParserRULE_oct_integer)

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
		p.SetState(94)

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
	p.EnterRule(localctx, 18, BundParserRULE_integer)

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
		p.SetState(96)

		var _m = p.Match(BundParserINTEGER)

		localctx.(*IntegerContext).value = _m
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
	p.EnterRule(localctx, 20, BundParserRULE_name_term)

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
		p.SetState(98)

		var _m = p.Match(BundParserNAME)

		localctx.(*Name_termContext).value = _m
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
	p.EnterRule(localctx, 22, BundParserRULE_duplicate_term)

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
		p.SetState(100)

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
	p.EnterRule(localctx, 24, BundParserRULE_drop_term)

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
		p.SetState(102)

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
	p.EnterRule(localctx, 26, BundParserRULE_begin)

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
		p.SetState(104)

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
	p.EnterRule(localctx, 28, BundParserRULE_end)

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
		p.SetState(106)

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
	p.EnterRule(localctx, 30, BundParserRULE_block)

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
		p.SetState(108)
		p.Match(BundParserT__0)
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(109)

				var _x = p.Term()

				localctx.(*BlockContext)._term = _x
			}
			localctx.(*BlockContext).values = append(localctx.(*BlockContext).values, localctx.(*BlockContext)._term)

		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	{
		p.SetState(115)
		p.Match(BundParserT__0)
	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHead returns the head token.
	GetHead() antlr.Token

	// SetHead sets the head token.
	SetHead(antlr.Token)

	// GetTail returns the tail rule contexts.
	GetTail() ITermContext

	// SetTail sets the tail rule contexts.
	SetTail(ITermContext)

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	head   antlr.Token
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

func (s *PairContext) GetHead() antlr.Token { return s.head }

func (s *PairContext) SetHead(v antlr.Token) { s.head = v }

func (s *PairContext) GetTail() ITermContext { return s.tail }

func (s *PairContext) SetTail(v ITermContext) { s.tail = v }

func (s *PairContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *PairContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 32, BundParserRULE_pair)

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
		p.Match(BundParserT__1)
	}
	{
		p.SetState(118)

		var _m = p.Match(BundParserNAME)

		localctx.(*PairContext).head = _m
	}
	{
		p.SetState(119)

		var _x = p.Term()

		localctx.(*PairContext).tail = _x
	}
	{
		p.SetState(120)
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
	p.EnterRule(localctx, 34, BundParserRULE_directive)

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
		p.SetState(122)

		var _m = p.Match(BundParserDIR)

		localctx.(*DirectiveContext).op = _m
	}
	{
		p.SetState(123)

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
	p.EnterRule(localctx, 36, BundParserRULE_cmd)

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

		var _m = p.Match(BundParserCMD)

		localctx.(*CmdContext).command = _m
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
	p.EnterRule(localctx, 38, BundParserRULE_lambda)
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
		p.SetState(127)
		p.Match(BundParserT__3)
	}
	{
		p.SetState(128)

		var _m = p.Match(BundParserNAME)

		localctx.(*LambdaContext).name = _m
	}
	{
		p.SetState(129)
		p.Match(BundParserT__4)
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__1)|(1<<BundParserINTEGER)|(1<<BundParserOCT_INTEGER)|(1<<BundParserHEX_INTEGER)|(1<<BundParserBIN_INTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserNAME)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND)|(1<<BundParserDIR)|(1<<BundParserCMD)|(1<<BundParserDUPLICATE)|(1<<BundParserDROP))) != 0 {
		{
			p.SetState(130)

			var _x = p.Term()

			localctx.(*LambdaContext)._term = _x
		}
		localctx.(*LambdaContext).body = append(localctx.(*LambdaContext).body, localctx.(*LambdaContext)._term)

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(136)
		p.Match(BundParserT__5)
	}

	return localctx
}

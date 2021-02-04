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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 33, 237,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 7, 2, 93, 10, 2, 12, 2, 14, 2, 96, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 122, 10, 3, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4,
	136, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 21, 3, 21, 7, 21, 173, 10, 21, 12, 21, 14, 21, 176, 11,
	21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3,
	27, 3, 27, 3, 27, 7, 27, 201, 10, 27, 12, 27, 14, 27, 204, 11, 27, 3, 27,
	3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 212, 10, 28, 12, 28, 14, 28,
	215, 11, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 2, 2, 33, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
	28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62,
	2, 2, 2, 270, 2, 94, 3, 2, 2, 2, 4, 121, 3, 2, 2, 2, 6, 135, 3, 2, 2, 2,
	8, 137, 3, 2, 2, 2, 10, 139, 3, 2, 2, 2, 12, 141, 3, 2, 2, 2, 14, 143,
	3, 2, 2, 2, 16, 145, 3, 2, 2, 2, 18, 147, 3, 2, 2, 2, 20, 149, 3, 2, 2,
	2, 22, 151, 3, 2, 2, 2, 24, 153, 3, 2, 2, 2, 26, 155, 3, 2, 2, 2, 28, 157,
	3, 2, 2, 2, 30, 159, 3, 2, 2, 2, 32, 162, 3, 2, 2, 2, 34, 164, 3, 2, 2,
	2, 36, 166, 3, 2, 2, 2, 38, 168, 3, 2, 2, 2, 40, 170, 3, 2, 2, 2, 42, 179,
	3, 2, 2, 2, 44, 184, 3, 2, 2, 2, 46, 187, 3, 2, 2, 2, 48, 191, 3, 2, 2,
	2, 50, 193, 3, 2, 2, 2, 52, 196, 3, 2, 2, 2, 54, 207, 3, 2, 2, 2, 56, 218,
	3, 2, 2, 2, 58, 222, 3, 2, 2, 2, 60, 226, 3, 2, 2, 2, 62, 231, 3, 2, 2,
	2, 64, 93, 5, 22, 12, 2, 65, 93, 5, 24, 13, 2, 66, 93, 5, 20, 11, 2, 67,
	93, 5, 18, 10, 2, 68, 93, 5, 16, 9, 2, 69, 93, 5, 14, 8, 2, 70, 93, 5,
	12, 7, 2, 71, 93, 5, 10, 6, 2, 72, 93, 5, 8, 5, 2, 73, 93, 5, 28, 15, 2,
	74, 93, 5, 30, 16, 2, 75, 93, 5, 32, 17, 2, 76, 93, 5, 34, 18, 2, 77, 93,
	5, 36, 19, 2, 78, 93, 5, 38, 20, 2, 79, 93, 5, 42, 22, 2, 80, 93, 5, 40,
	21, 2, 81, 93, 5, 44, 23, 2, 82, 93, 5, 46, 24, 2, 83, 93, 5, 48, 25, 2,
	84, 93, 5, 50, 26, 2, 85, 93, 5, 26, 14, 2, 86, 93, 5, 58, 30, 2, 87, 93,
	5, 56, 29, 2, 88, 93, 5, 62, 32, 2, 89, 93, 5, 60, 31, 2, 90, 93, 5, 52,
	27, 2, 91, 93, 5, 54, 28, 2, 92, 64, 3, 2, 2, 2, 92, 65, 3, 2, 2, 2, 92,
	66, 3, 2, 2, 2, 92, 67, 3, 2, 2, 2, 92, 68, 3, 2, 2, 2, 92, 69, 3, 2, 2,
	2, 92, 70, 3, 2, 2, 2, 92, 71, 3, 2, 2, 2, 92, 72, 3, 2, 2, 2, 92, 73,
	3, 2, 2, 2, 92, 74, 3, 2, 2, 2, 92, 75, 3, 2, 2, 2, 92, 76, 3, 2, 2, 2,
	92, 77, 3, 2, 2, 2, 92, 78, 3, 2, 2, 2, 92, 79, 3, 2, 2, 2, 92, 80, 3,
	2, 2, 2, 92, 81, 3, 2, 2, 2, 92, 82, 3, 2, 2, 2, 92, 83, 3, 2, 2, 2, 92,
	84, 3, 2, 2, 2, 92, 85, 3, 2, 2, 2, 92, 86, 3, 2, 2, 2, 92, 87, 3, 2, 2,
	2, 92, 88, 3, 2, 2, 2, 92, 89, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 91,
	3, 2, 2, 2, 93, 96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2,
	95, 3, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 97, 122, 5, 24, 13, 2, 98, 122,
	5, 22, 12, 2, 99, 122, 5, 20, 11, 2, 100, 122, 5, 18, 10, 2, 101, 122,
	5, 16, 9, 2, 102, 122, 5, 14, 8, 2, 103, 122, 5, 12, 7, 2, 104, 122, 5,
	10, 6, 2, 105, 122, 5, 8, 5, 2, 106, 122, 5, 28, 15, 2, 107, 122, 5, 32,
	17, 2, 108, 122, 5, 34, 18, 2, 109, 122, 5, 58, 30, 2, 110, 122, 5, 56,
	29, 2, 111, 122, 5, 62, 32, 2, 112, 122, 5, 60, 31, 2, 113, 122, 5, 36,
	19, 2, 114, 122, 5, 38, 20, 2, 115, 122, 5, 42, 22, 2, 116, 122, 5, 40,
	21, 2, 117, 122, 5, 44, 23, 2, 118, 122, 5, 48, 25, 2, 119, 122, 5, 50,
	26, 2, 120, 122, 5, 26, 14, 2, 121, 97, 3, 2, 2, 2, 121, 98, 3, 2, 2, 2,
	121, 99, 3, 2, 2, 2, 121, 100, 3, 2, 2, 2, 121, 101, 3, 2, 2, 2, 121, 102,
	3, 2, 2, 2, 121, 103, 3, 2, 2, 2, 121, 104, 3, 2, 2, 2, 121, 105, 3, 2,
	2, 2, 121, 106, 3, 2, 2, 2, 121, 107, 3, 2, 2, 2, 121, 108, 3, 2, 2, 2,
	121, 109, 3, 2, 2, 2, 121, 110, 3, 2, 2, 2, 121, 111, 3, 2, 2, 2, 121,
	112, 3, 2, 2, 2, 121, 113, 3, 2, 2, 2, 121, 114, 3, 2, 2, 2, 121, 115,
	3, 2, 2, 2, 121, 116, 3, 2, 2, 2, 121, 117, 3, 2, 2, 2, 121, 118, 3, 2,
	2, 2, 121, 119, 3, 2, 2, 2, 121, 120, 3, 2, 2, 2, 122, 5, 3, 2, 2, 2, 123,
	136, 5, 22, 12, 2, 124, 136, 5, 24, 13, 2, 125, 136, 5, 20, 11, 2, 126,
	136, 5, 18, 10, 2, 127, 136, 5, 16, 9, 2, 128, 136, 5, 14, 8, 2, 129, 136,
	5, 12, 7, 2, 130, 136, 5, 10, 6, 2, 131, 136, 5, 8, 5, 2, 132, 136, 5,
	28, 15, 2, 133, 136, 5, 44, 23, 2, 134, 136, 5, 48, 25, 2, 135, 123, 3,
	2, 2, 2, 135, 124, 3, 2, 2, 2, 135, 125, 3, 2, 2, 2, 135, 126, 3, 2, 2,
	2, 135, 127, 3, 2, 2, 2, 135, 128, 3, 2, 2, 2, 135, 129, 3, 2, 2, 2, 135,
	130, 3, 2, 2, 2, 135, 131, 3, 2, 2, 2, 135, 132, 3, 2, 2, 2, 135, 133,
	3, 2, 2, 2, 135, 134, 3, 2, 2, 2, 136, 7, 3, 2, 2, 2, 137, 138, 7, 21,
	2, 2, 138, 9, 3, 2, 2, 2, 139, 140, 7, 22, 2, 2, 140, 11, 3, 2, 2, 2, 141,
	142, 7, 20, 2, 2, 142, 13, 3, 2, 2, 2, 143, 144, 7, 19, 2, 2, 144, 15,
	3, 2, 2, 2, 145, 146, 7, 17, 2, 2, 146, 17, 3, 2, 2, 2, 147, 148, 7, 16,
	2, 2, 148, 19, 3, 2, 2, 2, 149, 150, 7, 15, 2, 2, 150, 21, 3, 2, 2, 2,
	151, 152, 7, 13, 2, 2, 152, 23, 3, 2, 2, 2, 153, 154, 7, 18, 2, 2, 154,
	25, 3, 2, 2, 2, 155, 156, 7, 32, 2, 2, 156, 27, 3, 2, 2, 2, 157, 158, 7,
	23, 2, 2, 158, 29, 3, 2, 2, 2, 159, 160, 7, 27, 2, 2, 160, 161, 7, 23,
	2, 2, 161, 31, 3, 2, 2, 2, 162, 163, 7, 30, 2, 2, 163, 33, 3, 2, 2, 2,
	164, 165, 7, 31, 2, 2, 165, 35, 3, 2, 2, 2, 166, 167, 7, 24, 2, 2, 167,
	37, 3, 2, 2, 2, 168, 169, 7, 25, 2, 2, 169, 39, 3, 2, 2, 2, 170, 174, 7,
	3, 2, 2, 171, 173, 5, 4, 3, 2, 172, 171, 3, 2, 2, 2, 173, 176, 3, 2, 2,
	2, 174, 172, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 177, 3, 2, 2, 2, 176,
	174, 3, 2, 2, 2, 177, 178, 7, 3, 2, 2, 178, 41, 3, 2, 2, 2, 179, 180, 7,
	4, 2, 2, 180, 181, 5, 28, 15, 2, 181, 182, 5, 6, 4, 2, 182, 183, 7, 5,
	2, 2, 183, 43, 3, 2, 2, 2, 184, 185, 7, 28, 2, 2, 185, 186, 7, 23, 2, 2,
	186, 45, 3, 2, 2, 2, 187, 188, 7, 27, 2, 2, 188, 189, 7, 28, 2, 2, 189,
	190, 7, 23, 2, 2, 190, 47, 3, 2, 2, 2, 191, 192, 7, 29, 2, 2, 192, 49,
	3, 2, 2, 2, 193, 194, 7, 27, 2, 2, 194, 195, 7, 29, 2, 2, 195, 51, 3, 2,
	2, 2, 196, 197, 7, 6, 2, 2, 197, 198, 7, 23, 2, 2, 198, 202, 7, 7, 2, 2,
	199, 201, 5, 4, 3, 2, 200, 199, 3, 2, 2, 2, 201, 204, 3, 2, 2, 2, 202,
	200, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 205, 3, 2, 2, 2, 204, 202,
	3, 2, 2, 2, 205, 206, 7, 8, 2, 2, 206, 53, 3, 2, 2, 2, 207, 208, 7, 9,
	2, 2, 208, 209, 7, 29, 2, 2, 209, 213, 7, 10, 2, 2, 210, 212, 5, 4, 3,
	2, 211, 210, 3, 2, 2, 2, 212, 215, 3, 2, 2, 2, 213, 211, 3, 2, 2, 2, 213,
	214, 3, 2, 2, 2, 214, 216, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 216, 217,
	7, 8, 2, 2, 217, 55, 3, 2, 2, 2, 218, 219, 7, 6, 2, 2, 219, 220, 7, 23,
	2, 2, 220, 221, 7, 11, 2, 2, 221, 57, 3, 2, 2, 2, 222, 223, 7, 12, 2, 2,
	223, 224, 7, 23, 2, 2, 224, 225, 7, 7, 2, 2, 225, 59, 3, 2, 2, 2, 226,
	227, 7, 27, 2, 2, 227, 228, 7, 6, 2, 2, 228, 229, 7, 23, 2, 2, 229, 230,
	7, 11, 2, 2, 230, 61, 3, 2, 2, 2, 231, 232, 7, 12, 2, 2, 232, 233, 7, 23,
	2, 2, 233, 234, 7, 7, 2, 2, 234, 235, 7, 27, 2, 2, 235, 63, 3, 2, 2, 2,
	9, 92, 94, 121, 135, 174, 202, 213,
}
var literalNames = []string{
	"", "'|'", "'{'", "'}'", "'['", "']'", "'.'", "'[['", "']]'", "'>'", "'<'",
	"", "", "", "", "", "", "", "", "", "", "", "':'", "';'", "'/'", "", "",
	"", "'^'", "'_'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "INTEGER", "DECIMAL_INTEGER",
	"OCT_INTEGER", "HEX_INTEGER", "BIN_INTEGER", "NTH", "FLOAT_NUMBER", "STRING",
	"TRUE", "FALSE", "NAME", "TOBEGIN", "TOEND", "SLASH", "SYS", "DIR", "CMD",
	"DUPLICATE", "DROP", "EXECUTE", "SKIP_",
}

var ruleNames = []string{
	"expressions", "term", "pvalue", "true_term", "false_term", "string_term",
	"float_term", "bin_integer", "hex_integer", "oct_integer", "integer", "nth",
	"exec", "name_term", "sys_name_term", "duplicate_term", "drop_term", "begin",
	"end", "block", "pair", "directive", "sys_directive", "cmd", "sys_cmd",
	"lambda", "lambda_cmd", "channel_out", "channel_in", "sys_channel_out",
	"sys_channel_in",
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
	BundParserT__6            = 7
	BundParserT__7            = 8
	BundParserT__8            = 9
	BundParserT__9            = 10
	BundParserINTEGER         = 11
	BundParserDECIMAL_INTEGER = 12
	BundParserOCT_INTEGER     = 13
	BundParserHEX_INTEGER     = 14
	BundParserBIN_INTEGER     = 15
	BundParserNTH             = 16
	BundParserFLOAT_NUMBER    = 17
	BundParserSTRING          = 18
	BundParserTRUE            = 19
	BundParserFALSE           = 20
	BundParserNAME            = 21
	BundParserTOBEGIN         = 22
	BundParserTOEND           = 23
	BundParserSLASH           = 24
	BundParserSYS             = 25
	BundParserDIR             = 26
	BundParserCMD             = 27
	BundParserDUPLICATE       = 28
	BundParserDROP            = 29
	BundParserEXECUTE         = 30
	BundParserSKIP_           = 31
)

// BundParser rules.
const (
	BundParserRULE_expressions     = 0
	BundParserRULE_term            = 1
	BundParserRULE_pvalue          = 2
	BundParserRULE_true_term       = 3
	BundParserRULE_false_term      = 4
	BundParserRULE_string_term     = 5
	BundParserRULE_float_term      = 6
	BundParserRULE_bin_integer     = 7
	BundParserRULE_hex_integer     = 8
	BundParserRULE_oct_integer     = 9
	BundParserRULE_integer         = 10
	BundParserRULE_nth             = 11
	BundParserRULE_exec            = 12
	BundParserRULE_name_term       = 13
	BundParserRULE_sys_name_term   = 14
	BundParserRULE_duplicate_term  = 15
	BundParserRULE_drop_term       = 16
	BundParserRULE_begin           = 17
	BundParserRULE_end             = 18
	BundParserRULE_block           = 19
	BundParserRULE_pair            = 20
	BundParserRULE_directive       = 21
	BundParserRULE_sys_directive   = 22
	BundParserRULE_cmd             = 23
	BundParserRULE_sys_cmd         = 24
	BundParserRULE_lambda          = 25
	BundParserRULE_lambda_cmd      = 26
	BundParserRULE_channel_out     = 27
	BundParserRULE_channel_in      = 28
	BundParserRULE_sys_channel_out = 29
	BundParserRULE_sys_channel_in  = 30
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

func (s *ExpressionsContext) AllExec() []IExecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExecContext)(nil)).Elem())
	var tst = make([]IExecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExecContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Exec(i int) IExecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExecContext)
}

func (s *ExpressionsContext) AllChannel_in() []IChannel_inContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IChannel_inContext)(nil)).Elem())
	var tst = make([]IChannel_inContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IChannel_inContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Channel_in(i int) IChannel_inContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChannel_inContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IChannel_inContext)
}

func (s *ExpressionsContext) AllChannel_out() []IChannel_outContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IChannel_outContext)(nil)).Elem())
	var tst = make([]IChannel_outContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IChannel_outContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Channel_out(i int) IChannel_outContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChannel_outContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IChannel_outContext)
}

func (s *ExpressionsContext) AllSys_channel_in() []ISys_channel_inContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISys_channel_inContext)(nil)).Elem())
	var tst = make([]ISys_channel_inContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISys_channel_inContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Sys_channel_in(i int) ISys_channel_inContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISys_channel_inContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISys_channel_inContext)
}

func (s *ExpressionsContext) AllSys_channel_out() []ISys_channel_outContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISys_channel_outContext)(nil)).Elem())
	var tst = make([]ISys_channel_outContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISys_channel_outContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Sys_channel_out(i int) ISys_channel_outContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISys_channel_outContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISys_channel_outContext)
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

func (s *ExpressionsContext) AllLambda_cmd() []ILambda_cmdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILambda_cmdContext)(nil)).Elem())
	var tst = make([]ILambda_cmdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILambda_cmdContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Lambda_cmd(i int) ILambda_cmdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILambda_cmdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILambda_cmdContext)
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
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__1)|(1<<BundParserT__3)|(1<<BundParserT__6)|(1<<BundParserT__9)|(1<<BundParserINTEGER)|(1<<BundParserOCT_INTEGER)|(1<<BundParserHEX_INTEGER)|(1<<BundParserBIN_INTEGER)|(1<<BundParserNTH)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserNAME)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND)|(1<<BundParserSYS)|(1<<BundParserDIR)|(1<<BundParserCMD)|(1<<BundParserDUPLICATE)|(1<<BundParserDROP)|(1<<BundParserEXECUTE))) != 0 {
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(62)
				p.Integer()
			}

		case 2:
			{
				p.SetState(63)
				p.Nth()
			}

		case 3:
			{
				p.SetState(64)
				p.Oct_integer()
			}

		case 4:
			{
				p.SetState(65)
				p.Hex_integer()
			}

		case 5:
			{
				p.SetState(66)
				p.Bin_integer()
			}

		case 6:
			{
				p.SetState(67)
				p.Float_term()
			}

		case 7:
			{
				p.SetState(68)
				p.String_term()
			}

		case 8:
			{
				p.SetState(69)
				p.False_term()
			}

		case 9:
			{
				p.SetState(70)
				p.True_term()
			}

		case 10:
			{
				p.SetState(71)
				p.Name_term()
			}

		case 11:
			{
				p.SetState(72)
				p.Sys_name_term()
			}

		case 12:
			{
				p.SetState(73)
				p.Duplicate_term()
			}

		case 13:
			{
				p.SetState(74)
				p.Drop_term()
			}

		case 14:
			{
				p.SetState(75)
				p.Begin()
			}

		case 15:
			{
				p.SetState(76)
				p.End()
			}

		case 16:
			{
				p.SetState(77)
				p.Pair()
			}

		case 17:
			{
				p.SetState(78)
				p.Block()
			}

		case 18:
			{
				p.SetState(79)
				p.Directive()
			}

		case 19:
			{
				p.SetState(80)
				p.Sys_directive()
			}

		case 20:
			{
				p.SetState(81)
				p.Cmd()
			}

		case 21:
			{
				p.SetState(82)
				p.Sys_cmd()
			}

		case 22:
			{
				p.SetState(83)
				p.Exec()
			}

		case 23:
			{
				p.SetState(84)
				p.Channel_in()
			}

		case 24:
			{
				p.SetState(85)
				p.Channel_out()
			}

		case 25:
			{
				p.SetState(86)
				p.Sys_channel_in()
			}

		case 26:
			{
				p.SetState(87)
				p.Sys_channel_out()
			}

		case 27:
			{
				p.SetState(88)
				p.Lambda()
			}

		case 28:
			{
				p.SetState(89)
				p.Lambda_cmd()
			}

		}

		p.SetState(94)
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

func (s *TermContext) Channel_in() IChannel_inContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChannel_inContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChannel_inContext)
}

func (s *TermContext) Channel_out() IChannel_outContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChannel_outContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChannel_outContext)
}

func (s *TermContext) Sys_channel_in() ISys_channel_inContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISys_channel_inContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISys_channel_inContext)
}

func (s *TermContext) Sys_channel_out() ISys_channel_outContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISys_channel_outContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISys_channel_outContext)
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

func (s *TermContext) Sys_cmd() ISys_cmdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISys_cmdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISys_cmdContext)
}

func (s *TermContext) Exec() IExecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExecContext)
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
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(95)
			p.Nth()
		}

	case 2:
		{
			p.SetState(96)
			p.Integer()
		}

	case 3:
		{
			p.SetState(97)
			p.Oct_integer()
		}

	case 4:
		{
			p.SetState(98)
			p.Hex_integer()
		}

	case 5:
		{
			p.SetState(99)
			p.Bin_integer()
		}

	case 6:
		{
			p.SetState(100)
			p.Float_term()
		}

	case 7:
		{
			p.SetState(101)
			p.String_term()
		}

	case 8:
		{
			p.SetState(102)
			p.False_term()
		}

	case 9:
		{
			p.SetState(103)
			p.True_term()
		}

	case 10:
		{
			p.SetState(104)
			p.Name_term()
		}

	case 11:
		{
			p.SetState(105)
			p.Duplicate_term()
		}

	case 12:
		{
			p.SetState(106)
			p.Drop_term()
		}

	case 13:
		{
			p.SetState(107)
			p.Channel_in()
		}

	case 14:
		{
			p.SetState(108)
			p.Channel_out()
		}

	case 15:
		{
			p.SetState(109)
			p.Sys_channel_in()
		}

	case 16:
		{
			p.SetState(110)
			p.Sys_channel_out()
		}

	case 17:
		{
			p.SetState(111)
			p.Begin()
		}

	case 18:
		{
			p.SetState(112)
			p.End()
		}

	case 19:
		{
			p.SetState(113)
			p.Pair()
		}

	case 20:
		{
			p.SetState(114)
			p.Block()
		}

	case 21:
		{
			p.SetState(115)
			p.Directive()
		}

	case 22:
		{
			p.SetState(116)
			p.Cmd()
		}

	case 23:
		{
			p.SetState(117)
			p.Sys_cmd()
		}

	case 24:
		{
			p.SetState(118)
			p.Exec()
		}

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
	p.SetState(133)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserINTEGER:
		{
			p.SetState(121)
			p.Integer()
		}

	case BundParserNTH:
		{
			p.SetState(122)
			p.Nth()
		}

	case BundParserOCT_INTEGER:
		{
			p.SetState(123)
			p.Oct_integer()
		}

	case BundParserHEX_INTEGER:
		{
			p.SetState(124)
			p.Hex_integer()
		}

	case BundParserBIN_INTEGER:
		{
			p.SetState(125)
			p.Bin_integer()
		}

	case BundParserFLOAT_NUMBER:
		{
			p.SetState(126)
			p.Float_term()
		}

	case BundParserSTRING:
		{
			p.SetState(127)
			p.String_term()
		}

	case BundParserFALSE:
		{
			p.SetState(128)
			p.False_term()
		}

	case BundParserTRUE:
		{
			p.SetState(129)
			p.True_term()
		}

	case BundParserNAME:
		{
			p.SetState(130)
			p.Name_term()
		}

	case BundParserDIR:
		{
			p.SetState(131)
			p.Directive()
		}

	case BundParserCMD:
		{
			p.SetState(132)
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
		p.SetState(135)

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
		p.SetState(137)

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
		p.SetState(139)

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
		p.SetState(141)

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
		p.SetState(143)

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
		p.SetState(145)

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
		p.SetState(147)

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
		p.SetState(149)

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
		p.SetState(151)

		var _m = p.Match(BundParserNTH)

		localctx.(*NthContext).value = _m
	}

	return localctx
}

// IExecContext is an interface to support dynamic dispatch.
type IExecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsExecContext differentiates from other interfaces.
	IsExecContext()
}

type ExecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyExecContext() *ExecContext {
	var p = new(ExecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_exec
	return p
}

func (*ExecContext) IsExecContext() {}

func NewExecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecContext {
	var p = new(ExecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_exec

	return p
}

func (s *ExecContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecContext) GetValue() antlr.Token { return s.value }

func (s *ExecContext) SetValue(v antlr.Token) { s.value = v }

func (s *ExecContext) EXECUTE() antlr.TerminalNode {
	return s.GetToken(BundParserEXECUTE, 0)
}

func (s *ExecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterExec(s)
	}
}

func (s *ExecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitExec(s)
	}
}

func (p *BundParser) Exec() (localctx IExecContext) {
	localctx = NewExecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BundParserRULE_exec)

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
		p.SetState(153)

		var _m = p.Match(BundParserEXECUTE)

		localctx.(*ExecContext).value = _m
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
	p.EnterRule(localctx, 26, BundParserRULE_name_term)

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
		p.SetState(155)

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
	p.EnterRule(localctx, 28, BundParserRULE_sys_name_term)

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
		p.SetState(157)

		var _m = p.Match(BundParserSYS)

		localctx.(*Sys_name_termContext).sys = _m
	}
	{
		p.SetState(158)

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
	p.EnterRule(localctx, 30, BundParserRULE_duplicate_term)

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
		p.SetState(160)

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
	p.EnterRule(localctx, 32, BundParserRULE_drop_term)

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
		p.SetState(162)

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
	p.EnterRule(localctx, 34, BundParserRULE_begin)

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
		p.SetState(164)

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
	p.EnterRule(localctx, 36, BundParserRULE_end)

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
		p.SetState(166)

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
	p.EnterRule(localctx, 38, BundParserRULE_block)

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
		p.SetState(168)
		p.Match(BundParserT__0)
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(169)

				var _x = p.Term()

				localctx.(*BlockContext)._term = _x
			}
			localctx.(*BlockContext).values = append(localctx.(*BlockContext).values, localctx.(*BlockContext)._term)

		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	{
		p.SetState(175)
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
	p.EnterRule(localctx, 40, BundParserRULE_pair)

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
		p.SetState(177)
		p.Match(BundParserT__1)
	}
	{
		p.SetState(178)

		var _x = p.Name_term()

		localctx.(*PairContext).head = _x
	}
	{
		p.SetState(179)

		var _x = p.Pvalue()

		localctx.(*PairContext).tail = _x
	}
	{
		p.SetState(180)
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
	p.EnterRule(localctx, 42, BundParserRULE_directive)

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
		p.SetState(182)

		var _m = p.Match(BundParserDIR)

		localctx.(*DirectiveContext).op = _m
	}
	{
		p.SetState(183)

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
	p.EnterRule(localctx, 44, BundParserRULE_sys_directive)

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
		p.SetState(185)

		var _m = p.Match(BundParserSYS)

		localctx.(*Sys_directiveContext).sys = _m
	}
	{
		p.SetState(186)

		var _m = p.Match(BundParserDIR)

		localctx.(*Sys_directiveContext).op = _m
	}
	{
		p.SetState(187)

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
	p.EnterRule(localctx, 46, BundParserRULE_cmd)

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
		p.SetState(189)

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
	p.EnterRule(localctx, 48, BundParserRULE_sys_cmd)

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
		p.SetState(191)

		var _m = p.Match(BundParserSYS)

		localctx.(*Sys_cmdContext).sys = _m
	}
	{
		p.SetState(192)

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
	p.EnterRule(localctx, 50, BundParserRULE_lambda)
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
		p.SetState(194)
		p.Match(BundParserT__3)
	}
	{
		p.SetState(195)

		var _m = p.Match(BundParserNAME)

		localctx.(*LambdaContext).name = _m
	}
	{
		p.SetState(196)
		p.Match(BundParserT__4)
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__1)|(1<<BundParserT__3)|(1<<BundParserT__9)|(1<<BundParserINTEGER)|(1<<BundParserOCT_INTEGER)|(1<<BundParserHEX_INTEGER)|(1<<BundParserBIN_INTEGER)|(1<<BundParserNTH)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserNAME)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND)|(1<<BundParserSYS)|(1<<BundParserDIR)|(1<<BundParserCMD)|(1<<BundParserDUPLICATE)|(1<<BundParserDROP)|(1<<BundParserEXECUTE))) != 0 {
		{
			p.SetState(197)

			var _x = p.Term()

			localctx.(*LambdaContext)._term = _x
		}
		localctx.(*LambdaContext).body = append(localctx.(*LambdaContext).body, localctx.(*LambdaContext)._term)

		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(203)
		p.Match(BundParserT__5)
	}

	return localctx
}

// ILambda_cmdContext is an interface to support dynamic dispatch.
type ILambda_cmdContext interface {
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

	// IsLambda_cmdContext differentiates from other interfaces.
	IsLambda_cmdContext()
}

type Lambda_cmdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyLambda_cmdContext() *Lambda_cmdContext {
	var p = new(Lambda_cmdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_lambda_cmd
	return p
}

func (*Lambda_cmdContext) IsLambda_cmdContext() {}

func NewLambda_cmdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lambda_cmdContext {
	var p = new(Lambda_cmdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_lambda_cmd

	return p
}

func (s *Lambda_cmdContext) GetParser() antlr.Parser { return s.parser }

func (s *Lambda_cmdContext) GetName() antlr.Token { return s.name }

func (s *Lambda_cmdContext) SetName(v antlr.Token) { s.name = v }

func (s *Lambda_cmdContext) Get_term() ITermContext { return s._term }

func (s *Lambda_cmdContext) Set_term(v ITermContext) { s._term = v }

func (s *Lambda_cmdContext) GetBody() []ITermContext { return s.body }

func (s *Lambda_cmdContext) SetBody(v []ITermContext) { s.body = v }

func (s *Lambda_cmdContext) CMD() antlr.TerminalNode {
	return s.GetToken(BundParserCMD, 0)
}

func (s *Lambda_cmdContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *Lambda_cmdContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Lambda_cmdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lambda_cmdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lambda_cmdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterLambda_cmd(s)
	}
}

func (s *Lambda_cmdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitLambda_cmd(s)
	}
}

func (p *BundParser) Lambda_cmd() (localctx ILambda_cmdContext) {
	localctx = NewLambda_cmdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BundParserRULE_lambda_cmd)
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
		p.SetState(205)
		p.Match(BundParserT__6)
	}
	{
		p.SetState(206)

		var _m = p.Match(BundParserCMD)

		localctx.(*Lambda_cmdContext).name = _m
	}
	{
		p.SetState(207)
		p.Match(BundParserT__7)
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__1)|(1<<BundParserT__3)|(1<<BundParserT__9)|(1<<BundParserINTEGER)|(1<<BundParserOCT_INTEGER)|(1<<BundParserHEX_INTEGER)|(1<<BundParserBIN_INTEGER)|(1<<BundParserNTH)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserNAME)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND)|(1<<BundParserSYS)|(1<<BundParserDIR)|(1<<BundParserCMD)|(1<<BundParserDUPLICATE)|(1<<BundParserDROP)|(1<<BundParserEXECUTE))) != 0 {
		{
			p.SetState(208)

			var _x = p.Term()

			localctx.(*Lambda_cmdContext)._term = _x
		}
		localctx.(*Lambda_cmdContext).body = append(localctx.(*Lambda_cmdContext).body, localctx.(*Lambda_cmdContext)._term)

		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(214)
		p.Match(BundParserT__5)
	}

	return localctx
}

// IChannel_outContext is an interface to support dynamic dispatch.
type IChannel_outContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsChannel_outContext differentiates from other interfaces.
	IsChannel_outContext()
}

type Channel_outContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyChannel_outContext() *Channel_outContext {
	var p = new(Channel_outContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_channel_out
	return p
}

func (*Channel_outContext) IsChannel_outContext() {}

func NewChannel_outContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Channel_outContext {
	var p = new(Channel_outContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_channel_out

	return p
}

func (s *Channel_outContext) GetParser() antlr.Parser { return s.parser }

func (s *Channel_outContext) GetName() antlr.Token { return s.name }

func (s *Channel_outContext) SetName(v antlr.Token) { s.name = v }

func (s *Channel_outContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Channel_outContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Channel_outContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Channel_outContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterChannel_out(s)
	}
}

func (s *Channel_outContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitChannel_out(s)
	}
}

func (p *BundParser) Channel_out() (localctx IChannel_outContext) {
	localctx = NewChannel_outContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BundParserRULE_channel_out)

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
		p.SetState(216)
		p.Match(BundParserT__3)
	}
	{
		p.SetState(217)

		var _m = p.Match(BundParserNAME)

		localctx.(*Channel_outContext).name = _m
	}
	{
		p.SetState(218)
		p.Match(BundParserT__8)
	}

	return localctx
}

// IChannel_inContext is an interface to support dynamic dispatch.
type IChannel_inContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsChannel_inContext differentiates from other interfaces.
	IsChannel_inContext()
}

type Channel_inContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyChannel_inContext() *Channel_inContext {
	var p = new(Channel_inContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_channel_in
	return p
}

func (*Channel_inContext) IsChannel_inContext() {}

func NewChannel_inContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Channel_inContext {
	var p = new(Channel_inContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_channel_in

	return p
}

func (s *Channel_inContext) GetParser() antlr.Parser { return s.parser }

func (s *Channel_inContext) GetName() antlr.Token { return s.name }

func (s *Channel_inContext) SetName(v antlr.Token) { s.name = v }

func (s *Channel_inContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Channel_inContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Channel_inContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Channel_inContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterChannel_in(s)
	}
}

func (s *Channel_inContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitChannel_in(s)
	}
}

func (p *BundParser) Channel_in() (localctx IChannel_inContext) {
	localctx = NewChannel_inContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BundParserRULE_channel_in)

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
		p.SetState(220)
		p.Match(BundParserT__9)
	}
	{
		p.SetState(221)

		var _m = p.Match(BundParserNAME)

		localctx.(*Channel_inContext).name = _m
	}
	{
		p.SetState(222)
		p.Match(BundParserT__4)
	}

	return localctx
}

// ISys_channel_outContext is an interface to support dynamic dispatch.
type ISys_channel_outContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSys returns the sys token.
	GetSys() antlr.Token

	// GetName returns the name token.
	GetName() antlr.Token

	// SetSys sets the sys token.
	SetSys(antlr.Token)

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsSys_channel_outContext differentiates from other interfaces.
	IsSys_channel_outContext()
}

type Sys_channel_outContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	sys    antlr.Token
	name   antlr.Token
}

func NewEmptySys_channel_outContext() *Sys_channel_outContext {
	var p = new(Sys_channel_outContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_sys_channel_out
	return p
}

func (*Sys_channel_outContext) IsSys_channel_outContext() {}

func NewSys_channel_outContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sys_channel_outContext {
	var p = new(Sys_channel_outContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_sys_channel_out

	return p
}

func (s *Sys_channel_outContext) GetParser() antlr.Parser { return s.parser }

func (s *Sys_channel_outContext) GetSys() antlr.Token { return s.sys }

func (s *Sys_channel_outContext) GetName() antlr.Token { return s.name }

func (s *Sys_channel_outContext) SetSys(v antlr.Token) { s.sys = v }

func (s *Sys_channel_outContext) SetName(v antlr.Token) { s.name = v }

func (s *Sys_channel_outContext) SYS() antlr.TerminalNode {
	return s.GetToken(BundParserSYS, 0)
}

func (s *Sys_channel_outContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Sys_channel_outContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sys_channel_outContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sys_channel_outContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterSys_channel_out(s)
	}
}

func (s *Sys_channel_outContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitSys_channel_out(s)
	}
}

func (p *BundParser) Sys_channel_out() (localctx ISys_channel_outContext) {
	localctx = NewSys_channel_outContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BundParserRULE_sys_channel_out)

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
		p.SetState(224)

		var _m = p.Match(BundParserSYS)

		localctx.(*Sys_channel_outContext).sys = _m
	}
	{
		p.SetState(225)
		p.Match(BundParserT__3)
	}
	{
		p.SetState(226)

		var _m = p.Match(BundParserNAME)

		localctx.(*Sys_channel_outContext).name = _m
	}
	{
		p.SetState(227)
		p.Match(BundParserT__8)
	}

	return localctx
}

// ISys_channel_inContext is an interface to support dynamic dispatch.
type ISys_channel_inContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetSys returns the sys token.
	GetSys() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetSys sets the sys token.
	SetSys(antlr.Token)

	// IsSys_channel_inContext differentiates from other interfaces.
	IsSys_channel_inContext()
}

type Sys_channel_inContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	sys    antlr.Token
}

func NewEmptySys_channel_inContext() *Sys_channel_inContext {
	var p = new(Sys_channel_inContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_sys_channel_in
	return p
}

func (*Sys_channel_inContext) IsSys_channel_inContext() {}

func NewSys_channel_inContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sys_channel_inContext {
	var p = new(Sys_channel_inContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_sys_channel_in

	return p
}

func (s *Sys_channel_inContext) GetParser() antlr.Parser { return s.parser }

func (s *Sys_channel_inContext) GetName() antlr.Token { return s.name }

func (s *Sys_channel_inContext) GetSys() antlr.Token { return s.sys }

func (s *Sys_channel_inContext) SetName(v antlr.Token) { s.name = v }

func (s *Sys_channel_inContext) SetSys(v antlr.Token) { s.sys = v }

func (s *Sys_channel_inContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Sys_channel_inContext) SYS() antlr.TerminalNode {
	return s.GetToken(BundParserSYS, 0)
}

func (s *Sys_channel_inContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sys_channel_inContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sys_channel_inContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterSys_channel_in(s)
	}
}

func (s *Sys_channel_inContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitSys_channel_in(s)
	}
}

func (p *BundParser) Sys_channel_in() (localctx ISys_channel_inContext) {
	localctx = NewSys_channel_inContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, BundParserRULE_sys_channel_in)

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
		p.SetState(229)
		p.Match(BundParserT__9)
	}
	{
		p.SetState(230)

		var _m = p.Match(BundParserNAME)

		localctx.(*Sys_channel_inContext).name = _m
	}
	{
		p.SetState(231)
		p.Match(BundParserT__4)
	}
	{
		p.SetState(232)

		var _m = p.Match(BundParserSYS)

		localctx.(*Sys_channel_inContext).sys = _m
	}

	return localctx
}

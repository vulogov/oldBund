// Code generated from Bund.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Bund

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBundListener is a complete listener for a parse tree produced by BundParser.
type BaseBundListener struct{}

var _ BundListener = &BaseBundListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBundListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBundListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBundListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBundListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BaseBundListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BaseBundListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseBundListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseBundListener) ExitTerm(ctx *TermContext) {}

// EnterPvalue is called when production pvalue is entered.
func (s *BaseBundListener) EnterPvalue(ctx *PvalueContext) {}

// ExitPvalue is called when production pvalue is exited.
func (s *BaseBundListener) ExitPvalue(ctx *PvalueContext) {}

// EnterTrue_term is called when production true_term is entered.
func (s *BaseBundListener) EnterTrue_term(ctx *True_termContext) {}

// ExitTrue_term is called when production true_term is exited.
func (s *BaseBundListener) ExitTrue_term(ctx *True_termContext) {}

// EnterFalse_term is called when production false_term is entered.
func (s *BaseBundListener) EnterFalse_term(ctx *False_termContext) {}

// ExitFalse_term is called when production false_term is exited.
func (s *BaseBundListener) ExitFalse_term(ctx *False_termContext) {}

// EnterString_term is called when production string_term is entered.
func (s *BaseBundListener) EnterString_term(ctx *String_termContext) {}

// ExitString_term is called when production string_term is exited.
func (s *BaseBundListener) ExitString_term(ctx *String_termContext) {}

// EnterFloat_term is called when production float_term is entered.
func (s *BaseBundListener) EnterFloat_term(ctx *Float_termContext) {}

// ExitFloat_term is called when production float_term is exited.
func (s *BaseBundListener) ExitFloat_term(ctx *Float_termContext) {}

// EnterBin_integer is called when production bin_integer is entered.
func (s *BaseBundListener) EnterBin_integer(ctx *Bin_integerContext) {}

// ExitBin_integer is called when production bin_integer is exited.
func (s *BaseBundListener) ExitBin_integer(ctx *Bin_integerContext) {}

// EnterHex_integer is called when production hex_integer is entered.
func (s *BaseBundListener) EnterHex_integer(ctx *Hex_integerContext) {}

// ExitHex_integer is called when production hex_integer is exited.
func (s *BaseBundListener) ExitHex_integer(ctx *Hex_integerContext) {}

// EnterOct_integer is called when production oct_integer is entered.
func (s *BaseBundListener) EnterOct_integer(ctx *Oct_integerContext) {}

// ExitOct_integer is called when production oct_integer is exited.
func (s *BaseBundListener) ExitOct_integer(ctx *Oct_integerContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseBundListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseBundListener) ExitInteger(ctx *IntegerContext) {}

// EnterNth is called when production nth is entered.
func (s *BaseBundListener) EnterNth(ctx *NthContext) {}

// ExitNth is called when production nth is exited.
func (s *BaseBundListener) ExitNth(ctx *NthContext) {}

// EnterName_term is called when production name_term is entered.
func (s *BaseBundListener) EnterName_term(ctx *Name_termContext) {}

// ExitName_term is called when production name_term is exited.
func (s *BaseBundListener) ExitName_term(ctx *Name_termContext) {}

// EnterDuplicate_term is called when production duplicate_term is entered.
func (s *BaseBundListener) EnterDuplicate_term(ctx *Duplicate_termContext) {}

// ExitDuplicate_term is called when production duplicate_term is exited.
func (s *BaseBundListener) ExitDuplicate_term(ctx *Duplicate_termContext) {}

// EnterDrop_term is called when production drop_term is entered.
func (s *BaseBundListener) EnterDrop_term(ctx *Drop_termContext) {}

// ExitDrop_term is called when production drop_term is exited.
func (s *BaseBundListener) ExitDrop_term(ctx *Drop_termContext) {}

// EnterBegin is called when production begin is entered.
func (s *BaseBundListener) EnterBegin(ctx *BeginContext) {}

// ExitBegin is called when production begin is exited.
func (s *BaseBundListener) ExitBegin(ctx *BeginContext) {}

// EnterEnd is called when production end is entered.
func (s *BaseBundListener) EnterEnd(ctx *EndContext) {}

// ExitEnd is called when production end is exited.
func (s *BaseBundListener) ExitEnd(ctx *EndContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseBundListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseBundListener) ExitBlock(ctx *BlockContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseBundListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseBundListener) ExitPair(ctx *PairContext) {}

// EnterDirective is called when production directive is entered.
func (s *BaseBundListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BaseBundListener) ExitDirective(ctx *DirectiveContext) {}

// EnterCmd is called when production cmd is entered.
func (s *BaseBundListener) EnterCmd(ctx *CmdContext) {}

// ExitCmd is called when production cmd is exited.
func (s *BaseBundListener) ExitCmd(ctx *CmdContext) {}

// EnterLambda is called when production lambda is entered.
func (s *BaseBundListener) EnterLambda(ctx *LambdaContext) {}

// ExitLambda is called when production lambda is exited.
func (s *BaseBundListener) ExitLambda(ctx *LambdaContext) {}

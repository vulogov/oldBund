// Code generated from Bund.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Bund

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BundListener is a complete listener for a parse tree produced by BundParser.
type BundListener interface {
	antlr.ParseTreeListener

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterPvalue is called when entering the pvalue production.
	EnterPvalue(c *PvalueContext)

	// EnterTrue_term is called when entering the true_term production.
	EnterTrue_term(c *True_termContext)

	// EnterFalse_term is called when entering the false_term production.
	EnterFalse_term(c *False_termContext)

	// EnterString_term is called when entering the string_term production.
	EnterString_term(c *String_termContext)

	// EnterFloat_term is called when entering the float_term production.
	EnterFloat_term(c *Float_termContext)

	// EnterBin_integer is called when entering the bin_integer production.
	EnterBin_integer(c *Bin_integerContext)

	// EnterHex_integer is called when entering the hex_integer production.
	EnterHex_integer(c *Hex_integerContext)

	// EnterOct_integer is called when entering the oct_integer production.
	EnterOct_integer(c *Oct_integerContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterNth is called when entering the nth production.
	EnterNth(c *NthContext)

	// EnterExec is called when entering the exec production.
	EnterExec(c *ExecContext)

	// EnterName_term is called when entering the name_term production.
	EnterName_term(c *Name_termContext)

	// EnterSys_name_term is called when entering the sys_name_term production.
	EnterSys_name_term(c *Sys_name_termContext)

	// EnterDuplicate_term is called when entering the duplicate_term production.
	EnterDuplicate_term(c *Duplicate_termContext)

	// EnterDrop_term is called when entering the drop_term production.
	EnterDrop_term(c *Drop_termContext)

	// EnterBegin is called when entering the begin production.
	EnterBegin(c *BeginContext)

	// EnterEnd is called when entering the end production.
	EnterEnd(c *EndContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterDirective is called when entering the directive production.
	EnterDirective(c *DirectiveContext)

	// EnterSys_directive is called when entering the sys_directive production.
	EnterSys_directive(c *Sys_directiveContext)

	// EnterCmd is called when entering the cmd production.
	EnterCmd(c *CmdContext)

	// EnterSys_cmd is called when entering the sys_cmd production.
	EnterSys_cmd(c *Sys_cmdContext)

	// EnterLambda is called when entering the lambda production.
	EnterLambda(c *LambdaContext)

	// EnterChannel_out is called when entering the channel_out production.
	EnterChannel_out(c *Channel_outContext)

	// EnterChannel_in is called when entering the channel_in production.
	EnterChannel_in(c *Channel_inContext)

	// EnterSys_channel_out is called when entering the sys_channel_out production.
	EnterSys_channel_out(c *Sys_channel_outContext)

	// EnterSys_channel_in is called when entering the sys_channel_in production.
	EnterSys_channel_in(c *Sys_channel_inContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitPvalue is called when exiting the pvalue production.
	ExitPvalue(c *PvalueContext)

	// ExitTrue_term is called when exiting the true_term production.
	ExitTrue_term(c *True_termContext)

	// ExitFalse_term is called when exiting the false_term production.
	ExitFalse_term(c *False_termContext)

	// ExitString_term is called when exiting the string_term production.
	ExitString_term(c *String_termContext)

	// ExitFloat_term is called when exiting the float_term production.
	ExitFloat_term(c *Float_termContext)

	// ExitBin_integer is called when exiting the bin_integer production.
	ExitBin_integer(c *Bin_integerContext)

	// ExitHex_integer is called when exiting the hex_integer production.
	ExitHex_integer(c *Hex_integerContext)

	// ExitOct_integer is called when exiting the oct_integer production.
	ExitOct_integer(c *Oct_integerContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitNth is called when exiting the nth production.
	ExitNth(c *NthContext)

	// ExitExec is called when exiting the exec production.
	ExitExec(c *ExecContext)

	// ExitName_term is called when exiting the name_term production.
	ExitName_term(c *Name_termContext)

	// ExitSys_name_term is called when exiting the sys_name_term production.
	ExitSys_name_term(c *Sys_name_termContext)

	// ExitDuplicate_term is called when exiting the duplicate_term production.
	ExitDuplicate_term(c *Duplicate_termContext)

	// ExitDrop_term is called when exiting the drop_term production.
	ExitDrop_term(c *Drop_termContext)

	// ExitBegin is called when exiting the begin production.
	ExitBegin(c *BeginContext)

	// ExitEnd is called when exiting the end production.
	ExitEnd(c *EndContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitDirective is called when exiting the directive production.
	ExitDirective(c *DirectiveContext)

	// ExitSys_directive is called when exiting the sys_directive production.
	ExitSys_directive(c *Sys_directiveContext)

	// ExitCmd is called when exiting the cmd production.
	ExitCmd(c *CmdContext)

	// ExitSys_cmd is called when exiting the sys_cmd production.
	ExitSys_cmd(c *Sys_cmdContext)

	// ExitLambda is called when exiting the lambda production.
	ExitLambda(c *LambdaContext)

	// ExitChannel_out is called when exiting the channel_out production.
	ExitChannel_out(c *Channel_outContext)

	// ExitChannel_in is called when exiting the channel_in production.
	ExitChannel_in(c *Channel_inContext)

	// ExitSys_channel_out is called when exiting the sys_channel_out production.
	ExitSys_channel_out(c *Sys_channel_outContext)

	// ExitSys_channel_in is called when exiting the sys_channel_in production.
	ExitSys_channel_in(c *Sys_channel_inContext)
}

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

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)
}

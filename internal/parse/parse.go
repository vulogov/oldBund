package parse

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/vulogov/Bund/internal/parser"
)

type bundListener struct {
	*parser.BaseBundListener
}

func Parse(code string) {
	_input := antlr.NewInputStream(code)
	lexer := parser.NewBundLexer(_input)
	// stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// p := parser.NewBundParser(stream)
	// antlr.ParseTreeWalkerDefault.Walk(&bundListener{}, p.Expressions())
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}

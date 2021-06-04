// Code generated from plural.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // plural

import "github.com/antlr/antlr4/runtime/Go/antlr"

// pluralListener is a complete listener for a parse tree produced by pluralParser.
type pluralListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)
}

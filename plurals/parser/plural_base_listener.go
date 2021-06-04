// Code generated from plural.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // plural

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasepluralListener is a complete listener for a parse tree produced by pluralParser.
type BasepluralListener struct{}

var _ pluralListener = &BasepluralListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasepluralListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasepluralListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasepluralListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasepluralListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BasepluralListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BasepluralListener) ExitStart(ctx *StartContext) {}

// EnterExp is called when production exp is entered.
func (s *BasepluralListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BasepluralListener) ExitExp(ctx *ExpContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BasepluralListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BasepluralListener) ExitPrimary(ctx *PrimaryContext) {}

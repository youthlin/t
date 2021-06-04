// Code generated from plural.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 31, 137,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	27, 3, 27, 3, 28, 3, 28, 3, 29, 6, 29, 127, 10, 29, 13, 29, 14, 29, 128,
	3, 30, 6, 30, 132, 10, 30, 13, 30, 14, 30, 133, 3, 30, 3, 30, 2, 2, 31,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41,
	22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59,
	31, 3, 2, 4, 3, 2, 50, 59, 4, 2, 11, 11, 34, 34, 2, 138, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2,
	2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3,
	2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35,
	3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2,
	43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2,
	2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2,
	2, 2, 59, 3, 2, 2, 2, 3, 61, 3, 2, 2, 2, 5, 64, 3, 2, 2, 2, 7, 67, 3, 2,
	2, 2, 9, 69, 3, 2, 2, 2, 11, 71, 3, 2, 2, 2, 13, 73, 3, 2, 2, 2, 15, 75,
	3, 2, 2, 2, 17, 77, 3, 2, 2, 2, 19, 79, 3, 2, 2, 2, 21, 81, 3, 2, 2, 2,
	23, 84, 3, 2, 2, 2, 25, 87, 3, 2, 2, 2, 27, 89, 3, 2, 2, 2, 29, 91, 3,
	2, 2, 2, 31, 94, 3, 2, 2, 2, 33, 97, 3, 2, 2, 2, 35, 100, 3, 2, 2, 2, 37,
	103, 3, 2, 2, 2, 39, 105, 3, 2, 2, 2, 41, 107, 3, 2, 2, 2, 43, 109, 3,
	2, 2, 2, 45, 112, 3, 2, 2, 2, 47, 115, 3, 2, 2, 2, 49, 117, 3, 2, 2, 2,
	51, 119, 3, 2, 2, 2, 53, 121, 3, 2, 2, 2, 55, 123, 3, 2, 2, 2, 57, 126,
	3, 2, 2, 2, 59, 131, 3, 2, 2, 2, 61, 62, 7, 45, 2, 2, 62, 63, 7, 45, 2,
	2, 63, 4, 3, 2, 2, 2, 64, 65, 7, 47, 2, 2, 65, 66, 7, 47, 2, 2, 66, 6,
	3, 2, 2, 2, 67, 68, 7, 45, 2, 2, 68, 8, 3, 2, 2, 2, 69, 70, 7, 47, 2, 2,
	70, 10, 3, 2, 2, 2, 71, 72, 7, 128, 2, 2, 72, 12, 3, 2, 2, 2, 73, 74, 7,
	35, 2, 2, 74, 14, 3, 2, 2, 2, 75, 76, 7, 44, 2, 2, 76, 16, 3, 2, 2, 2,
	77, 78, 7, 49, 2, 2, 78, 18, 3, 2, 2, 2, 79, 80, 7, 39, 2, 2, 80, 20, 3,
	2, 2, 2, 81, 82, 7, 64, 2, 2, 82, 83, 7, 64, 2, 2, 83, 22, 3, 2, 2, 2,
	84, 85, 7, 62, 2, 2, 85, 86, 7, 62, 2, 2, 86, 24, 3, 2, 2, 2, 87, 88, 7,
	64, 2, 2, 88, 26, 3, 2, 2, 2, 89, 90, 7, 62, 2, 2, 90, 28, 3, 2, 2, 2,
	91, 92, 7, 64, 2, 2, 92, 93, 7, 63, 2, 2, 93, 30, 3, 2, 2, 2, 94, 95, 7,
	62, 2, 2, 95, 96, 7, 63, 2, 2, 96, 32, 3, 2, 2, 2, 97, 98, 7, 63, 2, 2,
	98, 99, 7, 63, 2, 2, 99, 34, 3, 2, 2, 2, 100, 101, 7, 35, 2, 2, 101, 102,
	7, 63, 2, 2, 102, 36, 3, 2, 2, 2, 103, 104, 7, 40, 2, 2, 104, 38, 3, 2,
	2, 2, 105, 106, 7, 96, 2, 2, 106, 40, 3, 2, 2, 2, 107, 108, 7, 126, 2,
	2, 108, 42, 3, 2, 2, 2, 109, 110, 7, 40, 2, 2, 110, 111, 7, 40, 2, 2, 111,
	44, 3, 2, 2, 2, 112, 113, 7, 126, 2, 2, 113, 114, 7, 126, 2, 2, 114, 46,
	3, 2, 2, 2, 115, 116, 7, 65, 2, 2, 116, 48, 3, 2, 2, 2, 117, 118, 7, 60,
	2, 2, 118, 50, 3, 2, 2, 2, 119, 120, 7, 42, 2, 2, 120, 52, 3, 2, 2, 2,
	121, 122, 7, 43, 2, 2, 122, 54, 3, 2, 2, 2, 123, 124, 7, 112, 2, 2, 124,
	56, 3, 2, 2, 2, 125, 127, 9, 2, 2, 2, 126, 125, 3, 2, 2, 2, 127, 128, 3,
	2, 2, 2, 128, 126, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 58, 3, 2, 2,
	2, 130, 132, 9, 3, 2, 2, 131, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133,
	131, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 136,
	8, 30, 2, 2, 136, 60, 3, 2, 2, 2, 5, 2, 128, 133, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'++'", "'--'", "'+'", "'-'", "'~'", "'!'", "'*'", "'/'", "'%'", "'>>'",
	"'<<'", "'>'", "'<'", "'>='", "'<='", "'=='", "'!='", "'&'", "'^'", "'|'",
	"'&&'", "'||'", "'?'", "':'", "'('", "')'", "'n'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "INT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "INT", "WS",
}

type pluralLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewpluralLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *pluralLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewpluralLexer(input antlr.CharStream) *pluralLexer {
	l := new(pluralLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "plural.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// pluralLexer tokens.
const (
	pluralLexerT__0  = 1
	pluralLexerT__1  = 2
	pluralLexerT__2  = 3
	pluralLexerT__3  = 4
	pluralLexerT__4  = 5
	pluralLexerT__5  = 6
	pluralLexerT__6  = 7
	pluralLexerT__7  = 8
	pluralLexerT__8  = 9
	pluralLexerT__9  = 10
	pluralLexerT__10 = 11
	pluralLexerT__11 = 12
	pluralLexerT__12 = 13
	pluralLexerT__13 = 14
	pluralLexerT__14 = 15
	pluralLexerT__15 = 16
	pluralLexerT__16 = 17
	pluralLexerT__17 = 18
	pluralLexerT__18 = 19
	pluralLexerT__19 = 20
	pluralLexerT__20 = 21
	pluralLexerT__21 = 22
	pluralLexerT__22 = 23
	pluralLexerT__23 = 24
	pluralLexerT__24 = 25
	pluralLexerT__25 = 26
	pluralLexerT__26 = 27
	pluralLexerINT   = 28
	pluralLexerWS    = 29
)

// Code generated from plural.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type pluralLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PluralLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func plurallexerLexerInit() {
	staticData := &PluralLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'++'", "'--'", "'+'", "'-'", "'~'", "'!'", "'*'", "'/'", "'%'",
		"'>>'", "'<<'", "'>'", "'<'", "'>='", "'<='", "'=='", "'!='", "'&'",
		"'^'", "'|'", "'&&'", "'||'", "'?'", "':'", "'('", "')'", "'n'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "INT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "INT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 29, 135, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 4, 27, 125, 8, 27, 11, 27, 12, 27,
		126, 1, 28, 4, 28, 130, 8, 28, 11, 28, 12, 28, 131, 1, 28, 1, 28, 0, 0,
		29, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21,
		11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39,
		20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57,
		29, 1, 0, 2, 1, 0, 48, 57, 2, 0, 9, 9, 32, 32, 136, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0,
		11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0,
		0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0,
		0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0,
		0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1,
		0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49,
		1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0,
		57, 1, 0, 0, 0, 1, 59, 1, 0, 0, 0, 3, 62, 1, 0, 0, 0, 5, 65, 1, 0, 0, 0,
		7, 67, 1, 0, 0, 0, 9, 69, 1, 0, 0, 0, 11, 71, 1, 0, 0, 0, 13, 73, 1, 0,
		0, 0, 15, 75, 1, 0, 0, 0, 17, 77, 1, 0, 0, 0, 19, 79, 1, 0, 0, 0, 21, 82,
		1, 0, 0, 0, 23, 85, 1, 0, 0, 0, 25, 87, 1, 0, 0, 0, 27, 89, 1, 0, 0, 0,
		29, 92, 1, 0, 0, 0, 31, 95, 1, 0, 0, 0, 33, 98, 1, 0, 0, 0, 35, 101, 1,
		0, 0, 0, 37, 103, 1, 0, 0, 0, 39, 105, 1, 0, 0, 0, 41, 107, 1, 0, 0, 0,
		43, 110, 1, 0, 0, 0, 45, 113, 1, 0, 0, 0, 47, 115, 1, 0, 0, 0, 49, 117,
		1, 0, 0, 0, 51, 119, 1, 0, 0, 0, 53, 121, 1, 0, 0, 0, 55, 124, 1, 0, 0,
		0, 57, 129, 1, 0, 0, 0, 59, 60, 5, 43, 0, 0, 60, 61, 5, 43, 0, 0, 61, 2,
		1, 0, 0, 0, 62, 63, 5, 45, 0, 0, 63, 64, 5, 45, 0, 0, 64, 4, 1, 0, 0, 0,
		65, 66, 5, 43, 0, 0, 66, 6, 1, 0, 0, 0, 67, 68, 5, 45, 0, 0, 68, 8, 1,
		0, 0, 0, 69, 70, 5, 126, 0, 0, 70, 10, 1, 0, 0, 0, 71, 72, 5, 33, 0, 0,
		72, 12, 1, 0, 0, 0, 73, 74, 5, 42, 0, 0, 74, 14, 1, 0, 0, 0, 75, 76, 5,
		47, 0, 0, 76, 16, 1, 0, 0, 0, 77, 78, 5, 37, 0, 0, 78, 18, 1, 0, 0, 0,
		79, 80, 5, 62, 0, 0, 80, 81, 5, 62, 0, 0, 81, 20, 1, 0, 0, 0, 82, 83, 5,
		60, 0, 0, 83, 84, 5, 60, 0, 0, 84, 22, 1, 0, 0, 0, 85, 86, 5, 62, 0, 0,
		86, 24, 1, 0, 0, 0, 87, 88, 5, 60, 0, 0, 88, 26, 1, 0, 0, 0, 89, 90, 5,
		62, 0, 0, 90, 91, 5, 61, 0, 0, 91, 28, 1, 0, 0, 0, 92, 93, 5, 60, 0, 0,
		93, 94, 5, 61, 0, 0, 94, 30, 1, 0, 0, 0, 95, 96, 5, 61, 0, 0, 96, 97, 5,
		61, 0, 0, 97, 32, 1, 0, 0, 0, 98, 99, 5, 33, 0, 0, 99, 100, 5, 61, 0, 0,
		100, 34, 1, 0, 0, 0, 101, 102, 5, 38, 0, 0, 102, 36, 1, 0, 0, 0, 103, 104,
		5, 94, 0, 0, 104, 38, 1, 0, 0, 0, 105, 106, 5, 124, 0, 0, 106, 40, 1, 0,
		0, 0, 107, 108, 5, 38, 0, 0, 108, 109, 5, 38, 0, 0, 109, 42, 1, 0, 0, 0,
		110, 111, 5, 124, 0, 0, 111, 112, 5, 124, 0, 0, 112, 44, 1, 0, 0, 0, 113,
		114, 5, 63, 0, 0, 114, 46, 1, 0, 0, 0, 115, 116, 5, 58, 0, 0, 116, 48,
		1, 0, 0, 0, 117, 118, 5, 40, 0, 0, 118, 50, 1, 0, 0, 0, 119, 120, 5, 41,
		0, 0, 120, 52, 1, 0, 0, 0, 121, 122, 5, 110, 0, 0, 122, 54, 1, 0, 0, 0,
		123, 125, 7, 0, 0, 0, 124, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126,
		124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 56, 1, 0, 0, 0, 128, 130, 7,
		1, 0, 0, 129, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 129, 1, 0, 0,
		0, 131, 132, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 6, 28, 0, 0, 134,
		58, 1, 0, 0, 0, 3, 0, 126, 131, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// pluralLexerInit initializes any static state used to implement pluralLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewpluralLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PluralLexerInit() {
	staticData := &PluralLexerLexerStaticData
	staticData.once.Do(plurallexerLexerInit)
}

// NewpluralLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewpluralLexer(input antlr.CharStream) *pluralLexer {
	PluralLexerInit()
	l := new(pluralLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PluralLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
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

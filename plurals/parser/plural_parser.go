// Code generated from plural.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // plural

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type pluralParser struct {
	*antlr.BaseParser
}

var PluralParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pluralParserInit() {
	staticData := &PluralParserStaticData
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
		"start", "exp", "primary",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 71, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 15, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		5, 1, 58, 8, 1, 10, 1, 12, 1, 61, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 69, 8, 2, 1, 2, 0, 1, 2, 3, 0, 2, 4, 0, 9, 1, 0, 1, 4, 1, 0, 5,
		6, 1, 0, 7, 9, 1, 0, 3, 4, 1, 0, 10, 11, 1, 0, 12, 13, 1, 0, 14, 15, 1,
		0, 16, 17, 1, 0, 1, 2, 84, 0, 6, 1, 0, 0, 0, 2, 14, 1, 0, 0, 0, 4, 68,
		1, 0, 0, 0, 6, 7, 3, 2, 1, 0, 7, 1, 1, 0, 0, 0, 8, 9, 6, 1, -1, 0, 9, 15,
		3, 4, 2, 0, 10, 11, 7, 0, 0, 0, 11, 15, 3, 2, 1, 14, 12, 13, 7, 1, 0, 0,
		13, 15, 3, 2, 1, 13, 14, 8, 1, 0, 0, 0, 14, 10, 1, 0, 0, 0, 14, 12, 1,
		0, 0, 0, 15, 59, 1, 0, 0, 0, 16, 17, 10, 12, 0, 0, 17, 18, 7, 2, 0, 0,
		18, 58, 3, 2, 1, 13, 19, 20, 10, 11, 0, 0, 20, 21, 7, 3, 0, 0, 21, 58,
		3, 2, 1, 12, 22, 23, 10, 10, 0, 0, 23, 24, 7, 4, 0, 0, 24, 58, 3, 2, 1,
		11, 25, 26, 10, 9, 0, 0, 26, 27, 7, 5, 0, 0, 27, 58, 3, 2, 1, 10, 28, 29,
		10, 8, 0, 0, 29, 30, 7, 6, 0, 0, 30, 58, 3, 2, 1, 9, 31, 32, 10, 7, 0,
		0, 32, 33, 7, 7, 0, 0, 33, 58, 3, 2, 1, 8, 34, 35, 10, 6, 0, 0, 35, 36,
		5, 18, 0, 0, 36, 58, 3, 2, 1, 7, 37, 38, 10, 5, 0, 0, 38, 39, 5, 19, 0,
		0, 39, 58, 3, 2, 1, 6, 40, 41, 10, 4, 0, 0, 41, 42, 5, 20, 0, 0, 42, 58,
		3, 2, 1, 5, 43, 44, 10, 3, 0, 0, 44, 45, 5, 21, 0, 0, 45, 58, 3, 2, 1,
		4, 46, 47, 10, 2, 0, 0, 47, 48, 5, 22, 0, 0, 48, 58, 3, 2, 1, 3, 49, 50,
		10, 1, 0, 0, 50, 51, 5, 23, 0, 0, 51, 52, 3, 2, 1, 0, 52, 53, 5, 24, 0,
		0, 53, 54, 3, 2, 1, 1, 54, 58, 1, 0, 0, 0, 55, 56, 10, 15, 0, 0, 56, 58,
		7, 8, 0, 0, 57, 16, 1, 0, 0, 0, 57, 19, 1, 0, 0, 0, 57, 22, 1, 0, 0, 0,
		57, 25, 1, 0, 0, 0, 57, 28, 1, 0, 0, 0, 57, 31, 1, 0, 0, 0, 57, 34, 1,
		0, 0, 0, 57, 37, 1, 0, 0, 0, 57, 40, 1, 0, 0, 0, 57, 43, 1, 0, 0, 0, 57,
		46, 1, 0, 0, 0, 57, 49, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58, 61, 1, 0, 0,
		0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 3, 1, 0, 0, 0, 61, 59, 1,
		0, 0, 0, 62, 63, 5, 25, 0, 0, 63, 64, 3, 2, 1, 0, 64, 65, 5, 26, 0, 0,
		65, 69, 1, 0, 0, 0, 66, 69, 5, 27, 0, 0, 67, 69, 5, 28, 0, 0, 68, 62, 1,
		0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 67, 1, 0, 0, 0, 69, 5, 1, 0, 0, 0, 4,
		14, 57, 59, 68,
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

// pluralParserInit initializes any static state used to implement pluralParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewpluralParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PluralParserInit() {
	staticData := &PluralParserStaticData
	staticData.once.Do(pluralParserInit)
}

// NewpluralParser produces a new parser instance for the optional input antlr.TokenStream.
func NewpluralParser(input antlr.TokenStream) *pluralParser {
	PluralParserInit()
	this := new(pluralParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PluralParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "plural.g4"

	return this
}

// pluralParser tokens.
const (
	pluralParserEOF   = antlr.TokenEOF
	pluralParserT__0  = 1
	pluralParserT__1  = 2
	pluralParserT__2  = 3
	pluralParserT__3  = 4
	pluralParserT__4  = 5
	pluralParserT__5  = 6
	pluralParserT__6  = 7
	pluralParserT__7  = 8
	pluralParserT__8  = 9
	pluralParserT__9  = 10
	pluralParserT__10 = 11
	pluralParserT__11 = 12
	pluralParserT__12 = 13
	pluralParserT__13 = 14
	pluralParserT__14 = 15
	pluralParserT__15 = 16
	pluralParserT__16 = 17
	pluralParserT__17 = 18
	pluralParserT__18 = 19
	pluralParserT__19 = 20
	pluralParserT__20 = 21
	pluralParserT__21 = 22
	pluralParserT__22 = 23
	pluralParserT__23 = 24
	pluralParserT__24 = 25
	pluralParserT__25 = 26
	pluralParserT__26 = 27
	pluralParserINT   = 28
	pluralParserWS    = 29
)

// pluralParser rules.
const (
	pluralParserRULE_start   = 0
	pluralParserRULE_exp     = 1
	pluralParserRULE_primary = 2
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pluralParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pluralParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pluralParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pluralListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pluralListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *pluralParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, pluralParserRULE_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(6)
		p.exp(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPrefix returns the prefix token.
	GetPrefix() antlr.Token

	// GetBop returns the bop token.
	GetBop() antlr.Token

	// GetPostfix returns the postfix token.
	GetPostfix() antlr.Token

	// SetPrefix sets the prefix token.
	SetPrefix(antlr.Token)

	// SetBop sets the bop token.
	SetBop(antlr.Token)

	// SetPostfix sets the postfix token.
	SetPostfix(antlr.Token)

	// Getter signatures
	Primary() IPrimaryContext
	AllExp() []IExpContext
	Exp(i int) IExpContext

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	prefix  antlr.Token
	bop     antlr.Token
	postfix antlr.Token
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pluralParserRULE_exp
	return p
}

func InitEmptyExpContext(p *ExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pluralParserRULE_exp
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pluralParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) GetPrefix() antlr.Token { return s.prefix }

func (s *ExpContext) GetBop() antlr.Token { return s.bop }

func (s *ExpContext) GetPostfix() antlr.Token { return s.postfix }

func (s *ExpContext) SetPrefix(v antlr.Token) { s.prefix = v }

func (s *ExpContext) SetBop(v antlr.Token) { s.bop = v }

func (s *ExpContext) SetPostfix(v antlr.Token) { s.postfix = v }

func (s *ExpContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *ExpContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pluralListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pluralListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *pluralParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *pluralParser) exp(_p int) (localctx IExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, pluralParserRULE_exp, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(14)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case pluralParserT__24, pluralParserT__26, pluralParserINT:
		{
			p.SetState(9)
			p.Primary()
		}

	case pluralParserT__0, pluralParserT__1, pluralParserT__2, pluralParserT__3:
		{
			p.SetState(10)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpContext).prefix = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpContext).prefix = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(11)
			p.exp(14)
		}

	case pluralParserT__4, pluralParserT__5:
		{
			p.SetState(12)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpContext).prefix = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == pluralParserT__4 || _la == pluralParserT__5) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpContext).prefix = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(13)
			p.exp(13)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pluralParserRULE_exp)
				p.SetState(16)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(17)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&896) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(18)
					p.exp(13)
				}

			case 2:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pluralParserRULE_exp)
				p.SetState(19)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(20)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == pluralParserT__2 || _la == pluralParserT__3) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(21)
					p.exp(12)
				}

			case 3:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pluralParserRULE_exp)
				p.SetState(22)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(23)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == pluralParserT__9 || _la == pluralParserT__10) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(24)
					p.exp(11)
				}

			case 4:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pluralParserRULE_exp)
				p.SetState(25)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(26)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == pluralParserT__11 || _la == pluralParserT__12) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(27)
					p.exp(10)
				}

			case 5:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pluralParserRULE_exp)
				p.SetState(28)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(29)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == pluralParserT__13 || _la == pluralParserT__14) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(30)
					p.exp(9)
				}

			case 6:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pluralParserRULE_exp)
				p.SetState(31)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(32)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == pluralParserT__15 || _la == pluralParserT__16) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(33)
					p.exp(8)
				}

			case 7:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pluralParserRULE_exp)
				p.SetState(34)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(35)

					var _m = p.Match(pluralParserT__17)

					localctx.(*ExpContext).bop = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(36)
					p.exp(7)
				}

			case 8:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pluralParserRULE_exp)
				p.SetState(37)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(38)

					var _m = p.Match(pluralParserT__18)

					localctx.(*ExpContext).bop = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(39)
					p.exp(6)
				}

			case 9:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pluralParserRULE_exp)
				p.SetState(40)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(41)

					var _m = p.Match(pluralParserT__19)

					localctx.(*ExpContext).bop = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(42)
					p.exp(5)
				}

			case 10:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pluralParserRULE_exp)
				p.SetState(43)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(44)

					var _m = p.Match(pluralParserT__20)

					localctx.(*ExpContext).bop = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(45)
					p.exp(4)
				}

			case 11:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pluralParserRULE_exp)
				p.SetState(46)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(47)

					var _m = p.Match(pluralParserT__21)

					localctx.(*ExpContext).bop = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(48)
					p.exp(3)
				}

			case 12:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pluralParserRULE_exp)
				p.SetState(49)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(50)

					var _m = p.Match(pluralParserT__22)

					localctx.(*ExpContext).bop = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(51)
					p.exp(0)
				}
				{
					p.SetState(52)
					p.Match(pluralParserT__23)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(53)
					p.exp(1)
				}

			case 13:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pluralParserRULE_exp)
				p.SetState(55)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(56)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).postfix = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == pluralParserT__0 || _la == pluralParserT__1) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).postfix = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext
	INT() antlr.TerminalNode

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pluralParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = pluralParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = pluralParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *PrimaryContext) INT() antlr.TerminalNode {
	return s.GetToken(pluralParserINT, 0)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pluralListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pluralListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (p *pluralParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, pluralParserRULE_primary)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case pluralParserT__24:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Match(pluralParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.exp(0)
		}
		{
			p.SetState(64)
			p.Match(pluralParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case pluralParserT__26:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Match(pluralParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case pluralParserINT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Match(pluralParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *pluralParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *pluralParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 15)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

// Code generated from plural.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // plural

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 31, 73, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 5, 3, 17, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 60, 10,
	3, 12, 3, 14, 3, 63, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 71,
	10, 4, 3, 4, 2, 3, 4, 5, 2, 4, 6, 2, 11, 3, 2, 3, 6, 3, 2, 7, 8, 3, 2,
	9, 11, 3, 2, 5, 6, 3, 2, 12, 13, 3, 2, 14, 15, 3, 2, 16, 17, 3, 2, 18,
	19, 3, 2, 3, 4, 2, 86, 2, 8, 3, 2, 2, 2, 4, 16, 3, 2, 2, 2, 6, 70, 3, 2,
	2, 2, 8, 9, 5, 4, 3, 2, 9, 3, 3, 2, 2, 2, 10, 11, 8, 3, 1, 2, 11, 17, 5,
	6, 4, 2, 12, 13, 9, 2, 2, 2, 13, 17, 5, 4, 3, 16, 14, 15, 9, 3, 2, 2, 15,
	17, 5, 4, 3, 15, 16, 10, 3, 2, 2, 2, 16, 12, 3, 2, 2, 2, 16, 14, 3, 2,
	2, 2, 17, 61, 3, 2, 2, 2, 18, 19, 12, 14, 2, 2, 19, 20, 9, 4, 2, 2, 20,
	60, 5, 4, 3, 15, 21, 22, 12, 13, 2, 2, 22, 23, 9, 5, 2, 2, 23, 60, 5, 4,
	3, 14, 24, 25, 12, 12, 2, 2, 25, 26, 9, 6, 2, 2, 26, 60, 5, 4, 3, 13, 27,
	28, 12, 11, 2, 2, 28, 29, 9, 7, 2, 2, 29, 60, 5, 4, 3, 12, 30, 31, 12,
	10, 2, 2, 31, 32, 9, 8, 2, 2, 32, 60, 5, 4, 3, 11, 33, 34, 12, 9, 2, 2,
	34, 35, 9, 9, 2, 2, 35, 60, 5, 4, 3, 10, 36, 37, 12, 8, 2, 2, 37, 38, 7,
	20, 2, 2, 38, 60, 5, 4, 3, 9, 39, 40, 12, 7, 2, 2, 40, 41, 7, 21, 2, 2,
	41, 60, 5, 4, 3, 8, 42, 43, 12, 6, 2, 2, 43, 44, 7, 22, 2, 2, 44, 60, 5,
	4, 3, 7, 45, 46, 12, 5, 2, 2, 46, 47, 7, 23, 2, 2, 47, 60, 5, 4, 3, 6,
	48, 49, 12, 4, 2, 2, 49, 50, 7, 24, 2, 2, 50, 60, 5, 4, 3, 5, 51, 52, 12,
	3, 2, 2, 52, 53, 7, 25, 2, 2, 53, 54, 5, 4, 3, 2, 54, 55, 7, 26, 2, 2,
	55, 56, 5, 4, 3, 3, 56, 60, 3, 2, 2, 2, 57, 58, 12, 17, 2, 2, 58, 60, 9,
	10, 2, 2, 59, 18, 3, 2, 2, 2, 59, 21, 3, 2, 2, 2, 59, 24, 3, 2, 2, 2, 59,
	27, 3, 2, 2, 2, 59, 30, 3, 2, 2, 2, 59, 33, 3, 2, 2, 2, 59, 36, 3, 2, 2,
	2, 59, 39, 3, 2, 2, 2, 59, 42, 3, 2, 2, 2, 59, 45, 3, 2, 2, 2, 59, 48,
	3, 2, 2, 2, 59, 51, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 60, 63, 3, 2, 2, 2,
	61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 5, 3, 2, 2, 2, 63, 61, 3, 2,
	2, 2, 64, 65, 7, 27, 2, 2, 65, 66, 5, 4, 3, 2, 66, 67, 7, 28, 2, 2, 67,
	71, 3, 2, 2, 2, 68, 71, 7, 29, 2, 2, 69, 71, 7, 30, 2, 2, 70, 64, 3, 2,
	2, 2, 70, 68, 3, 2, 2, 2, 70, 69, 3, 2, 2, 2, 71, 7, 3, 2, 2, 2, 6, 16,
	59, 61, 70,
}
var literalNames = []string{
	"", "'++'", "'--'", "'+'", "'-'", "'~'", "'!'", "'*'", "'/'", "'%'", "'>>'",
	"'<<'", "'>'", "'<'", "'>='", "'<='", "'=='", "'!='", "'&'", "'^'", "'|'",
	"'&&'", "'||'", "'?'", "':'", "'('", "')'", "'n'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "INT", "WS",
}

var ruleNames = []string{
	"start", "exp", "primary",
}

type pluralParser struct {
	*antlr.BaseParser
}

// NewpluralParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *pluralParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewpluralParser(input antlr.TokenStream) *pluralParser {
	this := new(pluralParser)
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

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pluralParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pluralParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

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

func (p *pluralParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, pluralParserRULE_start)

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
		p.SetState(6)
		p.exp(0)
	}

	return localctx
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

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	prefix  antlr.Token
	bop     antlr.Token
	postfix antlr.Token
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pluralParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *ExpContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ExpContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

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

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(14)
	p.GetErrorHandler().Sync(p)

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

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<pluralParserT__0)|(1<<pluralParserT__1)|(1<<pluralParserT__2)|(1<<pluralParserT__3))) != 0) {
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
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, pluralParserRULE_exp)
				p.SetState(16)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(17)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<pluralParserT__6)|(1<<pluralParserT__7)|(1<<pluralParserT__8))) != 0) {
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
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
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
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
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
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
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
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
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
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
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
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(35)

					var _m = p.Match(pluralParserT__17)

					localctx.(*ExpContext).bop = _m
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
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(38)

					var _m = p.Match(pluralParserT__18)

					localctx.(*ExpContext).bop = _m
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
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(41)

					var _m = p.Match(pluralParserT__19)

					localctx.(*ExpContext).bop = _m
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
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(44)

					var _m = p.Match(pluralParserT__20)

					localctx.(*ExpContext).bop = _m
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
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(47)

					var _m = p.Match(pluralParserT__21)

					localctx.(*ExpContext).bop = _m
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
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(50)

					var _m = p.Match(pluralParserT__22)

					localctx.(*ExpContext).bop = _m
				}
				{
					p.SetState(51)
					p.exp(0)
				}
				{
					p.SetState(52)
					p.Match(pluralParserT__23)
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
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
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

			}

		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pluralParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pluralParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

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

	p.SetState(68)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case pluralParserT__24:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Match(pluralParserT__24)
		}
		{
			p.SetState(63)
			p.exp(0)
		}
		{
			p.SetState(64)
			p.Match(pluralParserT__25)
		}

	case pluralParserT__26:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Match(pluralParserT__26)
		}

	case pluralParserINT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Match(pluralParserINT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
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

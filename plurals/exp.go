package plurals

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/antlr4-go/antlr/v4"
	"github.com/youthlin/t/errors"
	parse2 "github.com/youthlin/t/plurals/parse"
	"github.com/youthlin/t/plurals/parser"
)

type ParserEngine string

const (
	ParserEngineAuto        ParserEngine = "auto"
	ParserEngineHandwritten ParserEngine = "handwritten"
	ParserEngineANTLR       ParserEngine = "antlr"
)

var defaultParserEngine atomic.Value

const ctxKeyParserEngine = ctxKey("plural-parser-engine")

func init() {
	defaultParserEngine.Store(ParserEngineAuto)
}

func WithParserEngine(ctx context.Context, engine ParserEngine) context.Context {
	return context.WithValue(ctx, ctxKeyParserEngine, engine.normalize())
}

func SetDefaultParserEngine(engine ParserEngine) {
	defaultParserEngine.Store(engine.normalize())
}

func DefaultParserEngine() ParserEngine {
	if v, ok := defaultParserEngine.Load().(ParserEngine); ok {
		return v.normalize()
	}
	return ParserEngineAuto
}

func (e ParserEngine) normalize() ParserEngine {
	switch e {
	case ParserEngineANTLR, ParserEngineHandwritten:
		return e
	default:
		return ParserEngineAuto
	}
}

func parserEngineFromContext(ctx context.Context) ParserEngine {
	if ctx != nil {
		if v, ok := ctx.Value(ctxKeyParserEngine).(ParserEngine); ok {
			return v.normalize()
		}
	}
	return DefaultParserEngine()
}

// Eval 传入复数表达式，返回 n 应该使用哪种复数形式
func Eval(ctx context.Context, exp string, n int64) (result int64, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.Errorf("unexpected error: %v", e)
		}
	}()

	// 1 常见表达式直接走函数
	commonExp := strings.ReplaceAll(exp, " ", "")
	fun, ok := commons[commonExp]
	if ok {
		return fun(n), nil
	}

	switch parserEngineFromContext(ctx) {
	case ParserEngineANTLR:
		return evalANTLR(ctx, exp, n)
	case ParserEngineHandwritten:
		return evalHandwritten(exp, n)
	default:
		result, err := evalHandwritten(exp, n)
		if err == nil {
			return result, nil
		}
		result2, err2 := evalANTLR(ctx, exp, n)
		if err2 == nil {
			return result2, nil
		}
		return 0, errors.Wrapf(errors.WithSecondaryError(err, err2), "parse plural expression")
	}
}

func evalHandwritten(exp string, n int64) (int64, error) {
	fn, err := parse2.String(exp)
	if err != nil {
		return 0, errors.Wrapf(err, "handwritten parser")
	}
	return fn(n), nil
}

func evalANTLR(ctx context.Context, exp string, n int64) (result int64, err error) {
	input := antlr.NewInputStream(exp)
	lexer := parser.NewpluralLexer(input)
	errListener := new(errorListener)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewpluralParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)

	l := newListener(ctx, n)
	tree := p.Start_()
	antlr.ParseTreeWalkerDefault.Walk(l, tree)
	if errListener.err == nil && stream.LA(1) != antlr.TokenEOF {
		errListener.addError(errors.Errorf("unexpected trailing token: %q", stream.LT(1).GetText()))
	}
	return l.result, errListener.err
}

// gettext plural expressions follow C-like operator syntax.
// In particular, binary '^' means XOR and unary '~' means bitwise NOT.
// Do not confuse this with Go's unary '^', which is also bitwise NOT.
// c syntax accept int(0) as false, none-zero as true
const (
	_TRUE  = 1
	_FALSE = 0
)

// myListener is a listener when walk the ast can be do something
// see https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/
// 遍历语法树时操作一个栈来计算表达式
type myListener struct {
	*parser.BasepluralListener
	stack  Int64Stack
	ctx    context.Context
	n      int64
	result int64
}

func newListener(ctx context.Context, n int64) *myListener {
	return &myListener{
		ctx:   ctx,
		n:     n,
		stack: make(Int64Stack, 0),
	}
}

func (s *myListener) ExitStart(ctx *parser.StartContext) {
	s.result, _ = s.stack.Pop()
}

func (s *myListener) ExitExp(ctx *parser.ExpContext) {
	var (
		prefix  = ctx.GetPrefix()
		bop     = ctx.GetBop()
		postfix = ctx.GetPostfix()
	)
	s.handleExpPrefix(prefix)
	s.handleExpBop(bop)
	s.handleExpPostfix(postfix)
}

func (s *myListener) handleExpPrefix(prefix antlr.Token) {
	if prefix == nil {
		return
	}
	prefixText := prefix.GetText()
	s.debug("prefix: %v stack=%v", prefixText, s.stack)
	switch prefixText {
	case "+":
	case "-":
		num, _ := s.stack.Pop()
		s.stack.Push(-num)
	case "++":
		num, _ := s.stack.Pop()
		s.stack.Push(num + 1)
	case "--":
		num, _ := s.stack.Pop()
		s.stack.Push(num - 1)
	case "~":
		num, _ := s.stack.Pop()
		s.stack.Push(^num)
	case "!":
		num, _ := s.stack.Pop()
		if num == _TRUE {
			s.stack.Push(_FALSE)
		} else {
			s.stack.Push(_TRUE)
		}
	default:
		panic("assert error: unexpected text(prefix): " + prefixText)
	}
}

func (s *myListener) handleExpBop(bop antlr.Token) {
	if bop == nil {
		return
	}
	bopText := bop.GetText()
	s.debug("bop   : %v stack=%v", bopText, s.stack)
	switch bopText {
	case "*":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		s.stack.Push(left * right)
	case "/":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		s.stack.Push(left / right)
	case "%":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		s.stack.Push(left % right)
	case "+":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		s.stack.Push(left + right)
	case "-":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		s.stack.Push(left - right)
	case ">>":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		s.stack.Push(left >> right)
	case "<<":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		s.stack.Push(left << right)
	case ">":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		if left > right {
			s.stack.Push(_TRUE)
		} else {
			s.stack.Push(_FALSE)
		}
	case "<":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		if left < right {
			s.stack.Push(_TRUE)
		} else {
			s.stack.Push(_FALSE)
		}
	case ">=":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		if left >= right {
			s.stack.Push(_TRUE)
		} else {
			s.stack.Push(_FALSE)
		}
	case "<=":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		if left <= right {
			s.stack.Push(_TRUE)
		} else {
			s.stack.Push(_FALSE)
		}
	case "==":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		if left == right {
			s.stack.Push(_TRUE)
		} else {
			s.stack.Push(_FALSE)
		}
	case "!=":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		if left != right {
			s.stack.Push(_TRUE)
		} else {
			s.stack.Push(_FALSE)
		}
	case "&":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		s.stack.Push(left & right)
	case "|":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		s.stack.Push(left | right)
	case "^":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		s.stack.Push(left ^ right)
	case "&&":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		if left == _TRUE && right == _TRUE {
			s.stack.Push(_TRUE)
		} else {
			s.stack.Push(_FALSE)
		}
	case "||":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		if left == _TRUE || right == _TRUE {
			s.stack.Push(_TRUE)
		} else {
			s.stack.Push(_FALSE)
		}
	case "?":
		var right, _ = s.stack.Pop()
		var left, _ = s.stack.Pop()
		var cond, _ = s.stack.Pop()
		if cond == _TRUE {
			s.stack.Push(left)
		} else {
			s.stack.Push(right)
		}
	default:
		panic("assert error: unexpected text(bop): " + bopText)
	}
	s.debug("bop done stack=%v", s.stack)
}

func (s *myListener) handleExpPostfix(postfix antlr.Token) {
	if postfix == nil {
		return
	}
	postText := postfix.GetText()
	s.debug("postfix: %v stack=%v", postText, s.stack)
	switch postText {
	case "++":
		num, _ := s.stack.Pop()
		s.stack.Push(num + 1)
	case "--":
		num, _ := s.stack.Pop()
		s.stack.Push(num - 1)
	default:
		panic("assert error: unexpected text(postfix): " + postText)
	}
}

func (s *myListener) ExitPrimary(ctx *parser.PrimaryContext) {
	start := ctx.GetStart()
	switch start.GetText() {
	case "n":
		s.stack.Push(s.n)
	case "(":
		s.debug("primary: (exp) stack=%v", s.stack)
	default:
		num := ctx.GetText()
		iNum, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			panic("assert error: not a number: " + num)
		}
		s.stack.Push(iNum)
	}
	s.debug("primary: %v stack=%v", ctx.GetText(), s.stack)
}

func (s *myListener) debug(format string, args ...interface{}) {
	if s.ctx.Value(ctxKeyDebug) != nil {
		fmt.Printf(format+"\n", args...)
	}
}

package plurals

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/pkg/errors"
	"github.com/youthlin/t/plurals/parser"
)

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

	// 2 词法解析
	input := antlr.NewInputStream(exp)
	lexer := parser.NewpluralLexer(input)
	errListener := new(errorListener)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errListener)

	// 3 语法树生成
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewpluralParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)

	// 4 遍历语法树计算表达式
	l := newListener(ctx, n)
	tree := p.Start()
	antlr.ParseTreeWalkerDefault.Walk(l, tree)
	return l.result, errListener.err
}

// c syntax accept int(0) as false, none-zero as true
const (
	_TRUE  = 1
	_FALSE = 0
)

// myListener is a listener when walk the ast can be do something
// see https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/
// 遍历语法树时操作一个栈来计算表达式
type myListener struct {
	// 实现这个父类上感兴趣的方法(相当于一个 adaptor)
	*parser.BasepluralListener
	stack  Int64Stack      // 操作数栈
	ctx    context.Context // 外部传入的 ctx， 用于 debug
	n      int64           // 参数 n
	result int64           // 结果
}

func newListener(ctx context.Context, n int64) *myListener {
	return &myListener{
		ctx:   ctx,
		n:     n,
		stack: make(Int64Stack, 0),
	}
}

// ExitStart override
func (s *myListener) ExitStart(ctx *parser.StartContext) {
	s.result, _ = s.stack.Pop()
}

// ExitExp override
func (s *myListener) ExitExp(ctx *parser.ExpContext) {
	var (
		prefix  = ctx.GetPrefix()
		bop     = ctx.GetBop()
		postfix = ctx.GetPostfix()
	)
	if prefix != nil {
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
			s.stack.Push(^num) // go 使用 ^表示按位取反
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
	if bop != nil {
		bopText := bop.GetText()
		s.debug("bop   : %v stack=%v", bopText, s.stack)
		switch bopText {
		case "*":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			s.stack.Push(left * right)
		case "/":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			s.stack.Push(left / right)
		case "%":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			s.stack.Push(left % right)
		case "+":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			s.stack.Push(left + right)
		case "-":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			s.stack.Push(left - right)
		case ">>":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			s.stack.Push(left >> right)
		case "<<":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			s.stack.Push(left << right)
		case ">":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			if left > right {
				s.stack.Push(_TRUE)
			} else {
				s.stack.Push(_FALSE)
			}
		case "<":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			if left < right {
				s.stack.Push(_TRUE)
			} else {
				s.stack.Push(_FALSE)
			}
		case ">=":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			if left >= right {
				s.stack.Push(_TRUE)
			} else {
				s.stack.Push(_FALSE)
			}
		case "<=":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			if left <= right {
				s.stack.Push(_TRUE)
			} else {
				s.stack.Push(_FALSE)
			}
		case "==":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			if left == right {
				s.stack.Push(_TRUE)
			} else {
				s.stack.Push(_FALSE)
			}
		case "!=":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			if left != right {
				s.stack.Push(_TRUE)
			} else {
				s.stack.Push(_FALSE)
			}
		case "&":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			s.stack.Push(left & right)
		case "|":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			s.stack.Push(left | right)
		case "^":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			s.stack.Push(left & right) // 作为二元操作符，是异或
		case "&&":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			if left == _TRUE && right == _TRUE {
				s.stack.Push(_TRUE)
			} else {
				s.stack.Push(_FALSE)
			}
		case "||":
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
			)
			if left == _TRUE || right == _TRUE {
				s.stack.Push(_TRUE)
			} else {
				s.stack.Push(_FALSE)
			}
		case "?":
			s.debug("%v", s.stack)
			var (
				right, _ = s.stack.Pop()
				left, _  = s.stack.Pop()
				cond, _  = s.stack.Pop()
			)
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
	if postfix != nil {
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
	// prefix, bop, postfix all nil, then it's a primary rule
}

// ExitPrimary override
func (s *myListener) ExitPrimary(ctx *parser.PrimaryContext) {
	// primary: '(' exp ')' | 'n' | INT;
	start := ctx.GetStart()
	switch start.GetText() {
	case "n":
		s.stack.Push(s.n) // the only variable n
	case "(":
		// 不用出入栈
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

// debug if ctx has a debug key, then print msg.
// see DebugContext
func (s *myListener) debug(format string, args ...interface{}) {
	if s.ctx.Value(ctxKeyDebug) != nil {
		fmt.Printf(format+"\n", args...)
	}
}

package parse

import (
	"fmt"
	"strconv"
	"sync"
)

var cache sync.Map

func Eval(input string, n int64) int64 {
	if val, ok := cache.Load(input); ok {
		if f, ok := val.(Fn); ok {
			return f(n)
		}
	}
	f, err := String(input)
	if err != nil {
		f = zero
	}
	cache.Store(input, f)
	return f(n)
}

type Fn func(n int64) int64

func zero(n int64) int64 { return 0 }

func String(input string) (Fn, error) {
	tokens, err := Lex(input)
	if err != nil {
		return zero, err
	}
	return Tokens(tokens)
}

func Tokens(input []*Token) (Fn, error) {
	tree, err := ParseToken(input)
	if err != nil {
		return zero, err
	}
	return Ast(tree)
}

func Ast(tree *Node) (Fn, error) {
	if tree == nil {
		return zero, fmt.Errorf("nil ast")
	}
	return compile(tree)
}

func compile(node *Node) (Fn, error) {
	if node == nil {
		return zero, fmt.Errorf("nil node")
	}
	switch node.Name {
	case nameInt:
		if node.Token == nil {
			return zero, fmt.Errorf("int node without token")
		}
		v, err := strconv.ParseInt(node.Token.Value, 10, 64)
		if err != nil {
			return zero, err
		}
		return func(n int64) int64 { return v }, nil
	case nameID:
		if node.Token == nil {
			return zero, fmt.Errorf("id node without token")
		}
		if node.Token.Value != "n" {
			return zero, fmt.Errorf("unexpected identifier %q, only \"n\" is allowed", node.Token.Value)
		}
		return func(n int64) int64 { return n }, nil
	case nameGroup:
		if len(node.Children) != 1 {
			return zero, fmt.Errorf("group node expects 1 child, got %d", len(node.Children))
		}
		return compile(node.Children[0])
	case namePrefix:
		if node.Token == nil || len(node.Children) != 1 {
			return zero, fmt.Errorf("invalid prefix node")
		}
		child, err := compile(node.Children[0])
		if err != nil {
			return zero, err
		}
		switch node.Token.Type {
		case TokenTypePlus:
			return child, nil
		case TokenTypeMinus:
			return func(n int64) int64 { return -child(n) }, nil
		case TokenTypeIncr:
			return func(n int64) int64 { return child(n) + 1 }, nil
		case TokenTypeDecr:
			return func(n int64) int64 { return child(n) - 1 }, nil
		case TokenTypeBitNot:
			return func(n int64) int64 { return ^child(n) }, nil
		case TokenTypeNot:
			return func(n int64) int64 { return toBool(child(n) == 0) }, nil
		default:
			return zero, fmt.Errorf("unsupported prefix operator %q", node.Token.Type)
		}
	case namePostfix:
		if node.Token == nil || len(node.Children) != 1 {
			return zero, fmt.Errorf("invalid postfix node")
		}
		child, err := compile(node.Children[0])
		if err != nil {
			return zero, err
		}
		switch node.Token.Type {
		case TokenTypeIncr:
			return func(n int64) int64 { return child(n) + 1 }, nil
		case TokenTypeDecr:
			return func(n int64) int64 { return child(n) - 1 }, nil
		default:
			return zero, fmt.Errorf("unsupported postfix operator %q", node.Token.Type)
		}
	case nameBinary:
		if node.Token == nil || len(node.Children) != 2 {
			return zero, fmt.Errorf("invalid binary node")
		}
		left, err := compile(node.Children[0])
		if err != nil {
			return zero, err
		}
		right, err := compile(node.Children[1])
		if err != nil {
			return zero, err
		}
		switch node.Token.Type {
		case TokenTypeTimes:
			return func(n int64) int64 { return left(n) * right(n) }, nil
		case TokenTypeOver:
			return func(n int64) int64 { return left(n) / right(n) }, nil
		case TokenTypeMod:
			return func(n int64) int64 { return left(n) % right(n) }, nil
		case TokenTypePlus:
			return func(n int64) int64 { return left(n) + right(n) }, nil
		case TokenTypeMinus:
			return func(n int64) int64 { return left(n) - right(n) }, nil
		case TokenTypeShiftL:
			return func(n int64) int64 { return left(n) << right(n) }, nil
		case TokenTypeShiftR:
			return func(n int64) int64 { return left(n) >> right(n) }, nil
		case TokenTypeGt:
			return func(n int64) int64 { return toBool(left(n) > right(n)) }, nil
		case TokenTypeLt:
			return func(n int64) int64 { return toBool(left(n) < right(n)) }, nil
		case TokenTypeGe:
			return func(n int64) int64 { return toBool(left(n) >= right(n)) }, nil
		case TokenTypeLe:
			return func(n int64) int64 { return toBool(left(n) <= right(n)) }, nil
		case TokenTypeEq:
			return func(n int64) int64 { return toBool(left(n) == right(n)) }, nil
		case TokenTypeNe:
			return func(n int64) int64 { return toBool(left(n) != right(n)) }, nil
		case TokenTypeBitAnd:
			return func(n int64) int64 { return left(n) & right(n) }, nil
		case TokenTypeBitXor:
			return func(n int64) int64 { return left(n) ^ right(n) }, nil
		case TokenTypeBitOr:
			return func(n int64) int64 { return left(n) | right(n) }, nil
		case TokenTypeAnd:
			return func(n int64) int64 { return toBool(left(n) != 0 && right(n) != 0) }, nil
		case TokenTypeOr:
			return func(n int64) int64 { return toBool(left(n) != 0 || right(n) != 0) }, nil
		default:
			return zero, fmt.Errorf("unsupported binary operator %q", node.Token.Type)
		}
	case nameTernary:
		if len(node.Children) != 3 {
			return zero, fmt.Errorf("ternary node expects 3 children, got %d", len(node.Children))
		}
		cond, err := compile(node.Children[0])
		if err != nil {
			return zero, err
		}
		thenFn, err := compile(node.Children[1])
		if err != nil {
			return zero, err
		}
		elseFn, err := compile(node.Children[2])
		if err != nil {
			return zero, err
		}
		return func(n int64) int64 {
			if cond(n) != 0 {
				return thenFn(n)
			}
			return elseFn(n)
		}, nil
	default:
		return zero, fmt.Errorf("unsupported node type %q", node.Name)
	}
}

func toBool(ok bool) int64 {
	if ok {
		return 1
	}
	return 0
}

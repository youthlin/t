package parse

import (
	"fmt"
	"io"
	"strings"
)

func ParseInput(input string) (tree *Node, err error) {
	tokens, err := Lex(input)
	if err != nil {
		return
	}
	return ParseToken(tokens)
}

func ParseToken(input []*Token) (tree *Node, err error) {
	p := &parser{
		tokens: input,
		size:   len(input),
	}
	return p.parse()
}

type Node struct {
	// A->a 非终极符 A, 终极符
	Name     string
	Token    *Token // 终极符
	Parent   *Node
	Children []*Node
	Value    int64
}

const (
	nameExp     = "exp"
	namePrimary = "primary"
	nameExpMore = "expMore"
)

func (n *Node) add(child *Node) {
	if child.Name == nameExpMore && len(child.Children) == 0 {
		return
	}
	child.Parent = n
	n.Children = append(n.Children, child)
}

func (n *Node) String() string {
	var sb strings.Builder
	print(&sb, n, 0)
	return sb.String()
}

func print(w io.StringWriter, n *Node, depth int) {
	prefix := strings.Repeat("| ", depth)
	w.WriteString(fmt.Sprintf("%s[Name=%v Value=%v Token=%v\n",
		prefix, n.Name, n.Value, n.Token))
	for _, child := range n.Children {
		print(w, child, depth+1)
	}
	w.WriteString(fmt.Sprintf("%sEndOf=%v]\n", prefix, n.Name))
}

type parser struct {
	tokens []*Token
	index  int
	size   int
}

func (p *parser) parse() (tree *Node, err error) {
	return p.exp()
}

func (p *parser) exp() (node *Node, err error) {
	node = &Node{Name: nameExp}
	var (
		child *Node
		tok   = p.peekToken()
	)
	debug("exp. peek=%v", tok)
	if tok = p.peekToken(); tok != nil {
		tt := tok.Type
		switch tt {
		case TokenTypeParenL, TokenTypeID, TokenTypeInt: // primary
			child, err = p.primary()
			if err != nil {
				return nil, err
			}
			node.add(child)

			child, err = p.expMore()
			if err != nil {
				return nil, err
			}
			node.add(child)
		case TokenTypePlus, TokenTypeMinus,
			TokenTypeIncr, TokenTypeDecr,
			TokenTypeBitNot, TokenTypeNot: // 一元前缀操作符
			child, err = p.match(tt)
			if err != nil {
				return nil, err
			}
			node.add(child)

			child, err = p.exp()
			if err != nil {
				return nil, fmt.Errorf("expected exp of unary operator %q: %w", tt, err)
			}
			if tt == TokenTypeIncr || tt == TokenTypeDecr { // 自增自减操作符 只能作用在标识符上
				if len(child.Children) == 0 {
					return nil, fmt.Errorf("prefix operator %v can only before ID, got %v", tt, child)
				}
				pri := child.Children[0]
				if pri.Name != namePrimary {
					return nil, fmt.Errorf("prefix operator %v can only before ID, got %v", tt, child)
				}
				if len(pri.Children) == 0 {
					return nil, fmt.Errorf("prefix operator %v can only before ID, got %v", tt, child)
				}
				t := pri.Children[0].Token
				if t == nil {
					return nil, fmt.Errorf("prefix operator %v can only before ID, got %v", tt, child)
				}
				if t.Type != TokenTypeID {
					return nil, fmt.Errorf("prefix operator %v can only before ID, got %v", tt, child)
				}
			}
			node.add(child)

			child, err = p.expMore()
			if err != nil {
				return nil, err
			}
			node.add(child)
		default:
			return nil, fmt.Errorf(`unexpected %q when parse exp, expected primary["(", ID, INT] or unaray opertor`, tt)
		}
	} else {
		return nil, fmt.Errorf(`unexpected EOF when parse exp, expected primary["(", ID, INT]  or unaray operator`)
	}
	return node, nil
}

func (p *parser) primary() (node *Node, err error) {
	node = &Node{Name: namePrimary}
	var (
		child *Node
		tok   = p.peekToken()
	)
	debug("primary. peek=%v", tok)
	if tok = p.peekToken(); tok != nil {
		tt := tok.Type
		switch tt {
		case TokenTypeParenL:
			child, err = p.match(TokenTypeParenL)
			if err != nil {
				return nil, err
			}
			node.add(child)

			child, err = p.exp()
			if err != nil {
				return nil, fmt.Errorf(`failed to parse primary "(exp)", expected exp after "(": %w`, err)
			}
			node.add(child)

			child, err = p.match(TokenTypeParenR)
			if err != nil {
				return nil, err
			}
			node.add(child)
		case TokenTypeID, TokenTypeInt:
			child, err = p.match(tt)
			if err != nil {
				return nil, err
			}
			node.add(child)
		default:
			return nil, fmt.Errorf(`unexpected %q when parse primary, expected "(", ID, INT`, tt)
		}
	} else {
		return nil, fmt.Errorf(`unexpected EOF when parse primary, expected "(" or ID or INT`)
	}
	return node, nil
}

func (p *parser) expMore() (node *Node, err error) {
	node = &Node{Name: nameExpMore}
	var (
		child *Node
		tok   = p.peekToken()
	)
	debug("expMore. peek=%v", tok)
	if tok != nil {
		tt := tok.Type
		switch tt {
		case TokenTypeIncr, TokenTypeDecr: // 后缀 ++ --
			pre := p.previous()
			if pre.Type != TokenTypeID {
				return nil, fmt.Errorf("postfix operator %v can only after ID, previous token is %v", tt, pre)
			}
			child, err = p.match(tt)
			if err != nil {
				return nil, err
			}
			node.add(child)

			child, err = p.expMore()
			if err != nil {
				return nil, err
			}
			node.add(child)
		case TokenTypeTimes, TokenTypeOver, TokenTypeMod,
			TokenTypePlus, TokenTypeMinus,
			TokenTypeShiftL, TokenTypeShiftR,
			TokenTypeGt, TokenTypeLt,
			TokenTypeGe, TokenTypeLe,
			TokenTypeEq, TokenTypeNe,
			TokenTypeBitAnd, TokenTypeBitXor, TokenTypeBitOr,
			TokenTypeAnd, TokenTypeOr: // 中缀操作符
			child, err = p.match(tt)
			if err != nil {
				return nil, err
			}
			node.add(child)

			child, err = p.exp()
			if err != nil {
				return nil, fmt.Errorf("expected right exp of operator %q: %w", tt, err)
			}
			node.add(child)

			child, err = p.expMore()
			if err != nil {
				return nil, err
			}
			node.add(child)
		case TokenTypeQst: // 三元操作符 ? :
			child, err = p.match(tt)
			if err != nil {
				return nil, err
			}
			node.add(child)

			child, err = p.exp()
			if err != nil {
				return nil, fmt.Errorf(`expected first exp of operator "?:" after "?": %w`, err)
			}
			node.add(child)

			child, err = p.match(TokenTypeCol)
			if err != nil {
				return nil, err
			}
			node.add(child)

			child, err = p.exp()
			if err != nil {
				return nil, fmt.Errorf(`expected second exp of operator "?:" after ":": %w`, err)
			}
			node.add(child)

			child, err = p.expMore()
			if err != nil {
				return nil, err
			}
			node.add(child)
		}
	}
	return node, nil
}

func (p *parser) match(expected TokenType) (node *Node, err error) {
	tok := p.peekToken()
	debug("expected=%v. peek=%v", expected, tok)
	if tok != nil {
		if tok.Type != expected {
			return nil, fmt.Errorf("unexpected token: %q, expected: %q", tok, expected)
		}
		p.index++
		return &Node{Name: tok.Type.String(), Token: tok}, nil
	}
	return nil, io.EOF
}

func (p *parser) previous() *Token {
	i := p.index - 1
	if i >= 0 && i < p.size {
		return p.tokens[i]
	}
	return nil
}

func (p *parser) peekToken() *Token {
	if p.index < p.size {
		return p.tokens[p.index]
	}
	return nil
}

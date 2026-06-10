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
	Name     string
	Token    *Token
	Parent   *Node
	Children []*Node
	Value    int64
}

const (
	nameInt     = "int"
	nameID      = "id"
	nameGroup   = "group"
	namePrefix  = "prefix"
	namePostfix = "postfix"
	nameBinary  = "binary"
	nameTernary = "ternary"
)

func (n *Node) add(child *Node) {
	if child == nil {
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
	tree, err = p.parseConditional()
	if err != nil {
		return nil, err
	}
	if tok := p.peekToken(); tok != nil {
		return nil, fmt.Errorf("unexpected token %q after expression", tok)
	}
	return tree, nil
}

func (p *parser) parseConditional() (*Node, error) {
	cond, err := p.parseBinary(0)
	if err != nil {
		return nil, err
	}
	if tok := p.peekToken(); tok != nil && tok.Type == TokenTypeQst {
		op, err := p.match(TokenTypeQst)
		if err != nil {
			return nil, err
		}
		thenNode, err := p.parseConditional()
		if err != nil {
			return nil, fmt.Errorf(`expected expression after "?": %w`, err)
		}
		if _, err := p.match(TokenTypeCol); err != nil {
			return nil, fmt.Errorf(`expected ":" in conditional expression: %w`, err)
		}
		elseNode, err := p.parseConditional()
		if err != nil {
			return nil, fmt.Errorf(`expected expression after ":": %w`, err)
		}
		node := &Node{Name: nameTernary, Token: op.Token}
		node.add(cond)
		node.add(thenNode)
		node.add(elseNode)
		return node, nil
	}
	return cond, nil
}

func (p *parser) parseBinary(minPrec int) (*Node, error) {
	left, err := p.parsePrefix()
	if err != nil {
		return nil, err
	}
	for {
		tok := p.peekToken()
		if tok == nil {
			return left, nil
		}
		prec := binaryPrecedence(tok.Type)
		if prec < minPrec {
			return left, nil
		}
		op, err := p.match(tok.Type)
		if err != nil {
			return nil, err
		}
		right, err := p.parseBinary(prec + 1)
		if err != nil {
			return nil, fmt.Errorf("expected right expression of operator %q: %w", tok.Type, err)
		}
		node := &Node{Name: nameBinary, Token: op.Token}
		node.add(left)
		node.add(right)
		left = node
	}
}

func (p *parser) parsePrefix() (*Node, error) {
	tok := p.peekToken()
	if tok == nil {
		return nil, fmt.Errorf("unexpected EOF when parse expression")
	}
	switch tok.Type {
	case TokenTypePlus, TokenTypeMinus, TokenTypeIncr, TokenTypeDecr, TokenTypeBitNot, TokenTypeNot:
		op, err := p.match(tok.Type)
		if err != nil {
			return nil, err
		}
		child, err := p.parsePrefix()
		if err != nil {
			return nil, fmt.Errorf("expected expression of unary operator %q: %w", tok.Type, err)
		}
		if (tok.Type == TokenTypeIncr || tok.Type == TokenTypeDecr) && !child.isID() {
			return nil, fmt.Errorf("prefix operator %v can only before ID", tok.Type)
		}
		node := &Node{Name: namePrefix, Token: op.Token}
		node.add(child)
		return node, nil
	default:
		return p.parsePostfix()
	}
}

func (p *parser) parsePostfix() (*Node, error) {
	node, err := p.parsePrimary()
	if err != nil {
		return nil, err
	}
	for {
		tok := p.peekToken()
		if tok == nil || (tok.Type != TokenTypeIncr && tok.Type != TokenTypeDecr) {
			return node, nil
		}
		if !node.isID() {
			return nil, fmt.Errorf("postfix operator %v can only after ID", tok.Type)
		}
		op, err := p.match(tok.Type)
		if err != nil {
			return nil, err
		}
		parent := &Node{Name: namePostfix, Token: op.Token}
		parent.add(node)
		node = parent
	}
}

func (p *parser) parsePrimary() (*Node, error) {
	tok := p.peekToken()
	if tok == nil {
		return nil, fmt.Errorf(`unexpected EOF when parse primary, expected "(" or ID or INT`)
	}
	switch tok.Type {
	case TokenTypeParenL:
		if _, err := p.match(TokenTypeParenL); err != nil {
			return nil, err
		}
		child, err := p.parseConditional()
		if err != nil {
			return nil, fmt.Errorf(`failed to parse primary "(exp)": %w`, err)
		}
		if _, err := p.match(TokenTypeParenR); err != nil {
			return nil, err
		}
		node := &Node{Name: nameGroup}
		node.add(child)
		return node, nil
	case TokenTypeID:
		node, err := p.match(TokenTypeID)
		if err != nil {
			return nil, err
		}
		node.Name = nameID
		return node, nil
	case TokenTypeInt:
		node, err := p.match(TokenTypeInt)
		if err != nil {
			return nil, err
		}
		node.Name = nameInt
		return node, nil
	default:
		return nil, fmt.Errorf(`unexpected %q when parse primary, expected "(", ID, INT`, tok.Type)
	}
}

func (n *Node) isID() bool {
	return n != nil && n.Name == nameID && n.Token != nil && n.Token.Type == TokenTypeID
}

func binaryPrecedence(tt TokenType) int {
	switch tt {
	case TokenTypeTimes, TokenTypeOver, TokenTypeMod:
		return 9
	case TokenTypePlus, TokenTypeMinus:
		return 8
	case TokenTypeShiftL, TokenTypeShiftR:
		return 7
	case TokenTypeGt, TokenTypeLt, TokenTypeGe, TokenTypeLe:
		return 6
	case TokenTypeEq, TokenTypeNe:
		return 5
	case TokenTypeBitAnd:
		return 4
	case TokenTypeBitXor:
		return 3
	case TokenTypeBitOr:
		return 2
	case TokenTypeAnd:
		return 1
	case TokenTypeOr:
		return 0
	default:
		return -1
	}
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

func (p *parser) peekToken() *Token {
	if p.index < p.size {
		return p.tokens[p.index]
	}
	return nil
}

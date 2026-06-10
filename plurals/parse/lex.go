package parse

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type TokenType int

const (
	TokenTypeInvalid TokenType = iota

	TokenTypeID     // 标识符 只能是 n
	TokenTypeInt    // 非负整数
	TokenTypeParenL // (
	TokenTypeParenR // )
	TokenTypeIncr   // ++
	TokenTypeDecr   // --
	TokenTypePlus   // +
	TokenTypeMinus  // -
	TokenTypeTimes  // *
	TokenTypeOver   // /
	TokenTypeMod    // %
	TokenTypeNot    // !
	TokenTypeBitNot // ~
	TokenTypeShiftL // <<
	TokenTypeShiftR // >>
	TokenTypeGt     // >
	TokenTypeLt     // <
	TokenTypeGe     // >=
	TokenTypeLe     // <=
	TokenTypeEq     // ==
	TokenTypeNe     // !=
	TokenTypeBitAnd // &
	TokenTypeBitOr  // |
	TokenTypeBitXor // ^
	TokenTypeAnd    // &&
	TokenTypeOr     // ||
	TokenTypeQst    // ?
	TokenTypeCol    // :
)

func (tt TokenType) String() string {
	if name, ok := ttName[tt]; ok {
		return name
	}
	return fmt.Sprintf("TokenType(%d)", tt)
}

var ttName = map[TokenType]string{
	TokenTypeInvalid: "Invalid",
	TokenTypeID:      "ID",
	TokenTypeInt:     "Int",
	TokenTypeParenL:  "(",
	TokenTypeParenR:  ")",
	TokenTypeIncr:    "++",
	TokenTypeDecr:    "--",
	TokenTypePlus:    "+",
	TokenTypeMinus:   "-",
	TokenTypeTimes:   "*",
	TokenTypeOver:    "/",
	TokenTypeMod:     "%",
	TokenTypeNot:     "!",
	TokenTypeBitNot:  "~",
	TokenTypeShiftL:  "<<",
	TokenTypeShiftR:  ">>",
	TokenTypeGt:      ">",
	TokenTypeLt:      "<",
	TokenTypeGe:      ">=",
	TokenTypeLe:      "<=",
	TokenTypeEq:      "==",
	TokenTypeNe:      "!=",
	TokenTypeBitAnd:  "&",
	TokenTypeBitOr:   "|",
	TokenTypeBitXor:  "^",
	TokenTypeAnd:     "&&",
	TokenTypeOr:      "||",
	TokenTypeQst:     "?",
	TokenTypeCol:     ":",
}

type Pos struct {
	Line   int
	Column int
}

func (p Pos) String() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Column)
}

type Token struct {
	Type  TokenType
	Value string
	Start Pos
	End   Pos
}

func (t *Token) String() string {
	return fmt.Sprintf("{Type=%v, Value=%q, Start=%v, End=%v}",
		t.Type, t.Value, t.Start, t.End)
}

func Lex(input string) (tokens []*Token, err error) {
	r := bufio.NewReader(strings.NewReader(input))
	var (
		token *Token
		l     = &lexer{r: r, line: 1, column: 1}
	)
	for {
		token, err = l.getToken()
		if err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
			}
			return
		}
		tokens = append(tokens, token)
	}
}

type lexer struct {
	r      *bufio.Reader
	start  Pos
	line   int
	column int
	lastCo int
}

// state 状态机
type state int

const (
	stateNormal state = iota
	stateID           // 标识符
	stateInt          // 整数
	statePlus         // + ++
	stateMinus        // - --
	stateGt           // > >> >=
	stateLt           // < << <=
	stateNot          // ! !=
	stateAnd          // & &&
	stateOr           // | ||
	stateEq           // ==
)

func (l *lexer) getToken() (*Token, error) {
	var (
		s  state
		sb strings.Builder
	)
	l.start = Pos{Line: l.line, Column: l.column}
	for {
		// 读取一个字符
		ch, _, err := l.read()
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		} // EOF 也继续

		// 根据当前状态继续分析
		switch s {
		case stateNormal: // 初始状态
			if errors.Is(err, io.EOF) {
				return nil, err // 输入结束
			}
			if unicode.IsSpace(ch) {
				l.start = Pos{Line: l.line, Column: l.column}
				continue // 跳过空白字符
			}
			sb.WriteRune(ch)
			switch {
			case isAlpha(ch): // 读取到字母进入标识符状态
				s = stateID
				continue
			case isDigit(ch): // 读取到数字进入整数状态
				s = stateInt
				continue
			}
			switch ch {
			case '+':
				s = statePlus
			case '-':
				s = stateMinus
			case '>':
				s = stateGt
			case '<':
				s = stateLt
			case '!':
				s = stateNot
			case '&':
				s = stateAnd
			case '|':
				s = stateOr
			case '=':
				s = stateEq
			case '(':
				return l.newToken(TokenTypeParenL, "("), nil
			case ')':
				return l.newToken(TokenTypeParenR, ")"), nil
			case '*':
				return l.newToken(TokenTypeTimes, "*"), nil
			case '/':
				return l.newToken(TokenTypeOver, "/"), nil
			case '%':
				return l.newToken(TokenTypeMod, "%"), nil
			case '~':
				return l.newToken(TokenTypeBitNot, "~"), nil
			case '^':
				return l.newToken(TokenTypeBitXor, "^"), nil
			case '?':
				return l.newToken(TokenTypeQst, "?"), nil
			case ':':
				return l.newToken(TokenTypeCol, ":"), nil
			default:
				return nil, fmt.Errorf("unexpected char: %q", ch)
			}
		case stateID: // 标识符状态
			if isAlpha(ch) || isDigit(ch) {
				sb.WriteRune(ch)
			} else { // 已经不是标识符了
				l.unread(ch)
				return l.newToken(TokenTypeID, sb.String()), nil
			}
		case stateInt: // 整数
			if isDigit(ch) {
				sb.WriteRune(ch)
			} else {
				l.unread(ch)
				return l.newToken(TokenTypeInt, sb.String()), nil
			}
		case statePlus: // + ++
			if ch == '+' {
				return l.newToken(TokenTypeIncr, "++"), nil
			}
			l.unread(ch)
			return l.newToken(TokenTypePlus, "+"), nil
		case stateMinus: // - --
			if ch == '-' {
				return l.newToken(TokenTypeDecr, "--"), nil
			}
			l.unread(ch)
			return l.newToken(TokenTypeMinus, "-"), nil
		case stateGt: // > >= >>
			switch ch {
			case '>':
				return l.newToken(TokenTypeShiftR, ">>"), nil
			case '=':
				return l.newToken(TokenTypeGe, ">="), nil
			default:
				l.unread(ch)
				return l.newToken(TokenTypeGt, ">"), nil
			}
		case stateLt: // < << <=
			switch ch {
			case '<':
				return l.newToken(TokenTypeShiftL, "<<"), nil
			case '=':
				return l.newToken(TokenTypeLe, "<="), nil
			default:
				l.unread(ch)
				return l.newToken(TokenTypeLt, "<"), nil
			}
		case stateNot: // ! !=
			if ch == '=' {
				return l.newToken(TokenTypeNe, "!="), nil
			}
			l.unread(ch)
			return l.newToken(TokenTypeNot, "!"), nil
		case stateAnd: // & &&
			if ch == '&' {
				return l.newToken(TokenTypeAnd, "&&"), nil
			}
			l.unread(ch)
			return l.newToken(TokenTypeBitAnd, "&"), nil
		case stateOr: // | ||
			if ch == '|' {
				return l.newToken(TokenTypeOr, "||"), nil
			}
			l.unread(ch)
			return l.newToken(TokenTypeBitOr, "|"), nil
		case stateEq:
			if ch == '=' {
				return l.newToken(TokenTypeEq, "=="), nil
			}
			return nil, fmt.Errorf("unexpected char: %q", ch)
		default:
			return nil, fmt.Errorf("unexpected state: %v", s)
		}
	}
}

func (l *lexer) newToken(typ TokenType, val string) *Token {
	return &Token{
		Type:  typ,
		Value: val,
		Start: l.start,
		End:   Pos{Line: l.line, Column: l.column},
	}
}

func (l *lexer) read() (ch rune, size int, err error) {
	ch, size, err = l.r.ReadRune()
	debug("read: %q, err=%v", ch, err)
	if err != nil {
		return
	}
	if ch == '\n' {
		l.line++
		l.lastCo = l.column
		l.column = 1
	} else {
		l.column++
	}
	return
}

func (l *lexer) unread(ch rune) {
	if ch == 0 {
		return
	}
	debug("unread: %q", ch)
	l.r.UnreadRune()
	if ch == '\n' {
		l.line--
		l.column = l.lastCo
	} else {
		l.column--
	}
}

func debug(msg string, args ...any) {
	fmt.Printf(msg+"\n", args...)
}

func isAlpha(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

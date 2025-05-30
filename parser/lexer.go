// Package parser implements parser for anko.
package parser

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/mattn/anko/ast"
)

const (
	// EOF is short for End of file.
	EOF = -1
	// EOL is short for End of line.
	EOL = '\n'
	// ILLEGAL is used for illegal tokens.
	ILLEGAL = -2
)

// Error is a parse error.
type Error struct {
	Message  string
	Pos      ast.Position
	Filename string
	Fatal    bool
}

// Error returns the parse error message.
func (e *Error) Error() string {
	return e.Message
}

// Scanner stores informations for lexer.
type Scanner struct {
	src      []rune
	offset   int
	lineHead int
	line     int
}

// opName is correction of operation names.
var opName = map[string]int{
	"func":     FUNC,
	"return":   RETURN,
	"var":      VAR,
	"throw":    THROW,
	"if":       IF,
	"for":      FOR,
	"break":    BREAK,
	"continue": CONTINUE,
	"in":       IN,
	"else":     ELSE,
	"new":      NEW,
	"true":     TRUE,
	"false":    FALSE,
	"nil":      NIL,
	"module":   MODULE,
	"try":      TRY,
	"catch":    CATCH,
	"finally":  FINALLY,
	"switch":   SWITCH,
	"case":     CASE,
	"default":  DEFAULT,
	"go":       GO,
	"chan":     CHAN,
	"struct":   STRUCT,
	"make":     MAKE,
	"type":     TYPE,
	"len":      LEN,
	"delete":   DELETE,
	"close":    CLOSE,
	"map":      MAP,
	"import":   IMPORT,
}

var (
	nilValue   = reflect.New(reflect.TypeOf((*interface{})(nil)).Elem()).Elem()
	trueValue  = reflect.ValueOf(true)
	falseValue = reflect.ValueOf(false)
	oneLiteral = &ast.LiteralExpr{Literal: reflect.ValueOf(int64(1))}
)

// Init resets code to scan.
func (s *Scanner) Init(src string) {
	s.src = []rune(src)
}

// Scan analyses token, and decide identify or literals.
func (s *Scanner) Scan() (tok int, lit string, pos ast.Position, err error) {
retry:
	s.skipBlank()
	pos = s.pos()
	switch ch := s.peek(); {
	case isLetter(ch):
		lit, err = s.scanIdentifier()
		if err != nil {
			return
		}
		if name, ok := opName[lit]; ok {
			tok = name
		} else {
			tok = IDENT
		}
	case isDigit(ch):
		tok = NUMBER
		lit, err = s.scanNumber()
		if err != nil {
			return
		}
	case ch == '"':
		tok = STRING
		lit, err = s.scanString('"')
		if err != nil {
			return
		}
	case ch == '\'':
		tok = STRING
		lit, err = s.scanString('\'')
		if err != nil {
			return
		}
	case ch == '`':
		tok = STRING
		lit, err = s.scanRawString('`')
		if err != nil {
			return
		}
	default:
		switch ch {
		case EOF:
			tok = EOF
		case '#':
			for !isEOL(s.peek()) {
				s.next()
			}
			goto retry
		case '!':
			s.next()
			switch s.peek() {
			case '=':
				tok = NEQ
				lit = "!="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '=':
			s.next()
			switch s.peek() {
			case '=':
				tok = EQEQ
				lit = "=="
			case ' ':
				if s.peekPlus(1) == '<' && s.peekPlus(2) == '-' {
					s.next()
					s.next()
					tok = EQOPCHAN
					lit = "= <-"
				} else {
					s.back()
					tok = int(ch)
					lit = string(ch)
				}
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '?':
			s.next()
			switch s.peek() {
			case '?':
				tok = NILCOALESCE
				lit = "??"
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '+':
			s.next()
			switch s.peek() {
			case '+':
				tok = PLUSPLUS
				lit = "++"
			case '=':
				tok = PLUSEQ
				lit = "+="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '-':
			s.next()
			switch s.peek() {
			case '-':
				tok = MINUSMINUS
				lit = "--"
			case '=':
				tok = MINUSEQ
				lit = "-="
			default:
				s.back()
				tok = int(ch)
				lit = "-"
			}
		case '*':
			s.next()
			switch s.peek() {
			case '=':
				tok = MULEQ
				lit = "*="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '/':
			s.next()
			switch s.peek() {
			case '=':
				tok = DIVEQ
				lit = "/="
			case '/':
				for !isEOL(s.peek()) {
					s.next()
				}
				goto retry
			case '*':
				for {
					_, err = s.scanRawString('*')
					if err != nil {
						return
					}

					if s.peek() == '/' {
						s.next()
						goto retry
					}

					s.back()
				}
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '>':
			s.next()
			switch s.peek() {
			case '=':
				tok = GE
				lit = ">="
			case '>':
				tok = SHIFTRIGHT
				lit = ">>"
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '<':
			s.next()
			switch s.peek() {
			case '-':
				tok = OPCHAN
				lit = "<-"
			case '=':
				tok = LE
				lit = "<="
			case '<':
				tok = SHIFTLEFT
				lit = "<<"
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '|':
			s.next()
			switch s.peek() {
			case '|':
				tok = OROR
				lit = "||"
			case '=':
				tok = OREQ
				lit = "|="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '&':
			s.next()
			switch s.peek() {
			case '&':
				tok = ANDAND
				lit = "&&"
			case '=':
				tok = ANDEQ
				lit = "&="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '.':
			s.next()
			if s.peek() == '.' {
				s.next()
				if s.peek() == '.' {
					tok = VARARG
				} else {
					err = fmt.Errorf("syntax error on '%v' at %v:%v", string(ch), pos.Line, pos.Column)
					return
				}
			} else {
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '\n', '(', ')', ':', ';', '%', '{', '}', '[', ']', ',', '^':
			tok = int(ch)
			lit = string(ch)
		default:
			err = fmt.Errorf("syntax error on '%v' at %v:%v", string(ch), pos.Line, pos.Column)
			tok = int(ch)
			lit = string(ch)
			return
		}
		s.next()
	}
	return
}

// isLetter returns true if the rune is a letter for identity.
func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

// isDigit returns true if the rune is a number.
func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

// isHex returns true if the rune is a hex digits.
func isHex(ch rune) bool {
	return ('0' <= ch && ch <= '9') || ('a' <= ch && ch <= 'f') || ('A' <= ch && ch <= 'F')
}

// isEOL returns true if the rune is at end-of-line or end-of-file.
func isEOL(ch rune) bool {
	return ch == '\n' || ch == -1
}

// isBlank returns true if the rune is empty character..
func isBlank(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\r'
}

// peek returns current rune in the code.
func (s *Scanner) peek() rune {
	if s.reachEOF() {
		return EOF
	}
	return s.src[s.offset]
}

// peek returns current rune plus i in the code.
func (s *Scanner) peekPlus(i int) rune {
	if len(s.src) <= s.offset+i {
		return EOF
	}
	return s.src[s.offset+i]
}

// next moves offset to next.
func (s *Scanner) next() {
	if !s.reachEOF() {
		if s.peek() == '\n' {
			s.lineHead = s.offset + 1
			s.line++
		}
		s.offset++
	}
}

// current returns the current offset.
func (s *Scanner) current() int {
	return s.offset
}

// offset sets the offset value.
func (s *Scanner) set(o int) {
	s.offset = o
}

// back moves back offset once to top.
func (s *Scanner) back() {
	s.offset--
}

// reachEOF returns true if offset is at end-of-file.
func (s *Scanner) reachEOF() bool {
	return len(s.src) <= s.offset
}

// pos returns the position of current.
func (s *Scanner) pos() ast.Position {
	return ast.Position{Line: s.line + 1, Column: s.offset - s.lineHead + 1}
}

// skipBlank moves position into non-black character.
func (s *Scanner) skipBlank() {
	for isBlank(s.peek()) {
		s.next()
	}
}

// scanIdentifier returns identifier beginning at current position.
func (s *Scanner) scanIdentifier() (string, error) {
	var ret []rune
	for {
		if !isLetter(s.peek()) && !isDigit(s.peek()) {
			break
		}
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret), nil
}

// scanNumber returns number beginning at current position.
func (s *Scanner) scanNumber() (string, error) {
	result := []rune{s.peek()}
	s.next()

	if result[0] == '0' && (s.peek() == 'x' || s.peek() == 'X') {
		// hex
		result = append(result, 'x')
		s.next()
		for isHex(s.peek()) {
			result = append(result, s.peek())
			s.next()
		}
	} else {
		// non-hex
		found := false
		for {
			if isDigit(s.peek()) {
				// is digit
				result = append(result, s.peek())
				s.next()
				continue
			}

			if s.peek() == '.' {
				// is .
				result = append(result, '.')
				s.next()
				continue
			}

			if s.peek() == 'e' || s.peek() == 'E' {
				// is e
				if found {
					return "", errors.New("unexpected " + string(s.peek()))
				}
				found = true
				s.next()

				// check if + or -
				if s.peek() == '+' || s.peek() == '-' {
					// add e with + or -
					result = append(result, 'e')
					result = append(result, s.peek())
					s.next()
				} else {
					// add e, but next char not + or -
					result = append(result, 'e')
				}
				continue
			}

			// not digit, e, nor .
			break
		}
	}

	if isLetter(s.peek()) {
		return "", errors.New("identifier starts immediately after numeric literal")
	}

	return string(result), nil
}

// scanRawString returns raw-string starting at current position.
func (s *Scanner) scanRawString(l rune) (string, error) {
	var ret []rune
	for {
		s.next()
		if s.peek() == EOF {
			return "", errors.New("unexpected EOF")
		}
		if s.peek() == l {
			s.next()
			break
		}
		ret = append(ret, s.peek())
	}
	return string(ret), nil
}

// scanString returns string starting at current position.
// This handles backslash escaping.
func (s *Scanner) scanString(l rune) (string, error) {
	var ret []rune
eos:
	for {
		s.next()
		switch s.peek() {
		case EOL:
			return "", errors.New("unexpected EOL")
		case EOF:
			return "", errors.New("unexpected EOF")
		case l:
			s.next()
			break eos
		case '\\':
			s.next()
			switch s.peek() {
			case 'b':
				ret = append(ret, '\b')
				continue
			case 'f':
				ret = append(ret, '\f')
				continue
			case 'r':
				ret = append(ret, '\r')
				continue
			case 'n':
				ret = append(ret, '\n')
				continue
			case 't':
				ret = append(ret, '\t')
				continue
			case 'u':
				tok, val := s.scanEscape()
				if tok == ILLEGAL {
					return "", errors.New(val.(string))
				}
				ret = append(ret, []rune(val.(string))...)
				continue
			}
			ret = append(ret, s.peek())
			continue
		default:
			ret = append(ret, s.peek())
		}
	}
	return string(ret), nil
}

// scanEscape handles escape sequences like \uXXXX
func (s *Scanner) scanEscape() (tok int, val interface{}) {
	if s.offset >= len(s.src) { // Fix invalid operation
		return ILLEGAL, "incomplete escape sequence"
	}

	c := s.src[s.offset]
	s.offset++

	if c != 'u' {
		switch c {
		case 'n':
			return STRING, "\n"
		case 't':
			return STRING, "\t"
		case 'r':
			return STRING, "\r"
		case '\\':
			return STRING, "\\"
		default:
			return ILLEGAL, fmt.Sprintf("invalid escape \\%c", c)
		}
	}

	if s.offset+4 > len(s.src) { // Fix invalid operation
		return ILLEGAL, "incomplete Unicode escape"
	}

	hex := string(s.src[s.offset : s.offset+4])
	s.offset += 4

	for _, h := range hex {
		if !strings.ContainsRune("0123456789ABCDEFabcdef", h) {
			return ILLEGAL, fmt.Sprintf("invalid Unicode escape \\u%s", hex)
		}
	}

	n, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return ILLEGAL, fmt.Sprintf("invalid Unicode escape \\u%s: %v", hex, err)
	}

	r := rune(n)
	if !utf8.ValidRune(r) {
		return ILLEGAL, fmt.Sprintf("invalid Unicode rune \\u%s", hex)
	}

	return STRING, string(r)
}

// Lexer provides interface to parse codes.
type Lexer struct {
	s    *Scanner
	lit  string
	pos  ast.Position
	e    error
	stmt ast.Stmt
}

// Lex scans the token and literals.
func (l *Lexer) Lex(lval *yySymType) int {
	tok, lit, pos, err := l.s.Scan()
	if err != nil {
		l.e = &Error{Message: err.Error(), Pos: pos, Fatal: true}
	}
	lval.tok = ast.Token{Tok: tok, Lit: lit}
	lval.tok.SetPosition(pos)
	l.lit = lit
	l.pos = pos
	return tok
}

// Error sets parse error.
func (l *Lexer) Error(msg string) {
	l.e = &Error{Message: msg, Pos: l.pos, Fatal: false}
}

// Parse provides way to parse the code using Scanner.
func Parse(s *Scanner) (ast.Stmt, error) {
	l := Lexer{s: s}
	if yyParse(&l) != 0 {
		return nil, l.e
	}
	return l.stmt, l.e
}

// EnableErrorVerbose enabled verbose errors from the parser
func EnableErrorVerbose() {
	yyErrorVerbose = true
}

// EnableDebug enabled debug from the parser
func EnableDebug(level int) {
	yyDebug = level
}

// ParseSrc provides way to parse the code from source.
func ParseSrc(src string) (ast.Stmt, error) {
	scanner := &Scanner{
		src: []rune(src),
	}
	return Parse(scanner)
}

func toNumber(numString string) (reflect.Value, error) {
	// hex
	if len(numString) > 2 && numString[0:2] == "0x" {
		i, err := strconv.ParseInt(numString[2:], 16, 64)
		if err != nil {
			return nilValue, err
		}
		return reflect.ValueOf(i), nil
	}

	// hex
	if len(numString) > 3 && numString[0:3] == "-0x" {
		i, err := strconv.ParseInt("-"+numString[3:], 16, 64)
		if err != nil {
			return nilValue, err
		}
		return reflect.ValueOf(i), nil
	}

	// float
	if strings.Contains(numString, ".") || strings.Contains(numString, "e") {
		f, err := strconv.ParseFloat(numString, 64)
		if err != nil {
			return nilValue, err
		}
		return reflect.ValueOf(f), nil
	}

	// int
	i, err := strconv.ParseInt(numString, 10, 64)
	if err != nil {
		return nilValue, err
	}
	return reflect.ValueOf(i), nil
}

func stringToValue(aString string) reflect.Value {
	return reflect.ValueOf(aString)
}

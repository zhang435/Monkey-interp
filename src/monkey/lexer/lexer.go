package lexer

import (
	"fmt"
	"monkey/token"
)

// Lexer : lexer object
type Lexer struct {
	input        string
	position     int  // current position of the input
	readPosition int  // current reading position in input
	ch           byte // current char
}

// New : function to specify the new input
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar : read in the input
// input will read as whole and been chop into piece
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

// NextToken  : get the next Token
// convert the next char to token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()
	fmt.Print(l.position)
	switch l.ch {
	case '=':
		if '=' == l.peekChar() {
			l.readChar()
			tok = token.Token{token.EQ, "=="}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{token.NEQ, "!="}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case 0:
		tok = token.Token{token.EOF, ""}
	default:
		if isLetter(l.ch) {
			ident := l.readIdentifier()
			t := token.LookupIdent(ident)
			tok = token.Token{t, ident}
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNum()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// readIdentifier : read in the char until it's reaching non letter char
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	fmt.Println(l.input[position:l.position])
	return l.input[position:l.position]
}

// newToken contrcut to create a new Token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{tokenType, string(ch)}
}

// isLetter : check if the even char is a letter
func isLetter(chr byte) bool {
	return (chr >= 'a' && chr <= 'z') || (chr >= 'A' && chr <= 'Z') || chr == '_'
}

//  isDigit : check if the given char is a number
func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func (l *Lexer) readNum() string {
	posn := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[posn:l.position]
}

// skipWhitespace : used to skip white space
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// return the char at the top
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

package lexer

// Traverse to each input string and convert them to series of token

import (
	"github.com/slier81/monkey/token"
)

type Lexer struct {
	input           string
	currentPosition int  // current position in input (points to current char)
	readPosition    int  // current reading position in input (point to the character after the current char)
	character       byte // current char under examination
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhiteSpace()
	var tok token.Token

	switch l.character {
	// Token using single character
	case ';':
		tok = newToken(token.SEMICOLON, l.character)
	case '(':
		tok = newToken(token.LPAREN, l.character)
	case ')':
		tok = newToken(token.RPAREN, l.character)
	case ',':
		tok = newToken(token.COMMA, l.character)
	case '+':
		tok = newToken(token.PLUS, l.character)
	case '{':
		tok = newToken(token.LBRACE, l.character)
	case '}':
		tok = newToken(token.RBRACE, l.character)
	case '-':
		tok = newToken(token.MINUS, l.character)
	case '/':
		tok = newToken(token.SLASH, l.character)
	case '*':
		tok = newToken(token.ASTERISK, l.character)
	case '<':
		tok = newToken(token.LT, l.character)
	case '>':
		tok = newToken(token.GT, l.character)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""

	// Token using double character
	case '=':
		if l.peekChar() == '=' {
			ch := l.character
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.character)}
		} else {
			tok = newToken(token.ASSIGN, l.character)
		}

	case '!':
		if l.peekChar() == '=' {
			ch := l.character
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.character)}
		} else {
			tok = newToken(token.BANG, l.character)
		}

	// Token using more than 2 character
	default:
		if isLetter(l.character) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		}

		if isDigit(l.character) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		}

		tok = newToken(token.ILLEGAL, l.character)
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.character = 0
	} else {
		l.character = l.input[l.readPosition]
	}

	l.currentPosition = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	// Store starting position of the identifier
	position := l.currentPosition

	// 0123456789
	// ||||||||||||||||
	// let name = "joe"

	for isLetter(l.character) {
		l.readChar()
	}

	return l.input[position:l.currentPosition]
}

func (l *Lexer) skipWhiteSpace() {
	// Advance the position to the non whitespace character
	for l.character == ' ' || l.character == '\t' || l.character == '\n' || l.character == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	// Store starting position of the identifier
	position := l.currentPosition

	for isDigit(l.character) {
		l.readChar()
	}

	return l.input[position:l.currentPosition]
}

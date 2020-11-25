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

func (lex *Lexer) NextToken() token.Token {
	lex.skipWhiteSpace()
	var tok token.Token

	switch lex.character {
	// Token using single character
	case ';':
		tok = newToken(token.SEMICOLON, lex.character)
	case '(':
		tok = newToken(token.LPAREN, lex.character)
	case ')':
		tok = newToken(token.RPAREN, lex.character)
	case ',':
		tok = newToken(token.COMMA, lex.character)
	case '+':
		tok = newToken(token.PLUS, lex.character)
	case '{':
		tok = newToken(token.LBRACE, lex.character)
	case '}':
		tok = newToken(token.RBRACE, lex.character)
	case '-':
		tok = newToken(token.MINUS, lex.character)
	case '/':
		tok = newToken(token.SLASH, lex.character)
	case '*':
		tok = newToken(token.ASTERISK, lex.character)
	case '<':
		tok = newToken(token.LT, lex.character)
	case '>':
		tok = newToken(token.GT, lex.character)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""

	// Token using double character
	case '=':
		if lex.peekChar() == '=' {
			ch := lex.character
			lex.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(lex.character)}
		} else {
			tok = newToken(token.ASSIGN, lex.character)
		}

	case '!':
		if lex.peekChar() == '=' {
			ch := lex.character
			lex.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(lex.character)}
		} else {
			tok = newToken(token.BANG, lex.character)
		}

	// Token using more than 2 character
	default:
		if isLetter(lex.character) {
			tok.Literal = lex.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		}

		if isDigit(lex.character) {
			tok.Literal = lex.readNumber()
			tok.Type = token.INT
			return tok
		}

		tok = newToken(token.ILLEGAL, lex.character)
	}

	lex.readChar()
	return tok
}

func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		lex.character = 0
	} else {
		lex.character = lex.input[lex.readPosition]
	}

	lex.currentPosition = lex.readPosition
	lex.readPosition++
}

func (lex *Lexer) peekChar() byte {
	if lex.readPosition >= len(lex.input) {
		return 0
	}

	return lex.input[lex.readPosition]
}

func (lex *Lexer) readIdentifier() string {
	// Store starting position of the identifier
	position := lex.currentPosition

	// 0123456789
	// ||||||||||||||||
	// let name = "joe"

	for isLetter(lex.character) {
		lex.readChar()
	}

	return lex.input[position:lex.currentPosition]
}

func (lex *Lexer) skipWhiteSpace() {
	// Advance the position to the non whitespace character
	for lex.character == ' ' || lex.character == '\t' || lex.character == '\n' || lex.character == '\r' {
		lex.readChar()
	}
}

func (lex *Lexer) readNumber() string {
	// Store starting position of the identifier
	position := lex.currentPosition

	for isDigit(lex.character) {
		lex.readChar()
	}

	return lex.input[position:lex.currentPosition]
}

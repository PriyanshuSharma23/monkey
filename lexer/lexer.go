package lexer

import "github.com/PriyanshuSharma23/monkey/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar() // set the position, readPosition, and ch
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readSequence(cond func(byte) bool) string {
	postion := l.position
	for cond(l.ch) {
		l.readChar()
	}
	return l.input[postion:l.position]
}

func (l *Lexer) readIdentifier() string {
	return l.readSequence(isLetter)
}

func (l *Lexer) readNumber() string {
	return l.readSequence(isDigit)
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = token.New(token.ASSIGN, string(l.ch))
	case ',':
		tok = token.New(token.COMMA, string(l.ch))
	case '+':
		tok = token.New(token.PLUS, string(l.ch))
	case ';':
		tok = token.New(token.SEMICOLON, string(l.ch))
	case '{':
		tok = token.New(token.LBRACE, string(l.ch))
	case '}':
		tok = token.New(token.RBRACE, string(l.ch))
	case '(':
		tok = token.New(token.LPAREN, string(l.ch))
	case ')':
		tok = token.New(token.RPAREN, string(l.ch))
	case '!':
		tok = token.New(token.BANG, string(l.ch))
	case '-':
		tok = token.New(token.MINUS, string(l.ch))
	case '*':
		tok = token.New(token.ASTERISK, string(l.ch))
	case '/':
		tok = token.New(token.SLASH, string(l.ch))
	case '<':
		tok = token.New(token.LT, string(l.ch))
	case '>':
		tok = token.New(token.GT, string(l.ch))
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = token.New(token.ILLEGAL, string(l.ch))
		}
	}

	l.readChar()
	return tok
}

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

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
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
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.EQ, string([]byte{ch, l.ch}))
		} else {
			tok = token.New(token.ASSIGN, string(l.ch))
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.NOT_EQ, string([]byte{ch, l.ch}))
		} else {
			tok = token.New(token.BANG, string(l.ch))
		}
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
	case 0:
		tok = token.New(token.EOF, string(l.ch))
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

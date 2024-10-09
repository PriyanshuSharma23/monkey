package lexer

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\n' || ch == '\t' || ch == '\r'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
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

func (l *Lexer) peekChar() byte {
	var ch byte

	if l.readPosition >= len(l.input) {
		ch = 0
	} else {
		ch = l.input[l.readPosition]
	}

	return ch
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

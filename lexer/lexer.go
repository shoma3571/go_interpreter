package lexer

import "github.com/shoma3571/go_interpreter/token"

type lexer struct {
	input string
	position int // 入力における現在の位置(現在の位置を指し示す)
	readPosition int // これから読み込む位置(現在の文字の次)
	ch byte // 現在検査中の文字
}

func New(input string) *lexer {
	// inputのみを定義して、他はゼロ値で設定
	l := &lexer{input: input}
	// とりあえず最初の文字を読んでおく
	l.readChar()
	return l
}

// ポインタメソッド
// 次の一文字を読んで、現在位置を進める
func (l *lexer) readChar() {
	// 次に読み込むものがあるかないかを判定
	if l.readPosition >= len(l.input) {
		// 終端に到達した場合 0 にする
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	// 位置の更新
	l.position = l.readPosition
	l.readPosition += 1
}


// 現在検査中の文字 l.ch を見て、それに応じてトークンを返す
// 返す前に入力のポインタを進めて、次に読んだ時に位置が更新されるようにする
func (l *lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

// token初期化の役割を果たす
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
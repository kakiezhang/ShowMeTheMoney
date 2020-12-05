package token

type Token struct {
	Type string
	Text string
}

// token 流
type TokenReader interface {
	Peek() *Token        // 只看不消耗
	Read() *Token        // 消耗一个 token
	Unread()             // 将消耗的 token 退回去
	GetPosition() int    // 获取消耗点
	SetPosition(pos int) // 设置消耗点
}

// Implement TokenReader
type SimpleTokenReader struct {
	pos    int
	tokens []*Token
}

func NewSimpleTokenReader(tokens []*Token) *SimpleTokenReader {
	return &SimpleTokenReader{
		tokens: tokens,
	}
}

func (str *SimpleTokenReader) Peek() *Token {
	if str.pos >= len(str.tokens) {
		return nil
	}
	return str.tokens[str.pos]
}

func (str *SimpleTokenReader) Read() *Token {
	if str.pos >= len(str.tokens) {
		return nil
	}
	token := str.tokens[str.pos]
	str.pos += 1
	return token
}

func (str *SimpleTokenReader) Unread() {
	if str.pos > 0 {
		str.pos -= 1
	}
}

func (str *SimpleTokenReader) GetPosition() int {
	return str.pos
}

func (str *SimpleTokenReader) SetPosition(pos int) {
	if str.pos >= 0 && str.pos <= len(str.tokens) {
		str.pos = pos
	}
}

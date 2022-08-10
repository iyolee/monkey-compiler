package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// 遇到未知的词法单元或字符
	ILLEGAL = "ILLEGAL"
	// 文件结尾 End Of File
	EOF = "EOF"

	// 标识符+字面量
	IDENT = "IDENT"
	INT   = "INT"

	// 运算符
	ASSIGN = "="
	PLUS   = "+"

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

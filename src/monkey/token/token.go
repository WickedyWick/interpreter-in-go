package token

type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"
    
    // Identifiers + Literals
    IDENT = "IDENT" // add, foobar, x, y 
    INT = "INT" // 123456 
    
    // Operators
    ASSIGN = "="
    PLUS = "+"
    
    // Delimiters
    COMMA = ","
    SEMICOLON = ";"
    
    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"
    
    // Keywords
    FUNCTION = "FUNCTION"
    LET = "LET" 
)

var keywords = map[string]TokenType{
  "fn": FUNCTION,
  "let": LET,
}

func LookupIdent(ident string) TokenType {
  // We declare tokenType (tok) variable and "ok" is true if value exists and false if it doesnt
  // with so basically we first declare variables and then check for values after semicolon
  // if its not one of our keywords its identifier (variable name) 
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENT
}
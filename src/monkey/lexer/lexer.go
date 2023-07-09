package lexer

import "monkey/token"
type Lexer struct {
  input           string
  position        int  // current position inm input (points to current char
  readPosition    int  // current reading position in input (after current char)
  ch              byte // current char under examination
}

func New(input string) *Lexer {
  l := &Lexer{input: input}
  l.readChar()
  return l
}

func (l *Lexer) NextToken() token.Token {
  var tok token.Token
  // Need to skip whitespace because in our language its only used as separator of tokens and doesn't have a meaning 
  l.skipWhitespace()
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
    default: 
      if isLetter(l.ch) {
        tok.Literal = l.readIdentifier()
        return tok
      } else {
        tok = newToken(token.ILLEGAL, l.ch)
      }
  }
  
  l.readChar()
  return tok
} 

func (l *Lexer) readChar() {
  if l.readPosition >= len(l.input) {
      // 0 is NUL ASCII representation so we can use it for EOF
    l.ch = 0
  } else {
    l.ch = l.input[l.readPosition]
  }
  l.position = l.readPosition
  l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
  return token.Token{Type: tokenType, Literal: string(ch)}
}

// We could probably generalize this by passing in the character`s identifying functions as arguments, but won’t,
// for simplicity’s sake and ease of understanding ----------- readIdentifier & readNumber
func (l *Lexer) readIdentifier() string {
  position := l.position
  for isLetter(l.ch) {
    l.readChar()
  }
  return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
  position := l.position
  for isDigit(l.ch) {
    l.readChar()
  }
  return l.input[position:l.position]
}
func isLetter(ch byte) bool {
  return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
  return '0' <= ch && ch <= '9'
}
func skipWhitespace(l* Lexer) {
  // I think in go you can write conditional for loop or I guess this is whileDoLoop
  for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
    l.readChar() 
  }
}

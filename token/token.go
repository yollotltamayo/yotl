package token


type TokenType string

type Token struct{
    Type TokenType
    Literal string
}
const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    //variables
    IDENT = "IDENT" //x , y,k ...
    INT = "INT" //numbers

    //operadores

    ASSIGN   = "="
    PLUS     = "+"
    MINUS    = "-"
    BANG     = "!"
    ASTERISK = "*"
    SLASH    = "/"

    LT = "<"
    GT = ">"

    EQ     = "=="
    NOT_EQ = "!="

    //delimitadores
    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"
    //key-words
    FUNCTION = "FUNCTION"
    LET = "LET"
    RETURN = "RETURN"
    IF = "IF"
    ELSE = "ELSE"
    TRUE = "TRUE"
    FALSE = "FALSE"
)
var keywords = map[string]TokenType{
    "fn" : FUNCTION,
    "let" : LET,
    "if" :IF,
    "return":RETURN,
    "else" :ELSE,
    "false":FALSE,
    "true":TRUE,
}
func LookUpIdent(ident string) TokenType{
    if tok,ok := keywords[ident];ok{ //asignar y preguntar ; pregunta
        return  tok
    }
    return IDENT
}



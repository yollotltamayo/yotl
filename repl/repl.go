package repl

import (
    "bufio"
    "io"
    "fmt"
    "yotl/lexer"
    "yotl/token"
)

const PROMPT = ">>>"
func Start(in io.Reader ,out io.Writer){
    scanner := bufio.NewScanner(in)
    fmt.Printf("S N E L L Y by F.Y.T.H released under M.I.T. license c2020-2020\n")
    for {
        fmt.Printf(PROMPT)
        scanned := scanner.Scan()
        if !scanned {
            return
        }
        line := scanner.Text()
        l := lexer.New(line)
        for tok := l.NextToken();tok.Type !=  token.EOF ;tok = l.NextToken() {
            fmt.Printf("%+v\n",tok)
        }
    }
}


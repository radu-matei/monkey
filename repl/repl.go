package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/radu-matei/monkey/lexer"
	"github.com/radu-matei/monkey/token"
)

const prompt = ">>"

// Start starts a REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(prompt)

		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

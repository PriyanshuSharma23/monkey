package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/PriyanshuSharma23/monkey/lexer"
	"github.com/PriyanshuSharma23/monkey/token"
)

const PROMPT = "-> "

func Start(in io.Reader, out io.Writer) {
	reader := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := reader.Scan()

		if !scanned {
			return
		}

		text := reader.Text()

		l := lexer.New(text)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

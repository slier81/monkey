package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/slier81/monkey/lexer"
	"github.com/slier81/monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		result := scanner.Scan()

		if !result {
			return
		}

		text := scanner.Text()
		lex := lexer.New(text)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}

}

package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"uakari/lexer"
	"uakari/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		if line == "quit" || line == "q" {
			fmt.Fprint(out, "Bye bye!\n")
			os.Exit(0)
		}

		l := lexer.New(line)

		for tk := l.NextToken(); tk.Type != token.EOF; tk = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tk)
		}
	}
}

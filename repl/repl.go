package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/vbardakos/fython/lexer"
	"github.com/vbardakos/fython/token"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		lxr := lexer.New(line)

		for tkn := lxr.NextToken(); tkn.Token != token.EOF; tkn = lxr.NextToken() {
			fmt.Fprintf(out, "%v\n", tkn)
		}
	}
}

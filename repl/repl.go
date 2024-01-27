package repl

import (
	"bufio"
	"evasion/evaluator"
	"evasion/lexer"
	"evasion/object"
	"evasion/parser"
	"evasion/token"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	io.WriteString(out, EVASION_LOGO)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		env := object.NewEnvironment()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}

const EVASION_LOGO = `
                           .__               
  _______  _______    _____|__| ____   ____  
_/ __ \  \/ /\__  \  /  ___/  |/  _ \ /    \ 
\  ___/\   /  / __ \_\___ \|  (  <_> )   |  \
 \___  >\_/  (____  /____  >__|\____/|___|  /
     \/           \/     \/               \/ 

`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

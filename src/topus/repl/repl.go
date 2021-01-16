package repl

import (
	"bufio"
	"fmt"
	"io"

	"errors"
	"os"
	"strings"
	"topus/src/topus/evaluator"
	"topus/src/topus/lexer"
	"topus/src/topus/object"
	"topus/src/topus/parser"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		// lineNumber := 1
		// lineNumber = lineNumber + 1

		// prompt := fmt.Sprintf("topus\x1b[91m:\x1b[0m%d\x1b[91m>\x1b[0m ", lineNumber)

		prompt := "topus\x1b[91m:\x1b[0m1\x1b[91m>\x1b[0m "

		fmt.Printf(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "topus") {
			cmd, err := getSimplyCommand(line)
			if err != nil {
				printParserErrors(out, []string{cmd})
				return
			}
			if cmd == "run" {
				filename, err := getFilenameArg(line)
				if err != nil {
					printParserErrors(out, []string{cmd})
					return
				}
				line, err = readSourceFromFile(filename)
				if err != nil {
					printParserErrors(out, []string{cmd})
					return
				}
			}
		}

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluator.Eval(program, env)
		// if evaluator != nil {
		// 	io.WriteString(out, evaluator.Inspect())
		// 	io.WriteString(out, "\n")
		// }
	}
}

func getSimplyCommand(input string) (string, error) {
	arr := strings.Split(input, " ")
	if len(arr) < 2 {
		return "", errors.New("Invalid command supplied")
	}
	cmd := arr[1]

	return cmd, nil
}

func getFilenameArg(input string) (string, error) {
	arr := strings.Split(input, " ")
	if len(arr) < 3 {
		return "", errors.New("No file supplied")
	}
	fileName := arr[2]

	return fileName, nil
}

func readSourceFromFile(fileName string) (content string, err error) {
	f, err := os.Open(fileName)
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		content += scanner.Text()
	}
	return
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

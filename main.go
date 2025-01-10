package main

import (
	"fmt"
	"github.com/eenees/slow/lexer"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename as an argument")
		return
	}

	filename := os.Args[1]

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	lxr := lexer.NewLexer(string(content))

  for t := lxr.NextToken(); t.type != lexer.EOF; tok = lxr.NextToken() {
    fmt.Println(t.Literal)
  }

}

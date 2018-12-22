package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	var files []string

	_ = filepath.Walk("/usr/share/go-1.10/src/fmt", getWalkFunc(&files))

	var fds []*fInfo

	for _, path := range files {

		buf, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		// positions are relative to fset
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, "", buf, parser.ParseComments)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Check if first char name is up (exported function)
		fds = append(fds, inspect(node, buf, true)...)
	}

	term := newLinuxTerminal()

	quiz(fds, term, 10)
}

func getWalkFunc(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
			*files = append(*files, path)
		}

		return nil
	}
}

func isUpper(b byte) bool {
	return b <= 'Z'
}

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/golang"
	ctags "github.com/sourcegraph/go-ctags"
)

func main() {
	filesTxt, err := ioutil.ReadFile("files.txt")
	if err != nil {
		panic(err)
	}
	filenames := strings.Split(strings.TrimSpace(string(filesTxt)), "\n")
	files := map[string][]byte{}
	for _, filename := range filenames {
		contents, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		files[filename] = contents
	}

	ctagsParser, err := ctags.New(ctags.Options{Bin: "ctags"})
	if err != nil {
		panic(err)
	}

	start := time.Now()
	for filename, contents := range files {
		ctagsParser.Parse(filename, contents)
	}
	fmt.Printf("ctags %v\n", time.Since(start))

	sitterParser := sitter.NewParser()
	sitterParser.SetLanguage(golang.GetLanguage())
	start2 := time.Now()
	for _, contents := range files {
		sitterParser.ParseCtx(context.Background(), nil, contents)
	}
	fmt.Printf("tree-sitter %v\n", time.Since(start2))
}

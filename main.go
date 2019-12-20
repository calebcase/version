package main

import "github.com/calebcase/version/cmd"

func main() {
	cmd.Main()
}

/*
import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"

	"github.com/calebcase/version/lib/version"
)

func main() {
	path := "."
	if len(os.Args) == 2 {
		path = os.Args[1]
	}

	r, err := git.PlainOpen(path)
	if err != nil {
		panic(err)
	}

	v, err := version.Repo(r, &version.ByFile{
		BasePath: path,
		FileName: "VERSION",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(v)
}
*/

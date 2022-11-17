package main

import (
	"fmt"

	. "github.com/Xuanwo/gg"
)

func main() {
	f := NewGroup()
	f.AddPackage("main")
	f.NewImport().
		AddPath("fmt")
	f.NewFunction("main").AddBody(
		String(`fmt.Println("%s")`, "Hello, World!"),
	)
	f.NewFunction().AddBody()
	fmt.Println(f.String())
}

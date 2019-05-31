package main

import (
	"github.com/go-modules-by-example/submodules/b"
	"github.com/go-modules-by-example/submodules/c"
	"fmt"
)

const Name = b.Name

func main() {
	fmt.Println(Name)
	fmt.Println(c.Name)
}

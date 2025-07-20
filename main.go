package main

import (
	"fmt"
	"runtime"

	"github.com/lstrumer/godesde0/variables"
	_ "github.com/lstrumer/godesde0/variables"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(variables.Hola())
}

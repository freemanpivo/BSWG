package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProgramName := os.Args
	onlyArgs := os.Args[1:]
	thirdArg := os.Args[3]

	fmt.Println(argsWithProgramName)
	fmt.Println(onlyArgs)
	fmt.Println(thirdArg)
}

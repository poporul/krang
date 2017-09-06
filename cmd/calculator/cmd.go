package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/poporul/calculator"
)

func main() {
	var result int

	if len(os.Args) >= 2 {
		result = calculator.Eval(os.Args[1])
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		source := strings.TrimSpace(string(bytes))
		result = calculator.Eval(source)
	}

	fmt.Println(result)
}

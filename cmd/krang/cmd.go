package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/poporul/krang"
)

func main() {
	var result int

	if len(os.Args) >= 2 {
		result = krang.Eval(os.Args[1])
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		source := strings.TrimSpace(string(bytes))
		result = krang.Eval(source)
	}

	fmt.Println(result)
}

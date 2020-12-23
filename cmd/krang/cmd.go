package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/poporul/krang"
)

func main() {
	var result *krang.Value
	var err error
	g := krang.MathGrammar()
	for k, v := range krang.LogicalGrammar().Operators {
		g.Operators[k] = v
	}
	if len(os.Args) >= 2 {
		result, err = g.Eval(os.Args[1])
	} else {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			bytes, _ := ioutil.ReadAll(os.Stdin)
			source := strings.TrimSpace(string(bytes))
			result, err = g.Eval(source)
		}
	}
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result.Val)
	}
}

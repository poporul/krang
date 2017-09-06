package krang

import "log"

func Process(tokens []Token) int {
	var stack TokenStack

	for _, t := range tokens {
		switch token := t.(type) {
		case Operator:
			second := stack.Pop()
			first := stack.Pop()

			result := token.Calculate(first.(Number), second.(Number))
			stack.Push(result)
		case Number:
			stack.Push(token)
		}
	}

	return stack.Pop().(Number).Value
}

func Eval(value string) int {
	tokens, err := Tokenize(value)
	if err != nil {
		log.Fatal(err)
	}

	rpn := Parse(tokens)
	return Process(rpn)
}

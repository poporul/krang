package krang

import (
	"errors"
)


var (
	UnsupportedOperator = errors.New("unsupported operator")
	InvalidExpression = errors.New("invalid expression")
)

func (g *Grammar) Process(tokens []Token) (*Value, error) {
	var stack TokenStack

	for _, t := range tokens {
		switch token := t.(type) {
		case Operator:
			var result interface{}
			var err error
			switch calc := token.(type) {
			case Unary:
				operand, ok := stack.Pop().(*Value)
				if !ok {
					return nil, InvalidExpression
				}
				result, err = calc.Calculate(*operand)
			case Binary:
				second, ok := stack.Pop().(*Value)
				if !ok {
					return nil, InvalidExpression
				}
				first, ok := stack.Pop().(*Value)
				if !ok {
					return  nil, InvalidExpression
				}
				result, err = calc.Calculate(*first, *second)
			default:
				err = UnsupportedOperator
			}
			if err != nil {
				return nil, err
			}
			stack.Push(result)
		case *Value:
			stack.Push(token)
		}
	}

	if len(stack) > 1 {
		return nil, InvalidExpression
	}
	return stack.Pop().(*Value), nil
}

func (g *Grammar) Eval(value string) (*Value, error) {
	tokens, err := g.Tokenize(value)
	if err != nil {
		return nil, err
	}

	rpn := Parse(tokens)
	return g.Process(rpn)
}

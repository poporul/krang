package krang

import "math"

func NumericBinary(precedence int, associativity Associativity, fn func(int, int) (int, error)) Operator {
	return numericOperator{
		operator: operator{
			precedence:    precedence,
			associativity: associativity,
		},
		fn: fn,
	}
}

type numericOperator struct {
	operator
	fn func(first, second int) (int, error)
}

func asNum(first, second Value) (int, int, error) {
	f, err := first.AsInt()
	if err != nil {
		return 0, 0, err
	}
	s, err := second.AsInt()
	if err != nil {
		return 0, 0, err
	}
	return f, s, nil
}

func (o numericOperator) Calculate(first, second Value) (*Value, error) {
	f, s, err := asNum(first, second)
	if err != nil {
		return nil, err
	}
	res, err := o.fn(f, s)
	if err != nil {
		return nil, err
	}
	return FromNum(res), nil
}

func MathGrammar() Grammar {
	pow := NumericBinary(4, "right", func(first int, second int) (int, error) {
		return int(math.Pow(float64(first), float64(second))), nil
	})
	mult := NumericBinary(3, "left", func(first int, second int) (int, error) {
		return first * second, nil
	})
	div := NumericBinary(3, "left", func(first int, second int) (int, error) {
		if second == 0 {
			return 0, DivideByZero
		}
		return first / second, nil
	})
	sub := NumericBinary(2, "left", func(first int, second int) (int, error) {
		return first - second, nil
	})
	add := NumericBinary(2, "left", func(first int, second int) (int, error) {
		return first + second, nil
	})
	g := Grammar{Operators: map[string]Operator{
		"^": pow,
		"*": mult,
		"×": mult,
		"/": div,
		"÷": div,
		"-": sub,
		"+": add,
	}}
	return g
}
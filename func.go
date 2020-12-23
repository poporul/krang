package krang

func UnaryFunc(fn func(v interface{}) (interface{}, error)) Operator {
	return &funCallOperator{operator: operator{
		precedence:    9,
		associativity: Right,
	},
		fn: fn,
	}
}

type funCallOperator struct {
	operator
	fn func(v interface{}) (interface{}, error)
}

func (o funCallOperator) Calculate(v Value) (*Value, error) {
	res, err := o.fn(v.Val)
	if err != nil {
		return nil, err
	}
	return &Value{Val: res}, nil
}
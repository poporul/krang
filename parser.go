package krang

func Parse(tokens []Token) []Token {
	var result []Token
	var operators TokenStack

	for _, t := range tokens {
		switch token := t.(type) {
		case *Value:
			result = append(result, token)

		case LeftBracket:
			operators.Push(token)

		case RightBracket:
			for {
				top := operators.Pop()
				if _, ok := top.(LeftBracket); ok {
					break
				}

				result = append(result, top)
			}

		case Operator:
			for {
				top := operators.Peek()

				if top == nil {
					break
				}
				if _, ok := top.(LeftBracket); ok {
					break
				}

				if top.(Operator).Precedence() >= token.Precedence() && token.Associativity() == Left {
					result = append(result, operators.Pop())
				} else {
					break
				}
			}

			operators.Push(token)
		}
	}

	for {
		if token := operators.Pop(); token != nil {
			result = append(result, token)
		} else {
			break
		}
	}

	return result
}

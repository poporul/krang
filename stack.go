package calculator

type TokenStack []Token

func (ts *TokenStack) Push(token Token) {
	*ts = append(*ts, token)
}

func (ts *TokenStack) Pop() Token {
	if len(*ts) == 0 {
		return nil
	}

	token := ts.Peek()
	*ts = (*ts)[:len(*ts)-1]

	return token
}

func (ts *TokenStack) Peek() Token {
	if len(*ts) == 0 {
		return nil
	}

	return (*ts)[len(*ts)-1]
}

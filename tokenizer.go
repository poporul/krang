package krang

import (
	"errors"
	"sort"
)

var (
	InvalidInputError = errors.New("invalid input")
)

type LeftBracket struct{}
type RightBracket struct{}

type Token interface{}


func (g *Grammar) Tokenize(value string) ([]Token, error) {
	var result []Token
	var err error
	var token string
	var isQuoted bool
	addToken := func() {
		if len(token) > 0 {
			result = append(result, FromString(token))
			token = ""
		}
	}
	var ops []string
	for k, _ := range g.Operators {
		ops = append(ops, k)
	}
	//sort operator tokens so that && precedes &
	sort.Slice(ops, func(i, j int) bool {
		if len(ops[i]) != len(ops[j]) {
			return len(ops[i]) > len(ops[j])
		}
		return ops[i] > ops[j]
	})
scan:
	for pos := 0; pos<len(value); pos++ {
		c := value[pos]
		if isQuoted {
			if c == '\'' {
				isQuoted = false
				addToken()
				continue
			}
			token += string(c)
		} else {
			if c == '\'' {
				addToken()
				isQuoted = true
				continue
			}
			if c == ' ' {
				addToken()
				continue
			}

			if c == '(' {
				addToken()
				result = append(result, LeftBracket{})
				continue
			}

			if c == ')' {
				addToken()
				result = append(result, RightBracket{})
				continue
			}
			for _, k := range ops {
				op := g.Operators[k]
				if pos+len(k) <=len(value) &&
					k == value[pos:pos+len(k)] {
					addToken()
					result = append(result, op)
					pos += len(k)-1
					continue scan
				}
			}
			token += string(c)
		}
	}

	if len(token) > 0 {
		result = append(result, FromString(token))
	}

	return result, err
}

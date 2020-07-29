package catsort

import "unicode"

// Eval is a evaluater for Parse.
type Eval interface {
	Number(string)
	Letter(string)
	Symbol(string)
}

// Parse divides the string s into chunks of same category runes,
// applies eval to them.
func Parse(s string, eval Eval) {
	rs := []rune(s)
	x := len(rs)
	for i := 0; i < x; {
		r := rs[i]
		j := i + 1
		var k int
		switch {
		case unicode.IsDigit(r):
			k = j + skipFunc(rs[j:], unicode.IsDigit)
			eval.Number(string(rs[i:k]))
		case unicode.IsLetter(r):
			k = j + skipFunc(rs[j:], unicode.IsLetter)
			eval.Letter(string(rs[i:k]))
		default:
			k = j + skipFunc(rs[j:], isSymbol)
			eval.Symbol(string(rs[i:k]))
		}
		i = k
	}
}

func skipFunc(rs []rune, f func(rune) bool) int {
	x := len(rs)
	for i := 0; i < x; i++ {
		if !f(rs[i]) {
			return i
		}
	}
	return x
}

func isSymbol(r rune) bool {
	return !(unicode.IsDigit(r) || unicode.IsLetter(r))
}

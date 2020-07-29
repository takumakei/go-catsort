package catsort

import (
	"strings"
)

// Cats is an array of Cat.
type Cats []Cat

// NewCats returns Cats that s is parsed into.
func NewCats(s string) Cats {
	var e Parser
	Parse(s, &e)
	return e.Cats
}

func (cats Cats) ID() string {
	var b strings.Builder
	for _, v := range cats {
		b.WriteString(v.ID())
	}
	return b.String()
}

// CompareCats returns -1 if lhs is less than rhs,
// returns +1 if lhs is greater than rhs,
// otherwise returns 0 that means equal.
//
// Cats is a sequence of Cat.
// Comparing Cats compares each Cat in Cats first to last.
//
// In case of the number of components of each Cats is
// different and the common parts are the same,
// the shorter one is less than the longer one.
func CompareCats(lhs, rhs Cats) int {
	xl, xr := len(lhs), len(rhs)
	for i, l := range lhs[:minInt(xl, xr)] {
		r := rhs[i]
		if c := l.Compare(r); c != 0 {
			return c
		}
	}
	return compareInt(xl, xr)
}

// Parser is Eval for NewCats.
type Parser struct {
	Cats Cats
}

var _ Eval = (*Parser)(nil)

func (eval *Parser) Number(s string) { eval.Cats = append(eval.Cats, NewNumber(s)) }
func (eval *Parser) Letter(s string) { eval.Cats = append(eval.Cats, Letter(s)) }
func (eval *Parser) Symbol(s string) { eval.Cats = append(eval.Cats, Symbol(s)) }

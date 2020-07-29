package catsort

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

const (
	// Place Number before Letter
	compareNumberToLetter = -1
	compareLetterToNumber = -compareNumberToLetter

	// Place Number after Symbol
	compareNumberToSymbol = 1
	compareSymbolToNumber = -compareNumberToSymbol

	// Place Letter after Symbol
	compareLetterToSymbol = 1
	compareSymbolToLetter = -compareLetterToSymbol

	// Place anything after nil
	compareToNil = 1

	// Place anything before unknown
	compareToAny = -1
)

// Cat is chunk of same category runes.
type Cat interface {
	// ID returns the identification string.
	// For example ID returns "2" for Number("002").
	ID() string

	// Source string
	String() string

	// Compare returns -1 if this instance is less than rhs,
	// returns +1 if this instance is greater than rhs,
	// otherwise returns 0 that means equal.
	Compare(rhs interface{}) int
}

// Number is a chunk of the string contains only digits.
type Number struct {
	idU uint64
	idB *big.Int
	str string
}

func NewNumber(s string) *Number {
	idU, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		return &Number{idU: idU, str: s}
	}
	idB, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic(fmt.Sprintf("NewNumber(%q)", s))
	}
	return &Number{idB: idB, str: s}
}

// ID returns the identification string.
//
// For example it returns "2" for Number from "002".
func (cat *Number) ID() string {
	if cat.idB != nil {
		return cat.idB.Text(10)
	}
	return strconv.FormatUint(cat.idU, 10)
}

// String returns the source string.
//
// For example it returns "002" for Number from "002".
func (cat *Number) String() string { return cat.str }

func (cat *Number) Compare(rhs interface{}) int {
	switch v := rhs.(type) {
	case *Number:
		return cat.compare(v)
	case Letter:
		return compareNumberToLetter
	case Symbol:
		return compareNumberToSymbol
	case nil:
		return compareToNil
	}
	return compareToAny
}

func (cat *Number) compare(rhs *Number) int {
	rhsB := rhs.idB
	if lhsB := cat.idB; lhsB != nil {
		if rhsB != nil {
			return lhsB.Cmp(rhsB)
		} else {
			return lhsB.Cmp(new(big.Int).SetUint64(rhs.idU))
		}
	} else {
		if rhsB != nil {
			return new(big.Int).SetUint64(cat.idU).Cmp(rhsB)
		} else {
			return compareUint64(cat.idU, rhs.idU)
		}
	}
}

// Letter is a chunk of the string contains only letters.
type Letter string

func (cat Letter) ID() string     { return string(cat) }
func (cat Letter) String() string { return string(cat) }

func (cat Letter) Compare(rhs interface{}) int {
	switch v := rhs.(type) {
	case *Number:
		return compareLetterToNumber
	case Letter:
		return strings.Compare(string(cat), string(v))
	case Symbol:
		return compareLetterToSymbol
	case nil:
		return compareToNil
	}
	return compareToAny
}

// Symbol is a chunk of the string contains only non-digit nor non-letters.
type Symbol string

func (cat Symbol) ID() string     { return string(cat) }
func (cat Symbol) String() string { return string(cat) }

func (cat Symbol) Compare(rhs interface{}) int {
	switch v := rhs.(type) {
	case *Number:
		return compareSymbolToNumber
	case Letter:
		return compareSymbolToLetter
	case Symbol:
		return strings.Compare(string(cat), string(v))
	case nil:
		return compareToNil
	}
	return compareToAny
}

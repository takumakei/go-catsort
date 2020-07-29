package catsort

import (
	"math"
	"strconv"
	"testing"
)

func TestNewNumber(t *testing.T) {
	b0 := strconv.FormatUint(math.MaxUint64, 10) + "00"
	b1 := strconv.FormatUint(math.MaxUint64, 10) + "01"
	ts := []struct {
		l string
		a string
		b string
		w int
	}{
		{"ii12", "1", "2", -1},
		{"ii21", "2", "1", +1},

		{"iB", "1", b0, -1},
		{"Bi", b0, "1", +1},

		{"BB01", b0, b1, -1},
		{"BB10", b1, b0, +1},
	}
	for _, v := range ts {
		t.Run(v.l, func(t *testing.T) {
			a := NewNumber(v.a)
			b := NewNumber(v.b)
			r := a.Compare(b)
			if r != v.w {
				t.Errorf("%q : %q = %d, want %d", v.a, v.b, r, v.w)
			}
		})
	}
}

func TestNumberID(t *testing.T) {
	b0 := strconv.FormatUint(math.MaxUint64, 10) + "00"
	ts := []struct {
		s string
		w string
	}{
		{"1", "1"},
		{"001", "1"},
		{b0, b0},
		{"00" + b0, b0},
	}
	for _, v := range ts {
		n := NewNumber(v.s)
		r := n.ID()
		if r != v.w {
			t.Errorf("NewNumber(%q).ID() = %q, want %q", v.s, r, v.w)
		}
	}
}

func TestNumberString(t *testing.T) {
	b0 := strconv.FormatUint(math.MaxUint64, 10) + "00"
	ts := []struct {
		s string
		w string
	}{
		{"1", "1"},
		{"001", "001"},
		{b0, b0},
		{"00" + b0, "00" + b0},
	}
	for _, v := range ts {
		n := NewNumber(v.s)
		r := n.String()
		if r != v.w {
			t.Errorf("NewNumber(%q).String() = %q, want %q", v.s, r, v.w)
		}
	}
}

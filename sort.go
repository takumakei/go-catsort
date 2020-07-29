package catsort

import "sort"

// Strings sorts the list of array of strings.
func Strings(list []string) {
	a := ToValues(list)
	Values(a)
	for i, v := range a {
		list[i] = v.String
	}
}

// Values sorts the list of array of *Value.
func Values(list []*Value) {
	sort.Slice(list, func(i, j int) bool {
		vi, vj := list[i], list[j]
		c := CompareValue(vi, vj)
		if c == 0 {
			return vi.String < vj.String
		}
		return c == -1
	})
}

// Indexes returns the array of indexes.
// The result of sorting is reported by its index and the list is unchanged.
func Indexes(list []string) []int {
	type Index struct {
		index int
		value *Value
	}
	indexes := make([]*Index, len(list))
	for i, v := range list {
		indexes[i] = &Index{index: i, value: NewValue(v)}
	}
	sort.Slice(indexes, func(i, j int) bool {
		vi, vj := indexes[i], indexes[j]
		c := CompareValue(vi.value, vj.value)
		if c == 0 {
			return vi.value.String < vj.value.String
		}
		return c == -1
	})
	a := make([]int, len(indexes))
	for i, v := range indexes {
		a[i] = v.index
	}
	return a
}

// ToStrings converts the list of *Value to array of string.
func ToStrings(list []*Value) []string {
	a := make([]string, len(list))
	for i, v := range list {
		a[i] = v.String
	}
	return a
}

// ToValues conerts the list of strings to array of *Value.
func ToValues(list []string) []*Value {
	a := make([]*Value, len(list))
	for i, v := range list {
		a[i] = NewValue(v)
	}
	return a
}

// IndexOfConsectiveID returns the first index of *Value which has
// a Cats.ID() equals to the a Cats.ID() of next *Value in the list.
// It returns -1 if it is not found.
// That is `CompareValue(list[i], list[i+1]) == 0`.
// The list must be sorted.
func IndexOfConsectiveID(list []*Value) int {
	for i, v := range list[1:] {
		// v is list[i+1] because range starts from 1
		// compare list[i] to list[i+1]
		if CompareValue(list[i], v) == 0 {
			return i
		}
	}
	return -1
}

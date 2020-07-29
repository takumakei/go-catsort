package catsort

import (
	"fmt"
	"testing"
)

func TestStrings(t *testing.T) {
	list := []string{
		"_",
		"0",
		"a",
	}

	sorted := make([]string, len(list))
	copy(sorted, list)

	Strings(sorted)

	for i, v := range sorted {
		if v != list[i] {
			t.Error(v, list[i])
		}
	}
}

func ExampleIndexes() {
	list := []string{
		// These are sorted by sort.Strings.
		"0.0.1",
		"0.0.100",
		"0.0.2",
		"0.0.20",
		"0.1",
		"0.1.1",
		"0.1.2",
		"0.10",
		"0.2",
		"0.2.1",

		"2020-02-02.1.hello(1).sql",
		"2020-02-02.1.hello(100).sql",
		"2020-02-02.1.hello(20).sql",
		"2020-02-02.100.melon(1).sql",
		"2020-02-02.100.melon(100).sql",
		"2020-02-02.100.melon(20).sql",
		"2020-02-02.20.world(1).sql",
		"2020-02-02.20.world(100).sql",
		"2020-02-02.20.world(20).sql",

		"catsort.2020-02-11",
		"catsort.2020-03-21",
		"catsort.2020-10-9",
		"catsort.2020-2-2",
	}

	indexes := Indexes(list)

	for _, i := range indexes {
		fmt.Println(list[i])
	}

	// output:
	// 0.0.1
	// 0.0.2
	// 0.0.20
	// 0.0.100
	// 0.1
	// 0.1.1
	// 0.1.2
	// 0.2
	// 0.2.1
	// 0.10
	// 2020-02-02.1.hello(1).sql
	// 2020-02-02.1.hello(20).sql
	// 2020-02-02.1.hello(100).sql
	// 2020-02-02.20.world(1).sql
	// 2020-02-02.20.world(20).sql
	// 2020-02-02.20.world(100).sql
	// 2020-02-02.100.melon(1).sql
	// 2020-02-02.100.melon(20).sql
	// 2020-02-02.100.melon(100).sql
	// catsort.2020-2-2
	// catsort.2020-02-11
	// catsort.2020-03-21
	// catsort.2020-10-9
}

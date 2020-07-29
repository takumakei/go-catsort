package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/takumakei/go-catsort"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	flagUnique := false

	flag.Usage = usage
	flag.BoolVar(&flagUnique, "u", false, "Only output lines that are not repeated in the input.")
	flag.Parse()

	list, err := readList(flag.Args())
	if err != nil {
		return err
	}

	catsort.Strings(list)

	if flagUnique {
		line := list[0]
		fmt.Println(line)
		for _, v := range list[1:] {
			if line != v {
				line = v
				fmt.Println(line)
			}
		}
	} else {
		for _, v := range list {
			fmt.Println(v)
		}
	}

	return nil
}

func usage() {
	fmt.Printf("Usage: %s [-u] [file ...]\n", filepath.Base(os.Args[0]))
	os.Exit(0)
}

func readList(files []string) (list []string, err error) {
	list = make([]string, 0, 1024)

	if len(files) == 0 {
		list, err = readListReader(list, os.Stdin)
		return
	}

	for _, file := range files {
		if list, err = readListFile(list, file); err != nil {
			break
		}
	}

	return
}

func readListReader(list []string, src io.Reader) ([]string, error) {
	b := bufio.NewReader(src)
	for {
		line, err := b.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return list, err
		}
		list = append(list, strings.TrimRight(line, "\r\n"))
	}
}

func readListFile(list []string, file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return list, err
	}
	defer f.Close()
	return readListReader(list, f)
}

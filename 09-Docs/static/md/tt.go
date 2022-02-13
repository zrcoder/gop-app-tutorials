package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

var pattern = []byte(`](/reference#`)

func main() {
	infos, err := ioutil.ReadDir(".")
	panicErr(err)
	for _, info := range infos {
		if !strings.HasSuffix(info.Name(), ".md") {
			continue
		}
		fmt.Println(info.Name())
		data, err := ioutil.ReadFile(info.Name())
		panicErr(err)
		for {
			i := bytes.LastIndex(data, pattern)
			if i == -1 {
				break
			}
			j := i - 1
			for j >= 0 && data[j] != '[' {
				j--
			}
			if j == i-1 || j < 0 {
				break
			}
			k := i + len(pattern)
			for k < len(data) && data[k] != ')' {
				k++
			}
			if k == i+len(pattern) || k == len(data) {
				break
			}
			src := data[j : k+1]
			dest := append([]byte{'`'}, data[j+1:i]...)
			dest = append(dest, '`')
			data = bytes.ReplaceAll(data, src, dest)
		}
		err = ioutil.WriteFile(info.Name(), data, 0600)
		panicErr(err)
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func TestFormatterFormat(t *testing.T) {
	var buf bytes.Buffer
	data := make(map[string]int)
	f := NewFormatter()

	data["a"] = 1
	data["aa"] = 20
	data["aaa"] = 300
	data["bbbb"] = 40
	data["bbbbb"] = 4

	f.Format(&buf, data)

	longest := 0

	for k := range data {
		if len(k) > longest {
			longest = len(k)
		}
	}

	lines := []string{}

	for k, v := range data {
		min := 4
		width := longest - len(k) + min
		line := k

		for i := 0; i < width; i++ {
			line += " "
		}

		line += strconv.Itoa(v)
		line += "\n"
		lines = append(lines, line)
	}

	for _, v := range lines {
		if !strings.Contains(buf.String(), v) {
			t.Fatalf("did not find line %q in output", v)
		}
	}

}

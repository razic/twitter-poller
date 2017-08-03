package main

import (
	"io"
	"strconv"
	"text/tabwriter"
)

// Formatter data struct
type Formatter struct{}

// NewFormatter returns a new Formatter
func NewFormatter() *Formatter {
	return &Formatter{}
}

// Format takes a map of strings to ints, and formats them for human reading
func (f *Formatter) Format(writer io.Writer, data map[string]int) {
	w := tabwriter.NewWriter(writer, 0, 0, 4, ' ', tabwriter.TabIndent)

	for k, v := range data {
		w.Write([]byte(k + "\t" + strconv.Itoa(v) + "\n"))
	}

	w.Flush()
}

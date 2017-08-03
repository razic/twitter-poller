package main

import (
	"bufio"
	"io"
)

// URLScanner data struct
type URLScanner struct {
	s *bufio.Scanner
}

// NewURLScanner takes a io.Reader, and returns a URLScanner
func NewURLScanner(r io.Reader) *URLScanner {
	return &URLScanner{
		s: bufio.NewScanner(r),
	}
}

// Scan takes a channel of string, on which is passed each scanned token
func (s *URLScanner) Scan(pollers chan Poller) {
	for s.s.Scan() {
		pollers <- &URL{raw: s.s.Text()}
	}
}

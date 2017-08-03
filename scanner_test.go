package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestURLScannerScan(t *testing.T) {
	reader := strings.NewReader("a\nb\nc")
	scanner := NewURLScanner(reader)
	urls := make(chan Poller)

	go scanner.Scan(urls)

	a := <-urls
	b := <-urls
	c := <-urls

	if !reflect.DeepEqual(a, &URL{raw: "a"}) {
		t.Fatalf("unexpected url: %v", a)
	}

	if !reflect.DeepEqual(b, &URL{raw: "b"}) {
		t.Fatalf("unexpected url: %v", b)
	}

	if !reflect.DeepEqual(c, &URL{raw: "c"}) {
		t.Fatalf("unexpected url: %v", c)
	}
}

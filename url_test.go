package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestURLString(t *testing.T) {
	u := NewURL("a")
	if str := u.String(); str != "http://a" {
		t.Fatalf("expected %q got %q", "a", str)
	}
	u = NewURL("http://b")
	if str := u.String(); str != "http://b" {
		t.Fatalf("expected %q got %q", "b", str)
	}
}

func TestURLPoll(t *testing.T) {
	s := Status{Application: "a", Version: "1"}
	url := NewURL("b")
	recorder := httptest.NewRecorder()
	getter := &FakeGetter{get: func(u string) (*http.Response, error) {
		return recorder.Result(), nil
	}}
	byt, _ := json.Marshal(s)

	recorder.Write(byt)
	status, err := url.Poll(getter)

	if err != nil {
		t.Fatalf("unexpected error")
	}

	// valid status json
	if !reflect.DeepEqual(status, s) {
		t.Fatalf("%v %v", status, s)
	}

	recorder = httptest.NewRecorder()
	getter = &FakeGetter{get: func(u string) (*http.Response, error) {
		return recorder.Result(), nil
	}}
	status, err = url.Poll(getter)

	// invalid json
	if !reflect.DeepEqual(status, Status{}) {
		t.Fatalf("%v %v", status, s)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

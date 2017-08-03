package main

import (
	"reflect"
	"testing"
)

func TestPoll(t *testing.T) {
	getter := &FakeGetter{}
	statuses := make(chan Status)
	pollers := make(chan Poller)

	go Poll(getter, pollers, statuses)

	s := Status{Application: "a", Version: "1", SuccessCount: 100}
	pollers <- &FakePoller{poll: func(g Getter) (Status, error) {
		return s, nil
	}}
	status := <-statuses
	close(pollers)

	if !reflect.DeepEqual(s, status) {
		t.Fatalf("expected %v, got %v", s, status)
	}
}

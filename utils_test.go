package main

import "net/http"

type FakeGetter struct {
	Getter
	get func(string) (*http.Response, error)
}

func (f *FakeGetter) Get(u string) (*http.Response, error) {
	return f.get(u)
}

type FakePoller struct {
	Poller
	poll func(getter Getter) (Status, error)
}

func (f *FakePoller) Poll(getter Getter) (Status, error) {
	return f.poll(getter)
}

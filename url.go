package main

import (
	"encoding/json"
	"io/ioutil"
)

// URL data struct
type URL struct {
	raw string
}

// NewURL initializes and returns a new URL
func NewURL(u string) *URL {
	return &URL{u}
}

// String gets the url string from a url
func (u *URL) String() string {
	return u.raw
}

// Poll polls a url
func (u *URL) Poll(getter Getter) (status Status, err error) {
	resp, err := getter.Get(u.raw)

	if err != nil {
		return
	}

	byt, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	err = json.Unmarshal(byt, &status)

	if err != nil {
		return
	}

	return
}

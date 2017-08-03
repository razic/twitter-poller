package main

import "net/http"

// Getter defines something that gets a url
type Getter interface {
	Get(uri string) (*http.Response, error)
}

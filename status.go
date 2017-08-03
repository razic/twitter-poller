package main

// Status data struct
type Status struct {
	Application  string
	Version      string
	SuccessCount int `json:"Success_Count"`
}

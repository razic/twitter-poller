package main

import "testing"

func TestAggregatorAggregate(t *testing.T) {
	a := NewAggregator()
	statuses := make(chan Status)

	go a.Aggregate(statuses)

	statuses <- Status{Application: "a", Version: "1", SuccessCount: 10}
	statuses <- Status{Application: "a", Version: "1", SuccessCount: 20}
	statuses <- Status{Application: "a", Version: "1", SuccessCount: 30}
	statuses <- Status{Application: "a", Version: "2", SuccessCount: 30}
	statuses <- Status{Application: "b", Version: "2", SuccessCount: 8}

	if a.Data["a1"] != 60 {
		t.Fatalf("expected %d, got %d", 60, a.Data["a1"])
	}

	if a.Data["a2"] != 30 {
		t.Fatalf("expected %d, got %d", 30, a.Data["a2"])
	}

	if a.Data["b2"] != 8 {
		t.Fatalf("expected %d, got %d", 8, a.Data["b2"])
	}
}

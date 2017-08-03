package main

// Aggregator data struct
type Aggregator struct {
	Data map[string]int
}

// Returns and initialized a new aggregator
func NewAggregator() *Aggregator {
	return &Aggregator{Data: make(map[string]int)}
}

// Aggregate takes a channel of statuses, on which it listens and aggregates
// the success count
func (a *Aggregator) Aggregate(statuses chan Status) {
	for s := range statuses {
		if s.Application != "" && s.Version != "" {
			a.Data[s.Application+s.Version] += s.SuccessCount
		}
	}
}

package main

import "log"

// Poller interface represents a poller
type Poller interface {
	Poll(Getter) (Status, error)
}

// Poll ranges over pollers chan, calling Poll on the poller, and passing its
// Status to the statuses chan
func Poll(getter Getter, pollers chan Poller, statuses chan Status) {
	for p := range pollers {
		status, err := p.Poll(getter)

		if err != nil {
			log.Printf("%v\n", err)
			continue
		}

		statuses <- status
	}
}

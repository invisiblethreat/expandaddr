package expandaddr

import "sync"

// SingleTarget is an atomic entity to attempt a connection
type SingleTarget struct {
	Addr  string `json:"addr"`
	Port  int    `json:"port"`
	Proto string `json:"proto"`
}

// AllTargets holds the exploded arguments which are used for the Cartesian
// product to generate the set of atomic SingleTargets
type AllTargets struct {
	Addrs  []string
	Ports  []int
	Protos []string
}

// Load builds out atomic targets
func (at *AllTargets) Load(output chan SingleTarget, wg *sync.WaitGroup) {
	for _, proto := range at.Protos {
		for _, port := range at.Ports {
			for _, addr := range at.Addrs {
				output <- SingleTarget{Addr: addr, Port: port, Proto: proto}
				wg.Add(1)
			}
		}
	}
}

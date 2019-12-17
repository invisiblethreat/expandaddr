package expandaddr

// SingleTarget is an atomic entity to attempt a connection
type SingleTarget struct {
	Addr  string `json:"addr"`
	Port  int    `json:"port"`
	Proto string `json:"proto"`
}

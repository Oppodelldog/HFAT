package server

// ForwardingTarget defines a target HFAT will forward received http requests to
type ForwardingTarget struct {
	Server  string
	Port    int
	Primary bool
}

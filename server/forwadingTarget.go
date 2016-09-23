package server

type ForwardingTarget struct {
	Server string
	Port   int
	Primary bool
}

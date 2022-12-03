package net

type NetAddr string

//The messages sent over the transport layer
type RPC struct {
	From NetAddr
	Payload []byte
}

type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	Send(NetAddr, []byte) error
	Addr() NetAddr
}
package main

import (
	"time"

	"github.com/notemptylist/sprat/net"
)

// Server container
// Transport layer => tcp, udp, websockets

// block
// Transaction
// Keypair

func main() {
	trLocal := net.NewLocalTransport("local")
	trRemote := net.NewLocalTransport("remote")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.Send(trLocal.Addr(), []byte("hello"))
			time.Sleep(1 * time.Second)
		}
	}()
	opts := net.ServerOpts{
		Transports: []net.Transport{trLocal},
	}
	s := net.NewServer(opts)
	s.Start()
}

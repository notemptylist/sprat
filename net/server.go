package net

import (
	"fmt"
	"time"
)

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	ServerOpts
	rpcCh  chan RPC
	quitCh chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		rpcCh:      make(chan RPC),
		quitCh:     make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	tick := time.NewTicker(3 * time.Second)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			//handler
			fmt.Printf("%+v\n", rpc)
		case <-s.quitCh:
			break free
		case <-tick.C:
			fmt.Println("Tick")
		}
	}
	fmt.Println("Server shutting down")
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				// handler
				s.rpcCh <- rpc
			}
		}(tr)
	}
}

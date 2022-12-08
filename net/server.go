package net

import (
	"fmt"
	"time"

	"github.com/notemptylist/sprat/core"
	"github.com/notemptylist/sprat/crypto"
)

type ServerOpts struct {
	Transports []Transport
	BlockTime  time.Duration
	PrivateKey *crypto.PrivateKey
}

type Server struct {
	ServerOpts
	blockTime   time.Duration
	memPool     *TxPool
	isValidator bool
	rpcCh       chan RPC
	quitCh      chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts:  opts,
		blockTime:   opts.BlockTime,
		memPool:     NewTxPool(64),
		isValidator: opts.PrivateKey != nil,
		rpcCh:       make(chan RPC),
		quitCh:      make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	tick := time.NewTicker(s.blockTime)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			//handler
			fmt.Printf("%+v\n", rpc)
		case <-s.quitCh:
			break free
		case <-tick.C:
			if s.isValidator {
				s.createNewBlock()
			}
		}
	}
	fmt.Println("Server shutting down")
}

func (s *Server) handleTransaction(tx *core.Transaction) error {
	if err := tx.Verify(); err != nil {
		return err
	}
	hash := tx.Hash(core.TxHasher{})
	if s.memPool.Has(hash) {
		return nil

	}
	return s.memPool.Add(tx)
}

func (s *Server) createNewBlock() error {
	fmt.Println("creating new block")
	return nil
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

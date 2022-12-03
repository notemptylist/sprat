package net

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	peer1 := NewLocalTransport("peer1")
	peer2 := NewLocalTransport("peer2")

	peer1.Connect(peer2)
	peer2.Connect(peer1)

	assert.Equal(t, peer1.peers[peer2.addr], peer2)
	assert.Equal(t, peer2.peers[peer1.addr], peer1)
}

func TestSend(t *testing.T) {
	peer1 := NewLocalTransport("peer1")
	peer2 := NewLocalTransport("peer2")

	peer1.Connect(peer2)
	peer2.Connect(peer1)

	payload := []byte("FOOBAR")
	err := peer1.Send(peer2.addr, payload)
	assert.Equal(t, err, nil)

	// Use the Public method
	received := <-peer2.Consume()
	assert.Equal(t, payload, received.Payload)
	assert.Equal(t, received.From, peer1.Addr())
}

package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr      NetAddr
	consumeCh chan PRC
	lock      sync.RWMutex
	peers     map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) Transport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan PRC, 1024),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

// Consume returns a channel that can be used to consume messages from the transport.

func (t *LocalTransport) Consume() <-chan PRC {
	return t.consumeCh
}
func (t *LocalTransport) Produce(rpc PRC) {
	t.consumeCh <- rpc
}

// Connect connects to another transport.
func (t *LocalTransport) Connect(tr Transport) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.peers[tr.Addr()] = tr.(*LocalTransport)
	return nil
}

func (t *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	t.lock.RLock()
	defer t.lock.RUnlock()
	peer, ok := t.peers[to]

	if !ok {
		return fmt.Errorf("%s could not send message to %s: peer not found", t.addr, to)
	}

	peer.consumeCh <- PRC{
		From:    t.addr,
		Payload: payload,
		To:      string(to),
	}
	return nil
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}

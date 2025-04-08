package network

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestConnect(t *testing.T) {
// 	tra := NewLocalTransport("A")
// 	trb := NewLocalTransport("B")
// 	tra.Connect(trb)
// 	trb.Connect(tra)
// 	assert.Equal(t, tra.peers[trb.Addr()], trb)
// 	assert.Equal(t, trb.peers[tra.Addr()], tra)
// }

// func TestSendMessage(t *testing.T) {
// 	var message []byte
// 	tra := NewLocalTransport("A")
// 	trb := NewLocalTransport("B")
// 	tra.Connect(trb)
// 	trb.Connect(tra)
// 	message = []byte("Hello,World!")
// 	assert.Nil(t, tra.SendMessage(trb.Addr(), message))
// 	// Consume the message from trb
// 	rpc := <-trb.Consume()
// 	assert.Equal(t, rpc.Payload, message)
// 	assert.Equal(t, rpc.From, tra.Addr())
// 	assert.Equal(t, rpc.To, trb.Addr())
// }

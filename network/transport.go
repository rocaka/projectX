package network

type NetAddr string

// PRC is a transport that uses the PRC protocol.
type PRC struct {
	From    NetAddr // The sender of the message
	To      string  // The receiver of the message
	Payload []byte  // The payload of the message
}

// Transport is an interface that defines the methods for a transport.
type Transport interface {
	Consume() <-chan PRC // Consume messages from the transport
	// Produce() chan<- PRC               // Produce messages to the transport
	Connect(Transport) error           // Connect to another transport
	SendMessage(NetAddr, []byte) error // Send a message to a specific address
	Addr() NetAddr                     // Get the address of the transport

}

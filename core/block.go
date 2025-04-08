package core

import (
	"encoding/binary"
	"io"

	"github.com/rocaka/projectX/types"
)

type Header struct {
	Version   uint32
	PrevBlock types.Hash
	Timestamp uint64
	Height    uint32
	Nonce     uint64
}

func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	binary.Write(w, binary.LittleEndian, &h.PrevBlock)
	binary.Write(w, binary.LittleEndian, &h.Timestamp)
	binary.Write(w, binary.LittleEndian, &h.Height)
	binary.Write(w, binary.LittleEndian, &h.Nonce)

	return nil
}

func (h *Header) DecodeBinary(r io.Reader) error {

	return nil
}

type Block struct {
	Header
	Transactions []*Transcation
}

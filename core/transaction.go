package core

import "io"

type Transcation struct {
	Data []byte
}

func (tx *Transcation) DecodeBinary(r io.Reader) error {
	return nil
}

func (tx *Transcation) EncodeBinary(w io.Writer) error {
	return nil
}

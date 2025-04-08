package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/rocaka/projectX/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode(t *testing.T) {
	h := Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     100,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h.Version, hDecode.Version)
	assert.Equal(t, h.PrevBlock, hDecode.PrevBlock)
	assert.Equal(t, h.Timestamp, hDecode.Timestamp)
	assert.Equal(t, h.Height, hDecode.Height)
	assert.Equal(t, h.Nonce, hDecode.Nonce)
}

func TestBlock_Encode_Decode(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     100,
		},
		Transcations: nil,
	}
	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))

	assert.Equal(t, b.Version, bDecode.Version)
	assert.Equal(t, b.PrevBlock, bDecode.PrevBlock)
	assert.Equal(t, b.Timestamp, bDecode.Timestamp)
	assert.Equal(t, b.Height, bDecode.Height)
	assert.Equal(t, b.Nonce, bDecode.Nonce)
}

func TestBlock_hash_test(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     100,
		},
		Transcations: nil,
	}
	buf := &bytes.Buffer{}

	assert.Nil(t, b.Header.EncodeBinary(buf))

}

func TestBlockHash(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     100,
		},
		Transcations: []Transcation{},
	}

	h := b.Hash()
	fmt.Println(h)
	assert.False(t, h.IsZero())
}

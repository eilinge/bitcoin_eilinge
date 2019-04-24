package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
)

// IntToHex ...
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	// Write func(w io.Writer, order ByteOrder, data interface{}) error
	if err := binary.Write(buff, binary.BigEndian, num); err != nil {
		log.Panic(err)
	}
	// Bytes func() []byte
	return buff.Bytes()
}

// DataToHash ...
func DataToHash(data []byte) []byte {
	// Sum256 func(data []byte) [Size]byte
	hash := sha256.Sum256(data)
	return hash[:]
}

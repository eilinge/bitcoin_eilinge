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
	if err := binary.Write(buff, binary.BigEndian, num); err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// DataToHash ...
func DataToHash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

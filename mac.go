package macpot

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strings"
)

const (
	macByteLength    = 6
	intValueByteSize = 8
	ouiByteLength    = 3
	nicByteLength    = 3
	firstBitOn       = 0x01
	firstBitOff      = 0xfe
	secondBitOn      = 0x02
	secondBitOff     = 0xfd
)

// first 3 octets are the OUI (Organisationally Unique Identifier)
// last 3 octets are the NIC (Network interface controler specific)
//
// first octet:
//   [d][d][d][d][d][d][gl][um]
//
//   um = 0: unicast; 1: multicast
//   gl = 0: global; 1: local
//
// For IPv4 we can use a deterministic value:
//   OUI: host machine specific
//   NIC: the last 3 bytes of the IP address

type MAC struct {
	bytes []byte
}

func New(options ...Option) (MAC, error) {
	mac := MAC{bytes: make([]byte, macByteLength)}

	if err := mac.FillRandom(); err != nil {
		return mac, err
	}

	for _, option := range options {
		err := option(&mac)
		if err != nil {
			return mac, err
		}
	}

	return mac, nil
}

func NewFromBytes(value []byte) MAC {
	bytes := make([]byte, macByteLength)

	copy(bytes, value)

	return MAC{bytes: bytes}
}

func NewFromUint64(value uint64) MAC {
	bytes := make([]byte, intValueByteSize)

	binary.BigEndian.PutUint64(bytes, value)

	return MAC{bytes: bytes[(intValueByteSize - macByteLength):]}
}

func (mac MAC) ToString() string {
	output := []string{}

	for _, b := range mac.bytes {
		output = append(output, fmt.Sprintf("%02x", b))
	}

	return strings.Join(output, ":")
}

func (mac *MAC) FillRandom() error {
	if _, err := rand.Read(mac.bytes); err != nil {
		return fmt.Errorf("unable to generate random data: %w", err)
	}

	return nil
}

func (mac MAC) Bytes() []byte {
	return mac.bytes
}

func (mac MAC) ToUint64() uint64 {
	bytes := make([]byte, intValueByteSize)

	copy(bytes[intValueByteSize-macByteLength:], mac.Bytes())

	return binary.BigEndian.Uint64(bytes)
}

func (mac MAC) Next() MAC {
	bytes := make([]byte, intValueByteSize)

	copy(bytes[intValueByteSize-macByteLength:], mac.Bytes())

	intValue := binary.BigEndian.Uint64(bytes)

	return NewFromUint64(intValue + 1)
}

func (mac *MAC) SetLocal() {
	mac.bytes[0] |= secondBitOn
}

func (mac *MAC) SetGlobal() {
	mac.bytes[0] &= secondBitOff
}

func (mac *MAC) SetMulticast() {
	mac.bytes[0] |= firstBitOn
}

func (mac *MAC) SetUnicast() {
	mac.bytes[0] &= firstBitOff
}

func (mac MAC) IsLocal() bool {
	return mac.bytes[0]&secondBitOn == secondBitOn
}

func (mac MAC) IsGlobal() bool {
	return !mac.IsLocal()
}

func (mac MAC) IsMulticast() bool {
	return mac.bytes[0]&firstBitOn == firstBitOn
}

func (mac MAC) IsUnicast() bool {
	return !mac.IsMulticast()
}

func (mac *MAC) SetOctet(index int, value byte) error {
	if index < 0 || index > macByteLength {
		return OutOfBoundError{TargetIndex: index}
	}

	mac.bytes[index] = value

	return nil
}

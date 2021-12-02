package macpot

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strings"
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

// MAC address.
type MAC struct {
	bytes []byte
}

// New creates a new MAC address with Options.
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

// NewFromBytes creates a new MAC address from bytes.
func NewFromBytes(value []byte) MAC {
	bytes := make([]byte, macByteLength)

	copy(bytes, value)

	return MAC{bytes: bytes}
}

// NewFromUint64 creates a new MAC address from Uint64.
func NewFromUint64(value uint64) MAC {
	bytes := make([]byte, intValueByteSize)

	binary.BigEndian.PutUint64(bytes, value)

	return MAC{bytes: bytes[(intValueByteSize - macByteLength):]}
}

// ToString returns the well known string representation of the MAC address.
func (mac MAC) ToString() string {
	output := []string{}

	for _, b := range mac.bytes {
		output = append(output, fmt.Sprintf("%02x", b))
	}

	return strings.Join(output, ":")
}

// FillRandom popupates the MAC address with random values.
func (mac *MAC) FillRandom() error {
	if _, err := rand.Read(mac.bytes); err != nil {
		return fmt.Errorf("unable to generate random data: %w", err)
	}

	return nil
}

// Bytes of the MAC address.
func (mac MAC) Bytes() []byte {
	return mac.bytes
}

// ToUint64 returns with the Uint64 representation of the MAC address.
func (mac MAC) ToUint64() uint64 {
	bytes := make([]byte, intValueByteSize)

	copy(bytes[intValueByteSize-macByteLength:], mac.Bytes())

	return binary.BigEndian.Uint64(bytes)
}

// Next generates the next MAC address.
func (mac MAC) Next() MAC {
	bytes := make([]byte, intValueByteSize)

	copy(bytes[intValueByteSize-macByteLength:], mac.Bytes())

	intValue := binary.BigEndian.Uint64(bytes)

	return NewFromUint64(intValue + 1)
}

// SetLocal forces the MAC address to be a Locally Administered address.
func (mac *MAC) SetLocal() {
	mac.bytes[0] |= secondBitOn
}

// SetGlobal forces the MAC address to be a Globally Unique address.
func (mac *MAC) SetGlobal() {
	mac.bytes[0] &= secondBitOff
}

// SetMulticast forces the MAC address to be a Multicast address.
func (mac *MAC) SetMulticast() {
	mac.bytes[0] |= firstBitOn
}

// SetUnicast forces the MAC address to be a Unicast address.
func (mac *MAC) SetUnicast() {
	mac.bytes[0] &= firstBitOff
}

// IsLocal checks if the address is a Locally Administered address.
func (mac MAC) IsLocal() bool {
	return mac.bytes[0]&secondBitOn == secondBitOn
}

// IsGlobal checks if the address is a Globally Unique address.
func (mac MAC) IsGlobal() bool {
	return !mac.IsLocal()
}

// IsMulticast checks if the address is a Multicast address.
func (mac MAC) IsMulticast() bool {
	return mac.bytes[0]&firstBitOn == firstBitOn
}

// IsUnicast checks if the address is a Unicast address.
func (mac MAC) IsUnicast() bool {
	return !mac.IsMulticast()
}

// SetOctet sets a specific octet in the MAC address.
func (mac *MAC) SetOctet(index int, value byte) error {
	if index < 0 || index > macByteLength {
		return OutOfBoundError{TargetIndex: index}
	}

	mac.bytes[index] = value

	return nil
}

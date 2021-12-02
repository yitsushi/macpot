package macpot

import (
	"fmt"
	"strconv"
	"strings"
)

// WithNIC is an Option for New() that sets the generated MAC address with fixed
// NIC.
func WithNIC(nic string) Option {
	parts := strings.Split(nic, ":")

	return func(mac *MAC) error {
		if len(parts) != nicByteLength {
			return NICError{Message: fmt.Sprintf("invalid NIC: %s", nic)}
		}

		for idx := 0; idx < nicByteLength; idx++ {
			value, err := strconv.ParseInt(parts[idx], base16, ipIntSize)
			if err != nil {
				return NICError{Message: err.Error()}
			}

			if err := mac.SetOctet(ouiByteLength+idx, byte(value)); err != nil {
				return err
			}
		}

		return nil
	}
}

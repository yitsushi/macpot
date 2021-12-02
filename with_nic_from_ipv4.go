package macpot

import (
	"fmt"
	"strconv"
	"strings"
)

// WithNICFromIPv4 is an Option for New() that sets the generated MAC address with fixed
// NIC based on the provided IPv4 address. It uses the last 3 bytes of the
// address.
func WithNICFromIPv4(ip string) Option {
	parts := strings.Split(ip, ".")

	return func(mac *MAC) error {
		// We will skip the first part.
		if len(parts) != nicByteLength+1 {
			return IPv4Error{Message: fmt.Sprintf("invalid IPv4 address: %s", ip)}
		}

		for idx := 0; idx < nicByteLength; idx++ {
			value, err := strconv.ParseInt(parts[idx+1], base10, ipIntSize)
			if err != nil {
				return IPv4Error{Message: err.Error()}
			}

			if err := mac.SetOctet(ouiByteLength+idx, byte(value)); err != nil {
				return err
			}
		}

		return nil
	}
}

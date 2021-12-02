package macpot

import (
	"fmt"
	"strconv"
	"strings"
)

// WithFullAddress is an Option for New() to set the full MAC address.
// Good for validation.
func WithFullAddress(address string) Option {
	parts := strings.Split(address, ":")

	return func(mac *MAC) error {
		if len(parts) != macByteLength {
			return AddressError{Message: fmt.Sprintf("invalid MAC address length: %s", address)}
		}

		for idx := 0; idx < macByteLength; idx++ {
			value, err := strconv.ParseInt(parts[idx], base16, ipIntSize)
			if err != nil {
				return AddressError{Message: err.Error()}
			}

			if err := mac.SetOctet(idx, byte(value)); err != nil {
				return err
			}
		}

		return nil
	}
}

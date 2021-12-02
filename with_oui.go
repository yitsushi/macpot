package macpot

import (
	"fmt"
	"strconv"
	"strings"
)

// WithOUI is an Option for New() that sets the generated MAC address with fixed
// OUI.
func WithOUI(oui string) Option {
	parts := strings.Split(oui, ":")

	return func(mac *MAC) error {
		if len(parts) != ouiByteLength {
			return OUIError{Message: fmt.Sprintf("invalid OUI: %s", oui)}
		}

		for idx := 0; idx < ouiByteLength; idx++ {
			value, err := strconv.ParseInt(parts[idx], base16, ipIntSize)
			if err != nil {
				return OUIError{Message: err.Error()}
			}

			if err := mac.SetOctet(idx, byte(value)); err != nil {
				return err
			}
		}

		return nil
	}
}

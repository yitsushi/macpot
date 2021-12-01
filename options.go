package macpot

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

const (
	base10    = 10
	ipIntSize = 16
)

type Option func(*MAC) error

func AsLocal() Option {
	return func(mac *MAC) error {
		mac.SetLocal()

		return nil
	}
}

func AsGlobal() Option {
	return func(mac *MAC) error {
		mac.SetGlobal()

		return nil
	}
}

func AsUnicast() Option {
	return func(mac *MAC) error {
		mac.SetUnicast()

		return nil
	}
}

func AsMulticast() Option {
	return func(mac *MAC) error {
		mac.SetMulticast()

		return nil
	}
}

func WithOUI(oui string) Option {
	parts := strings.Split(oui, ":")

	return func(mac *MAC) error {
		if len(parts) != ouiByteLength {
			return OUIError{Message: fmt.Sprintf("invalid OUI: %s", oui)}
		}

		for idx := 0; idx < ouiByteLength; idx++ {
			value, err := hex.DecodeString(parts[idx])
			if err != nil {
				return OUIError{Message: err.Error()}
			}

			if err := mac.SetOctet(idx, value[0]); err != nil {
				return err
			}
		}

		return nil
	}
}

func WithNIC(nic string) Option {
	parts := strings.Split(nic, ":")

	return func(mac *MAC) error {
		if len(parts) != nicByteLength {
			return NICError{Message: fmt.Sprintf("invalid NIC: %s", nic)}
		}

		for idx := 0; idx < nicByteLength; idx++ {
			value, err := hex.DecodeString(parts[idx])
			if err != nil {
				return NICError{Message: err.Error()}
			}

			if err := mac.SetOctet(ouiByteLength+idx, value[0]); err != nil {
				return err
			}
		}

		return nil
	}
}

func WithNICFromIPv4(ip string) Option {
	parts := strings.Split(ip, ".")

	return func(mac *MAC) error {
		// We will skip the first part.
		if len(parts) != nicByteLength+1 {
			return NICError{Message: fmt.Sprintf("invalid NIC: %s", ip)}
		}

		for idx := 0; idx < nicByteLength; idx++ {
			value, err := strconv.ParseInt(parts[idx+1], base10, ipIntSize)
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

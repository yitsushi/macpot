package macpot

// Option that can change the behaviour of the New() function.
type Option func(*MAC) error

// AsLocal is an Option for New() that sets the generated MAC address to be a
// Locally Administered address.
func AsLocal() Option {
	return func(mac *MAC) error {
		mac.SetLocal()

		return nil
	}
}

// AsGlobal is an Option for New() that sets the generated MAC address to be a
// Globally Unique address.
func AsGlobal() Option {
	return func(mac *MAC) error {
		mac.SetGlobal()

		return nil
	}
}

// AsUnicast is an Option for New() that sets the generated MAC address to be a
// Unicast address.
func AsUnicast() Option {
	return func(mac *MAC) error {
		mac.SetUnicast()

		return nil
	}
}

// AsMulticast is an Option for New() that sets the generated MAC address to be a
// Multicast address.
func AsMulticast() Option {
	return func(mac *MAC) error {
		mac.SetMulticast()

		return nil
	}
}

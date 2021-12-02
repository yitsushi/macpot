package macpot

import "fmt"

// OutOfBoundError occurs when we try to access an unknown octet.
type OutOfBoundError struct {
	TargetIndex int
}

func (e OutOfBoundError) Error() string {
	return fmt.Sprintf("unable to set octet %d", e.TargetIndex)
}

// OUIError occurs when something goes wrong parsing the provided OUI.
type OUIError struct {
	Message string
}

func (e OUIError) Error() string {
	return e.Message
}

// NICError occurs when something goes wrong parsing the provided NIC.
type NICError struct {
	Message string
}

func (e NICError) Error() string {
	return e.Message
}

// IPv4Error occurs when something goes wrong parsing the provided IPv4.
type IPv4Error struct {
	Message string
}

func (e IPv4Error) Error() string {
	return e.Message
}

// AddressError occurs when something goes wrong parsing the provided Address.
type AddressError struct {
	Message string
}

func (e AddressError) Error() string {
	return e.Message
}

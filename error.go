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

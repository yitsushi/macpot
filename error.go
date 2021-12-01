package macpot

import "fmt"

type OutOfBoundError struct {
	TargetIndex int
}

func (e OutOfBoundError) Error() string {
	return fmt.Sprintf("unable to set octet %d", e.TargetIndex)
}

type OUIError struct {
	Message string
}

func (e OUIError) Error() string {
	return e.Message
}

type NICError struct {
	Message string
}

func (e NICError) Error() string {
	return e.Message
}

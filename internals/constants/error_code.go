package constants

type ErrorCode int

const (
	INTERNAL ErrorCode = iota + 1
	BAD_REQUEST
)

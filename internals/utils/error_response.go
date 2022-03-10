package utils

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/constants"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/runtime/protoiface"
)

// CustomError ...
type CustomError struct {
	Status       *status.Status
	Reason       string
	Metadata     map[string]string
	ErrorDetails []protoiface.MessageV1
}

// AddMetadata ...
func (e *CustomError) AddMetadata(k string, v string) {
	e.Metadata[k] = v
}

// AddErrorDetails ...
func (e *CustomError) AddErrorDetails(ed protoiface.MessageV1) {
	e.ErrorDetails = append(e.ErrorDetails, ed)
}

// Err ..
func (e *CustomError) Err() error {
	s, err := e.Status.WithDetails(e.ErrorDetails...)

	if err != nil {
		return e.Status.Err()
	}

	return s.Err()
}

// NewError ...
func NewError(c constants.ErrorCode, msg string, gc ...codes.Code) *CustomError {
	grpcCode := codes.Unknown
	if len(gc) > 0 {
		grpcCode = gc[0]
	}

	s := status.New(grpcCode, msg)
	er := &CustomError{
		Status: s,
		Reason: c.String(),
	}

	errorInfo := &errdetails.ErrorInfo{
		Reason:   er.Reason,
		Metadata: er.Metadata,
	}
	er.AddErrorDetails(errorInfo)

	return er
}

// Error ...
func Error(c constants.ErrorCode, msg string, gc ...codes.Code) error {
	grpcCode := codes.Unknown
	if len(gc) > 0 {
		grpcCode = gc[0]
	}

	return NewError(c, msg, grpcCode).Err()
}

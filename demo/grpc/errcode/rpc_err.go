package errcode

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TogRPCError(err Error) error {
	s := status.New(ToRpcCode(err.code), err.msg)
	return s.Err()
}

func ToRpcCode(code int) codes.Code {
	var statusCode codes.Code
	switch code {
	case Fail.code:
		statusCode = codes.Internal
	}

	return statusCode
}

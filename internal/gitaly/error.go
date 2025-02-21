package gitaly

import (
	"errors"
	"fmt"

	"gitlab.com/gitlab-org/cluster-integration/gitlab-agent/v16/internal/gitaly/vendored/gitalypb"
	"google.golang.org/grpc/status"
)

type ErrorCode int

const (
	// UnknownError - what happened is unknown
	UnknownError ErrorCode = iota
	// NotFound - file/directory/ref was not found
	NotFound
	// FileTooBig - file is too big
	FileTooBig
	// RpcError - gRPC returned an error
	RpcError
	// ProtocolError - protocol violation, an unexpected situation occurred.
	ProtocolError
	// UnexpectedTreeEntryType - returned when TreeEntryResponse has an unexpected type.
	UnexpectedTreeEntryType
	// InvalidArgument - returned when Gitaly responds with InvalidArgument status code.
	InvalidArgument
)

func (e ErrorCode) String() string {
	switch e {
	case UnknownError:
		return "UnknownErr"
	case NotFound:
		return "NotFound"
	case FileTooBig:
		return "FileTooBig"
	case RpcError:
		return "RpcError"
	case ProtocolError:
		return "ProtocolError"
	case UnexpectedTreeEntryType:
		return "UnexpectedTreeEntryType"
	case InvalidArgument:
		return "InvalidArgument"
	default:
		return fmt.Sprintf("invalid ErrorCode: %d", e)
	}
}

type Error struct {
	Code    ErrorCode
	Cause   error
	Message string
	// RpcName is the name of gRPC method that failed.
	RpcName string
	// Path contains name of the file or directory the operation was being carried on.
	Path string
}

func NewNotFoundError(rpcName, path string) error {
	return &Error{
		Code:    NotFound,
		Message: "file/directory/ref not found",
		RpcName: rpcName,
		Path:    path,
	}
}

func NewFileTooBigError(err error, rpcName, path string) error {
	return &Error{
		Code:    FileTooBig,
		Cause:   err,
		Message: "file is too big",
		RpcName: rpcName,
		Path:    path,
	}
}

func NewUnexpectedTreeEntryTypeError(rpcName, path string) error {
	return &Error{
		Code:    UnexpectedTreeEntryType,
		Message: "file is not a usual file",
		RpcName: rpcName,
		Path:    path,
	}
}

func NewRpcError(err error, rpcName, path string) error {
	return &Error{
		Code:    RpcError,
		Cause:   err,
		Message: "RPC failed",
		RpcName: rpcName,
		Path:    path,
	}
}

func NewProtocolError(err error, message, rpcName, path string) error {
	return &Error{
		Code:    ProtocolError,
		Cause:   err,
		Message: message,
		RpcName: rpcName,
		Path:    path,
	}
}

func NewInvalidArgument(err error, rpcName, path string) error {
	return &Error{
		Code:    InvalidArgument,
		Cause:   err,
		Message: "invalid argument",
		RpcName: rpcName,
		Path:    path,
	}
}

func (e *Error) Error() string {
	format := "%s"
	args := []interface{}{e.Code}
	if e.RpcName != "" {
		format += ": %s"
		args = append(args, e.RpcName)
	}
	format += ": %s"
	args = append(args, e.Message)
	if e.Path != "" {
		format += ": %s"
		args = append(args, e.Path)
	}
	if e.Cause != nil {
		format += ": %v"
		args = append(args, e.Cause)
	}
	return fmt.Sprintf(format, args...)
}

func (e *Error) Unwrap() error {
	return e.Cause
}

func ErrorCodeFromError(err error) ErrorCode {
	var e *Error
	if !errors.As(err, &e) {
		return UnknownError
	}
	return e.Code
}

func isInvalidRevisionOrPath(err error) bool {
	s := status.Convert(err)
	for _, d := range s.Details() {
		if e, ok := d.(*gitalypb.GetTreeEntriesError); ok {
			if _, ok = e.Error.(*gitalypb.GetTreeEntriesError_ResolveTree); ok {
				return true
			}
		}
	}
	return false
}

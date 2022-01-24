// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: internal/module/agent_tracker/rpc/rpc.proto

package rpc

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GetConnectedAgentsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetConnectedAgentsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetConnectedAgentsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetConnectedAgentsRequestMultiError, or nil if none found.
func (m *GetConnectedAgentsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetConnectedAgentsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.Request.(type) {

	case *GetConnectedAgentsRequest_ProjectId:
		// no validation rules for ProjectId

	case *GetConnectedAgentsRequest_AgentId:
		// no validation rules for AgentId

	default:
		err := GetConnectedAgentsRequestValidationError{
			field:  "Request",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return GetConnectedAgentsRequestMultiError(errors)
	}

	return nil
}

// GetConnectedAgentsRequestMultiError is an error wrapping multiple validation
// errors returned by GetConnectedAgentsRequest.ValidateAll() if the
// designated constraints aren't met.
type GetConnectedAgentsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetConnectedAgentsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetConnectedAgentsRequestMultiError) AllErrors() []error { return m }

// GetConnectedAgentsRequestValidationError is the validation error returned by
// GetConnectedAgentsRequest.Validate if the designated constraints aren't met.
type GetConnectedAgentsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetConnectedAgentsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetConnectedAgentsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetConnectedAgentsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetConnectedAgentsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetConnectedAgentsRequestValidationError) ErrorName() string {
	return "GetConnectedAgentsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetConnectedAgentsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetConnectedAgentsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetConnectedAgentsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetConnectedAgentsRequestValidationError{}

// Validate checks the field values on GetConnectedAgentsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetConnectedAgentsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetConnectedAgentsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetConnectedAgentsResponseMultiError, or nil if none found.
func (m *GetConnectedAgentsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetConnectedAgentsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetAgents() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetConnectedAgentsResponseValidationError{
						field:  fmt.Sprintf("Agents[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetConnectedAgentsResponseValidationError{
						field:  fmt.Sprintf("Agents[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetConnectedAgentsResponseValidationError{
					field:  fmt.Sprintf("Agents[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetConnectedAgentsResponseMultiError(errors)
	}

	return nil
}

// GetConnectedAgentsResponseMultiError is an error wrapping multiple
// validation errors returned by GetConnectedAgentsResponse.ValidateAll() if
// the designated constraints aren't met.
type GetConnectedAgentsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetConnectedAgentsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetConnectedAgentsResponseMultiError) AllErrors() []error { return m }

// GetConnectedAgentsResponseValidationError is the validation error returned
// by GetConnectedAgentsResponse.Validate if the designated constraints aren't met.
type GetConnectedAgentsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetConnectedAgentsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetConnectedAgentsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetConnectedAgentsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetConnectedAgentsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetConnectedAgentsResponseValidationError) ErrorName() string {
	return "GetConnectedAgentsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetConnectedAgentsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetConnectedAgentsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetConnectedAgentsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetConnectedAgentsResponseValidationError{}

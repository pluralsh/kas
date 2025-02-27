// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: internal/module/notifications/rpc/rpc.proto

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

// Validate checks the field values on GitPushEventRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GitPushEventRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GitPushEventRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GitPushEventRequestMultiError, or nil if none found.
func (m *GitPushEventRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GitPushEventRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetEvent() == nil {
		err := GitPushEventRequestValidationError{
			field:  "Event",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetEvent()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GitPushEventRequestValidationError{
					field:  "Event",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GitPushEventRequestValidationError{
					field:  "Event",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEvent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GitPushEventRequestValidationError{
				field:  "Event",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GitPushEventRequestMultiError(errors)
	}

	return nil
}

// GitPushEventRequestMultiError is an error wrapping multiple validation
// errors returned by GitPushEventRequest.ValidateAll() if the designated
// constraints aren't met.
type GitPushEventRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GitPushEventRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GitPushEventRequestMultiError) AllErrors() []error { return m }

// GitPushEventRequestValidationError is the validation error returned by
// GitPushEventRequest.Validate if the designated constraints aren't met.
type GitPushEventRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GitPushEventRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GitPushEventRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GitPushEventRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GitPushEventRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GitPushEventRequestValidationError) ErrorName() string {
	return "GitPushEventRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GitPushEventRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGitPushEventRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GitPushEventRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GitPushEventRequestValidationError{}

// Validate checks the field values on GitPushEventResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GitPushEventResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GitPushEventResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GitPushEventResponseMultiError, or nil if none found.
func (m *GitPushEventResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GitPushEventResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GitPushEventResponseMultiError(errors)
	}

	return nil
}

// GitPushEventResponseMultiError is an error wrapping multiple validation
// errors returned by GitPushEventResponse.ValidateAll() if the designated
// constraints aren't met.
type GitPushEventResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GitPushEventResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GitPushEventResponseMultiError) AllErrors() []error { return m }

// GitPushEventResponseValidationError is the validation error returned by
// GitPushEventResponse.Validate if the designated constraints aren't met.
type GitPushEventResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GitPushEventResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GitPushEventResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GitPushEventResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GitPushEventResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GitPushEventResponseValidationError) ErrorName() string {
	return "GitPushEventResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GitPushEventResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGitPushEventResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GitPushEventResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GitPushEventResponseValidationError{}

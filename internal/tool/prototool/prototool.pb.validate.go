// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: internal/tool/prototool/prototool.proto

package prototool

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on Values with the rules defined in the
// proto definition for this message. If any rules are violated, an error is
// returned. When asked to return all errors, validation continues after first
// violation, and the result is a list of violation errors wrapped in
// ValuesMultiError, or nil if none found. Otherwise, only the first error is
// returned, if any.
func (m *Values) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ValuesMultiError(errors)
	}
	return nil
}

// ValuesMultiError is an error wrapping multiple validation errors returned by
// Values.Validate(true) if the designated constraints aren't met.
type ValuesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ValuesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ValuesMultiError) AllErrors() []error { return m }

// ValuesValidationError is the validation error returned by Values.Validate if
// the designated constraints aren't met.
type ValuesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValuesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValuesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValuesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValuesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValuesValidationError) ErrorName() string { return "ValuesValidationError" }

// Error satisfies the builtin error interface
func (e ValuesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValues.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValuesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValuesValidationError{}

// Validate checks the field values on HttpRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned. When asked to return all errors, validation continues after
// first violation, and the result is a list of violation errors wrapped in
// HttpRequestMultiError, or nil if none found. Otherwise, only the first
// error is returned, if any.
func (m *HttpRequest) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetMethod()) < 1 {
		err := HttpRequestValidationError{
			field:  "Method",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for key, val := range m.GetHeader() {
		_ = val

		// no validation rules for Header[key]

		if v, ok := interface{}(val).(interface{ Validate(bool) error }); ok {
			if err := v.Validate(all); err != nil {
				err = HttpRequestValidationError{
					field:  fmt.Sprintf("Header[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}
		}

	}

	if utf8.RuneCountInString(m.GetUrlPath()) < 1 {
		err := HttpRequestValidationError{
			field:  "UrlPath",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for key, val := range m.GetQuery() {
		_ = val

		// no validation rules for Query[key]

		if v, ok := interface{}(val).(interface{ Validate(bool) error }); ok {
			if err := v.Validate(all); err != nil {
				err = HttpRequestValidationError{
					field:  fmt.Sprintf("Query[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}
		}

	}

	if len(errors) > 0 {
		return HttpRequestMultiError(errors)
	}
	return nil
}

// HttpRequestMultiError is an error wrapping multiple validation errors
// returned by HttpRequest.Validate(true) if the designated constraints aren't met.
type HttpRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HttpRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HttpRequestMultiError) AllErrors() []error { return m }

// HttpRequestValidationError is the validation error returned by
// HttpRequest.Validate if the designated constraints aren't met.
type HttpRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpRequestValidationError) ErrorName() string { return "HttpRequestValidationError" }

// Error satisfies the builtin error interface
func (e HttpRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpRequestValidationError{}

// Validate checks the field values on HttpResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned. When asked to return all errors, validation continues after
// first violation, and the result is a list of violation errors wrapped in
// HttpResponseMultiError, or nil if none found. Otherwise, only the first
// error is returned, if any.
func (m *HttpResponse) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StatusCode

	// no validation rules for Status

	for key, val := range m.GetHeader() {
		_ = val

		// no validation rules for Header[key]

		if v, ok := interface{}(val).(interface{ Validate(bool) error }); ok {
			if err := v.Validate(all); err != nil {
				err = HttpResponseValidationError{
					field:  fmt.Sprintf("Header[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}
		}

	}

	if len(errors) > 0 {
		return HttpResponseMultiError(errors)
	}
	return nil
}

// HttpResponseMultiError is an error wrapping multiple validation errors
// returned by HttpResponse.Validate(true) if the designated constraints
// aren't met.
type HttpResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HttpResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HttpResponseMultiError) AllErrors() []error { return m }

// HttpResponseValidationError is the validation error returned by
// HttpResponse.Validate if the designated constraints aren't met.
type HttpResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpResponseValidationError) ErrorName() string { return "HttpResponseValidationError" }

// Error satisfies the builtin error interface
func (e HttpResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpResponseValidationError{}

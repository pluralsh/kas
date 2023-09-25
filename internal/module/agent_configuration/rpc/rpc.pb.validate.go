// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: internal/module/agent_configuration/rpc/rpc.proto

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

// Validate checks the field values on ConfigurationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ConfigurationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConfigurationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConfigurationRequestMultiError, or nil if none found.
func (m *ConfigurationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ConfigurationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CommitId

	if all {
		switch v := interface{}(m.GetAgentMeta()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfigurationRequestValidationError{
					field:  "AgentMeta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfigurationRequestValidationError{
					field:  "AgentMeta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAgentMeta()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigurationRequestValidationError{
				field:  "AgentMeta",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for SkipRegister

	if len(errors) > 0 {
		return ConfigurationRequestMultiError(errors)
	}

	return nil
}

// ConfigurationRequestMultiError is an error wrapping multiple validation
// errors returned by ConfigurationRequest.ValidateAll() if the designated
// constraints aren't met.
type ConfigurationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigurationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigurationRequestMultiError) AllErrors() []error { return m }

// ConfigurationRequestValidationError is the validation error returned by
// ConfigurationRequest.Validate if the designated constraints aren't met.
type ConfigurationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigurationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigurationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigurationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigurationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigurationRequestValidationError) ErrorName() string {
	return "ConfigurationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ConfigurationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigurationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigurationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigurationRequestValidationError{}

// Validate checks the field values on ConfigurationResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ConfigurationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConfigurationResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConfigurationResponseMultiError, or nil if none found.
func (m *ConfigurationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ConfigurationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetConfiguration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfigurationResponseValidationError{
					field:  "Configuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfigurationResponseValidationError{
					field:  "Configuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfiguration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigurationResponseValidationError{
				field:  "Configuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetCommitId()) < 1 {
		err := ConfigurationResponseValidationError{
			field:  "CommitId",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ConfigurationResponseMultiError(errors)
	}

	return nil
}

// ConfigurationResponseMultiError is an error wrapping multiple validation
// errors returned by ConfigurationResponse.ValidateAll() if the designated
// constraints aren't met.
type ConfigurationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigurationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigurationResponseMultiError) AllErrors() []error { return m }

// ConfigurationResponseValidationError is the validation error returned by
// ConfigurationResponse.Validate if the designated constraints aren't met.
type ConfigurationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigurationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigurationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigurationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigurationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigurationResponseValidationError) ErrorName() string {
	return "ConfigurationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ConfigurationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigurationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigurationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigurationResponseValidationError{}

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/entity/entity.proto

package entity

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

// Validate checks the field values on AgentMeta with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AgentMeta) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AgentMeta with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AgentMetaMultiError, or nil
// if none found.
func (m *AgentMeta) ValidateAll() error {
	return m.validate(true)
}

func (m *AgentMeta) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Version

	// no validation rules for CommitId

	// no validation rules for PodNamespace

	// no validation rules for PodName

	if len(errors) > 0 {
		return AgentMetaMultiError(errors)
	}

	return nil
}

// AgentMetaMultiError is an error wrapping multiple validation errors returned
// by AgentMeta.ValidateAll() if the designated constraints aren't met.
type AgentMetaMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AgentMetaMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AgentMetaMultiError) AllErrors() []error { return m }

// AgentMetaValidationError is the validation error returned by
// AgentMeta.Validate if the designated constraints aren't met.
type AgentMetaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AgentMetaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AgentMetaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AgentMetaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AgentMetaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AgentMetaValidationError) ErrorName() string { return "AgentMetaValidationError" }

// Error satisfies the builtin error interface
func (e AgentMetaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAgentMeta.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AgentMetaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AgentMetaValidationError{}

// Validate checks the field values on GitalyInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GitalyInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GitalyInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GitalyInfoMultiError, or
// nil if none found.
func (m *GitalyInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *GitalyInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetAddress()) < 1 {
		err := GitalyInfoValidationError{
			field:  "Address",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Token

	// no validation rules for Features

	if len(errors) > 0 {
		return GitalyInfoMultiError(errors)
	}

	return nil
}

// GitalyInfoMultiError is an error wrapping multiple validation errors
// returned by GitalyInfo.ValidateAll() if the designated constraints aren't met.
type GitalyInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GitalyInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GitalyInfoMultiError) AllErrors() []error { return m }

// GitalyInfoValidationError is the validation error returned by
// GitalyInfo.Validate if the designated constraints aren't met.
type GitalyInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GitalyInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GitalyInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GitalyInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GitalyInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GitalyInfoValidationError) ErrorName() string { return "GitalyInfoValidationError" }

// Error satisfies the builtin error interface
func (e GitalyInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGitalyInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GitalyInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GitalyInfoValidationError{}

// Validate checks the field values on GitalyRepository with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GitalyRepository) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GitalyRepository with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GitalyRepositoryMultiError, or nil if none found.
func (m *GitalyRepository) ValidateAll() error {
	return m.validate(true)
}

func (m *GitalyRepository) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetStorageName()) < 1 {
		err := GitalyRepositoryValidationError{
			field:  "StorageName",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetRelativePath()) < 1 {
		err := GitalyRepositoryValidationError{
			field:  "RelativePath",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for GitObjectDirectory

	if len(m.GetGlRepository()) < 1 {
		err := GitalyRepositoryValidationError{
			field:  "GlRepository",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetGlProjectPath()) < 1 {
		err := GitalyRepositoryValidationError{
			field:  "GlProjectPath",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GitalyRepositoryMultiError(errors)
	}

	return nil
}

// GitalyRepositoryMultiError is an error wrapping multiple validation errors
// returned by GitalyRepository.ValidateAll() if the designated constraints
// aren't met.
type GitalyRepositoryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GitalyRepositoryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GitalyRepositoryMultiError) AllErrors() []error { return m }

// GitalyRepositoryValidationError is the validation error returned by
// GitalyRepository.Validate if the designated constraints aren't met.
type GitalyRepositoryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GitalyRepositoryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GitalyRepositoryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GitalyRepositoryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GitalyRepositoryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GitalyRepositoryValidationError) ErrorName() string { return "GitalyRepositoryValidationError" }

// Error satisfies the builtin error interface
func (e GitalyRepositoryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGitalyRepository.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GitalyRepositoryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GitalyRepositoryValidationError{}

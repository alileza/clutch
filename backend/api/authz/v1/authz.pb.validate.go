// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: authz/v1/authz.proto

package authzv1

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

	apiv1 "github.com/lyft/clutch/backend/api/api/v1"
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

	_ = apiv1.ActionType(0)
)

// Validate checks the field values on Subject with the rules defined in the
// proto definition for this message. If any rules are violated, an error is
// returned. When asked to return all errors, validation continues after first
// violation, and the result is a list of violation errors wrapped in
// SubjectMultiError, or nil if none found. Otherwise, only the first error is
// returned, if any.
func (m *Subject) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for User

	if len(errors) > 0 {
		return SubjectMultiError(errors)
	}
	return nil
}

// SubjectMultiError is an error wrapping multiple validation errors returned
// by Subject.Validate(true) if the designated constraints aren't met.
type SubjectMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubjectMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubjectMultiError) AllErrors() []error { return m }

// SubjectValidationError is the validation error returned by Subject.Validate
// if the designated constraints aren't met.
type SubjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubjectValidationError) ErrorName() string { return "SubjectValidationError" }

// Error satisfies the builtin error interface
func (e SubjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubjectValidationError{}

// Validate checks the field values on CheckRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned. When asked to return all errors, validation continues after
// first violation, and the result is a list of violation errors wrapped in
// CheckRequestMultiError, or nil if none found. Otherwise, only the first
// error is returned, if any.
func (m *CheckRequest) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetSubject() == nil {
		err := CheckRequestValidationError{
			field:  "Subject",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if v, ok := interface{}(m.GetSubject()).(interface{ Validate(bool) error }); ok {
		if err := v.Validate(all); err != nil {
			err = CheckRequestValidationError{
				field:  "Subject",
				reason: "embedded message failed validation",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
	}

	// no validation rules for Method

	// no validation rules for ActionType

	// no validation rules for Resource

	if len(errors) > 0 {
		return CheckRequestMultiError(errors)
	}
	return nil
}

// CheckRequestMultiError is an error wrapping multiple validation errors
// returned by CheckRequest.Validate(true) if the designated constraints
// aren't met.
type CheckRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckRequestMultiError) AllErrors() []error { return m }

// CheckRequestValidationError is the validation error returned by
// CheckRequest.Validate if the designated constraints aren't met.
type CheckRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckRequestValidationError) ErrorName() string { return "CheckRequestValidationError" }

// Error satisfies the builtin error interface
func (e CheckRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckRequestValidationError{}

// Validate checks the field values on CheckResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned. When asked to return all errors, validation continues after
// first violation, and the result is a list of violation errors wrapped in
// CheckResponseMultiError, or nil if none found. Otherwise, only the first
// error is returned, if any.
func (m *CheckResponse) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Decision

	if len(errors) > 0 {
		return CheckResponseMultiError(errors)
	}
	return nil
}

// CheckResponseMultiError is an error wrapping multiple validation errors
// returned by CheckResponse.Validate(true) if the designated constraints
// aren't met.
type CheckResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckResponseMultiError) AllErrors() []error { return m }

// CheckResponseValidationError is the validation error returned by
// CheckResponse.Validate if the designated constraints aren't met.
type CheckResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResponseValidationError) ErrorName() string { return "CheckResponseValidationError" }

// Error satisfies the builtin error interface
func (e CheckResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResponseValidationError{}

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: chaos/experimentation/v1/create_experiment_data.proto

package experimentationv1

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

// Validate checks the field values on CreateExperimentData with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned. When asked to return all errors, validation
// continues after first violation, and the result is a list of violation
// errors wrapped in CreateExperimentDataMultiError, or nil if none found.
// Otherwise, only the first error is returned, if any.
func (m *CreateExperimentData) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetRunId()) > 100 {
		err := CreateExperimentDataValidationError{
			field:  "RunId",
			reason: "value length must be at most 100 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_CreateExperimentData_RunId_Pattern.MatchString(m.GetRunId()) {
		err := CreateExperimentDataValidationError{
			field:  "RunId",
			reason: "value does not match regex pattern \"^[A-Za-z0-9-._~]*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetConfig() == nil {
		err := CreateExperimentDataValidationError{
			field:  "Config",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if a := m.GetConfig(); a != nil {

	}

	if v, ok := interface{}(m.GetStartTime()).(interface{ Validate(bool) error }); ok {
		if err := v.Validate(all); err != nil {
			err = CreateExperimentDataValidationError{
				field:  "StartTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
	}

	if v, ok := interface{}(m.GetEndTime()).(interface{ Validate(bool) error }); ok {
		if err := v.Validate(all); err != nil {
			err = CreateExperimentDataValidationError{
				field:  "EndTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return CreateExperimentDataMultiError(errors)
	}
	return nil
}

// CreateExperimentDataMultiError is an error wrapping multiple validation
// errors returned by CreateExperimentData.Validate(true) if the designated
// constraints aren't met.
type CreateExperimentDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateExperimentDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateExperimentDataMultiError) AllErrors() []error { return m }

// CreateExperimentDataValidationError is the validation error returned by
// CreateExperimentData.Validate if the designated constraints aren't met.
type CreateExperimentDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateExperimentDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateExperimentDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateExperimentDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateExperimentDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateExperimentDataValidationError) ErrorName() string {
	return "CreateExperimentDataValidationError"
}

// Error satisfies the builtin error interface
func (e CreateExperimentDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateExperimentData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateExperimentDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateExperimentDataValidationError{}

var _CreateExperimentData_RunId_Pattern = regexp.MustCompile("^[A-Za-z0-9-._~]*$")

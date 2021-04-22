// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: config/service/chaos/experimentation/terminator/v1/termination.proto

package terminatorv1

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

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, an error is
// returned. When asked to return all errors, validation continues after first
// violation, and the result is a list of violation errors wrapped in
// ConfigMultiError, or nil if none found. Otherwise, only the first error is
// returned, if any.
func (m *Config) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for key, val := range m.GetPerConfigTypeConfiguration() {
		_ = val

		// no validation rules for PerConfigTypeConfiguration[key]

		if v, ok := interface{}(val).(interface{ Validate(bool) error }); ok {
			if err := v.Validate(all); err != nil {
				err = ConfigValidationError{
					field:  fmt.Sprintf("PerConfigTypeConfiguration[%v]", key),
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

	if d := m.GetOuterLoopInterval(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = ConfigValidationError{
				field:  "OuterLoopInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				err := ConfigValidationError{
					field:  "OuterLoopInterval",
					reason: "value must be greater than 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if d := m.GetPerExperimentCheckInterval(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = ConfigValidationError{
				field:  "PerExperimentCheckInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				err := ConfigValidationError{
					field:  "PerExperimentCheckInterval",
					reason: "value must be greater than 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if len(errors) > 0 {
		return ConfigMultiError(errors)
	}
	return nil
}

// ConfigMultiError is an error wrapping multiple validation errors returned by
// Config.Validate(true) if the designated constraints aren't met.
type ConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigMultiError) AllErrors() []error { return m }

// ConfigValidationError is the validation error returned by Config.Validate if
// the designated constraints aren't met.
type ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigValidationError) ErrorName() string { return "ConfigValidationError" }

// Error satisfies the builtin error interface
func (e ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigValidationError{}

// Validate checks the field values on MaxTimeTerminationCriterion with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned. When asked to return all errors, validation
// continues after first violation, and the result is a list of violation
// errors wrapped in MaxTimeTerminationCriterionMultiError, or nil if none
// found. Otherwise, only the first error is returned, if any.
func (m *MaxTimeTerminationCriterion) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if d := m.GetMaxDuration(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = MaxTimeTerminationCriterionValidationError{
				field:  "MaxDuration",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				err := MaxTimeTerminationCriterionValidationError{
					field:  "MaxDuration",
					reason: "value must be greater than 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if len(errors) > 0 {
		return MaxTimeTerminationCriterionMultiError(errors)
	}
	return nil
}

// MaxTimeTerminationCriterionMultiError is an error wrapping multiple
// validation errors returned by MaxTimeTerminationCriterion.Validate(true) if
// the designated constraints aren't met.
type MaxTimeTerminationCriterionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MaxTimeTerminationCriterionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MaxTimeTerminationCriterionMultiError) AllErrors() []error { return m }

// MaxTimeTerminationCriterionValidationError is the validation error returned
// by MaxTimeTerminationCriterion.Validate if the designated constraints
// aren't met.
type MaxTimeTerminationCriterionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MaxTimeTerminationCriterionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MaxTimeTerminationCriterionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MaxTimeTerminationCriterionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MaxTimeTerminationCriterionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MaxTimeTerminationCriterionValidationError) ErrorName() string {
	return "MaxTimeTerminationCriterionValidationError"
}

// Error satisfies the builtin error interface
func (e MaxTimeTerminationCriterionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMaxTimeTerminationCriterion.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MaxTimeTerminationCriterionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MaxTimeTerminationCriterionValidationError{}

// Validate checks the field values on Config_PerConfigTypeConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned. When asked to return all errors, validation
// continues after first violation, and the result is a list of violation
// errors wrapped in Config_PerConfigTypeConfigMultiError, or nil if none
// found. Otherwise, only the first error is returned, if any.
func (m *Config_PerConfigTypeConfig) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetTerminationCriteria()) < 1 {
		err := Config_PerConfigTypeConfigValidationError{
			field:  "TerminationCriteria",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetTerminationCriteria() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate(bool) error }); ok {
			if err := v.Validate(all); err != nil {
				err = Config_PerConfigTypeConfigValidationError{
					field:  fmt.Sprintf("TerminationCriteria[%v]", idx),
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
		return Config_PerConfigTypeConfigMultiError(errors)
	}
	return nil
}

// Config_PerConfigTypeConfigMultiError is an error wrapping multiple
// validation errors returned by Config_PerConfigTypeConfig.Validate(true) if
// the designated constraints aren't met.
type Config_PerConfigTypeConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Config_PerConfigTypeConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Config_PerConfigTypeConfigMultiError) AllErrors() []error { return m }

// Config_PerConfigTypeConfigValidationError is the validation error returned
// by Config_PerConfigTypeConfig.Validate if the designated constraints aren't met.
type Config_PerConfigTypeConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Config_PerConfigTypeConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Config_PerConfigTypeConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Config_PerConfigTypeConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Config_PerConfigTypeConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Config_PerConfigTypeConfigValidationError) ErrorName() string {
	return "Config_PerConfigTypeConfigValidationError"
}

// Error satisfies the builtin error interface
func (e Config_PerConfigTypeConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig_PerConfigTypeConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Config_PerConfigTypeConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Config_PerConfigTypeConfigValidationError{}

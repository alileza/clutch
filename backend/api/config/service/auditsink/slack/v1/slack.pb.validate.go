// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: config/service/auditsink/slack/v1/slack.proto

package slackv1

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

// Validate checks the field values on SlackConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned. When asked to return all errors, validation continues after
// first violation, and the result is a list of violation errors wrapped in
// SlackConfigMultiError, or nil if none found. Otherwise, only the first
// error is returned, if any.
func (m *SlackConfig) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetToken()) < 1 {
		err := SlackConfigValidationError{
			field:  "Token",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetChannel()) < 1 {
		err := SlackConfigValidationError{
			field:  "Channel",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if v, ok := interface{}(m.GetFilter()).(interface{ Validate(bool) error }); ok {
		if err := v.Validate(all); err != nil {
			err = SlackConfigValidationError{
				field:  "Filter",
				reason: "embedded message failed validation",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
	}

	for idx, item := range m.GetOverrides() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate(bool) error }); ok {
			if err := v.Validate(all); err != nil {
				err = SlackConfigValidationError{
					field:  fmt.Sprintf("Overrides[%v]", idx),
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
		return SlackConfigMultiError(errors)
	}
	return nil
}

// SlackConfigMultiError is an error wrapping multiple validation errors
// returned by SlackConfig.Validate(true) if the designated constraints aren't met.
type SlackConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SlackConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SlackConfigMultiError) AllErrors() []error { return m }

// SlackConfigValidationError is the validation error returned by
// SlackConfig.Validate if the designated constraints aren't met.
type SlackConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SlackConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SlackConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SlackConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SlackConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SlackConfigValidationError) ErrorName() string { return "SlackConfigValidationError" }

// Error satisfies the builtin error interface
func (e SlackConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSlackConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SlackConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SlackConfigValidationError{}

// Validate checks the field values on CustomMessage with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned. When asked to return all errors, validation continues after
// first violation, and the result is a list of violation errors wrapped in
// CustomMessageMultiError, or nil if none found. Otherwise, only the first
// error is returned, if any.
func (m *CustomMessage) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetFullMethod()) < 1 {
		err := CustomMessageValidationError{
			field:  "FullMethod",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetMessage()) < 1 {
		err := CustomMessageValidationError{
			field:  "Message",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CustomMessageMultiError(errors)
	}
	return nil
}

// CustomMessageMultiError is an error wrapping multiple validation errors
// returned by CustomMessage.Validate(true) if the designated constraints
// aren't met.
type CustomMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CustomMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CustomMessageMultiError) AllErrors() []error { return m }

// CustomMessageValidationError is the validation error returned by
// CustomMessage.Validate if the designated constraints aren't met.
type CustomMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CustomMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CustomMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CustomMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CustomMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CustomMessageValidationError) ErrorName() string { return "CustomMessageValidationError" }

// Error satisfies the builtin error interface
func (e CustomMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CustomMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CustomMessageValidationError{}

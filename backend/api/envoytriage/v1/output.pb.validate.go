// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoytriage/v1/output.proto

package envoytriagev1

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

// Validate checks the field values on HostStatus with the rules defined in the
// proto definition for this message. If any rules are violated, an error is
// returned. When asked to return all errors, validation continues after first
// violation, and the result is a list of violation errors wrapped in
// HostStatusMultiError, or nil if none found. Otherwise, only the first error
// is returned, if any.
func (m *HostStatus) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	// no validation rules for Healthy

	if len(errors) > 0 {
		return HostStatusMultiError(errors)
	}
	return nil
}

// HostStatusMultiError is an error wrapping multiple validation errors
// returned by HostStatus.Validate(true) if the designated constraints aren't met.
type HostStatusMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HostStatusMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HostStatusMultiError) AllErrors() []error { return m }

// HostStatusValidationError is the validation error returned by
// HostStatus.Validate if the designated constraints aren't met.
type HostStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HostStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HostStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HostStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HostStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HostStatusValidationError) ErrorName() string { return "HostStatusValidationError" }

// Error satisfies the builtin error interface
func (e HostStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHostStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HostStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HostStatusValidationError{}

// Validate checks the field values on ClusterStatus with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned. When asked to return all errors, validation continues after
// first violation, and the result is a list of violation errors wrapped in
// ClusterStatusMultiError, or nil if none found. Otherwise, only the first
// error is returned, if any.
func (m *ClusterStatus) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	for idx, item := range m.GetHostStatuses() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate(bool) error }); ok {
			if err := v.Validate(all); err != nil {
				err = ClusterStatusValidationError{
					field:  fmt.Sprintf("HostStatuses[%v]", idx),
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
		return ClusterStatusMultiError(errors)
	}
	return nil
}

// ClusterStatusMultiError is an error wrapping multiple validation errors
// returned by ClusterStatus.Validate(true) if the designated constraints
// aren't met.
type ClusterStatusMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClusterStatusMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClusterStatusMultiError) AllErrors() []error { return m }

// ClusterStatusValidationError is the validation error returned by
// ClusterStatus.Validate if the designated constraints aren't met.
type ClusterStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterStatusValidationError) ErrorName() string { return "ClusterStatusValidationError" }

// Error satisfies the builtin error interface
func (e ClusterStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterStatusValidationError{}

// Validate checks the field values on Clusters with the rules defined in the
// proto definition for this message. If any rules are violated, an error is
// returned. When asked to return all errors, validation continues after first
// violation, and the result is a list of violation errors wrapped in
// ClustersMultiError, or nil if none found. Otherwise, only the first error
// is returned, if any.
func (m *Clusters) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetClusterStatuses() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate(bool) error }); ok {
			if err := v.Validate(all); err != nil {
				err = ClustersValidationError{
					field:  fmt.Sprintf("ClusterStatuses[%v]", idx),
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
		return ClustersMultiError(errors)
	}
	return nil
}

// ClustersMultiError is an error wrapping multiple validation errors returned
// by Clusters.Validate(true) if the designated constraints aren't met.
type ClustersMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClustersMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClustersMultiError) AllErrors() []error { return m }

// ClustersValidationError is the validation error returned by
// Clusters.Validate if the designated constraints aren't met.
type ClustersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClustersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClustersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClustersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClustersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClustersValidationError) ErrorName() string { return "ClustersValidationError" }

// Error satisfies the builtin error interface
func (e ClustersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusters.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClustersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClustersValidationError{}

// Validate checks the field values on ConfigDump with the rules defined in the
// proto definition for this message. If any rules are violated, an error is
// returned. When asked to return all errors, validation continues after first
// violation, and the result is a list of violation errors wrapped in
// ConfigDumpMultiError, or nil if none found. Otherwise, only the first error
// is returned, if any.
func (m *ConfigDump) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if v, ok := interface{}(m.GetValue()).(interface{ Validate(bool) error }); ok {
		if err := v.Validate(all); err != nil {
			err = ConfigDumpValidationError{
				field:  "Value",
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
		return ConfigDumpMultiError(errors)
	}
	return nil
}

// ConfigDumpMultiError is an error wrapping multiple validation errors
// returned by ConfigDump.Validate(true) if the designated constraints aren't met.
type ConfigDumpMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigDumpMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigDumpMultiError) AllErrors() []error { return m }

// ConfigDumpValidationError is the validation error returned by
// ConfigDump.Validate if the designated constraints aren't met.
type ConfigDumpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigDumpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigDumpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigDumpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigDumpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigDumpValidationError) ErrorName() string { return "ConfigDumpValidationError" }

// Error satisfies the builtin error interface
func (e ConfigDumpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigDump.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigDumpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigDumpValidationError{}

// Validate checks the field values on ListenerStatus with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned. When asked to return all errors, validation continues after
// first violation, and the result is a list of violation errors wrapped in
// ListenerStatusMultiError, or nil if none found. Otherwise, only the first
// error is returned, if any.
func (m *ListenerStatus) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for LocalAddress

	if len(errors) > 0 {
		return ListenerStatusMultiError(errors)
	}
	return nil
}

// ListenerStatusMultiError is an error wrapping multiple validation errors
// returned by ListenerStatus.Validate(true) if the designated constraints
// aren't met.
type ListenerStatusMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListenerStatusMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListenerStatusMultiError) AllErrors() []error { return m }

// ListenerStatusValidationError is the validation error returned by
// ListenerStatus.Validate if the designated constraints aren't met.
type ListenerStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListenerStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListenerStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListenerStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListenerStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListenerStatusValidationError) ErrorName() string { return "ListenerStatusValidationError" }

// Error satisfies the builtin error interface
func (e ListenerStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListenerStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListenerStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListenerStatusValidationError{}

// Validate checks the field values on Listeners with the rules defined in the
// proto definition for this message. If any rules are violated, an error is
// returned. When asked to return all errors, validation continues after first
// violation, and the result is a list of violation errors wrapped in
// ListenersMultiError, or nil if none found. Otherwise, only the first error
// is returned, if any.
func (m *Listeners) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetListenerStatuses() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate(bool) error }); ok {
			if err := v.Validate(all); err != nil {
				err = ListenersValidationError{
					field:  fmt.Sprintf("ListenerStatuses[%v]", idx),
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
		return ListenersMultiError(errors)
	}
	return nil
}

// ListenersMultiError is an error wrapping multiple validation errors returned
// by Listeners.Validate(true) if the designated constraints aren't met.
type ListenersMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListenersMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListenersMultiError) AllErrors() []error { return m }

// ListenersValidationError is the validation error returned by
// Listeners.Validate if the designated constraints aren't met.
type ListenersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListenersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListenersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListenersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListenersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListenersValidationError) ErrorName() string { return "ListenersValidationError" }

// Error satisfies the builtin error interface
func (e ListenersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListeners.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListenersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListenersValidationError{}

// Validate checks the field values on Runtime with the rules defined in the
// proto definition for this message. If any rules are violated, an error is
// returned. When asked to return all errors, validation continues after first
// violation, and the result is a list of violation errors wrapped in
// RuntimeMultiError, or nil if none found. Otherwise, only the first error is
// returned, if any.
func (m *Runtime) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetEntries() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate(bool) error }); ok {
			if err := v.Validate(all); err != nil {
				err = RuntimeValidationError{
					field:  fmt.Sprintf("Entries[%v]", idx),
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
		return RuntimeMultiError(errors)
	}
	return nil
}

// RuntimeMultiError is an error wrapping multiple validation errors returned
// by Runtime.Validate(true) if the designated constraints aren't met.
type RuntimeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RuntimeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RuntimeMultiError) AllErrors() []error { return m }

// RuntimeValidationError is the validation error returned by Runtime.Validate
// if the designated constraints aren't met.
type RuntimeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RuntimeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RuntimeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RuntimeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RuntimeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RuntimeValidationError) ErrorName() string { return "RuntimeValidationError" }

// Error satisfies the builtin error interface
func (e RuntimeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRuntime.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RuntimeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RuntimeValidationError{}

// Validate checks the field values on ServerInfo with the rules defined in the
// proto definition for this message. If any rules are violated, an error is
// returned. When asked to return all errors, validation continues after first
// violation, and the result is a list of violation errors wrapped in
// ServerInfoMultiError, or nil if none found. Otherwise, only the first error
// is returned, if any.
func (m *ServerInfo) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if v, ok := interface{}(m.GetValue()).(interface{ Validate(bool) error }); ok {
		if err := v.Validate(all); err != nil {
			err = ServerInfoValidationError{
				field:  "Value",
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
		return ServerInfoMultiError(errors)
	}
	return nil
}

// ServerInfoMultiError is an error wrapping multiple validation errors
// returned by ServerInfo.Validate(true) if the designated constraints aren't met.
type ServerInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServerInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServerInfoMultiError) AllErrors() []error { return m }

// ServerInfoValidationError is the validation error returned by
// ServerInfo.Validate if the designated constraints aren't met.
type ServerInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerInfoValidationError) ErrorName() string { return "ServerInfoValidationError" }

// Error satisfies the builtin error interface
func (e ServerInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerInfoValidationError{}

// Validate checks the field values on Stats with the rules defined in the
// proto definition for this message. If any rules are violated, an error is
// returned. When asked to return all errors, validation continues after first
// violation, and the result is a list of violation errors wrapped in
// StatsMultiError, or nil if none found. Otherwise, only the first error is
// returned, if any.
func (m *Stats) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetStats() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate(bool) error }); ok {
			if err := v.Validate(all); err != nil {
				err = StatsValidationError{
					field:  fmt.Sprintf("Stats[%v]", idx),
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
		return StatsMultiError(errors)
	}
	return nil
}

// StatsMultiError is an error wrapping multiple validation errors returned by
// Stats.Validate(true) if the designated constraints aren't met.
type StatsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StatsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StatsMultiError) AllErrors() []error { return m }

// StatsValidationError is the validation error returned by Stats.Validate if
// the designated constraints aren't met.
type StatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatsValidationError) ErrorName() string { return "StatsValidationError" }

// Error satisfies the builtin error interface
func (e StatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatsValidationError{}

// Validate checks the field values on Runtime_Entry with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned. When asked to return all errors, validation continues after
// first violation, and the result is a list of violation errors wrapped in
// Runtime_EntryMultiError, or nil if none found. Otherwise, only the first
// error is returned, if any.
func (m *Runtime_Entry) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	switch m.Type.(type) {

	case *Runtime_Entry_Value:
		// no validation rules for Value

	}

	if len(errors) > 0 {
		return Runtime_EntryMultiError(errors)
	}
	return nil
}

// Runtime_EntryMultiError is an error wrapping multiple validation errors
// returned by Runtime_Entry.Validate(true) if the designated constraints
// aren't met.
type Runtime_EntryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Runtime_EntryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Runtime_EntryMultiError) AllErrors() []error { return m }

// Runtime_EntryValidationError is the validation error returned by
// Runtime_Entry.Validate if the designated constraints aren't met.
type Runtime_EntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Runtime_EntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Runtime_EntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Runtime_EntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Runtime_EntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Runtime_EntryValidationError) ErrorName() string { return "Runtime_EntryValidationError" }

// Error satisfies the builtin error interface
func (e Runtime_EntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRuntime_Entry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Runtime_EntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Runtime_EntryValidationError{}

// Validate checks the field values on Stats_Stat with the rules defined in the
// proto definition for this message. If any rules are violated, an error is
// returned. When asked to return all errors, validation continues after first
// violation, and the result is a list of violation errors wrapped in
// Stats_StatMultiError, or nil if none found. Otherwise, only the first error
// is returned, if any.
func (m *Stats_Stat) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Value

	if len(errors) > 0 {
		return Stats_StatMultiError(errors)
	}
	return nil
}

// Stats_StatMultiError is an error wrapping multiple validation errors
// returned by Stats_Stat.Validate(true) if the designated constraints aren't met.
type Stats_StatMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Stats_StatMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Stats_StatMultiError) AllErrors() []error { return m }

// Stats_StatValidationError is the validation error returned by
// Stats_Stat.Validate if the designated constraints aren't met.
type Stats_StatValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Stats_StatValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Stats_StatValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Stats_StatValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Stats_StatValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Stats_StatValidationError) ErrorName() string { return "Stats_StatValidationError" }

// Error satisfies the builtin error interface
func (e Stats_StatValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStats_Stat.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Stats_StatValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Stats_StatValidationError{}

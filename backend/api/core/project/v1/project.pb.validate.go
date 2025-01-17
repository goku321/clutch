// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: core/project/v1/project.proto

package projectv1

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

// Validate checks the field values on Project with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Project) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Tier

	for key, val := range m.GetData() {
		_ = val

		// no validation rules for Data[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProjectValidationError{
					field:  fmt.Sprintf("Data[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetDependencies()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectValidationError{
				field:  "Dependencies",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetOncall()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectValidationError{
				field:  "Oncall",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ProjectValidationError is the validation error returned by Project.Validate
// if the designated constraints aren't met.
type ProjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectValidationError) ErrorName() string { return "ProjectValidationError" }

// Error satisfies the builtin error interface
func (e ProjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectValidationError{}

// Validate checks the field values on ProjectDependencies with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProjectDependencies) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetUpstreams() {
		_ = val

		// no validation rules for Upstreams[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProjectDependenciesValidationError{
					field:  fmt.Sprintf("Upstreams[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for key, val := range m.GetDownstreams() {
		_ = val

		// no validation rules for Downstreams[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProjectDependenciesValidationError{
					field:  fmt.Sprintf("Downstreams[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ProjectDependenciesValidationError is the validation error returned by
// ProjectDependencies.Validate if the designated constraints aren't met.
type ProjectDependenciesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectDependenciesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectDependenciesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectDependenciesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectDependenciesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectDependenciesValidationError) ErrorName() string {
	return "ProjectDependenciesValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectDependenciesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectDependencies.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectDependenciesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectDependenciesValidationError{}

// Validate checks the field values on Dependency with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Dependency) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DependencyValidationError is the validation error returned by
// Dependency.Validate if the designated constraints aren't met.
type DependencyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DependencyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DependencyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DependencyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DependencyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DependencyValidationError) ErrorName() string { return "DependencyValidationError" }

// Error satisfies the builtin error interface
func (e DependencyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDependency.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DependencyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DependencyValidationError{}

// Validate checks the field values on OnCall with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OnCall) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPagerduty()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OnCallValidationError{
				field:  "Pagerduty",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OnCallValidationError is the validation error returned by OnCall.Validate if
// the designated constraints aren't met.
type OnCallValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OnCallValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OnCallValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OnCallValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OnCallValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OnCallValidationError) ErrorName() string { return "OnCallValidationError" }

// Error satisfies the builtin error interface
func (e OnCallValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOnCall.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OnCallValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OnCallValidationError{}

// Validate checks the field values on PagerDuty with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *PagerDuty) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// PagerDutyValidationError is the validation error returned by
// PagerDuty.Validate if the designated constraints aren't met.
type PagerDutyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PagerDutyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PagerDutyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PagerDutyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PagerDutyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PagerDutyValidationError) ErrorName() string { return "PagerDutyValidationError" }

// Error satisfies the builtin error interface
func (e PagerDutyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPagerDuty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PagerDutyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PagerDutyValidationError{}

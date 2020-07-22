// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: config/service/authn/v1/authn.proto

package authnv1

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _authn_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on OIDC with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *OIDC) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetIssuer()) < 1 {
		return OIDCValidationError{
			field:  "Issuer",
			reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetClientId()) < 1 {
		return OIDCValidationError{
			field:  "ClientId",
			reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetClientSecret()) < 1 {
		return OIDCValidationError{
			field:  "ClientSecret",
			reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetRedirectUrl()) < 1 {
		return OIDCValidationError{
			field:  "RedirectUrl",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// OIDCValidationError is the validation error returned by OIDC.Validate if the
// designated constraints aren't met.
type OIDCValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OIDCValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OIDCValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OIDCValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OIDCValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OIDCValidationError) ErrorName() string { return "OIDCValidationError" }

// Error satisfies the builtin error interface
func (e OIDCValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOIDC.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OIDCValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OIDCValidationError{}

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Config) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetSessionSecret()) < 1 {
		return ConfigValidationError{
			field:  "SessionSecret",
			reason: "value length must be at least 1 bytes",
		}
	}

	switch m.Type.(type) {

	case *Config_Oidc:

		if v, ok := interface{}(m.GetOidc()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigValidationError{
					field:  "Oidc",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

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

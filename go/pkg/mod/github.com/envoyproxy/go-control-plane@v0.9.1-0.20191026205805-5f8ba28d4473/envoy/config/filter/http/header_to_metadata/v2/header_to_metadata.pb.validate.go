// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/header_to_metadata/v2/header_to_metadata.proto

package envoy_config_filter_http_header_to_metadata_v2

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
var _header_to_metadata_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Config) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRequestRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigValidationError{
					field:  fmt.Sprintf("RequestRules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetResponseRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigValidationError{
					field:  fmt.Sprintf("ResponseRules[%v]", idx),
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

// Validate checks the field values on Config_KeyValuePair with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Config_KeyValuePair) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MetadataNamespace

	if len(m.GetKey()) < 1 {
		return Config_KeyValuePairValidationError{
			field:  "Key",
			reason: "value length must be at least 1 bytes",
		}
	}

	// no validation rules for Value

	// no validation rules for Type

	// no validation rules for Encode

	return nil
}

// Config_KeyValuePairValidationError is the validation error returned by
// Config_KeyValuePair.Validate if the designated constraints aren't met.
type Config_KeyValuePairValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Config_KeyValuePairValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Config_KeyValuePairValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Config_KeyValuePairValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Config_KeyValuePairValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Config_KeyValuePairValidationError) ErrorName() string {
	return "Config_KeyValuePairValidationError"
}

// Error satisfies the builtin error interface
func (e Config_KeyValuePairValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig_KeyValuePair.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Config_KeyValuePairValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Config_KeyValuePairValidationError{}

// Validate checks the field values on Config_Rule with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Config_Rule) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetHeader()) < 1 {
		return Config_RuleValidationError{
			field:  "Header",
			reason: "value length must be at least 1 bytes",
		}
	}

	if v, ok := interface{}(m.GetOnHeaderPresent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Config_RuleValidationError{
				field:  "OnHeaderPresent",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetOnHeaderMissing()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Config_RuleValidationError{
				field:  "OnHeaderMissing",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Remove

	return nil
}

// Config_RuleValidationError is the validation error returned by
// Config_Rule.Validate if the designated constraints aren't met.
type Config_RuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Config_RuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Config_RuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Config_RuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Config_RuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Config_RuleValidationError) ErrorName() string { return "Config_RuleValidationError" }

// Error satisfies the builtin error interface
func (e Config_RuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig_Rule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Config_RuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Config_RuleValidationError{}

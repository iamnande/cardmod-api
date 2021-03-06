// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: iamnande/cardmod/magic/v1/api.proto

package magicv1

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

// Validate checks the field values on GetMagicRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetMagicRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetMagicRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetMagicRequestMultiError, or nil if none found.
func (m *GetMagicRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetMagicRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetName()); l < 3 || l > 25 {
		err := GetMagicRequestValidationError{
			field:  "Name",
			reason: "value length must be between 3 and 25 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_GetMagicRequest_Name_Pattern.MatchString(m.GetName()) {
		err := GetMagicRequestValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[-, \\\\w]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetMagicRequestMultiError(errors)
	}

	return nil
}

// GetMagicRequestMultiError is an error wrapping multiple validation errors
// returned by GetMagicRequest.ValidateAll() if the designated constraints
// aren't met.
type GetMagicRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetMagicRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetMagicRequestMultiError) AllErrors() []error { return m }

// GetMagicRequestValidationError is the validation error returned by
// GetMagicRequest.Validate if the designated constraints aren't met.
type GetMagicRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMagicRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMagicRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMagicRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMagicRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMagicRequestValidationError) ErrorName() string { return "GetMagicRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetMagicRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMagicRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMagicRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMagicRequestValidationError{}

var _GetMagicRequest_Name_Pattern = regexp.MustCompile("^[-, \\w]+$")

// Validate checks the field values on ListMagicsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListMagicsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListMagicsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListMagicsRequestMultiError, or nil if none found.
func (m *ListMagicsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListMagicsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListMagicsRequestMultiError(errors)
	}

	return nil
}

// ListMagicsRequestMultiError is an error wrapping multiple validation errors
// returned by ListMagicsRequest.ValidateAll() if the designated constraints
// aren't met.
type ListMagicsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListMagicsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListMagicsRequestMultiError) AllErrors() []error { return m }

// ListMagicsRequestValidationError is the validation error returned by
// ListMagicsRequest.Validate if the designated constraints aren't met.
type ListMagicsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMagicsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMagicsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMagicsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMagicsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMagicsRequestValidationError) ErrorName() string {
	return "ListMagicsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListMagicsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMagicsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMagicsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMagicsRequestValidationError{}

// Validate checks the field values on ListMagicsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListMagicsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListMagicsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListMagicsResponseMultiError, or nil if none found.
func (m *ListMagicsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListMagicsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetMagics() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListMagicsResponseValidationError{
						field:  fmt.Sprintf("Magics[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListMagicsResponseValidationError{
						field:  fmt.Sprintf("Magics[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListMagicsResponseValidationError{
					field:  fmt.Sprintf("Magics[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListMagicsResponseMultiError(errors)
	}

	return nil
}

// ListMagicsResponseMultiError is an error wrapping multiple validation errors
// returned by ListMagicsResponse.ValidateAll() if the designated constraints
// aren't met.
type ListMagicsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListMagicsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListMagicsResponseMultiError) AllErrors() []error { return m }

// ListMagicsResponseValidationError is the validation error returned by
// ListMagicsResponse.Validate if the designated constraints aren't met.
type ListMagicsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMagicsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMagicsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMagicsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMagicsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMagicsResponseValidationError) ErrorName() string {
	return "ListMagicsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListMagicsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMagicsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMagicsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMagicsResponseValidationError{}

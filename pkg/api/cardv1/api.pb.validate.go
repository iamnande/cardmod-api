// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: iamnande/cardmod/card/v1/api.proto

package cardv1

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

// Validate checks the field values on GetCardRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetCardRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCardRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetCardRequestMultiError,
// or nil if none found.
func (m *GetCardRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCardRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetName()); l < 3 || l > 25 {
		err := GetCardRequestValidationError{
			field:  "Name",
			reason: "value length must be between 3 and 25 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_GetCardRequest_Name_Pattern.MatchString(m.GetName()) {
		err := GetCardRequestValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[-, \\\\w]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetCardRequestMultiError(errors)
	}
	return nil
}

// GetCardRequestMultiError is an error wrapping multiple validation errors
// returned by GetCardRequest.ValidateAll() if the designated constraints
// aren't met.
type GetCardRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCardRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCardRequestMultiError) AllErrors() []error { return m }

// GetCardRequestValidationError is the validation error returned by
// GetCardRequest.Validate if the designated constraints aren't met.
type GetCardRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCardRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCardRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCardRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCardRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCardRequestValidationError) ErrorName() string { return "GetCardRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetCardRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCardRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCardRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCardRequestValidationError{}

var _GetCardRequest_Name_Pattern = regexp.MustCompile("^[-, \\w]+$")

// Validate checks the field values on ListCardsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListCardsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCardsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCardsRequestMultiError, or nil if none found.
func (m *ListCardsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCardsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListCardsRequestMultiError(errors)
	}
	return nil
}

// ListCardsRequestMultiError is an error wrapping multiple validation errors
// returned by ListCardsRequest.ValidateAll() if the designated constraints
// aren't met.
type ListCardsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCardsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCardsRequestMultiError) AllErrors() []error { return m }

// ListCardsRequestValidationError is the validation error returned by
// ListCardsRequest.Validate if the designated constraints aren't met.
type ListCardsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCardsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCardsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCardsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCardsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCardsRequestValidationError) ErrorName() string { return "ListCardsRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListCardsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCardsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCardsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCardsRequestValidationError{}

// Validate checks the field values on ListCardsResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListCardsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCardsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCardsResponseMultiError, or nil if none found.
func (m *ListCardsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCardsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCards() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListCardsResponseValidationError{
						field:  fmt.Sprintf("Cards[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListCardsResponseValidationError{
						field:  fmt.Sprintf("Cards[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListCardsResponseValidationError{
					field:  fmt.Sprintf("Cards[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListCardsResponseMultiError(errors)
	}
	return nil
}

// ListCardsResponseMultiError is an error wrapping multiple validation errors
// returned by ListCardsResponse.ValidateAll() if the designated constraints
// aren't met.
type ListCardsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCardsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCardsResponseMultiError) AllErrors() []error { return m }

// ListCardsResponseValidationError is the validation error returned by
// ListCardsResponse.Validate if the designated constraints aren't met.
type ListCardsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCardsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCardsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCardsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCardsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCardsResponseValidationError) ErrorName() string {
	return "ListCardsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListCardsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCardsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCardsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCardsResponseValidationError{}

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: note.proto

package note_v1

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

// Validate checks the field values on CreateNoteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateNoteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateNoteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateNoteRequestMultiError, or nil if none found.
func (m *CreateNoteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateNoteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetTitle()); l < 3 || l > 100 {
		err := CreateNoteRequestValidationError{
			field:  "Title",
			reason: "value length must be between 3 and 100 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetText()); l < 10 || l > 10000 {
		err := CreateNoteRequestValidationError{
			field:  "Text",
			reason: "value length must be between 10 and 10000 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetAuthor()); l < 2 || l > 100 {
		err := CreateNoteRequestValidationError{
			field:  "Author",
			reason: "value length must be between 2 and 100 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		err = CreateNoteRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateNoteRequestMultiError(errors)
	}

	return nil
}

func (m *CreateNoteRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *CreateNoteRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// CreateNoteRequestMultiError is an error wrapping multiple validation errors
// returned by CreateNoteRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateNoteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateNoteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateNoteRequestMultiError) AllErrors() []error { return m }

// CreateNoteRequestValidationError is the validation error returned by
// CreateNoteRequest.Validate if the designated constraints aren't met.
type CreateNoteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNoteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNoteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNoteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNoteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNoteRequestValidationError) ErrorName() string {
	return "CreateNoteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNoteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNoteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNoteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNoteRequestValidationError{}

// Validate checks the field values on CreateNoteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateNoteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateNoteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateNoteResponseMultiError, or nil if none found.
func (m *CreateNoteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateNoteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateNoteResponseMultiError(errors)
	}

	return nil
}

// CreateNoteResponseMultiError is an error wrapping multiple validation errors
// returned by CreateNoteResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateNoteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateNoteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateNoteResponseMultiError) AllErrors() []error { return m }

// CreateNoteResponseValidationError is the validation error returned by
// CreateNoteResponse.Validate if the designated constraints aren't met.
type CreateNoteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNoteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNoteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNoteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNoteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNoteResponseValidationError) ErrorName() string {
	return "CreateNoteResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNoteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNoteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNoteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNoteResponseValidationError{}

// Validate checks the field values on GetNoteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetNoteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNoteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetNoteRequestMultiError,
// or nil if none found.
func (m *GetNoteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNoteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() < 1 {
		err := GetNoteRequestValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetNoteRequestMultiError(errors)
	}

	return nil
}

// GetNoteRequestMultiError is an error wrapping multiple validation errors
// returned by GetNoteRequest.ValidateAll() if the designated constraints
// aren't met.
type GetNoteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNoteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNoteRequestMultiError) AllErrors() []error { return m }

// GetNoteRequestValidationError is the validation error returned by
// GetNoteRequest.Validate if the designated constraints aren't met.
type GetNoteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNoteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNoteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNoteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNoteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNoteRequestValidationError) ErrorName() string { return "GetNoteRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetNoteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNoteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNoteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNoteRequestValidationError{}

// Validate checks the field values on GetNoteResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetNoteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNoteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetNoteResponseMultiError, or nil if none found.
func (m *GetNoteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNoteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Text

	// no validation rules for Author

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetNoteResponseValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetNoteResponseValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetNoteResponseValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetNoteResponseValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetNoteResponseValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetNoteResponseValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetNoteResponseMultiError(errors)
	}

	return nil
}

// GetNoteResponseMultiError is an error wrapping multiple validation errors
// returned by GetNoteResponse.ValidateAll() if the designated constraints
// aren't met.
type GetNoteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNoteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNoteResponseMultiError) AllErrors() []error { return m }

// GetNoteResponseValidationError is the validation error returned by
// GetNoteResponse.Validate if the designated constraints aren't met.
type GetNoteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNoteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNoteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNoteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNoteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNoteResponseValidationError) ErrorName() string { return "GetNoteResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetNoteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNoteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNoteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNoteResponseValidationError{}

// Validate checks the field values on GetListNoteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetListNoteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetListNoteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetListNoteRequestMultiError, or nil if none found.
func (m *GetListNoteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetListNoteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetListNoteRequestMultiError(errors)
	}

	return nil
}

// GetListNoteRequestMultiError is an error wrapping multiple validation errors
// returned by GetListNoteRequest.ValidateAll() if the designated constraints
// aren't met.
type GetListNoteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetListNoteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetListNoteRequestMultiError) AllErrors() []error { return m }

// GetListNoteRequestValidationError is the validation error returned by
// GetListNoteRequest.Validate if the designated constraints aren't met.
type GetListNoteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetListNoteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetListNoteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetListNoteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetListNoteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetListNoteRequestValidationError) ErrorName() string {
	return "GetListNoteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetListNoteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetListNoteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetListNoteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetListNoteRequestValidationError{}

// Validate checks the field values on GetListNoteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetListNoteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetListNoteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetListNoteResponseMultiError, or nil if none found.
func (m *GetListNoteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetListNoteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetNotes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetListNoteResponseValidationError{
						field:  fmt.Sprintf("Notes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetListNoteResponseValidationError{
						field:  fmt.Sprintf("Notes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetListNoteResponseValidationError{
					field:  fmt.Sprintf("Notes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetListNoteResponseMultiError(errors)
	}

	return nil
}

// GetListNoteResponseMultiError is an error wrapping multiple validation
// errors returned by GetListNoteResponse.ValidateAll() if the designated
// constraints aren't met.
type GetListNoteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetListNoteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetListNoteResponseMultiError) AllErrors() []error { return m }

// GetListNoteResponseValidationError is the validation error returned by
// GetListNoteResponse.Validate if the designated constraints aren't met.
type GetListNoteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetListNoteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetListNoteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetListNoteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetListNoteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetListNoteResponseValidationError) ErrorName() string {
	return "GetListNoteResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetListNoteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetListNoteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetListNoteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetListNoteResponseValidationError{}

// Validate checks the field values on UpdateNoteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateNoteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateNoteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateNoteRequestMultiError, or nil if none found.
func (m *UpdateNoteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateNoteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Text

	// no validation rules for Author

	if len(errors) > 0 {
		return UpdateNoteRequestMultiError(errors)
	}

	return nil
}

// UpdateNoteRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateNoteRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateNoteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateNoteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateNoteRequestMultiError) AllErrors() []error { return m }

// UpdateNoteRequestValidationError is the validation error returned by
// UpdateNoteRequest.Validate if the designated constraints aren't met.
type UpdateNoteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateNoteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateNoteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateNoteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateNoteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateNoteRequestValidationError) ErrorName() string {
	return "UpdateNoteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateNoteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateNoteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateNoteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateNoteRequestValidationError{}

// Validate checks the field values on DeleteNoteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteNoteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteNoteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteNoteRequestMultiError, or nil if none found.
func (m *DeleteNoteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteNoteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteNoteRequestMultiError(errors)
	}

	return nil
}

// DeleteNoteRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteNoteRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteNoteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteNoteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteNoteRequestMultiError) AllErrors() []error { return m }

// DeleteNoteRequestValidationError is the validation error returned by
// DeleteNoteRequest.Validate if the designated constraints aren't met.
type DeleteNoteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteNoteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteNoteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteNoteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteNoteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteNoteRequestValidationError) ErrorName() string {
	return "DeleteNoteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteNoteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteNoteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteNoteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteNoteRequestValidationError{}

// Validate checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Empty) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EmptyMultiError, or nil if none found.
func (m *Empty) ValidateAll() error {
	return m.validate(true)
}

func (m *Empty) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptyMultiError(errors)
	}

	return nil
}

// EmptyMultiError is an error wrapping multiple validation errors returned by
// Empty.ValidateAll() if the designated constraints aren't met.
type EmptyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyMultiError) AllErrors() []error { return m }

// EmptyValidationError is the validation error returned by Empty.Validate if
// the designated constraints aren't met.
type EmptyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyValidationError) ErrorName() string { return "EmptyValidationError" }

// Error satisfies the builtin error interface
func (e EmptyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmpty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyValidationError{}

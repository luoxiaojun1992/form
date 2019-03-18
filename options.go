// Copyright 2014~2019 Alvaro J. Genial. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package form

import "errors"

// The individual values for the standard set of options.
const (
	standardZeros    = false
	standardTolerant = false
	standardCaseless = false

	StandardDelimiter   = '.'
	StandardEscape      = '\\'
	StandardImplicitKey = "_"
	StandardOmittedKey  = "-"
)

// Options provides a way to control encoding & decoding behaviors.
type Options struct {
	unexported  int8
	Zeros       bool
	Tolerant    bool
	Caseless    bool
	Delimiter   rune
	Escape      rune
	ImplicitKey string
	OmittedKey  string
}

// MakeStandardOptions returns an Options copy with the standard set of values, which are constant;
// the result can be safely changed and used locally and will have no effect elsewhere.
func MakeStandardOptions() Options {
	return Options{
		Zeros:       standardZeros,
		Tolerant:    standardTolerant,
		Caseless:    standardCaseless,
		Delimiter:   StandardDelimiter,
		Escape:      StandardEscape,
		ImplicitKey: StandardImplicitKey,
		OmittedKey:  StandardOmittedKey,
	}
}

// Default is the set of current (global) options to choose when left unspecified; though not
// suggested, these options can be changed, but will globally affect default behavior; if changing
// these options, care must be taken to avoid concurrency issues and ensure deterministic access;
// most likely it will be safest to do so just once, initially--prior to reaching concurrent code.
var Default = MakeStandardOptions()

// The set of errors that can be produced as part of validation.
var (
	ErrInvalidOptions     = errors.New("invalid form options")
	ErrInvalidDelimiter   = errors.New("invalid form delimiter")
	ErrInvalidEscape      = errors.New("invalid form escape")
	ErrInvalidImplicitKey = errors.New("invalid form implicit key")
	ErrInvalidOmittedKey  = errors.New("invalid form omitted key")
)

// IsValid returns whether options o are considered valid; equivalent to Validate returning nil.
func (o Options) IsValid() bool {
	return o.Validate() == nil
}

// Validate returns an error for invalid conditions in options o.
func (o Options) Validate() error {
	if o == (Options{}) {
		return ErrInvalidOptions
	} else if o.Delimiter == 0 {
		return ErrInvalidDelimiter
	} else if o.Escape == 0 {
		return ErrInvalidEscape
	} else if o.ImplicitKey == "" {
		return ErrInvalidImplicitKey
	} else if o.OmittedKey == "" {
		return ErrInvalidOmittedKey
	}
	return nil
}

// MustValidate returns options o unchanged unless there's a validation error, in which case it panics.
func (o Options) MustValidate() Options {
	if err := o.Validate(); err != nil {
		panic(err)
	}
	return o
}

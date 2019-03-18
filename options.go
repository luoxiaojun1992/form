// Copyright 2014~2019 Alvaro J. Genial. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package form

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

var Default = MakeStandardOptions()

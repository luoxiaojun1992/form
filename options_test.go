// Copyright 2014~2019 Alvaro J. Genial. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package form

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	for _, c := range []struct {
		o Options
		b bool
	}{
		{Options{}, false},
		{Options{0, false, false, false, '\x00', '\\', "_", "-"}, false},
		{Options{0, false, false, false, '.', '\x00', "_", "-"}, false},
		{Options{0, false, false, false, '.', '\\', "", "-"}, false},
		{Options{0, false, false, false, '.', '\\', "_", ""}, false},
		{Options{0, false, false, false, '.', '\\', "_", "-"}, true},
		{Options{0, false, false, true, '.', '\\', "_", "-"}, true},
		{Options{0, false, true, false, '.', '\\', "_", "-"}, true},
		{Options{0, false, true, true, '.', '\\', "_", "-"}, true},
		{Options{0, true, false, false, '.', '\\', "_", "-"}, true},
		{Options{0, true, false, true, '.', '\\', "_", "-"}, true},
		{Options{0, true, true, false, '.', '\\', "_", "-"}, true},
		{Options{0, true, true, true, '.', '\\', "_", "-"}, true},
	} {
		if b := c.o.IsValid(); b != c.b {
			t.Errorf("%#v.IsValid()\n want (%#v)\n have (%#v)", c.o, c.b, b)
		}
	}
}

func TestValidate(t *testing.T) {
	for _, c := range []struct {
		o Options
		e error
	}{
		{Options{}, ErrInvalidOptions},
		{Options{0, false, false, false, '\x00', '\\', "_", "-"}, ErrInvalidDelimiter},
		{Options{0, false, false, false, '.', '\x00', "_", "-"}, ErrInvalidEscape},
		{Options{0, false, false, false, '.', '\\', "", "-"}, ErrInvalidImplicitKey},
		{Options{0, false, false, false, '.', '\\', "_", ""}, ErrInvalidOmittedKey},
		{Options{0, false, false, false, '.', '\\', "_", "-"}, nil},
		{Options{0, false, false, true, '.', '\\', "_", "-"}, nil},
		{Options{0, false, true, false, '.', '\\', "_", "-"}, nil},
		{Options{0, false, true, true, '.', '\\', "_", "-"}, nil},
		{Options{0, true, false, false, '.', '\\', "_", "-"}, nil},
		{Options{0, true, false, true, '.', '\\', "_", "-"}, nil},
		{Options{0, true, true, false, '.', '\\', "_", "-"}, nil},
		{Options{0, true, true, true, '.', '\\', "_", "-"}, nil},
	} {
		if e := c.o.Validate(); e != c.e {
			t.Errorf("%#v.Validate()\n want (%#v)\n have (%#v)", c.o, c.e, e)
		}
	}
}

// Based on ssh/terminal:
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package term

import "syscall"

const ioctlReadTermios = syscall.TCGETS

// Termios is the Terminal Input/Output structure
type Termios syscall.Termios

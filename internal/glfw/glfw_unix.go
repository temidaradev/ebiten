// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2012 The glfw3-go Authors
// SPDX-FileCopyrightText: 2023 The Ebitengine Authors

//go:build darwin || freebsd || linux || netbsd || openbsd

package glfw

//#include <stdlib.h>
//#define GLFW_INCLUDE_NONE
//#include "glfw3_unix.h"
import "C"

import (
	"errors"
	"unsafe"
)

// Init initializes the GLFW library. Before most GLFW functions can be used,
// GLFW must be initialized, and before a program terminates GLFW should be
// terminated in order to free any resources allocated during or after
// initialization.
//
// If this function fails, it calls Terminate before returning. If it succeeds,
// you should call Terminate before the program exits.
//
// Additional calls to this function after successful initialization but before
// termination will succeed but will do nothing.
//
// This function may take several seconds to complete on some systems, while on
// other systems it may take only a fraction of a second to complete.
//
// On Mac OS X, this function will change the current directory of the
// application to the Contents/Resources subdirectory of the application's
// bundle, if present.
//
// This function may only be called from the main thread.
func Init() error {
	C.glfwInit()
	if err := fetchErrorIgnoringPlatformError(); err != nil {
		return err
	}
	return nil
}

// Terminate destroys all remaining windows, frees any allocated resources and
// sets the library to an uninitialized state. Once this is called, you must
// again call Init successfully before you will be able to use most GLFW
// functions.
//
// If GLFW has been successfully initialized, this function should be called
// before the program exits. If initialization fails, there is no need to call
// this function, as it is called by Init before it returns failure.
//
// This function may only be called from the main thread.
func Terminate() error {
	C.glfwTerminate()
	if err := fetchErrorIgnoringPlatformError(); err != nil {
		return err
	}
	return nil
}

// InitHint function sets hints for the next initialization of GLFW.
//
// The values you set hints to are never reset by GLFW, but they only take
// effect during initialization. Once GLFW has been initialized, any values you
// set will be ignored until the library is terminated and initialized again.
//
// Some hints are platform specific. These may be set on any platform but they
// will only affect their specific platform. Other platforms will ignore them.
// Setting these hints requires no platform specific headers or functions.
//
// This function must only be called from the main thread.
func InitHint(hint Hint, value int) {
	C.glfwInitHint(C.int(hint), C.int(value))
}

// GetClipboardString returns the contents of the system clipboard, if it
// contains or is convertible to a UTF-8 encoded string.
//
// This function may only be called from the main thread.
func GetClipboardString() (string, error) {
	cs := C.glfwGetClipboardString(nil)
	if cs == nil {
		if err := fetchErrorIgnoringPlatformError(); err != nil {
			if errors.Is(err, FormatUnavailable) {
				return "", nil
			}
			return "", err
		}
		return "", nil
	}
	return C.GoString(cs), nil
}

// SetClipboardString sets the system clipboard to the specified UTF-8 encoded
// string.
//
// This function may only be called from the main thread.
func SetClipboardString(str string) error {
	cp := C.CString(str)
	defer C.free(unsafe.Pointer(cp))
	C.glfwSetClipboardString(nil, cp)
	return fetchErrorIgnoringPlatformError()
}

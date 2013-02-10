// Copyright 2013, Carnegie Mellon University. All Rights Reserved.
// Use of this code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Author: Alok Parlikar <aup@cs.cmu.edu>

package goflite

/*
 #include <flite.h>
*/
import "C"

import (
	"errors"
	"sync"
	"unsafe"
)

// Type for pointers to Flite Voice Structures
type flitevoice *C.cst_voice

const defaultVoiceName = "slt"

// Voice db
type voxbase struct {
	flitevox map[string]flitevoice // List of voices available for use
	mutex    sync.RWMutex
}

func newVoxBase() *voxbase {
	s := &voxbase{flitevox: make(map[string]flitevoice)}

	// Add Default Voice
	name := C.CString(defaultVoiceName)
	v := C.flite_voice_select(name)
	C.free(unsafe.Pointer(name))

	if v != nil {
		name := C.GoString(v.name)
		if name == defaultVoiceName {
			s.flitevox[defaultVoiceName] = v
		} else {
			C.delete_voice(v)
		}
	}

	return s
}

// Add a voice to list of available voices, given a name the voice
// will be known as, and the path to the flitevox file. Preferably use
// absolute voice paths.  If no voices are added, the "slt" voice is
// always supported
func (voices *voxbase) addVoice(name, path string) error {
	voices.mutex.Lock()
	defer voices.mutex.Unlock()
	_, present := voices.flitevox[name]
	if present {
		return errors.New("Voice with given name already present")
	}

	pathC := C.CString("file://" + path)
	defer C.free(unsafe.Pointer(pathC))

	v := C.flite_voice_select(pathC)
	if v == nil {
		return errors.New("Voice File could not be loaded")
	}

	voices.flitevox[name] = v
	return nil
}

// Voices stored in voxbase are C structures that should be freed
func (voices *voxbase) free() {
	voices.mutex.Lock()
	defer voices.mutex.Unlock()

	for name, voice := range voices.flitevox {
		if name != defaultVoiceName {
			// Default voice is linked in, don't remove it
			C.delete_voice(voice)
		}
	}
}

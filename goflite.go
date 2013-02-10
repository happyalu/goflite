// Copyright 2013, Carnegie Mellon University. All Rights Reserved.
// Use of this code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Author: Alok Parlikar <aup@cs.cmu.edu>

// Use the CMU Flite Text-To-Speech Engine from Go
package goflite

// +build linux,cgo

/*
 #cgo CFLAGS: -I. -I dep/flite/include
 #cgo linux,amd64 LDFLAGS: dep/flite/build/x86_64-linux-gnu/lib/libflite_cmu_us_slt.a dep/flite/build/x86_64-linux-gnu/lib/libflite_cmulex.a dep/flite/build/x86_64-linux-gnu/lib/libflite_usenglish.a dep/flite/build/x86_64-linux-gnu/lib/libflite.a -lm

 #cgo linux,386   LDFLAGS: dep/flite/build/i386-linux-gnu/lib/libflite_cmu_us_slt.a dep/flite/build/i386-linux-gnu/lib/libflite_cmulex.a dep/flite/build/i386-linux-gnu/lib/libflite_usenglish.a dep/flite/build/i386-linux-gnu/lib/libflite.a -lm

 #include <flitewrap.h>
 #include <flite.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

var voices *voxbase // List of available voices stored here

// Initialize Flite
func init() {
	C.flitewrap_init()
	voices = newVoxBase()
}

// Add a voice to list of available voices, given a name the voice
// will be known as, and the path to the flitevox file. Flitevox files
// can be dumped using the -voicedump option in flite. Preferably use
// absolute voice paths to specify location of flitevox files to add.
func AddVoice(name, path string) error {
	return voices.addVoice(name, path)
}

// Run Text to Speech on a given text with a selected voice and return
// Wave data. If voicename is empty, a default voice will be used for
// the speech synthesis.
func TextToWave(text, voicename string) (*Wave, error) {
	var (
		w      *Wave       // Waveform to Return
		v      flitevoice  // Voice to use
		cstwav *C.cst_wave // Flite's wave structure
	)

	if voicename == "" {
		// Choose default voice
		voicename = defaultVoiceName
	}

	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	voices.mutex.RLock()
	v, ok := voices.flitevox[voicename]
	voices.mutex.RUnlock()

	if !ok {
		return nil, errors.New("Requested voice not available")
	}

	cstwav = C.flite_text_to_wave(ctext, v)
	if cstwav == nil {
		return nil, errors.New("Speech synthesis failed")
	}

	num_samples := uint32(cstwav.num_samples)

	w = &Wave{
		SampleRate:  uint16(cstwav.sample_rate),
		NumSamples:  num_samples,
		NumChannels: uint16(cstwav.num_channels),
		Samples:     make([]uint16, num_samples),
	}

	C.copy_wav_into_slice(cstwav, (*C.short)(unsafe.Pointer(&(w.Samples[0]))))
	C.delete_wave(cstwav)

	return w, nil
}

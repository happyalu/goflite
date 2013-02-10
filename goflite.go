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

var voices *voxbase

// Initialize Flite
func init() {
	C.flitewrap_init()
	voices = newVoxBase()
}

// Run Text to Speech on a given text with a selected voice and return Wave data
func TextToWave(text, voicename string) (*Wave, error) {
	var (
		w      *Wave       // Waveform to Return
		v      flitevoice  // Voice to use
		cstwav *C.cst_wave // Flite's wave structure
	)

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

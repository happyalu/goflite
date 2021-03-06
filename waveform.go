// Copyright 2013, Carnegie Mellon University. All Rights Reserved.
// Use of this code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Author: Alok Parlikar <aup@cs.cmu.edu>

package goflite

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Structure for Waveform Data
type Wave struct {
	SampleRate  uint16
	NumSamples  uint32
	NumChannels uint16
	Samples     []uint16
}

// Get the Duration of Waveform in Seconds
func (w *Wave) Duration() float32 {
	if w.SampleRate == 0 {
		return 0.0
	}

	return float32(w.NumSamples) / float32(w.SampleRate)
}

// Write out the waveform, with RIFF headers
func (w *Wave) EncodeRIFF(out io.Writer) (err error) {

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	// File Type
	_, err = fmt.Fprintf(out, "%s", "RIFF")
	if err != nil {
		return
	}

	// Bytes in whole file
	binWrite(out, uint32(uint32(w.NumChannels)*w.NumSamples*2+8+16+12))

	_, err = fmt.Fprintf(out, "%s", "WAVE")
	if err != nil {
		return
	}

	_, err = fmt.Fprintf(out, "%s", "fmt ")
	if err != nil {
		return
	}

	binWrite(out, uint32(16))                           // Size of Header
	binWrite(out, uint16(0x0001))                       // Sample Type (RIFF_FORMAT_PCM)
	binWrite(out, w.NumChannels)                        // Number of Channels
	binWrite(out, uint32(w.SampleRate))                 // Sample Rate
	binWrite(out, uint32(w.SampleRate*w.NumChannels*2)) // Average Bytes Per Second
	binWrite(out, uint16(w.NumChannels*2))              // Block Align
	binWrite(out, uint16(16))                           // Bits per Sample

	_, err = fmt.Fprintf(out, "%s", "data")
	if err != nil {
		return
	}

	binWrite(out, uint32(uint32(w.NumChannels)*w.NumSamples*2)) // Bytes in Data

	// Data Bytes
	for _, v := range w.Samples {
		binWrite(out, v)
	}
	return
}

// Utility function to write binary LittleEndian output
func binWrite(w io.Writer, value interface{}) {
	err := binary.Write(w, binary.LittleEndian, value)
	if err != nil {
		panic(err)
	}
}

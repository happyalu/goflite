// Copyright 2013, Carnegie Mellon University. All Rights Reserved.
// Use of this code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Author: Alok Parlikar <aup@cs.cmu.edu>

package goflite

import "testing"

func TestSynthesisWithInvalidVoice(t *testing.T) {
	voicename := "invalid"
	_, err := TextToWave("Testing", voicename)
	if err == nil {
		t.Errorf("Synthesis should fail when voicename is invalid")
	}
}

func TestSynthesisWithDefaultVoice(t *testing.T) {
	voicename := defaultVoiceName
	w, err := TextToWave("Hello World", voicename)
	if err != nil {
		t.Errorf("Synthesis with default voice should not fail")
	}

	if w.Duration() == 0 {
		t.Errorf("Synthesis with default voice has empty waveform")
	}
}

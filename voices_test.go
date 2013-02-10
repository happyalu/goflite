// Copyright 2013, Carnegie Mellon University. All Rights Reserved.
// Use of this code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Author: Alok Parlikar <aup@cs.cmu.edu>

package goflite

import "testing"

import ()

func TestSLTVoiceAvailable(t *testing.T) {
	vb := newVoxBase()
	defer vb.free()
	_, present := vb.flitevox[defaultVoiceName]
	if !present {
		t.Errorf("The default voice (%s) is not available", defaultVoiceName)
	}

}

func TestAddNonExistingVoice(t *testing.T) {
	vb := newVoxBase()
	defer vb.free()
	err := vb.addVoice("none", "/none/none091.flitevox")
	if err == nil {
		t.Errorf("AddVoice should not accept invalid files")
	}
}

func TestValidVoice(t *testing.T) {
	vb := newVoxBase()
	defer vb.free()
	err := vb.addVoice("aup", "dep/cmu_us_aup.flitevox")
	if err != nil {
		t.Errorf("AddVoice unable to add dep/cmu_us_aup.flitevox")
	}
}

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

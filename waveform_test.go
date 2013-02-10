package goflite

import "testing"

import (
	"bytes"
	"errors"
)

type spanic struct{}

func (c *spanic) Write(p []byte) (n int, err error) {
	err = errors.New("Can't write to me")
	return
}

func TestWaveformWritePanic(t *testing.T) {

	var c spanic
	var w Wave

	err := w.DumpRIFF(&c)
	if err == nil {
		t.Errorf("Waveform writing not handling errors")
	}
}

func TestEmptyWaveform(t *testing.T) {
	var b bytes.Buffer
	var w Wave
	err := w.DumpRIFF(&b)
	if err != nil {
		t.Errorf("DumpRIFF Failed")
	}
	if b.Len() != 44 {
		t.Errorf("Empty RIFF header must be 44 bytes. Received %d", b.Len())
	}

	if w.Duration() != 0 {
		t.Errorf("Duration of empty waveform must be 0")
	}
}

func TestNonEmptyWaveform(t *testing.T) {
	var b bytes.Buffer

	w := Wave{
		SampleRate:  1,
		NumSamples:  10,
		NumChannels: 1,
		Samples:     make([]uint16, 10, 10),
	}

	err := w.DumpRIFF(&b)
	if err != nil {
		t.Errorf("DumpRIFF Failed")
	}
	if b.Len() != 64 {
		t.Errorf("Expected length of waveform 64 bytes. Received %d bytes", b.Len())
	}

	if w.Duration() != 10.0 {
		t.Errorf("Expected duration: 10. Received Duration %1.3f", w.Duration())
	}

}

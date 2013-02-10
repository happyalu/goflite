=======
GoFlite
=======

Access the Flite Speech Synthesizer from Go!

Usage
=====

import "github.com/happyalu/goflite"

See Example application at http://www.github.com/happyalu/gofliteweb

API
===


::

 func TextToWave(text, voicename string) (*Wave, error)
    Run Text to Speech on a given text with a selected voice and return Wave
    data. If voicename is empty, a default voice will be used for the speech
    synthesis.

::

 func (w *Wave) DumpRIFF(out io.Writer) (err error)
    Write out complete RIFF waveform, with headers

::

 func (w *Wave) Duration() float32
    Get the Duration of Waveform in Seconds

::

 func AddVoice(name, path string) error
    Optional: If you have your own flitevox files to use:
	Add a voice to list of available voices, given a name the voice will be
    known as, and the path to the flitevox file. Flitevox files can be
    dumped using the -voicedump option in flite. Preferably use absolute
    voice paths to specify location of flitevox files to add.


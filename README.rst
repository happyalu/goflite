=======
GoFlite
=======

Access the Flite Speech Synthesizer from Go!

Usage
=====

import "github.com/happyalu/goflite"

See Example application at http://www.github.com/happyalu/gofliteweb

Build / Install
===============

GoFlite depends upon the CMU Flite speech synthesis engine ( http://www.cmuflite.org ).
Running "make depflite" in the source directory will download all dependencies. 

After than, "go build" or "go install" will build the library properly. 

Alternatively, you could just run "make" to download dependencies, build and test goflite.


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
    If you have built flite voices and have the flitevox files generated,
    use this function to add them to goflite. Provide a name to the voice
    being added and a path to the location of the flitevox file. Prefer
    absolute pathname.


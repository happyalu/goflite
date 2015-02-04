# Copyright 2013, Carnegie Mellon University. All Rights Reserved.
# Use of this code is governed by a BSD-style license that can be
# found in the LICENSE file.
# Author: Alok Parlikar <aup@cs.cmu.edu>

TOP=.
FLITEDIR=$(TOP)/dep

.phony:  all depflite build test install clean

all: build test

depflite: $(FLITEDIR)/flite $(FLITEDIR)/cmu_us_aup.flitevox

$(FLITEDIR)/flite: $(FLITEDIR)/flite-2.0.0-release.tar.bz2
	cd $(FLITEDIR) && tar xvjf flite-2.0.0-release.tar.bz2 && \
		ln -s flite-2.0.0-release flite && cd flite && CFLAGS="-DCST_AUDIO_NONE -DCST_NO_SOCKETS" ./configure --with-pic --with-audio=none --with-mmap=none && make

$(FLITEDIR)/flite-2.0.0-release.tar.bz2:
	mkdir -p $(FLITEDIR)
	cd $(FLITEDIR) && wget "http://www.festvox.org/flite/packed/flite-2.0/flite-2.0.0-release.tar.bz2";

$(FLITEDIR)/cmu_us_aup.flitevox:
	mkdir -p $(FLITEDIR)
	cd $(FLITEDIR) && wget "http://www.festvox.org/flite/packed/flite-2.0/voices/cmu_us_aup.flitevox";

build: depflite
	go build

test: depflite
	go test

install: depflite
	go install

clean:
	rm -rf $(FLITEDIR)

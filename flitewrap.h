// Copyright 2013, Carnegie Mellon University. All Rights Reserved.
// Use of this code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Author: Alok Parlikar <aup@cs.cmu.edu>

#ifndef _FLITEWRAP_H
#define _FLITEWRAP_H
#include <flite.h>

void usenglish_init(cst_voice *v);
cst_lexicon *cmulex_init(void);

void cmu_indic_lang_init(cst_voice *v);
cst_lexicon *cmu_indic_lex_init(void);

cst_voice *register_cmu_us_slt(const char *voxdir);
cst_voice *unregister_cmu_us_slt();

void flitewrap_init();
void copy_wav_into_slice(const cst_wave *w, short *s);
#endif

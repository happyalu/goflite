#include <flite.h>
#include "flitewrap.h"

// Initialize Flite
void flitewrap_init() {
  flite_init();
  flite_add_lang("usenglish",usenglish_init,cmulex_init);
  flite_voice_list = val_reverse(cons_val(voice_val(register_cmu_us_slt(NULL)),flite_voice_list));
}

// Copy data from a flite waveform into short array.
// This is a helper function to copy cst_wave into go []uint16
void copy_wav_into_slice(const cst_wave *w, short *s) {
  int i;
  for(i=0; i<w->num_samples; i++) {
    s[i] = w->samples[i];
  }
}

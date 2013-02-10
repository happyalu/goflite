#ifndef _FLITEWRAP_H
#define _FLITEWRAP_H
#include <flite.h>

void usenglish_init(cst_voice *v);
cst_lexicon *cmulex_init(void);
cst_voice *register_cmu_us_slt(const char *voxdir);
cst_voice *unregister_cmu_us_slt();

void flitewrap_init();
void copy_wav_into_slice(cst_wave *w, short *s);
#endif

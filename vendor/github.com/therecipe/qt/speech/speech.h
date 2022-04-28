// +build !minimal

#pragma once

#ifndef GO_QTSPEECH_H
#define GO_QTSPEECH_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtSpeech_PackedString { char* data; long long len; void* ptr; };
struct QtSpeech_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif
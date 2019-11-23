#ifndef FACERECOGNIZER_LIB_INTERNAL_H
#define FACERECOGNIZER_LIB_INTERNAL_H

#include "seeta/seeta.h"

extern "C" _declspec(dllexport) void init();

extern "C" _declspec(dllexport) int addImage(char *name, char *path);

extern "C" _declspec(dllexport) char *recognizer(char *path);

#endif //FACERECOGNIZER_LIB_INTERNAL_H

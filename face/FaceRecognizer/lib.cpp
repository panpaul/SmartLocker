#include "lib_internal.h"

void init() {
    initSeetaEngine();
    loadData();
}

bool addImage(char *name, char *path) {
    int i = addFace(name, path);
    return i != -1;
}

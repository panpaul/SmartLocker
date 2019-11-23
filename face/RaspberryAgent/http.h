#ifndef RASPBERRYAGENT_HTTP_H
#define RASPBERRYAGENT_HTTP_H

#include "header.h"
#include <curl/curl.h>
#include "config.h"

struct MemoryStruct {
    char *memory;
    size_t size;
};

std::string post(const std::string &address, const std::string &param);

std::string postImg(const std::string &address, const std::string &imgPath);

void closeCurl();

#endif //RASPBERRYAGENT_HTTP_H

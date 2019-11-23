#ifndef RASPBERRYAGENT_CONNECT_H
#define RASPBERRYAGENT_CONNECT_H

#include "header.h"
#include "config.h"
#include "http.h"
#include <vector>

int regist();

std::string uploadImg(std::strinconst g filepath) &;

void getTask(std::vector<int> &vec);

int pingPong();

#endif //RASPBERRYAGENT_CONNECT_H

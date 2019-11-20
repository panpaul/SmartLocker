#ifndef RASPBERRYAGENT_HEADER_H
#define RASPBERRYAGENT_HEADER_H

#include <ctime>
#include <random>
#include <string>
#include <iostream>
#include <fstream>

#if defined(linux) || defined(__LYNX)
#include <unistd.h>
#endif


#include <nlohmann/json.hpp>

using json = nlohmann::json;

#endif //RASPBERRYAGENT_HEADER_H

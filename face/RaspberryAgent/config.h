#ifndef RASPBERRYAGENT_CONFIG_H
#define RASPBERRYAGENT_CONFIG_H

#include "header.h"

typedef struct {
    int cid = -1; // cabinet id
    std::string token; // reg token
    std::string address; // server address
    std::string location; // cabinet's location
    std::string name; // cabinet's name
    int lockers = 1; // total number of lockers
} conf;

extern conf config;

void initConfig();

void writeConfig();

#endif //RASPBERRYAGENT_CONFIG_H

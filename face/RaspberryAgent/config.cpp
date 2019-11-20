#include "config.h"

conf config;
json c;

void initConfig() {
    std::ifstream infile;
    infile.open("config.json");
    infile >> c;
    infile.close();

    config.cid = c["cid"].get<int>();
    config.address = c["address"].get<std::string>();
    config.token = c["token"].get<std::string>();
    config.location = c["location"].get<std::string>();
    config.name = c["name"].get<std::string>();
    config.lockers = c["lockers"].get<int>();
}

void writeConfig() {
    std::ofstream outfile;
    outfile.open("config.json");
    c["cid"] = config.cid;
    outfile << c.dump();
    outfile.close();
}
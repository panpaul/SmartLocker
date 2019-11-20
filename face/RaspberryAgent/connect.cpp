#include "connect.h"

int regist() {
    std::string token = config.token;
    std::string addr = config.address + "api/v1/cabinet/register";
    std::string param =
            "&regToken=" + token + "&name=" + config.name + "&location=" + config.location + "&num=" +
            std::to_string(config.lockers);
    std::string r = post(addr, param);
    std::cout << r << std::endl;
    auto j = json::parse(r);
    auto code = j["code"].get<int>();
    if (code != 0) {
        return -1;
    }
    return j["body"].get<int>();
}
#include "connect.h"

int regist() {
    std::string token = config.token;
    std::string addr = config.address + "api/v1/cabinet/register";
    std::string param =
            "&regToken=" + token + "&name=" + config.name + "&location=" + config.location + "&num=" +
            std::to_string(config.lockers);
    std::string r = post(addr, param);
    std::cout << "raw cabinet id:" << r << std::endl;
    auto j = json::parse(r);
    auto code = j["code"].get<int>();
    if (code != 0) {
        return -1;
    }
    return j["body"].get<int>();
}

std::string uploadImg(const std::string &filepath) {
    std::string addr = config.address + "api/v1/http/face/recognize";
    std::string r = postImg(addr, filepath);
    std::cout << "raw img name:" << r << std::endl;
    auto j = json::parse(r);
    auto code = j["code"].get<int>();
    if (code != 0) {
        return "";
    }
    return j["body"].get<std::string>();
}

void getTask(std::vector<int> &vec) {
    std::string param = "&cid=" + std::to_string(config.cid);
    std::string addr = config.address + "/api/v1/cabinet/task";
    std::string r = post(addr, param);
    std::cout << "raw tasks:" << r << std::endl;

    auto j = json::parse(r);
    auto code = j["code"].get<int>();
    if (code != 0) {
        return;
    }
    for (auto &element : j["body"]) {
        vec.push_back(element.get<int>());
    }
}

bool pingPong() {
    std::string param = "&cid=" + std::to_string(config.cid) + "&regToken=" + config.token;
    std::string addr = config.address + "/api/v1/cabinet/ping";
    std::string r = post(addr, param);
    std::cout << "raw ping pong:" << r << std::endl;

    auto j = json::parse(r);
    auto code = j["code"].get<int>();
    if (code != 0) {
        return false;
    }
    return true;
}
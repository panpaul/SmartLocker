#include "http.h"

static size_t WriteMemoryCallback(void *contents, size_t size, size_t nmemb, void *userp) {
    size_t realsize = size * nmemb;
    auto *mem = (struct MemoryStruct *) userp;

    char *ptr = static_cast<char *>(realloc(mem->memory, mem->size + realsize + 1));
    if (ptr == nullptr) {
        /* out of memory! */
        printf("not enough memory (realloc returned NULL)\n");
        return 0;
    }

    mem->memory = ptr;
    memcpy(&(mem->memory[mem->size]), contents, realsize);
    mem->size += realsize;
    mem->memory[mem->size] = 0;

    return realsize;
}

std::string post(const std::string &address, const std::string &param) {
    struct MemoryStruct chunk{};
    chunk.memory = static_cast<char *>(malloc(1));
    chunk.size = 0;

    CURL *curl;
    curl = curl_easy_init();

    if (!curl) {
        return "";
    }

    std::string ret;

    curl_easy_setopt(curl, CURLOPT_POSTFIELDS, param.c_str());
    curl_easy_setopt(curl, CURLOPT_URL, address.c_str());
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, WriteMemoryCallback);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, (void *) &chunk);
    CURLcode res = curl_easy_perform(curl);
    if (res != CURLE_OK) {
        fprintf(stderr, "curl_easy_perform() failed: %s\n",
                curl_easy_strerror(res));
    } else {
        ret = chunk.memory;
    }
    curl_easy_cleanup(curl);
    free(chunk.memory);
    return ret;
}

std::string postImg(const std::string &address, const std::string &imgPath) {
    struct MemoryStruct chunk{};
    chunk.memory = static_cast<char *>(malloc(1));
    chunk.size = 0;

    CURL *curl;
    curl = curl_easy_init();

    if (!curl) {
        return "";
    }

    // prepare to post
    curl_mime *mime;
    curl_mimepart *part;

    mime = curl_mime_init(curl);
    part = curl_mime_addpart(mime);
    curl_mime_filedata(part, imgPath.c_str());
    curl_mime_name(part, "img");
    // debug
    //curl_easy_setopt(curl, CURLOPT_PROXY, "http://127.0.0.1:8888/");
    curl_easy_setopt(curl, CURLOPT_MIMEPOST, mime);
    curl_easy_setopt(curl, CURLOPT_URL, address.c_str());
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, WriteMemoryCallback);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, (void *) &chunk);

    CURLcode res = curl_easy_perform(curl);
    std::string ret;

    if (res != CURLE_OK) {
        fprintf(stderr, "curl_easy_perform() failed: %s\n",
                curl_easy_strerror(res));
    } else {
        ret = chunk.memory;
    }
    curl_easy_cleanup(curl);
    curl_mime_free(mime);
    free(chunk.memory);
    return ret;
}

void closeCurl() {
    curl_global_cleanup();
}
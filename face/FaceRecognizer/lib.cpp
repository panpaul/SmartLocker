#include "lib_internal.h"

void init() {
    initSeetaEngine();
    loadData();
}

int addImage(char *name, char *path) {
    int i = addFace(name, path);
    return i != -1;
}

char *recognizer(char *path) {
    cv::Mat frame;
    frame = cv::imread(path);
    seeta::cv::ImageData image = frame;
    vector<SeetaFaceInfo> faces = engine.DetectFaces(image);

    for (SeetaFaceInfo &face : faces) {
        // Query top 1
        int64_t index = -1;
        float similarity = 0;
        auto points = engine.DetectPoints(image, face);

        std::string name;

        auto queried =
                engine.QueryTop(image, points.data(), 1, &index, &similarity);

        // no face queried from database
        if (queried < 1)
            continue;

        if (similarity > similarThreshold) {
            name = ImageIndexMap[index];
            char *writable = new char[name.size() + 1];
            std::copy(name.begin(), name.end(), writable);
            writable[name.size()] = '\0';
            return writable;
        }
    }
    return "";
}
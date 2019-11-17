#include "face.h"

char *getFace() {
    UltraFace ultraface("./model/RFB-320.bin", "./model/RFB-320.param", 320, 2, 0.75); // config model input

    cv::VideoCapture cap(0);

    cv::Mat frame;

    cap >> frame;
    cv::resize(frame, frame, cv::Size(320, 240), (0, 0), (0, 0), cv::INTER_LINEAR);
    ncnn::Mat inmat = ncnn::Mat::from_pixels(frame.data, ncnn::Mat::PIXEL_BGR2RGB, frame.cols, frame.rows);

    std::vector<FaceInfo> face_info;
    ultraface.detect(inmat, face_info);

    if (face_info.empty()) {
        return "0";
    }

    if (face_info.size() > 1) {
        return "1";
    }

    auto face = face_info[0];
    cv::Rect rect(face.x1, face.y1, face.x2 - face.x1, face.y2 - face.y1);

    srand(time(nullptr));
    std::string filename = "./img/" + std::to_string(rand()) + ".jpg";

    cv::imwrite(filename, frame(rect));

    char *ret = new char[filename.size() + 1];
    std::copy(filename.begin(), filename.end(), ret);
    ret[filename.size()] = '\0';

    return ret;
}

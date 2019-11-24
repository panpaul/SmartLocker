#include "header.h"
#include "config.h"
#include "connect.h"
#include <opencv2/opencv.hpp>
#include "UltraFace.h"

using namespace std;

// init random engine
default_random_engine randomEngine(time(nullptr));
// init NCNN engine
UltraFace ultraface("./model/RFB-320.bin", "./model/RFB-320.param", 320, 2, 0.75); // config model input

// get VideoCapture instance
cv::VideoCapture cap(0);
// the mat to store camera image
cv::Mat frame;

string getFace();

void processFace();

int main(int argc, char **argv) {
    // getConfig
    initConfig();
    // register
    int cid = regist();
    if (cid == -1) {
        fprintf(stderr, "registration failed\n");
        return -1;
    }
    config.cid = cid;
    // write back config
    writeConfig();

    while (true) {// main loop
        // ping pong
        auto ppStatus = pingPong();
        if (!ppStatus) {
            break;
        }

        processFace();
        cv::imshow("RaspberryAgent", frame);
        cv::waitKey(1);

        std::vector<int> tasks;
        getTask(tasks);
        std::cout << "total task:" << tasks.size() << std::endl;
        for (const auto &task : tasks) {
            std::cout << "task:" << task << std::endl;
        }

        // sleep for 2 seconds
#if defined(linux) || defined(__LYNX)
        sleep(2);
#endif
#if defined(_WIN32)
        Sleep(2000);
#endif
    }

    closeCurl();
    return 0;
}

void processFace() {
    // getFace
    string filename = getFace();
    if (filename.empty()) {
        return;
    }
    if (filename == "0") {
        cv::putText(frame, "no faces detected", cv::Point(0, 10), cv::FONT_HERSHEY_SIMPLEX, 0.5,
                    cv::Scalar(0, 0, 255));
        return;
    } else if (filename == "1") {
        cv::putText(frame, "only one face allowed", cv::Point(0, 10), cv::FONT_HERSHEY_SIMPLEX, 0.5,
                    cv::Scalar(0, 0, 255));
        return;
    }
    string username = uploadImg(filename);
    std::cout << "img name:" << username << std::endl;
    if (username.empty()) {
        cv::putText(frame, "Couldn't recognize. Try again", cv::Point(0, frame.rows - 5), cv::FONT_HERSHEY_SIMPLEX, 0.5,
                    cv::Scalar(0, 0, 255));
        return;
    }
    cv::putText(frame, username, cv::Point(0, frame.rows - 5), cv::FONT_HERSHEY_SIMPLEX, 0.5,
                cv::Scalar(0, 0, 255));
}

string getFace() {
    // get the frame
    cap >> frame;
    // resize to 320*240
    cv::resize(frame, frame, cv::Size(320, 240), (0, 0), (0, 0), cv::INTER_LINEAR);
    // convert NCNN mat
    ncnn::Mat inMat = ncnn::Mat::from_pixels(frame.data, ncnn::Mat::PIXEL_BGR2RGB, frame.cols, frame.rows);
    // a vector to hold the faces detected
    std::vector<FaceInfo> face_info;

    // detect the face
    ultraface.detect(inMat, face_info);

    if (face_info.empty()) { // no faces found
        return "0";
    }

    if (face_info.size() > 1) { // more than one
        return "1";
    }

    // now we have only one face
    auto face = face_info[0];
    // generate a rect to hold the position
    // we need a slightly bigger rect
    float x1 = MAX(face.x1 - 30, 0), y1 = MAX(face.y1 - 40, 0);
    float x2 = MIN(face.x2 + 30, frame.cols), y2 = MIN(face.y2 + 30, frame.rows);
    cv::Rect rect(x1, y1, x2 - x1, y2 - y1);
    // prepare the filename
    std::string filename = "./img/" + std::to_string(randomEngine()) + ".jpg";
    // save the face into a file
    cv::imwrite(filename, frame(rect));
    // now draw a rectangle on the frame
    cv::rectangle(frame, rect, cv::Scalar(0, 0, 255));
    return filename;
}
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


string getFace();

template<class T>
int length(T &arr) {
    return sizeof(arr) / sizeof(arr[0]);
}

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
        // getFace
        string filename = getFace();
        if (!filename.empty()) {
            string username = uploadImg(filename);
            std::cout << "img name:" << username << std::endl;
        }
        std::vector<int> tasks;
        getTask(tasks);
        std::cout << "total task:" << tasks.size() << std::endl;
        for (std::_Vector_iterator<std::_Vector_val<std::_Simple_types<int> > >::value_type &task : tasks) {
            std::cout << "task:" << task << std::endl;
        }


        // sleep for 2 seconds
#if defined(linux) || defined(__LYNX)
        sleep(2);
#endif
#if defined(_WIN32)
        Sleep(2000);
#endif
        //break;
    }

    closeCurl();
    return 0;
}

string getFace() {
    // the mat to store camera image
    cv::Mat frame;
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
        return "";
    }

    if (face_info.size() > 1) { // more than one
        return "";
    }

    // now we have only one face
    auto face = face_info[0];
    // generate a rect to hold the position
    cv::Rect rect(face.x1, face.y1, face.x2 - face.x1, face.y2 - face.y1);
    // prepare the filename
    std::string filename = "./img/" + std::to_string(randomEngine()) + ".jpg";
    // save the face into a file
    cv::imwrite(filename, frame(rect));
    return filename;
}
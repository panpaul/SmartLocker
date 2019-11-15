//
// Created by panyu on 2019/11/11.
//

#ifndef FACERECOGNIZER_HEADER_H
#define FACERECOGNIZER_HEADER_H

#include <map>
#include <iostream>
#include <fstream>
#include <vector>

using namespace std;

#ifdef EXPFS
	#include <experimental/filesystem>
	using namespace experimental::filesystem;
#else
	#include <filesystem>
	using namespace filesystem;
#endif


#include <opencv2/highgui/highgui.hpp>
#include <opencv2/imgproc/imgproc.hpp>

using namespace cv;

#endif //FACERECOGNIZER_HEADER_H

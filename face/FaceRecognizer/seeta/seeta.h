//
// Created by panyu on 2019/11/11.
//

#ifndef FACERECOGNIZER_SEETA_H
#define FACERECOGNIZER_SEETA_H

#include "FaceEngine.h"
#include "Struct_cv.h"
#include <seeta/Struct.h>
#include <seeta/QualityAssessor.h>
#include "../header.h"
#include "../read.h"

inline seeta::QualityAssessor QA;

extern vector<string> ImageFilename;
extern vector<string> ImageName;
extern map<int64_t, string> ImageIndexMap;

inline seeta::ModelSetting FD_model("./model/fd_2_00.dat", seeta::ModelSetting::CPU, 0);
inline seeta::ModelSetting PD_model("./model/pd_2_00_pts5.dat", seeta::ModelSetting::CPU, 0);
inline seeta::ModelSetting FR_model("./model/fr_2_10.dat", seeta::ModelSetting::CPU, 0);
extern inline seeta::FaceEngine engine(FD_model, PD_model, FR_model);

extern float similarThreshold;

void initSeetaEngine();

void loadData();

int addFace(string name, string filename);

#endif //FACERECOGNIZER_SEETA_H

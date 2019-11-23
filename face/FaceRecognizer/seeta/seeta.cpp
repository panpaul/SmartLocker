#include "seeta.h"

vector<string> ImageFilename;
vector<string> ImageName;
map<int64_t, string> ImageIndexMap;

float similarThreshold;

void initSeetaEngine() {
    similarThreshold = 0.7f;

    //set face detector's min face size
    engine.FD.set(seeta::FaceDetector::PROPERTY_MIN_FACE_SIZE, 80);
}

void loadData() {
    bool useDump = exists("./FaceBackend/data/dump.data") && exists("./FaceBackend/data/dump.inf");
    bool appendNew = exists("./FaceBackend/data/append.csv");

    if (!useDump && !appendNew) {
        //cout << "Invalid input" << endl;
        return;
    }

    ifstream file;
    if (useDump) {
        engine.Load("./FaceBackend/data/dump.data");
        file.open("./FaceBackend/data/dump.inf", ifstream::in);
    } else if (appendNew) {
        file.open("./FaceBackend/data/append.csv", ifstream::in);
    }

    if (!file) {
        string error_message = "No valid input file was given.";
        //CV_Error(cv::Error::StsBadArg, error_message);
        //cout << "No valid input file was given." << endl;
        return;
    }

    //read data
    read(file, ImageFilename, ImageName);

    if (useDump) {
        //just link the data
        for (size_t i = 0; i < ImageFilename.size(); ++i) {
            // save index and name pair
            ImageIndexMap.insert(make_pair(stoi(ImageFilename[i]), ImageName[i]));
        }
        ImageFilename.clear();
        ImageName.clear();
    }

    vector<int64_t> ImageIndex(ImageFilename.size());

    for (size_t i = 0; i < ImageFilename.size(); ++i) {
        //register face into face database
        string &filename = ImageFilename[i];
        int64_t &index = ImageIndex[i];
        //cout << "Registering: " << filename << " with name: " << ImageName[i] << endl;
        seeta::cv::ImageData image = cv::imread(filename);
        auto idImage = engine.Register(image);
        index = idImage;
        //cout << "Registered id = " << idImage << endl;
    }

    for (size_t i = 0; i < ImageIndex.size(); ++i) {
        // save index and name pair
        if (ImageIndex[i] < 0) continue;
        ImageIndexMap.insert(make_pair(ImageIndex[i], ImageName[i]));
    }

    //save data
    if (!(useDump && (!appendNew))) {
        //cout << "Saving" << endl;
        FILE *fpWrite = fopen("./FaceBackend/data/dump.inf", "a");
        if (fpWrite == nullptr) {
            cout << "Could not save" << endl;
            return;
        }
        for (size_t i = 0; i < ImageIndexMap.size(); ++i) {
            fprintf(fpWrite, "%zu,%s\n", i, ImageIndexMap[i].c_str());
        }
        fclose(fpWrite);
    }

    //save dump
    int e = remove("./FaceBackend/data/dump.data");
    if (e != 0) {
        //cout << "could not delete previous file" << endl;
        //return;
    }
    bool status = engine.Save("./FaceBackend/data/dump.data");
    if (!status) {
        //cout << "could not save database" << endl;
    }
}

int addFace(string name, string filename) {
    seeta::cv::ImageData image = cv::imread(filename);
    auto idImage = engine.Register(image);
    if (idImage == -1) {
        return -1;
    }
    ImageIndexMap.insert(make_pair(idImage, name));

    FILE *fpWrite = fopen("./FaceBackend/data/dump.inf", "a");
    if (fpWrite == nullptr) {
        //cout << "Could not save" << endl;
        return idImage;
    }
    fprintf(fpWrite, "%zu,%s\n", idImage, name.c_str());
    fclose(fpWrite);

    int e = remove("./FaceBackend/data/dump.data");
    if (e != 0) {
        //cout << "could not delete previous file" << endl;
        //return;
    }
    bool status = engine.Save("./FaceBackend/data/dump.data");
    if (!status) {
        //cout << "could not save database" << endl;
    }

    return idImage;
}

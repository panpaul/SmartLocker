#include "read.h"

void read(ifstream &file, vector<string> &GalleryImageFilename, vector<string> &GalleryImageName) {
    string line, path, classlabel;
    while (getline(file, line)) {
        stringstream liness(line);
        getline(liness, path, ',');
        getline(liness, classlabel);
        if (!path.empty() && !classlabel.empty()) {
            GalleryImageFilename.push_back(path);
            GalleryImageName.push_back(classlabel);
        }
    }
    file.close();
}
#include "seeta/seeta.h"

int main() {
  // prepare Seeta
  initSeetaEngine();
  loadData();

  // main loop
  cout << "----open camera----" << endl;
  // Open default camera
  cv::VideoCapture capture;
  capture.open(0);

  cv::Mat frame;

  while (capture.isOpened()) {
    // read an image
    capture >> frame;
    if (frame.empty())
      continue;

    // flip img
    //cv::flip(frame, frame, 0);
    //cv::flip(frame, frame, 1);

    seeta::cv::ImageData image = frame;

    // Detect all faces
    vector<SeetaFaceInfo> faces = engine.DetectFaces(image);

    for (SeetaFaceInfo &face : faces) {
      // Query top 1
      int64_t index = -1;
      float similarity = 0;

      auto points = engine.DetectPoints(image, face);

      rectangle(
          frame,
          cv::Rect(face.pos.x, face.pos.y, face.pos.width, face.pos.height),
          CV_RGB(128, 128, 255), 3);

      for (int i = 0; i < 5; ++i) {
        auto &point = points[i];
        circle(frame, cv::Point(int(point.x), int(point.y)), 2,
               CV_RGB(128, 255, 128), -1);
      }

      std::string name;

      auto score = QA.evaluate(image, face.pos, points.data());
	  cout << "QA: " << score << endl;
      if (score == 0) {
        name = "ignored";
      } else {
        auto queried =
            engine.QueryTop(image, points.data(), 1, &index, &similarity);

        // no face queried from database
        if (queried < 1)
          continue;

        if (similarity > similarThreshold) {
          name = ImageIndexMap[index];
          cout << "Find a possible face: " << name << " with similarity: " << similarity << endl;
        }
      }

      if (!name.empty()) {
        putText(frame, name,
                cv::Point(face.pos.x, face.pos.y - 5), 3, 1,
                CV_RGB(255, 128, 128));
      }
    }

    imshow("Frame", frame);

    auto key = waitKey(20);
    if (key == 27) {
      break;
    }
  }
}

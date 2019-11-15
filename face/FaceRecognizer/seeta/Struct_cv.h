#pragma once

#include <opencv2/core/core.hpp>

#include <seeta/CStruct.h>

namespace seeta::cv
    {
        // using namespace ::cv;
        class ImageData : public SeetaImageData {
        public:
            ImageData( const ::cv::Mat &mat )
                    : SeetaImageData(), cv_mat( mat.clone() ) {
                this->width = cv_mat.cols;
                this->height = cv_mat.rows;
                this->channels = cv_mat.channels();
                this->data = cv_mat.data;
            }

            ImageData( int width, int height, int channels = 3 )
                    : cv_mat( height, width, CV_8UC( channels ) ) {
                this->width = cv_mat.cols;
                this->height = cv_mat.rows;
                this->channels = cv_mat.channels();
                this->data = cv_mat.data;
            }
            ImageData( const SeetaImageData &img )
                    : SeetaImageData(), cv_mat( img.height, img.width, CV_8UC( img.channels ), img.data ) {
                this->width = cv_mat.cols;
                this->height = cv_mat.rows;
                this->channels = cv_mat.channels();
                this->data = cv_mat.data;
            }
            ImageData()
                    : SeetaImageData(), cv_mat() {
                this->width = cv_mat.cols;
                this->height = cv_mat.rows;
                this->channels = cv_mat.channels();
                this->data = cv_mat.data;
            }
            [[nodiscard]] bool empty() const {
                return cv_mat.empty();
            }
            explicit operator ::cv::Mat() const {
                return cv_mat.clone();
            }
            [[nodiscard]] ::cv::Mat toMat() const {
                return cv_mat.clone();
            }
        private:
            ::cv::Mat cv_mat;
        };
    }
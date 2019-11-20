cmake_minimum_required(VERSION 3.2)
set(CMAKE_CXX_STANDARD 14)

project(RaspberryAgent)

add_executable(RaspberryAgent main.cpp UltraFace.cpp UltraFace.h)

find_package(ncnn REQUIRED)
target_include_directories(RaspberryAgent PRIVATE ${ncnn_DIR}/../../../include)
target_link_libraries(RaspberryAgent ncnn)

find_package(OpenCV REQUIRED)
target_include_directories(RaspberryAgent PRIVATE ${OpenCV_INCLUDE_DIR})
target_link_libraries(RaspberryAgent ${OpenCV_LIBS})

find_package(CURL CONFIG REQUIRED)
target_include_directories(RaspberryAgent PRIVATE ${CURL_INCLUDE_DIR})
target_link_libraries(RaspberryAgent CURL::libcurl)

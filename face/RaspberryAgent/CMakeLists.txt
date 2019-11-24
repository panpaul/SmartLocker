cmake_minimum_required(VERSION 3.2)
set(CMAKE_CXX_STANDARD 14)

project(RaspberryAgent)

add_executable(RaspberryAgent main.cpp UltraFace.cpp UltraFace.h http.cpp http.h connect.cpp connect.h header.h config.cpp config.h)

find_package(ncnn REQUIRED)
target_include_directories(RaspberryAgent PRIVATE ${ncnn_DIR}/../../../include)
target_link_libraries(RaspberryAgent ncnn)

find_package(OpenCV REQUIRED)
target_include_directories(RaspberryAgent PRIVATE ${OpenCV_INCLUDE_DIR})
target_link_libraries(RaspberryAgent ${OpenCV_LIBS})

find_package(CURL REQUIRED)
target_include_directories(RaspberryAgent PRIVATE ${CURL_INCLUDE_DIR})
target_link_libraries(RaspberryAgent CURL::libcurl)

find_package(nlohmann_json REQUIRED)
target_include_directories(RaspberryAgent PRIVATE ${nlohmann_json_INCLUDE_DIR})
target_link_libraries(RaspberryAgent nlohmann_json nlohmann_json::nlohmann_json)

target_link_libraries(RaspberryAgent wiringPi)

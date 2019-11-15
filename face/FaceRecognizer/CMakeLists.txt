cmake_minimum_required(VERSION 2.8)
project(FaceRecognizer)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

add_executable(FaceRecognizer seeta/FaceEngine.h seeta/Struct_cv.h seeta/seeta.h read.cpp read.h header.h seeta/seeta.cpp main.cpp)

if(NOT MSVC)
    target_link_libraries(FaceRecognizer stdc++fs)
endif()

if(CMAKE_CXX_COMPILER_ID STREQUAL "GNU")
  if(CMAKE_CXX_COMPILER_VERSION VERSION_LESS "8.3.0")
    add_definitions("-DEXPFS")
  endif()
endif()

find_package(SeetaFace REQUIRED)
if (SeetaFace_FOUND)
    include_directories(${SeetaFace_INCLUDE_DIR})
    target_link_libraries(FaceRecognizer ${SeetaFace_LIBRARIES})
endif()

find_package(OpenCV REQUIRED)
if (OpenCV_FOUND)
    include_directories(${OpenCV_INCLUDE_DIR})
    target_link_libraries(FaceRecognizer ${OpenCV_LIBRARIES})
endif()

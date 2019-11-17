cmake_minimum_required(VERSION 3.2)
set(CMAKE_CXX_STANDARD 14)

project(libface)

add_library(libface STATIC face.cpp face.h UltraFace.cpp UltraFace.h)
set_target_properties(libface PROPERTIES
        PUBLIC_HEADER include/face.h)

find_package(ncnn REQUIRED)
target_include_directories(libface PRIVATE ${ncnn_DIR}/../../../include)
target_link_libraries(libface ncnn)

find_package(OpenCV REQUIRED)
target_include_directories(libface PRIVATE ${OpenCV_INCLUDE_DIR})
target_link_libraries(libface ${OpenCV_LIBS})

install(TARGETS libface
        ARCHIVE DESTINATION lib ${CMAKE_INSTALL_LIBDIR}
        PUBLIC_HEADER DESTINATION include ${CMAKE_INSTALL_INCLUDEDIR})

project(testface)

add_executable(testface main.cpp)
target_link_libraries(testface libface)

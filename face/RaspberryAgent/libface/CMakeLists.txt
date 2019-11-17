cmake_minimum_required(VERSION 2.8)
set(CMAKE_CXX_STANDARD 17)

project(libface)

add_library(libface STATIC face.cpp face.h UltraFace.cpp UltraFace.h)
set_target_properties(libface PROPERTIES PUBLIC_HEADER include/face.h)
target_include_directories(libface PRIVATE .)
install(TARGETS libface
        LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}
        PUBLIC_HEADER DESTINATION ${CMAKE_INSTALL_INCLUDEDIR})

find_package(ncnn REQUIRED)
target_include_directories(libface PRIVATE ${ncnn_DIR}/../../../include)
target_link_libraries(libface ncnn)

find_package(OpenCV REQUIRED)
target_include_directories(libface PRIVATE ${OpenCV_INCLUDE_DIR})
target_link_libraries(libface ${OpenCV_LIBS})


project(testface)

add_executable(testface main.cpp)
target_link_libraries(testface libface)


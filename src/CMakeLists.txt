# At LEAST 2.8 but newer is better
cmake_minimum_required(VERSION 3.2 FATAL_ERROR)
project(directory-thumbnails VERSION 0.1 LANGUAGES C)

# Compilation settings
set(CMAKE_C_STANDARD 11)
set(CMAKE_COMPILER_IS_GNUCC TRUE)

# Include Boost as an imported target
#find_package(Boost REQUIRED)
#add_library(boost INTERFACE IMPORTED)
#set_property(TARGET boost PROPERTY
#    INTERFACE_INCLUDE_DIRECTORIES ${Boost_INCLUDE_DIR})

# External 3rd party libs that we include
#add_subdirectory(libs/extlib EXCLUDE_FROM_ALL)

# Targets that we develop
#enable_testing()
add_subdirectory(thumbnail)

add_executable(directory-thumbnails main.c)
target_link_libraries(directory-thumbnails LINK_PUBLIC thumbnail)

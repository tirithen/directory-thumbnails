add_library(thumbnail get_thumbnails_for_directory.c)
target_include_directories(thumbnail PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})

#include "thumbnail.h"

int main() {
  char* thumbnail_paths = get_thumbnail_paths_for_directory("~/Hämtningar", 10);
  return 0;
}

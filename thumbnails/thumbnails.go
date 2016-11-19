package thumbnails

import (
	"bytes"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

var thumbnailPathRegexp = regexp.MustCompile("thumbnail::path:\\s*(.+)\\s*")

// GetForDirectory uses $ gvfs-info -a thumbnail::path to get thumbnails for the first files in a directory
func GetForDirectory(path string, count int) []string {
	thumbnailPaths := make([]string, 0)

	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() == false {
			command := exec.Command("gvfs-info", "-a", "thumbnail::path", path+"/"+file.Name())
			var output bytes.Buffer
			command.Stdout = &output
			err := command.Run()

			if err != nil {
				log.Fatal(err)
			}

			outputWithoutNewline := strings.Replace(output.String(), "\n", " ", -1)
			thumbnailMatch := thumbnailPathRegexp.FindStringSubmatch(outputWithoutNewline)

			if len(thumbnailMatch) > 1 {
				thumbnailPaths = append(thumbnailPaths, thumbnailMatch[1])
				if len(thumbnailPaths) >= count {
					break
				}
			}
		}
	}

	return thumbnailPaths
}

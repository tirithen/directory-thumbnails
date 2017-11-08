package thumbnails

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os/exec"
	"regexp"
	"strings"

	bimg "gopkg.in/h2non/bimg.v1"
)

var thumbnailPathRegexp = regexp.MustCompile("thumbnail::path:\\s*(.+)\\s*")

// GetForDirectory uses $ gio info -a thumbnail::path to get thumbnails for the first files in a directory
func GetForDirectory(path string, count int) []string {
	thumbnailPaths := make([]string, 0)

	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() == false {
			command := exec.Command("gio", "info", "-a", "thumbnail::path", path+"/"+file.Name())
			var output bytes.Buffer
			command.Stdout = &output
			err := command.Run()

			if err != nil {
				log.Fatal(err)
			}

			outputWithoutNewline := strings.Replace(output.String(), "\n", " ", -1)
			thumbnailMatch := thumbnailPathRegexp.FindStringSubmatch(outputWithoutNewline)

			if len(thumbnailMatch) > 1 {
				thumbnailPaths = append(thumbnailPaths, strings.Trim(thumbnailMatch[1], " "))
				if len(thumbnailPaths) >= count {
					break
				}
			}
		}
	}

	return thumbnailPaths
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func CreateFromImages(bottomBuffer, topBuffer []byte, imagePaths []string) (output []byte, err error) {
	spreadAngle := float64(45)
	thumbnailSize := 256

	bottomImage := bimg.NewImage(bottomBuffer)

	buffer, err := bottomImage.Thumbnail(thumbnailSize)
	if err != nil {
		return
	}
	bottomImage = bimg.NewImage(buffer)

	var images [][]byte
	for _, imagePath := range imagePaths {
		buffer, readErr := bimg.Read(imagePath)

		if readErr == nil {
			small, smallErr := bimg.NewImage(buffer).Enlarge(thumbnailSize/2, thumbnailSize/2)
			if smallErr == nil {
				images = append(images, small)
			}
		}
	}

	spreadStep := spreadAngle / float64(len(images)-1)

	//radius := float64(bottomImageSize.Height) / 4

	for index, image := range images {
		angle := spreadStep*float64(index) - 90 - spreadAngle/2
		//x := int(radius * math.Cos(degreesToRadians(angle)))
		//y := int(radius * math.Sin(degreesToRadians(angle)))

		imageSize, imageSizeErr := bimg.NewImage(image).Size()
		if imageSizeErr == nil {
			//y = imageSize.Height / 2
		}

		x := 0
		y := -1 + thumbnailSize - imageSize.Height
		fmt.Println("angle", angle, "x", x, "y", y)

		/*rotated, rotatedErr := bimg.NewImage(image).Rotate(70)
		if rotatedErr != nil {
			fmt.Println("rotatedErr", rotatedErr)
			rotated = image
		}*/

		buffer, watermarkErr := bottomImage.WatermarkImage(bimg.WatermarkImage{
			Left: x,
			Top:  y,
			Buf:  image,
		})
		if watermarkErr == nil {
			bottomImage = bimg.NewImage(buffer)
		}
	}

	buffer, watermarkErr := bottomImage.WatermarkImage(bimg.WatermarkImage{Left: 0, Top: 0, Buf: topBuffer})
	if watermarkErr == nil {
		bottomImage = bimg.NewImage(buffer)
	}

	output, err = bottomImage.Process(bimg.Options{
		Quality:   100,
		Interlace: true,
	})

	return
}

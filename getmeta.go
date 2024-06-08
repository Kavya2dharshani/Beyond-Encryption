package File

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

func GetMetaData() {
	// URL of the image
	imageURL := "https://raw.githubusercontent.com/ianare/exif-samples/master/jpg/Canon_40D.jpg"
	// imageURL := "https://commons.wikimedia.org/wiki/File:JPEG_example_down.jpg"
	// Fetch the image from the URL
	resp, err := http.Get(imageURL)
	if err != nil {
		fmt.Println("Error fetching image:", err)
		return
	}

	fmt.Println("resp ------->", resp)

	defer resp.Body.Close()

	// Read the image data
	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading image data:", err)
		return
	}

	fmt.Println("imageData ------->", len(imageData))

	// Create a bytes.Reader from image data
	reader := bytes.NewReader(imageData)

	// Decode the image data to extract metadata
	x, err := exif.Decode(reader)
	if err != nil {
		fmt.Println("Error decoding image data:", err)
		return
	}

	// Print EXIF metadata
	fmt.Println("Metadata:")
	x.Walk(exifWalker{})
}

// exifWalker implements exif.Walker interface
type exifWalker struct{}

func (w exifWalker) Walk(name exif.FieldName, tag *tiff.Tag) error {
	fmt.Printf("%s: %s\n", name, tag)
	return nil
}

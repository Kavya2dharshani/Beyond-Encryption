package File

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"net/http"
)

func RemoveMetaData() {
	// URL of the image
	imageURL := "https://raw.githubusercontent.com/ianare/exif-samples/master/jpg/Canon_40D.jpg"

	// Fetch the image from the URL
	resp, err := http.Get(imageURL)
	if err != nil {
		fmt.Println("Error fetching image:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("resp -------->", resp)

	// Read the image data
	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading image data:", err)
		return
	}

	fmt.Println("imageData ", len(imageData))

	// Remove EXIF metadata
	withoutExif, err := removeExif(imageData)
	if err != nil {
		fmt.Println("Error removing EXIF metadata:", err)
		return
	}

	fmt.Println("withoutExif", len(withoutExif))

	// Save the image without EXIF metadata
	err = ioutil.WriteFile("image_without_exif.jpg", withoutExif, 0644)
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}

	fmt.Println("Image saved without EXIF metadata.")
}

func removeExif(data []byte) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("error decoding image: %v", err)
	}
	// fmt.Println("buf", len(img))

	var buf bytes.Buffer

	err = jpeg.Encode(&buf, img, nil)
	if err != nil {
		return nil, fmt.Errorf("error encoding image: %v", err)
	}

	return buf.Bytes(), nil
}

package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func ProcessImage(imageURL string) error {
	resp, err := http.Get(imageURL)
	if err != nil {
		return fmt.Errorf("failed to download image: %v", err)
	}
	defer resp.Body.Close()

	fileName := "downloaded_image.png"

	out, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}

	fmt.Printf("Image saved successfully as %s\n", fileName)
	return nil
}

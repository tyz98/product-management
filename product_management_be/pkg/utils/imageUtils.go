package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// IsValidImage validates the file extension for image files
func IsValidImage(file *multipart.FileHeader) bool {
	allowedExtensions := []string{".jpg", ".jpeg", ".png"}
	ext := filepath.Ext(file.Filename)
	for _, validExt := range allowedExtensions {
		if ext == validExt {
			return true
		}
	}
	return false
}

// SaveImage saves the image to a specified path
func SaveImage(file *multipart.FileHeader, destination string) error {
	// Open the file for reading
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Create the file at the destination
	dst, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the file content
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

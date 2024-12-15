package utils

import (
	"errors"
	"net/url"
	"strings"
)

func ConvertGoogleDriveURL(gdriveURL string) (string, error) {
	parsedURL, err := url.Parse(gdriveURL)
	if err != nil {
		return "", err
	}

	segments := strings.Split(parsedURL.Path, "/")
	if len(segments) < 5 {
		return "", errors.New("invalid Google Drive URL")
	}

	fileID := segments[3]
	return fileID, nil
}

package util

import (
	"mime/multipart"
	"strings"
)

const (
	MB = 1 << 20
)

func IsFileValid(file *multipart.FileHeader) (bool, string) {
	name := file.Filename
	size := file.Size

	validExtension := []string{".jpeg", ".png", ".jpg", ".svg", ".apng"}
	var isFileValid bool
	var message string

	for _, format := range validExtension {
		if strings.Contains(name[len(name)-5:], format) {
			isFileValid = true
			message = ""
			break
		} else {
			isFileValid = false
			message = "File extension not allowed, only accept file in .jpeg, .png, .jpg, .svg, .apng"
		}
	}

	if size > int64(5*MB) {
		isFileValid = false
		if message != "" {
			message = message + " & " + "The File Size Exceeds the Limit Allowed, the Limit Allowed is under 5 MB"
		} else {
			message = "The File Size Exceeds the Limit Allowed, the Limit Allowed is under 5 MB"
		}
	}

	return isFileValid, message
}

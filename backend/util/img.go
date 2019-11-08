package util

import (
	"SmartLocker/e"
	"net/http"
	"os"
	"strings"
)

func VerifyImg(contentType string) (string, int) {
	//contentType should be "image/xxxx"
	split := strings.Split(contentType, "/")
	if len(split) != 2 || split[0] != "image" {
		return "", e.FileTypeMismatch
	}
	ext := []string{
		"bmp",
		"gif",
		"jpeg",
		"jpg",
		"png",
		"webp",
	}
	for i := 0; i < len(ext); i++ {
		if strings.Contains(ext[i], split[1]) {
			return ext[i], e.Success
		}
	}
	return "", e.FileTypeMismatch
}

func getFileContentType(out *os.File) (string, error) {
	// 只需要前 512 个字节就可以了
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

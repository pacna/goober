package zipper

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pacna/goober/internal/utility"
	"github.com/schollz/progressbar/v3"
)

func DownloadImages(outputPath string, userInputURL string, imgURLs []string, size int64) (string, error) {
	var fileName string = strconv.FormatInt(time.Now().Unix(), 10) + ".zip"
	destZipFilePath := filepath.Join(outputPath, fileName)
	zipFile, _ := os.Create(destZipFilePath)
	zipWriter := zip.NewWriter(zipFile)
	
	defer zipFile.Close()
	defer zipWriter.Close()

	bar := progressbar.Default(size)

	for index, imgURL := range imgURLs {
		var imgURLInSegments []string = strings.Split(imgURL, "/")
		appendImageToZip(userInputURL, *NewZipInfo(createFileName(imgURLInSegments[len(imgURLInSegments)-1], index), imgURL, zipWriter))

		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}

	return fileName, nil
}


func createFileName(fileNameFromURL string, index int) string {
	reg, _ := regexp.Compile(`[\d?]+`)
	fileNameFromURLSegments := strings.Split(fileNameFromURL, ".")

	cleanFileName := reg.ReplaceAllString(fileNameFromURLSegments[0], "")
	cleanFileExtension := reg.ReplaceAllString(fileNameFromURLSegments[1], "")

	if utility.IsEmpty(cleanFileName) {
		return fmt.Sprintf("%d.%s", index, cleanFileExtension)
	}

	return fmt.Sprintf("%s_%d.%s", cleanFileName, index, cleanFileExtension)
}

func appendImageToZip(userInputURL string, info zipInfo) error {
	image := GetResponseAsReader(createImgURL(userInputURL, info.imgURL))

	zipFileHeader := &zip.FileHeader{
		Name:   info.fileName,
		Method: zip.Deflate,
	}

	zipFile, _ := info.writer.CreateHeader(zipFileHeader)

	_, err := io.Copy(zipFile, image)

	if err != nil {
		return err
	}

	return nil
}

func GetResponseAsReader(inputURL string) io.Reader {
	var buffer bytes.Buffer
	response, err := http.Get(inputURL)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatal("Status is not returning a success code", response.StatusCode, response.Status)
		return nil
	}

	_, err = buffer.ReadFrom(response.Body)

	if err != nil {
		log.Fatal("Unable to read from buffer")
		return nil
	}

	imageBody := io.NopCloser(&buffer)

	return imageBody
}

func createImgURL(baseURL string, imgURL string) string {
	if utility.IsHttpURL(imgURL) {
		return imgURL
	}

	return baseURL + imgURL
}
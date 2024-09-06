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


func DownloadImages(outputPath string, userInputURL string, imgURLs []string, size int64) error {
	destZipFilePath := filepath.Join(outputPath, strconv.FormatInt(time.Now().Unix(), 10) + ".zip")
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

	return nil
}


func createFileName(fileNameFromURL string, index int) string {
	reg, _ := regexp.Compile(`[\d?]+`)
	fileNameFromURLSegments := strings.Split(fileNameFromURL, ".")

	cleanFileName := reg.ReplaceAllString(fileNameFromURLSegments[0], "")
	cleanFileExtension := reg.ReplaceAllString(fileNameFromURLSegments[1], "")

	fileName := ""

	if utility.IsEmpty(cleanFileName) {
		fileName = fmt.Sprintf("%d.%s", index, cleanFileExtension)
	} else {
		fileName = fmt.Sprintf("%s_%d.%s", cleanFileName, index, cleanFileExtension)
	}


	return fileName
}

func appendImageToZip(userInputURL string, info zipInfo) error {
	image := storeImage(userInputURL, info.imgURL)

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

func storeImage(baseURL string, imgURL string) io.Reader {
	var buffer bytes.Buffer
	response, err := http.Get(createImgURL(baseURL, imgURL))

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatal("Status is not returning a success code", response.StatusCode, response.Status)
	}

	_, err = buffer.ReadFrom(response.Body)

	if err != nil {
		log.Fatal("Unable to read from buffer")
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
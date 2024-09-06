package zipper

import "archive/zip"

type zipInfo struct {
	fileName string
	imgURL string
	writer *zip.Writer
}

func NewZipInfo(fileName string, imgURL string, writer *zip.Writer) *zipInfo {
	return &zipInfo{
		fileName: fileName,
		imgURL: imgURL,
		writer: writer,
	}
}
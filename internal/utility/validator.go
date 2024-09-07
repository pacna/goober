package utility

import (
	"net/url"
	"os"
	"reflect"
)


func IsEmpty[T any](value T) bool {
	return reflect.ValueOf(value).IsZero()
}

func IsHttpURL(value string) bool {
	if IsEmpty(value){
		return false
	}

	url , err := url.ParseRequestURI(value)

	if err != nil {
		return false
	}

	if url.Scheme == "" || url.Hostname() == "" {
		return false
	}

	if url.Scheme != "http" && url.Scheme != "https" {
		return false
	}

	return err == nil
}

func IsValidPath(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
package utility

import (
	"net/url"
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

	if url.Scheme == "" || url.Hostname() == "" {
		return false
	}

	return err == nil
}
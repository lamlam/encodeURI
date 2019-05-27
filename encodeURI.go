package main

import (
	"errors"
	"net/url"
)

func encodeURI(uri string) (string, error) {
	if len(uri) == 0 {
		return "", errors.New("URI must not be empty or null")
	}

	return url.QueryEscape(uri), nil
}

func decodeURI(uri string) (string, error) {
	if len(uri) == 0 {
		return "", errors.New("URI must not be empty or null")
	}

	d, err := url.QueryUnescape(uri)
	if err != nil {
		return "", err
	}

	return d, nil
}

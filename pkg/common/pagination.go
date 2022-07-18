package common

import (
	"encoding/base64"
	"errors"
	"strconv"
)

func EncodePageToken(offest string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(offest))
}

func EncodeToken(offset uint32) string {
	return EncodePageToken(strconv.FormatUint(uint64(offset), 10))
}

func DecodeToken(token string) (uint32, error) {
	offsetString, err := DecodePageToken(token)
	if err != nil {
		return 0, err
	}

	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		return 0, err
	}

	return uint32(offset), nil
}

func DecodePageToken(pageToken string) (token string, err error) {
	decodedToken, err := base64.RawURLEncoding.DecodeString(pageToken)
	if err != nil {
		return token, errors.New("Invalid page token")
	}

	token = string(decodedToken[:])

	if token == "" {
		token = "0"
	}

	return token, nil
}

func GetPreviousPageToken(offset, pageSize uint32) string {
	if offset <= 0 {
		return ""
	}

	var previousOffest uint32
	if offset >= pageSize {
		previousOffest = offset - pageSize
	} else {
		previousOffest = 0
	}

	return EncodeToken(previousOffest)
}

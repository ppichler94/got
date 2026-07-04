package objectdb

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Object interface {
	Write() error
	PrettyPrint() string
}

func Read(hash string) (Object, error) {
	path := filepath.Join(DbDir, hash[:2], hash[2:])
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	parts := strings.SplitN(string(data), "\000", 2)
	header := parts[0]
	content := parts[1]

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return nil, errors.New("invalid object header")
	}
	objectType := headerParts[0]
	lengthString := headerParts[1]
	length, err := strconv.Atoi(lengthString)
	if err != nil {
		return nil, err
	}

	if len(content) != length {
		return nil, errors.New("invalid object length")
	}

	switch objectType {
	case "blob":
		return NewBlob(content), nil
	default:
		return nil, errors.New("invalid object type")
	}
}

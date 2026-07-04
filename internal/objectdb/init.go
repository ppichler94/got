package objectdb

import (
	"os"
	"path/filepath"
)

const DbDir = ".got/objects/"

func Init(dir string) error {
	dbDir := filepath.Join(dir, DbDir)
	if err := os.MkdirAll(dbDir, 0750); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(dbDir, "pack"), 0750); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(dbDir, "info"), 0750); err != nil {
		return err
	}

	return nil
}

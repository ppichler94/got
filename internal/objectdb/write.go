package objectdb

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
)

func writeObject(hash string, content string) error {
	log.Infof("Writing object: '%s'", hash)
	prefix := hash[:2]
	filename := hash[2:]
	path := filepath.Join(DbDir, prefix)

	if err := os.MkdirAll(path, 0750); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(path, filename), []byte(content), 0640); err != nil {
		return err
	}
	return nil
}

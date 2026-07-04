package objectdb

import (
	"crypto/sha1"
	"fmt"
)

type Blob struct {
	Content string
	Hash    string
	header  string
	data    string
}

func (b *Blob) PrettyPrint() string {
	return b.Content
}

func NewBlob(content string) *Blob {
	blob := &Blob{Content: content}
	blob.header = fmt.Sprintf("blob %d\000", len(content))
	blob.data = blob.header + content
	sha := sha1.Sum([]byte(blob.data))
	blob.Hash = fmt.Sprintf("%x", sha)
	return blob

}

func (b *Blob) Write() error {
	if err := writeObject(b.Hash, b.data); err != nil {
		return err
	}
	return nil
}

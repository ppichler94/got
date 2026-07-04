package objectdb

type Object interface {
	Hash() string
	Write() error
}

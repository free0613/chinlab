package strategy

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
)

type Storagestrategy interface {
	Save(name string, data []byte) error
}

var strategy = map[string]Storagestrategy{
	"file":         &fileStorage{},
	"encrypt_file": &encryptfileStorage{},
}

type fileStorage struct{}

func (f *fileStorage) Save(name string, data []byte) error {
	return ioutil.WriteFile(name, data, os.ModeAppend)
}

type encryptfileStorage struct{}

func (e *encryptfileStorage) Save(name string, data []byte) error {
	encrypt_data, err := encrypt(data)
	if err != nil {
		return fmt.Errorf("error")
	}
	return ioutil.WriteFile(name, encrypt_data, 0777)
}

func encrypt(data []byte) ([]byte, error) {
	if data == nil {
		return nil, fmt.Errorf("error type")
	}

	rehash := sha1.New()
	rehash.Write(data)
	result := rehash.Sum(nil)
	return result, nil
}

func NewStorageStrategy(savetype string) (Storagestrategy, error) {
	val, ok := strategy[savetype]
	if !ok {
		return nil, fmt.Errorf("error")
	}
	return val, nil
}

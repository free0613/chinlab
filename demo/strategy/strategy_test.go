package strategy

import (
	"testing"
)

func Test_demo(t *testing.T) {
	data, ok := getData()
	strategyType := "file"
	if !ok {
		strategyType = "encrypt_file"
	}
	storageStrategy, err := NewStorageStrategy(strategyType)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	storageStrategy.Save("./test.txt", data)
}

func getData() ([]byte, bool) {
	return []byte("test data"), true
}

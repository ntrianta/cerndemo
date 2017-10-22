package infrastructure

import "io/ioutil"

type FileHandler struct {
	File string
}

func (file FileHandler) Store(writable []byte) error {
	ioutil.WriteFile(file.File, writable, 0644)
	return nil
}

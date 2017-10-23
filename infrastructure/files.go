package infrastructure

import (
	"fmt"
	"os"
)

type FileHandler struct {
	File string
}

func (file FileHandler) Store(writable string) error {

	fmt.Println("Infrastructure layer. File handler. Storage")

	f, _ := os.OpenFile(file.File, os.O_APPEND|os.O_WRONLY, 0644)

	_, _ = f.WriteString(writable)

	defer f.Close()

	return nil
}

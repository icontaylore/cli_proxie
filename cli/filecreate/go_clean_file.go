package filecreate

import (
	"os"
)

func GoCleanFile(file *os.File) error {
	if err := file.Truncate(0); err != nil {
		return err
	}

	// на начало файла
	file.Seek(0, 0)
	return nil
}

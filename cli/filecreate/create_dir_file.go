package filecreate

import (
	"os"
)

func CreateDir() bool {
	if err := os.Mkdir("proxie", 0666); err != nil {
		return false
	}
	return true
}

//
//func CreateFile() error {
//	path, err := os.Getwd()
//	if err != nil {
//		log.Printf("[ошибка получения пути]")
//		return err
//	}
//	filepath.Join(path, "proxie")
//
//	file, err := os.OpenFile("proxie/proxie.txt", os.O_CREATE|os.O_RDWR, 0644)
//	if err != nil {
//		log.Println("[ошибка получения файла proxie.txt]")
//		return err
//	}
//	file.Truncate(0)
//	file.Seek(0, 0)
//	defer file.Close()
//	return nil
//}

package fileCreate

import (
	"log"
	"os"
	"path/filepath"
)

func CreateFile(arr []string) {
	dir, err := os.Executable()
	if err != nil {
		log.Println("[ошибка в получении os - directory]")
	}
	// root path
	rootDir := filepath.Dir(dir)

	CreateNewFile(rootDir)
}

func CreateNewFile(rootDir string) {
	path := rootDir + "/proxy/proxy.txt"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("[ошибка в создании папки, файла]")
	}
	defer file.Close()
}

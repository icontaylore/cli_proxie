package getPath

import (
	"log"
	"os"
)

func Path() string {
	path, err := os.Getwd()
	if err != nil {
		log.Printf("ошибка в получении пути")
	}
	return path
}

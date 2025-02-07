package filecreate

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

type PaintyValue struct {
	mu sync.Mutex
	wg sync.WaitGroup
}

func PayloadInFile(arr *[]string) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	actualPath := filepath.Join(path, "proxie", "proxie.txt")

	file, err := os.OpenFile(actualPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("[ошибка открытия файла]")
		return err
	}
	defer file.Close()

	RunPaint(file, arr)
	return nil
}

func RunPaint(file *os.File, arr *[]string) {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(len(*arr) - 1)
	for i := 1; i < len(*arr); i++ {
		go func(index int) {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()

			file.WriteString((*arr)[index] + "\n")
		}(i)
	}
	wg.Wait()
}

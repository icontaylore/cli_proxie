package animated

import (
	"fmt"
	"time"
)

func Animated() {
	channel := make(chan string)

	arr := []string{"|", "/", "-", "\\"}

	go func() {
		for {
			for _, v := range arr {
				channel <- v
			}
		}
	}()

	red := "\033[31m"
	reset := "\033[0m"
	for v := range channel {
		fmt.Print("\r", red, v, reset)
		time.Sleep(100 * time.Millisecond)
	}
}

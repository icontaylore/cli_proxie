package main

import (
	"apiproxy/cli/command"
	"apiproxy/cli/getPath"
	"apiproxy/cli/getProxy"
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	filepath := getPath.Path()

	for {
		fmt.Printf("%s $ ", filepath)
		command := command.Command(scanner)

		if command == "" {
			continue
		}

		if command == "proxyget" {
			getProxy.GetProxy()
		}
	}
}

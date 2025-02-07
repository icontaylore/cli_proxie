package main

import (
	"apiproxy/cli/command"
	"apiproxy/cli/getpath"
	"apiproxy/cli/getproxie"
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	filepath := getpath.Path()

	for {
		fmt.Printf("%s $ ", filepath)
		command := command.Command(scanner)

		if command == "" {
			continue
		}

		if command == "proxyget" {
			//getproxie.GetProxy()
			getproxie.BodyUnmarshal()
		}
	}
}

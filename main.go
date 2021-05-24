package main

import (
	"fmt"
	"os"

	"github.com/Skarlso/gtui/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Failed to execute command: ", err)
		os.Exit(1)
	}
}

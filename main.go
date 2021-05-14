package main

import (
	"log"

	"github.com/Skarlso/gtui/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/Skarlso/gtui/pkg"
)

func main() {
	token := os.Getenv("GTUI_TOKEN")
	if token == "" {
		fmt.Println("Please provide GTUI_TOKEN to access github.")
		os.Exit(1)
	}
	gtui := pkg.NewGTUIClient(pkg.Config{
		Token: token,
	})

	if err := gtui.Start(); err != nil {
		fmt.Println("Error while starting gtui: ", err)
	}
}

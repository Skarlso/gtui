package main

import (
	"fmt"
	"os"

	"github.com/Skarlso/gtui/pkg/providers/github"

	"github.com/Skarlso/gtui/pkg"
)

func main() {
	token := os.Getenv("GTUI_TOKEN")
	if token == "" {
		fmt.Println("Please provide GTUI_TOKEN to access github.")
		os.Exit(1)
	}
	githubProvider := github.NewGithubProvider(github.Config{
		Token: token,
	})
	gtui := pkg.NewGTUIClient(pkg.Config{}, pkg.Dependencies{
		Github: githubProvider,
	})

	if err := gtui.Start(); err != nil {
		fmt.Println("Error while starting gtui: ", err)
	}
}

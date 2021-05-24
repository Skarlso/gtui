package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"

	"github.com/Skarlso/gtui/pkg/providers/github"

	"github.com/Skarlso/gtui/pkg"
)

func main() {
	token := os.Getenv("GTUI_TOKEN")
	if token == "" {
		fmt.Println("Please provide GTUI_TOKEN to access github.")
		os.Exit(1)
	}
	logger := zerolog.New(zerolog.ConsoleWriter{
		Out: os.Stderr,
	}).With().Timestamp().Logger()
	githubProvider := github.NewGithubProvider(github.Config{
		Token: token,
	}, logger)
	gtui := pkg.NewGTUIClient(pkg.Config{}, pkg.Dependencies{
		Github: githubProvider,
		Logger: logger,
	})

	if err := gtui.Start(); err != nil {
		fmt.Println("Error while starting gtui: ", err)
	}
}

package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/Skarlso/gtui/pkg"
	"github.com/Skarlso/gtui/pkg/providers/github"
)

const (
	// Prefix of all environment variables
	envKeyPrefix = "GTUI_"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gtui",
		Short: "GTUI a TUI for github Projects",
		Long:  "GTUI is a github Project management TUI",
		Run:   runRootCmd,
	}
	// rootArgs define the root arguments for all commands.
	rootArgs struct {
		Token        string
		Organization string
		Repository   string
		ProjectID    int64
	}
)

func init() {
	f := rootCmd.PersistentFlags()
	// Persistent flags
	token := getEnvOrDefault("TOKEN", "")
	f.StringVar(&rootArgs.Token, "token", token, "Token used to authenticate with github")
	f.StringVar(&rootArgs.Organization, "organization", "", "The organization / owner of the project or the repository to select")
	f.StringVar(&rootArgs.Repository, "repository", "", "The repository which contains projects to select")
	f.Int64Var(&rootArgs.ProjectID, "project-id", -1, "If provided, gtui will immediately open this project")

	if token == "" {
		fmt.Println("Token is empty. Please either provide one with --token or use GTUI_TOKEN environment property.")
		os.Exit(1)
	}
}

func runRootCmd(cmd *cobra.Command, args []string) {
	logger := zerolog.New(zerolog.ConsoleWriter{
		Out: os.Stderr,
	}).With().Timestamp().Logger()
	githubProvider := github.NewGithubProvider(github.Config{
		Token: rootArgs.Token,
	}, logger)
	gtui := pkg.NewGTUIClient(pkg.Config{
		Organization: rootArgs.Organization,
		Repository:   rootArgs.Repository,
		ProjectID:    rootArgs.ProjectID,
	}, pkg.Dependencies{
		Github: githubProvider,
		Logger: logger,
	})

	if err := gtui.Start(); err != nil {
		fmt.Println("Error while starting gtui: ", err)
		os.Exit(1)
	}
}

func getEnvOrDefault(key, def string) string {
	if v := os.Getenv(envKeyPrefix + key); v != "" {
		return v
	}
	return def
}

func Execute() error {
	return rootCmd.Execute()
}

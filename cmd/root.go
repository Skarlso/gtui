package cmd

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

const (
	// Prefix of all environment variables
	envKeyPrefix = "GTUI_"
)

var (
	gtuiCmd = &cobra.Command{
		Use:   "gtuictl",
		Short: "gtui main command",
		Long:  "GTUI handles terminal ui management for github projects and other things.",
		Run:   ShowUsage,
	}

	// CLILog is the gtuictl's logger.
	CLILog = zerolog.New(zerolog.ConsoleWriter{
		Out: os.Stderr,
	}).With().Timestamp().Logger()
	// gtuiArgs define the root arguments for all commands.
	gtuiArgs struct {
		Token     string
		Formatter string
		endpoint  string
	}
)

func getEnvOrDefault(key, def string) string {
	if v := os.Getenv(envKeyPrefix + key); v != "" {
		return v
	}
	return def
}

func init() {
	f := gtuiCmd.PersistentFlags()
	// Persistent flags
	token := getEnvOrDefault("TOKEN", "")
	f.StringVar(&gtuiArgs.Token, "token", token, "Token used to authenticate with the server")

	if token == "" {
		CLILog.Fatal().Msg("Token is empty. Please either provide one with --token or use GTUI_TOKEN environment property.")
	}
}

// ShowUsage shows usage of the given command on stdout.
func ShowUsage(cmd *cobra.Command, args []string) {
	_ = cmd.Usage()
}

// Execute runs the main gtui command.
func Execute() error {
	return gtuiCmd.Execute()
}

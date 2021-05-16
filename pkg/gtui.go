package pkg

import (
	"context"
	"fmt"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"

	"github.com/Skarlso/gtui/pkg/providers"
)

// Config contains configuration properties for GTUI.
type Config struct {
	Token string
}

type Dependencies struct {
	Github providers.Github
}

// GTUIClient defines a client for GTUI.
type GTUIClient struct {
	Config
}

func NewGTUIClient(cfg Config) *GTUIClient {
	return &GTUIClient{
		Config: cfg,
	}
}

// Start launches the GTUI App.
func (g *GTUIClient) Start() error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: g.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		return err
	}
	fmt.Println("repo: ", *repos[0].FullName)
	return nil
}

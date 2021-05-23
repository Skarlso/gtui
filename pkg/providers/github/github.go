package github

import (
	"context"
	"io/ioutil"
	"time"

	"github.com/Skarlso/gtui/models"
	"github.com/google/go-github/v35/github"
	"github.com/rs/zerolog"
	"golang.org/x/oauth2"
)

// Config .
type Config struct {
	Org   string
	Token string
	// TODO add some more
}

// GithubProvider .
type GithubProvider struct {
	Config
	Client *github.Client
	Logger zerolog.Logger
}

// NewGithubProvider .
func NewGithubProvider(cfg Config) *GithubProvider {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "token"},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(tc)
	return &GithubProvider{Config: cfg, Client: client}
}

// ListProjects lists all projects for an organization
func (g *GithubProvider) ListProjects(ctx context.Context, org string, o *models.ListOptions) ([]*github.Project, error) {
	result := make([]*github.Project, 0)
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	opts := &github.ProjectListOptions{
		ListOptions: github.ListOptions{
			Page:    o.Page,
			PerPage: o.PerPage,
		},
	}
	// list all projects for the authenticated user
	// TODO: do pagination
	for {
		projects, response, err := g.Client.Organizations.ListProjects(ctx, g.Org, opts)
		if err != nil {
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				g.Logger.Debug().Err(err).Msg("Failed to read message body.")
			}
			g.Logger.Debug().Err(err).Str("body", string(body)).Msg("Failed list projects")
			return nil, err
		}
		result = append(result, projects...)
		if response.NextPage == 0 {
			break
		}
	}
	return result, nil
}

// GetProject .
func GetProject(ctx context.Context, name string) (*models.Project, error) {
	return nil, nil
}

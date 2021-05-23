package github

import (
	"context"
	"io/ioutil"
	"time"

	"github.com/Skarlso/gtui/pkg/providers"

	"github.com/Skarlso/gtui/models"
	"github.com/google/go-github/v35/github"
	"github.com/rs/zerolog"
	"golang.org/x/oauth2"
)

// Config .
type Config struct {
	Token string
	// TODO add some more
}

// GithubProvider .
type GithubProvider struct {
	Config
	Client *github.Client
	Logger zerolog.Logger
}

var _ providers.Github = &GithubProvider{}

// NewGithubProvider .
func NewGithubProvider(cfg Config) *GithubProvider {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.Token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(tc)
	return &GithubProvider{Config: cfg, Client: client}
}

// ListProjects lists all projects for an organization
func (g *GithubProvider) ListProjects(ctx context.Context, org string, opts *models.ListOptions) ([]*models.Project, error) {
	result := make([]*models.Project, 0)
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	o := &github.ProjectListOptions{
		ListOptions: github.ListOptions{
			Page:    opts.Page,
			PerPage: opts.PerPage,
		},
	}
	for {
		projects, response, err := g.Client.Organizations.ListProjects(ctx, org, o)
		if err != nil {
			g.logGithubResponseBody(response)
			g.Logger.Debug().Err(err).Msg("Failed to list projects.")
			return nil, err
		}
		for _, p := range projects {
			proj := &models.Project{
				Name: providers.GetString(p.Name),
				ID:   providers.GetInt64(p.ID),
			}
			result = append(result, proj)
		}
		if response.NextPage == 0 {
			break
		}
	}
	return result, nil
}

// GetProject .
func (g *GithubProvider) GetProject(ctx context.Context, id int64) (*models.Project, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	project, response, err := g.Client.Projects.GetProject(ctx, id)
	if err != nil {
		g.logGithubResponseBody(response)
		g.Logger.Debug().Err(err).Msg("Failed get project.")
		return nil, err
	}
	p := &models.Project{
		Name: providers.GetString(project.Name),
		ID:   providers.GetInt64(project.ID),
	}
	return p, nil
}

// logGithubResponseBody logs a response if it's not nil.
func (g *GithubProvider) logGithubResponseBody(response *github.Response) {
	if response == nil {
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		g.Logger.Debug().Err(err).Msg("Failed to read message body.")
	}
	g.Logger.Debug().Err(err).Str("body", string(body)).Msg("The body of the failed response.")
}

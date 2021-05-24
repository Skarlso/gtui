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
	// TODO add some more like pagination and shit
}

// GithubProvider .
type GithubProvider struct {
	Config
	Client *github.Client
	Logger zerolog.Logger
}

var _ providers.Github = &GithubProvider{}

// NewGithubProvider .
func NewGithubProvider(cfg Config, logger zerolog.Logger) *GithubProvider {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.Token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(tc)
	return &GithubProvider{Config: cfg, Client: client, Logger: logger}
}

// ListRepositoryProjects lists all projects for a repository
func (g *GithubProvider) ListRepositoryProjects(ctx context.Context, repo, owner string, opts *models.ListOptions) ([]*models.Project, error) {
	result := make([]*models.Project, 0)
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	var (
		page    int
		perPage = 10
	)
	if opts != nil {
		page = opts.Page
		perPage = opts.PerPage
	}
	o := &github.ProjectListOptions{
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: perPage,
		},
	}
	for {
		projects, response, err := g.Client.Repositories.ListProjects(ctx, owner, repo, o)
		if err != nil {
			g.logGithubResponseBody(response)
			g.Logger.Debug().Err(err).Msg("Failed to list projects.")
			return nil, err
		}
		for _, p := range projects {
			proj := &models.Project{
				Name: p.GetName(),
				ID:   p.GetID(),
			}
			result = append(result, proj)
		}
		if response.NextPage == 0 {
			break
		}
		o.ListOptions.Page = response.NextPage
	}
	return result, nil
}

// ListOrganizationProjects lists all projects for an organization
func (g *GithubProvider) ListOrganizationProjects(ctx context.Context, org string, opts *models.ListOptions) ([]*models.Project, error) {
	result := make([]*models.Project, 0)
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	var (
		page    int
		perPage = 10
	)
	if opts != nil {
		page = opts.Page
		perPage = opts.PerPage
	}
	o := &github.ProjectListOptions{
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: perPage,
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
				Name: p.GetName(),
				ID:   p.GetID(),
			}
			result = append(result, proj)
		}
		if response.NextPage == 0 {
			break
		}
		o.ListOptions.Page = response.NextPage
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
		Name: project.GetName(),
		ID:   project.GetID(),
	}
	return p, nil
}

// GetProjectData returns all the columns with all the cards in the columns.
func (g *GithubProvider) GetProjectData(ctx context.Context, projectID int64) (*models.ProjectData, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	opts := &github.ListOptions{
		PerPage: 10,
	}
	allColumns := make([]*github.ProjectColumn, 0)

	for {
		columns, response, err := g.Client.Projects.ListProjectColumns(ctx, projectID, opts)
		if err != nil {
			g.logGithubResponseBody(response)
			g.Logger.Debug().Err(err).Msg("Failed get project columns.")
			return nil, err
		}
		allColumns = append(allColumns, columns...)
		if response.NextPage == 0 {
			break
		}
		opts.Page = response.NextPage
	}
	data := &models.ProjectData{
		ProjectColumns: make([]*models.ProjectColumn, 0),
	}
	for _, c := range allColumns {
		// get all cards
		listOpts := &github.ProjectCardListOptions{
			ListOptions: *opts,
		}
		allCards := make([]*github.ProjectCard, 0)
		for {
			cards, response, err := g.Client.Projects.ListProjectCards(ctx, c.GetID(), listOpts)
			if err != nil {
				g.logGithubResponseBody(response)
				g.Logger.Debug().Err(err).Int64("column_id", c.GetID()).Str("name", c.GetName()).Msg("Failed get column cards.")
				return nil, err
			}
			allCards = append(allCards, cards...)
			if response.NextPage == 0 {
				break
			}
			listOpts.Page = response.NextPage
		}
		col := &models.ProjectColumn{
			Name: c.GetName(),
			ID:   c.GetID(),
		}
		cards := make([]*models.ProjectColumnCard, 0)
		for _, card := range allCards {
			// TODO: get card content for preview. It's an issue -> get issue by ID. If it's a note, the note
			// field will be not nil.
			cards = append(cards, &models.ProjectColumnCard{
				ID: card.GetID(),
			})
		}
		col.ProjectColumnCards = cards
		data.ProjectColumns = append(data.ProjectColumns, col)
	}

	return data, nil
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

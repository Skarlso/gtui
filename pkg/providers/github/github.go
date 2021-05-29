package github

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"time"

	"github.com/google/go-github/v35/github"
	"github.com/rivo/tview"
	"github.com/rs/zerolog"
	"golang.org/x/oauth2"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"

	"github.com/Skarlso/gtui/models"
	"github.com/Skarlso/gtui/pkg/providers"
)

const itemsPerPage = 100

var repoExtract = regexp.MustCompile(`repos/(.*)/(.*)/issues/(\d+)`)

// Config .
type Config struct {
	Token       string
	MaxFetchers int64
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
func (g *GithubProvider) ListRepositoryProjects(ctx context.Context, owner, repo string, opts *models.ListOptions) ([]*models.Project, error) {
	result := make([]*models.Project, 0)
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	var (
		page    int
		perPage = itemsPerPage
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
		perPage = itemsPerPage
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
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	opts := &github.ListOptions{
		PerPage: itemsPerPage,
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
		allCards, response, err := g.Client.Projects.ListProjectCards(ctx, c.GetID(), listOpts)
		if err != nil {
			g.logGithubResponseBody(response)
			g.Logger.Debug().Err(err).Int64("column_id", c.GetID()).Str("name", c.GetName()).Msg("Failed get column cards.")
			return nil, err
		}
		col := &models.ProjectColumn{
			Name: c.GetName(),
			ID:   c.GetID(),
		}
		cards := make([]*models.ProjectColumnCard, 0)
		e, ctx := errgroup.WithContext(ctx)
		sem := semaphore.NewWeighted(g.MaxFetchers)
		for _, card := range allCards {
			card := card
			e.Go(func() error {
				if err := sem.Acquire(ctx, 1); err != nil {
					return err
				}
				defer sem.Release(1)

				if card.Note != nil {
					cards = append(cards, &models.ProjectColumnCard{
						ID:   card.GetID(),
						Note: card.Note,
					})
					return nil
				}

				m := repoExtract.FindAllStringSubmatch(card.GetContentURL(), -1)
				if len(m) == 0 {
					g.Logger.Error().Str("url", card.GetContentURL()).Msg("Failed to extract repo owner data for url.")
					return errors.New("failed to extract repo owner data for card")
				}
				if len(m[0]) < 3 {
					g.Logger.Error().Str("url", card.GetContentURL()).Strs("matches", m[0]).Msg("Match groups didn't match 3.")
					return errors.New("failed to match repo, owner, issue id from url")
				}
				owner, repo, issueID := m[0][1], m[0][2], m[0][3]
				iID, err := strconv.Atoi(issueID)
				if err != nil {
					g.Logger.Debug().Err(err).Str("id", issueID).Msg("Failed to convert issue ID number to string.")
					return err
				}
				issue, response, err := g.Client.Issues.Get(ctx, owner, repo, iID)
				if err != nil {
					g.logGithubResponseBody(response)
					return err
				}
				assignee := "-"
				if issue.Assignee != nil {
					assignee = issue.Assignee.GetLogin()
				}
				cards = append(cards, &models.ProjectColumnCard{
					ID:       card.GetID(),
					Note:     card.Note,
					Title:    issue.GetTitle(),
					Content:  issue.GetBody(),
					Author:   issue.GetUser().GetLogin(),
					Assignee: assignee,
					IssueID:  issue.GetID(),
				})
				return nil
			})
		}
		if err := e.Wait(); err != nil {
			return nil, fmt.Errorf("failed to wait for all workers to finish fetching cards: %w", err)
		}
		col.ProjectColumnCards = cards
		data.ProjectColumns = append(data.ProjectColumns, col)
	}

	return data, nil
}

// LoadRest will fetch the rest of the cards if there are any.
func (g *GithubProvider) LoadRest(ctx context.Context, columnID int64, list *tview.List) error {
	listOpts := &github.ProjectCardListOptions{
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: itemsPerPage,
		},
	}
	allCards := make([]*github.ProjectCard, 0)
	for {
		cards, response, err := g.Client.Projects.ListProjectCards(ctx, columnID, listOpts)
		if err != nil {
			g.logGithubResponseBody(response)
			g.Logger.Debug().Err(err).Int64("column_id", columnID).Msg("Failed get column cards.")
			return err
		}
		// load but skip the first page
		if listOpts.Page == 0 && response.NextPage == 0 {
			break
		}
		allCards = append(allCards, cards...)
		if response.NextPage == 0 {
			break
		}
		listOpts.Page = response.NextPage
	}
	if len(allCards) == 0 {
		// no more cards to add
		return nil
	}
	e, ctx := errgroup.WithContext(ctx)
	sem := semaphore.NewWeighted(g.MaxFetchers)
	for _, card := range allCards {
		card := card
		e.Go(func() error {
			if err := sem.Acquire(ctx, 1); err != nil {
				return err
			}
			defer sem.Release(1)

			if card.Note != nil {
				list.AddItem(*card.Note, "", 0, nil)
				return nil
			}

			m := repoExtract.FindAllStringSubmatch(card.GetContentURL(), -1)
			if len(m) == 0 {
				g.Logger.Error().Str("url", card.GetContentURL()).Msg("Failed to extract repo owner data for url.")
				return errors.New("failed to extract repo owner data for card")
			}
			if len(m[0]) < 3 {
				g.Logger.Error().Str("url", card.GetContentURL()).Strs("matches", m[0]).Msg("Match groups didn't match 3.")
				return errors.New("failed to match repo, owner, issue id from url")
			}
			owner, repo, issueID := m[0][1], m[0][2], m[0][3]
			iID, err := strconv.Atoi(issueID)
			if err != nil {
				g.Logger.Debug().Err(err).Str("id", issueID).Msg("Failed to convert issue ID number to string.")
				return err
			}
			issue, response, err := g.Client.Issues.Get(ctx, owner, repo, iID)
			if err != nil {
				g.logGithubResponseBody(response)
				return err
			}
			assignee := "-"
			if issue.Assignee != nil {
				assignee = issue.Assignee.GetLogin()
			}
			secondaryText := fmt.Sprintf("Author: %s, Assignee: %s, IssueID: %d", issue.GetUser().GetLogin(), assignee, issue.GetID())
			list.AddItem(issue.GetTitle(), secondaryText, 0, nil)
			return nil
		})
	}
	if err := e.Wait(); err != nil {
		return fmt.Errorf("failed to wait for all workers to finish fetching cards: %w", err)
	}
	return nil
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

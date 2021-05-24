package providers

import (
	"context"

	"github.com/Skarlso/gtui/models"
)

// Github provides github specific api
//go:generate counterfeiter -o fakes/fake_github.go . Github
type Github interface {
	// ListRepositoryProjects for when the project is on a repository.
	ListRepositoryProjects(ctx context.Context, owner, repo string, opts *models.ListOptions) ([]*models.Project, error)
	// ListOrganizationProjects for when the project is on an organization.
	ListOrganizationProjects(ctx context.Context, org string, opts *models.ListOptions) ([]*models.Project, error)
	// GetProject should work once the project ID is known.
	GetProject(ctx context.Context, id int64) (*models.Project, error)
	// GetProjectData returns all the data for a project to show the project management page.
	GetProjectData(ctx context.Context, id int64) (*models.ProjectData, error)
}

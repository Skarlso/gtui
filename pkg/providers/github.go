package providers

import (
	"context"

	"github.com/Skarlso/gtui/models"
)

// Github provides github specific api
//go:generate counterfeiter -o fakes/fake_github.go . Github
type Github interface {
	ListProjects(ctx context.Context, org string, opts models.ListOptions) ([]*models.Project, error)
	GetProject(ctx context.Context, name string) (*models.Project, error)
}

package github

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-github/v35/github"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/Skarlso/gtui/models"
)

type mockTransport struct {
	res *http.Response
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.res, nil
}

func TestGithubProvider_GetProject(t *testing.T) {
	content, err := ioutil.ReadFile(filepath.Join("testdata", "get_project.json"))
	assert.NoError(t, err)
	logger := zerolog.New(os.Stderr)
	client := &http.Client{
		Transport: &mockTransport{
			res: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(content)),
			},
		},
	}
	gClient := GithubProvider{
		Config: Config{
			Token: "token",
		},
		Client: github.NewClient(client),
		Logger: logger,
	}
	expected := &models.Project{
		Name: "Projects Documentation",
		ID:   1002604,
	}
	got, err := gClient.GetProject(context.Background(), 1002604)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGithubProvider_ListOrganizationProjects(t *testing.T) {
	content, err := ioutil.ReadFile(filepath.Join("testdata", "list_projects.json"))
	assert.NoError(t, err)
	logger := zerolog.New(os.Stderr)
	client := &http.Client{
		Transport: &mockTransport{
			res: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(content)),
			},
		},
	}
	gClient := GithubProvider{
		Config: Config{
			Token: "token",
		},
		Client: github.NewClient(client),
		Logger: logger,
	}
	expected := &models.Project{
		Name: "Organization Roadmap",
		ID:   1002605,
	}
	got, err := gClient.ListOrganizationProjects(context.Background(), "octocat", nil)
	assert.NoError(t, err)
	assert.Len(t, got, 1)
	assert.Equal(t, expected, got[0])
}

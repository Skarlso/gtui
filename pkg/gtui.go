package pkg

import (
	"context"
	"fmt"

	"github.com/Skarlso/gtui/pkg/providers"
	"github.com/rs/zerolog"
)

// Config contains configuration properties for GTUI.
type Config struct {
	Organization string
	Repository   string
	ProjectID    int64
}

type Dependencies struct {
	Github providers.Github
	Logger zerolog.Logger
}

// GTUIClient defines a client for GTUI.
type GTUIClient struct {
	Config
	Dependencies
}

// NewGTUIClient creates a tui client with all the configs and dependencies needed.
func NewGTUIClient(cfg Config, deps Dependencies) *GTUIClient {
	return &GTUIClient{
		Dependencies: deps,
		Config:       cfg,
	}
}

// Start launches the GTUI App.
func (g *GTUIClient) Start() error {

	// Show based on what's provided?
	if g.ProjectID != -1 {
		return g.showProjectData()
	}
	if g.Repository != "" && g.Organization != "" {
		if err := g.showRepositoryProjectSelector(); err != nil {
			g.Logger.Debug().Err(err).Msg("Failed to show repository project selector.")
			return err
		}
	} else if g.Organization != "" && g.Repository == "" {
		if err := g.showOrganizationProjectSelector(); err != nil {
			g.Logger.Debug().Err(err).Msg("Failed to show organization project selector.")
			return err
		}
	}
	project, err := g.Github.GetProjectData(context.Background(), g.ProjectID)
	if err != nil {
		g.Logger.Debug().Err(err).Msg("Failed to get project data.")
		return err
	}
	for _, c := range project.ProjectColumns {
		fmt.Println(c)
	}
	//app := tview.NewApplication()
	//flex := tview.NewFlex().
	//	AddItem(tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
	//	AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
	//		AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false).
	//		AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false).
	//		AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 5, 1, false), 0, 2, false).
	//	AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
	//if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
	//	return err
	//}
	return nil
}

func (g *GTUIClient) showRepositoryProjectSelector() error {
	// call show project data after user selected a project and projectID has been set on `g` receiver
	return nil
}

func (g *GTUIClient) showOrganizationProjectSelector() error {
	// call show project data after user selected a project and projectID has been set on `g` receiver
	return nil
}

func (g *GTUIClient) showProjectData() error {
	return nil
}

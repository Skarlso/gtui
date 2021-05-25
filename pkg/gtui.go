package pkg

import (
	"github.com/rivo/tview"
	"github.com/rs/zerolog"

	"github.com/Skarlso/gtui/pkg/providers"
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

	app    *tview.Application
	middle *tview.Box
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
	app := tview.NewApplication()
	middle := tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)")
	g.app = app
	g.middle = middle
	if g.ProjectID != -1 {
		if err := g.showProjectData(); err != nil {
			return err
		}
	} else if g.Repository != "" && g.Organization != "" {
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
	flex := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("GTUI"), 0, 1, false).
			AddItem(middle, 0, 3, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 5, 1, false), 0, 2, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
	g.app = app
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		return err
	}
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
	// update the app
	g.middle.SetTitle("Project board")
	return nil
}

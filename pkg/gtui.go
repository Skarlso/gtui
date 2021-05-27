package pkg

import (
	"time"

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

	app         *tview.Application
	middleFlex  *tview.Flex
	projectList *tview.List
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
	projectList := tview.NewList().ShowSecondaryText(false)
	middleFlex := tview.NewFlex()
	g.app = app
	g.middleFlex = middleFlex
	g.projectList = projectList

	go func() {
		time.Sleep(3 * time.Second)
		//g.middleFlex = tview.NewFlex()
		//list := tview.NewList().ShowSecondaryText(false)
		//list.AddItem("new stuff", "", 0, nil)
		//g.middleFlex.AddItem(list, 0, 4, true)
		g.middleFlex.SetTitle("UPDATED")
		g.middleFlex.Clear() // Remove all items instead of RemoveItem!
		list := tview.NewList()
		list.SetBorder(true).SetTitle("NewList")
		list.AddItem("new", "", 0, nil)
		g.middleFlex.AddItem(list, 0, 4, true)
		g.app.Draw()
	}()

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
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("GTUI"), 5, 1, false).
			AddItem(middleFlex, 0, 1, true), 0, 1, true)
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
	//g.middle.SetTitle("Project board")
	g.middleFlex.SetBorder(true).SetTitle("Project Selector")
	g.projectList.AddItem("project", "", 0, nil)
	g.middleFlex.AddItem(g.projectList, 0, 1, true)
	return nil
}

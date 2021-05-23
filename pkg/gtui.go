package pkg

import (
	"github.com/rivo/tview"

	"github.com/Skarlso/gtui/pkg/providers"
)

// Config contains configuration properties for GTUI.
type Config struct {
}

type Dependencies struct {
	Github providers.Github
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
	app := tview.NewApplication()
	flex := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 5, 1, false), 0, 2, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		return err
	}
	return nil
}

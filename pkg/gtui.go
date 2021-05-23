package pkg

import (
	"github.com/rivo/tview"

	"github.com/Skarlso/gtui/pkg/providers"
)

// Config contains configuration properties for GTUI.
type Config struct {
	Token string
}

type Dependencies struct {
	Github providers.Github
}

// GTUIClient defines a client for GTUI.
type GTUIClient struct {
	Config
}

func NewGTUIClient(cfg Config) *GTUIClient {
	return &GTUIClient{
		Config: cfg,
	}
}

// Start launches the GTUI App.
func (g *GTUIClient) Start() error {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	main := newPrimitive("Project")

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, false).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(main, 1, 0, 1, 3, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(main, 1, 1, 1, 1, 0, 100, false)

	if err := tview.NewApplication().SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		return err
	}
	return nil
}

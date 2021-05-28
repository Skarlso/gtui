package pkg

import (
	"context"
	"fmt"

	"github.com/gdamore/tcell/v2"
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

	app        *tview.Application
	middleFlex *tview.Flex
	status     *tview.TextView
	lists      []*tview.List
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
	middleFlex := tview.NewFlex()
	status := tview.NewTextView()
	status.SetTitle("[red]Welcome To GTUI")
	status.SetBorder(true)
	status.SetWordWrap(true)
	status.SetDynamicColors(true)
	g.app = app
	g.middleFlex = middleFlex
	g.status = status
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
			AddItem(status, 5, 1, false).
			AddItem(middleFlex, 0, 1, true), 0, 1, true)
	g.app = app
	if err := app.SetRoot(flex, true).EnableMouse(true).SetFocus(flex).Run(); err != nil {
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
	data, err := g.Github.GetProjectData(context.Background(), g.ProjectID)
	if err != nil {
		g.Logger.Debug().Err(err).Int64("project_id", g.ProjectID).Msg("Failed to get project data")
		return err
	}
	g.middleFlex.SetBorder(true).SetTitle(fmt.Sprintf("Project #%d", g.ProjectID))
	for _, c := range data.ProjectColumns {
		textView := tview.NewTextView()
		textView.SetWordWrap(true)
		list := tview.NewList()
		list.SetBorder(true)
		list.SetTitle(c.Name)
		list.SetWrapAround(true)
		list.SetMainTextColor(tcell.ColorRed)
		list.SetSecondaryTextColor(tcell.ColorLightGreen)
		list.SetTitleColor(tcell.ColorLightGoldenrodYellow)
		for _, card := range c.ProjectColumnCards {
			title := card.Title
			secondaryText := fmt.Sprintf("Author: %s, Assignee: %s, IssueID: %d", card.Author, card.Assignee, card.IssueID)
			if card.Note != nil {
				title = fmt.Sprintf("[gray]%s", *card.Note)
				secondaryText = ""
			}
			list.AddItem(title, secondaryText, 0, nil)
		}
		list.SetBorderColor(tcell.ColorMediumPurple)
		list.SetSelectedBackgroundColor(tcell.ColorYellow)
		list.SetSelectedFocusOnly(true)
		list.SetSelectedFunc(g.ListEnterHandler)
		list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTab {
				g.cycleFocus(false)
			} else if event.Key() == tcell.KeyBacktab {
				g.cycleFocus(true)
			}
			return event
		})
		g.lists = append(g.lists, list)
		g.middleFlex.AddItem(list, 0, 1, true)
	}
	return nil
}

// ListEnterHandler handles issue enter presses for a list.
func (g *GTUIClient) ListEnterHandler(i int, mainText string, secondaryText string, shortcut rune) {
	g.status.SetText(fmt.Sprintf("%s, %s", mainText, secondaryText))
}

func (g *GTUIClient) cycleFocus(reverse bool) {
	for i, el := range g.lists {
		if !el.HasFocus() {
			continue
		}

		if reverse {
			i = i - 1
			if i < 0 {
				i = len(g.lists) - 1
			}
		} else {
			i = i + 1
			i = i % len(g.lists)
		}

		g.app.SetFocus(g.lists[i])
		return
	}
}

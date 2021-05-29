package pkg

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rs/zerolog"

	"github.com/Skarlso/gtui/pkg/providers"
)

var issueID = regexp.MustCompile(`IssueID: \[yellow\](\d+)\[lightgreen\]`)

// Config contains configuration properties for GTUI.
type Config struct {
	Organization   string
	Repository     string
	ProjectID      int64
	ColumnsPerPage int
}

type Dependencies struct {
	Github providers.Github
	Logger zerolog.Logger
}

// GTUIClient defines a client for GTUI.
type GTUIClient struct {
	Config
	Dependencies

	app      *tview.Application
	status   *tview.TextView
	columns  []*tview.List
	issueMap map[int]string
	pages    *tview.Pages
}

// NewGTUIClient creates a tui client with all the configs and dependencies needed.
func NewGTUIClient(cfg Config, deps Dependencies) *GTUIClient {
	return &GTUIClient{
		Dependencies: deps,
		Config:       cfg,
		issueMap:     make(map[int]string),
	}
}

// Start launches the GTUI App.
func (g *GTUIClient) Start() error {
	// Show based on what's provided?
	app := tview.NewApplication()
	status := tview.NewTextView()
	status.SetTitle("[red]Welcome To GTUI")
	status.SetBorder(true)
	status.SetWordWrap(true)
	status.SetDynamicColors(true)
	pages := tview.NewPages()
	g.app = app
	g.status = status
	g.pages = pages
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
			AddItem(status, 10, 1, false).
			AddItem(g.pages, 0, 1, true), 0, 1, true)
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
	//flexes := make([]*tview.Flex, 0)
	//middleFlex := tview.NewFlex()
	//middleFlex.SetBorder(true).SetTitle(fmt.Sprintf("Project #%d", g.ProjectID))
	for _, c := range data.ProjectColumns {
		textView := tview.NewTextView()
		textView.SetWordWrap(true)
		list := tview.NewList()
		list.SetBorder(true)
		list.SetTitle(c.Name)
		list.SetWrapAround(true)
		//list.SetMainTextColor(tcell.ColorBlueViolet)
		list.SetSecondaryTextColor(tcell.ColorLightGreen)
		list.SetTitleColor(tcell.ColorLightGoldenrodYellow)
		for _, card := range c.ProjectColumnCards {
			g.issueMap[int(card.IssueID)] = card.Content
			title := card.Title
			secondaryText := fmt.Sprintf("Author: [yellow]%s[lightgreen], Assignee: [yellow]%s[lightgreen], IssueID: [yellow]%d[lightgreen]", card.Author, card.Assignee, card.IssueID)
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
		g.columns = append(g.columns, list)
		//middleFlex.AddItem(list, 0, 1, true)
		// launch background rest fetcher
		go func(id int64, name string, l *tview.List) {
			if err := g.Github.LoadRest(context.Background(), id, l); err != nil {
				g.status.SetText(fmt.Sprintf("[red]Failed to fetch cards in the background for column %d with name %s", id, name))
			}
		}(c.ID, c.Name, list)
	}
	// divi up the list and create groups of flexes and then create pages out of them
	pages := make([]*tview.Flex, 0)
	index := 0
	for {
		if index+g.ColumnsPerPage > len(g.columns) {
			middleFlex := tview.NewFlex()
			middleFlex.SetBorder(false)
			for _, l := range g.columns[index:] {
				middleFlex.AddItem(l, 0, 1, true)
			}
			pages = append(pages, middleFlex)
			break
		}
		list := g.columns[index : index+g.ColumnsPerPage]
		middleFlex := tview.NewFlex()
		middleFlex.SetBorder(false)
		for _, l := range list {
			middleFlex.AddItem(l, 0, 1, true)
		}
		pages = append(pages, middleFlex)
		index += g.ColumnsPerPage
	}

	for i, p := range pages {
		name := fmt.Sprintf("%d/%d", i+1, len(pages))
		g.pages.AddPage(name, p, true, true)
	}
	// focus on the first page
	g.pages.SwitchToPage(fmt.Sprintf("1/%d", len(pages)))
	g.pages.SetBorder(true)
	g.pages.SetTitle(fmt.Sprintf("Project #%d (%d)", g.ProjectID, len(pages)))
	return nil
}

// ListEnterHandler handles issue enter presses for a list.
func (g *GTUIClient) ListEnterHandler(i int, mainText string, secondaryText string, shortcut rune) {
	content := mainText
	if secondaryText != "" {
		m := issueID.FindAllStringSubmatch(secondaryText, -1)
		if len(m) == 0 {
			g.status.SetText("[red]failed to match out issue code")
			return
		}
		if len(m[0]) < 1 {
			g.status.SetText("[red]failed to parse out issue ID")
			return
		}
		i, err := strconv.Atoi(m[0][1])
		if err != nil {
			g.status.SetText(err.Error())
			return
		}
		content = g.issueMap[i]
	}
	g.status.SetText(content)
}

func (g *GTUIClient) cycleFocus(reverse bool) {
	for i, el := range g.columns {
		if !el.HasFocus() {
			continue
		}

		if reverse {
			i--
			if i < 0 {
				i = len(g.columns) - 1
			}
		} else {
			i++
			i = i % len(g.columns)
		}

		page := (i / g.ColumnsPerPage) + 1
		if page == 0 {
			page++
		}
		name := fmt.Sprintf("%d/%d", page, (len(g.columns)/g.ColumnsPerPage)+1)
		g.pages.SwitchToPage(name)
		g.app.SetFocus(g.columns[i])
		return
	}
}

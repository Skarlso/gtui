package pkg

import (
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
	// TODO: Use this https://github.com/rivo/tview/
	return nil
}

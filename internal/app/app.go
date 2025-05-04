package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/romanitalian/catch-my-file/internal/config"
)

// Application represents the main application
type Application struct {
	fyneApp fyne.App
	window  fyne.Window
	config  *config.Config
}

// New creates a new application instance
func New(cfg *config.Config) *Application {
	return &Application{
		fyneApp: app.New(),
		config:  cfg,
	}
}

// Run starts the application
func (a *Application) Run() {
	a.window = a.fyneApp.NewWindow(a.config.Title)
	a.setupUI()
	a.window.Resize(fyne.NewSize(float32(a.config.WindowWidth), float32(a.config.WindowHeight)))
	a.window.ShowAndRun()
}

// setupUI initializes the user interface
func (a *Application) setupUI() {
	// Create main content
	content := container.NewVBox(
		widget.NewLabel("Welcome to Catch My File!"),
		widget.NewButton("Select File", func() {
			// TODO: Implement file selection
		}),
	)

	a.window.SetContent(content)
}

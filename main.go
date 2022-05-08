package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	client := NewClient()
	// app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "UDP_CLIENT",
		Width:  500,
		Height: 768,
		Assets: assets,
		OnStartup:  client.startup,
    OnShutdown: client.shutdown,
		Bind: []interface{}{
			client,
			// app,
		},
	})

	if err != nil {
		println("Error:", err)
	}
}

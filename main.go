package main

import (
	// base

	"os"
	"time"

	// Discord Rich Presence Library
	"github.com/hugolgst/rich-go/client"

	// GUI library "Fyne"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	// Define required variables
	now := time.Now()
	DiscordAppID := os.Getenv("770195570173804575")

	// Login to Discord
	errLogin := client.Login(DiscordAppID)
	if errLogin != nil {
		panic(errLogin)
	}

	// Fyne window setup
	a := app.New()
	w := a.NewWindow("ぎゅっとe Discord RPC") // Window title

	// Body of window
	w.SetContent(widget.NewVBox(
		widget.NewButton("Reading", func() { // Reading
			client.SetActivity(client.Activity{
				State:      "Playing",
				Details:    "Reading",
				LargeImage: "e",
				LargeText:  "ぎゅっとe: eラーニングシステム",
				Timestamps: &client.Timestamps{
					Start: &now,
				},
			})
		}),
		widget.NewButton("Vocabulary", func() { // Vocabulary
			client.SetActivity(client.Activity{
				State:      "Playing",
				Details:    "Vocabulary",
				LargeImage: "e",
				LargeText:  "ぎゅっとe: eラーニングシステム",
				Timestamps: &client.Timestamps{
					Start: &now,
				},
			})
		}),
		widget.NewButton("Listening", func() { // Listening
			client.SetActivity(client.Activity{
				State:      "Playing",
				Details:    "Listening",
				LargeImage: "e",
				LargeText:  "ぎゅっとe: eラーニングシステム",
				Timestamps: &client.Timestamps{
					Start: &now,
				},
			})
		}),
		widget.NewButton("Grammar", func() { // Grammar
			client.SetActivity(client.Activity{
				State:      "Playing",
				Details:    "Grammar",
				LargeImage: "e",
				LargeText:  "ぎゅっとe: eラーニングシステム",
				Timestamps: &client.Timestamps{
					Start: &now,
				},
			})
		}),
		widget.NewButton("MaG", func() { // ぎゅっとMaG
			client.SetActivity(client.Activity{
				State:      "Playing",
				Details:    "ぎゅっとMaG",
				LargeImage: "e",
				LargeText:  "ぎゅっとe: eラーニングシステム",
				Timestamps: &client.Timestamps{
					Start: &now,
				},
			})
		}),
	))

	// Resize & Disable maximum button
	w.Resize(fyne.NewSize(300, 150))
	w.FixedSize()
	w.SetFixedSize(true)

	// Show window
	w.ShowAndRun()
}

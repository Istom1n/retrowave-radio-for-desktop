package main

import (
	"log"

	"github.com/0xAX/notificator"
	"github.com/andlabs/ui"
)

var notify *notificator.Notificator

func main() {
	err := ui.Main(func() {
		input := ui.NewEntry()
		button := ui.NewButton("Greet")
		greeting := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter your name:"), false)
		box.Append(input, false)
		box.Append(button, false)
		box.Append(greeting, false)
		window := ui.NewWindow("Hello", 200, 100, false)
		window.SetMargined(true)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			greeting.SetText("Hello, " + input.Text() + "!")
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}

	notificate("https://retrowave.ru/artwork/1ffecb1c64a70d2554e9db6a259165fe3680b58e.jpg", "Dead Astronauts – Witch Hunt")
}

func notificate(artworkURL string, title string) {
	notify = notificator.New(notificator.Options{
		DefaultIcon: artworkURL,
		AppName:     title,
	})

	notify.Push("title", "text", "/home/user/icon.png", notificator.UR_CRITICAL)

	// Check errorsgo run
	err := notify.Push("error", "ops =(", "/home/user/icon.png", notificator.UR_CRITICAL)

	if err != nil {
		log.Fatal(err)
	}
}

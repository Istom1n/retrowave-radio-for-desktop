package main

import (
	"log"

	"github.com/0xAX/notificator"
)

var notify *notificator.Notificator

func main() {
	// webview.Open("Retrowave Radio for Mac",
	// 	"https://retrowave.ru", 530, 400, true)
	notificate("https://retrowave.ru/artwork/1ffecb1c64a70d2554e9db6a259165fe3680b58e.jpg", "Dead Astronauts – Witch Hunt")
}

func notificate(artworkURL string, title string) {
	notify = notificator.New(notificator.Options{
		DefaultIcon: artworkURL,
		AppName:     title,
	})

	notify.Push("title", "text", "/home/user/icon.png", notificator.UR_CRITICAL)

	// Check errors
	err := notify.Push("error", "ops =(", "/home/user/icon.png", notificator.UR_CRITICAL)

	if err != nil {
		log.Fatal(err)
	}
}

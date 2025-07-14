package main

import "strings"

var player *Player
var rooms map[string]*Room

func initGame() {
	player = &Player{
		inventory:  make(map[string]bool),
		wearingBag: false,
		doorIsOpen: false,
	}
	initRooms()
	player.location = rooms["кухня"]
}

func handleCommand(input string) string {
	parts := strings.Split(input, " ")
	cmd := parts[0]
	args := parts[1:]

	switch cmd {
	case "осмотреться":
		return player.Look()
	case "идти":
		return player.Move(args)
	case "взять":
		return player.Take(args)
	case "надеть":
		return player.Wear(args)
	case "применить":
		return player.Apply(args)
	default:
		return "неизвестная команда"
	}

}

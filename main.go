package main

import "cart/game"

//go:export update
func update() {
	game.Update()
	game.Draw()
}

//go:export start
func start() {
	game.Initialize()
}

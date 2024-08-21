package main

import (
	"fmt"
	"game2048/game"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in safeCall:", r)
		}
	}()
	game := game.NewGame()
	game.Test()
	game.Start()

}

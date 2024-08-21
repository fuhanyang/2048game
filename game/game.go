package game

import (
	"fmt"
	"game2048/player"
	"game2048/resource"
)

var gameflag bool
var target int = 8

type level struct {
	wid int
	len int
}
type Game struct {
	opr   player.Operator
	board *resource.Board
	level *level
}

func (game *Game) Test() {
	game.board.OnScreen()
}

func NewGame() *Game {
	gameflag = true
	level := level{wid: 4, len: 4}
	randPos := resource.GetRand()
	randVal := resource.GetRand()

	game := &Game{
		opr:   *player.NewOpr(),
		board: resource.NewBoard(level.wid, level.len),
		level: &level,
	}
	game.board.BoardData[randPos] = (randVal%2 + 1) * 2
	return game
}
func (game *Game) Start() {
	for {
		doOpt(game, &game.opr)
		randPos := resource.GetRand()
		randVal := (resource.GetRand()%2 + 1) * 2
		game.board.AddNum(randPos, randVal)
		game.board.OnScreen()
		if !game.IfContinue() {
			fmt.Println("成功完成游戏")
			return
		}
	}

}
func (game *Game) IfContinue() bool {
	return gameflag
}

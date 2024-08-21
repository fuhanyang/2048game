package game

import (
	"game2048/player"
	"game2048/resource"
)

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
	}

}

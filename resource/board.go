package resource

import "fmt"

type Board struct {
	BoardData []int
	wid       int
	len       int
}

func NewBoard(wid int, len int) *Board {
	Data := make([]int, wid*len)
	board := &Board{
		BoardData: Data,
		wid:       wid,
		len:       len,
	}
	return board
}
func (board *Board) OnScreen() {
	for i := 0; i < board.wid; i++ {
		for j := 0; j < board.len; j++ {
			s := board.BoardData[i*board.len+j]
			fmt.Print("-", s, "-")
		}
		fmt.Print("\n")
	}
}
func (board *Board) AddNum(pos int, val int) {
	num := board.len * board.wid
	j := 0
	for i := 0; i < num; i++ {
		if board.BoardData[i] == 0 {
			j++
			if j == pos {
				board.BoardData[i] = val
			}
		}
	}
}

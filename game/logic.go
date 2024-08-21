package game

import (
	"game2048/player"
)

type Handle interface {
	onHandle(player.OPT)
}

func (handle *Game) up() {
	wid := handle.level.wid
	len := handle.level.len
	for i := 0; i < len; i++ {
		flag := -1
		for j := 0; j < wid; j++ {
			pos := j*len + i
			//找到第一个空位置记录
			if handle.board.BoardData[pos] == 0 && flag == -1 {
				flag = pos
			}
			//找到不空的并移动
			if handle.board.BoardData[pos] != 0 && flag != -1 {
				handle.board.BoardData[flag] = handle.board.BoardData[pos]
				handle.board.BoardData[pos] = 0
				flag += len
			}
		}
		top := i
		for j := 1; j < wid && handle.board.BoardData[j*len+i] != 0; j++ {
			pos := j*len + i
			//找到第一个能换的
			if handle.board.BoardData[pos] == handle.board.BoardData[top] {
				handle.board.BoardData[top] *= 2
				handle.board.BoardData[pos] = 0
				continue
			}
			//对齐
			if handle.board.BoardData[pos-len] == 0 {
				handle.board.BoardData[pos], handle.board.BoardData[pos-len] = 0, handle.board.BoardData[pos]
			}
			top += len

		}
	}
}
func (handle *Game) down() {
	wid := handle.level.wid
	len := handle.level.len
	for i := len - 1; i >= 0; i-- {
		flag := -1
		for j := wid - 1; j >= 0; j-- {
			pos := j*len + i
			//找到第一个空位置记录
			if handle.board.BoardData[pos] == 0 && flag == -1 {
				flag = pos
			}
			//找到不空的并移动
			if handle.board.BoardData[pos] != 0 && flag != -1 {
				handle.board.BoardData[flag] = handle.board.BoardData[pos]
				handle.board.BoardData[pos] = 0
				flag -= len
			}
		}
		top := i + len*(wid-1)
		for j := wid - 2; j >= 0 && handle.board.BoardData[j*len+i] != 0; j-- {
			pos := j*len + i
			//找到第一个能换的
			if handle.board.BoardData[pos] == handle.board.BoardData[top] {
				handle.board.BoardData[top] *= 2
				handle.board.BoardData[pos] = 0
				continue
			}
			//对齐
			if handle.board.BoardData[pos+len] == 0 {
				handle.board.BoardData[pos], handle.board.BoardData[pos+len] = 0, handle.board.BoardData[pos]
			}
			top -= len

		}
	}
}
func (handle *Game) left() {
	wid := handle.level.wid
	len := handle.level.len
	for i := 0; i < wid; i++ {
		flag := -1
		for j := 0; j < len; j++ {
			pos := j + i*len
			//找到第一个空位置记录
			if handle.board.BoardData[pos] == 0 && flag == -1 {
				flag = pos
			}
			//找到不空的并移动
			if handle.board.BoardData[pos] != 0 && flag != -1 {
				handle.board.BoardData[flag] = handle.board.BoardData[pos]
				handle.board.BoardData[pos] = 0
				flag += 1
			}
		}
		top := i * len
		for j := 1; j < len && handle.board.BoardData[j+top] != 0; j++ {
			pos := j + top
			//找到第一个能换的
			if handle.board.BoardData[pos] == handle.board.BoardData[top] {
				handle.board.BoardData[top] *= 2
				handle.board.BoardData[pos] = 0
				continue
			}
			//对齐
			if handle.board.BoardData[pos-1] == 0 {
				handle.board.BoardData[pos], handle.board.BoardData[pos-1] = 0, handle.board.BoardData[pos]
			}
			top += 1

		}
	}
}
func (handle *Game) right() {
	wid := handle.level.wid
	len := handle.level.len
	for i := wid - 1; i >= 0; i-- {
		flag := -1
		for j := len - 1; j >= 0; j-- {
			pos := j + i*len
			//找到第一个空位置记录
			if handle.board.BoardData[pos] == 0 && flag == -1 {
				flag = pos
			}
			//找到不空的并移动
			if handle.board.BoardData[pos] != 0 && flag != -1 {
				handle.board.BoardData[flag] = handle.board.BoardData[pos]
				handle.board.BoardData[pos] = 0
				flag -= 1
			}
		}
		top := i*len + wid - 1
		for j := len - 2; j >= 0 && handle.board.BoardData[j+i*len] != 0; j-- {
			pos := j + i*len
			//找到第一个能换的
			if handle.board.BoardData[pos] == handle.board.BoardData[top] {
				handle.board.BoardData[top] *= 2
				handle.board.BoardData[pos] = 0
				continue
			}
			//对齐
			if handle.board.BoardData[pos+1] == 0 {
				handle.board.BoardData[pos], handle.board.BoardData[pos+1] = 0, handle.board.BoardData[pos]
			}
			top -= 1

		}
	}
}
func (game *Game) onHandle(opt player.OPT) {
	if opt == player.UP {
		game.up()
	}
	if opt == player.DOWN {
		game.down()
	}
	if opt == player.LEFT {
		game.left()
	}
	if opt == player.RIGHT {
		game.right()
	}
}
func doOpt(handle Handle, opr player.Opr) { //这里的接口变量要接受指针
	opt, flag := opr.GetInput()
	for !flag {
		opt, flag = opr.GetInput()
	}
	handle.onHandle(opt)
}

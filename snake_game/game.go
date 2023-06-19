package main

import (
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	Screen    tcell.Screen
	snakeBody SnakeBody
}

func (g *Game) Run() {

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGreen)
	g.Screen.SetStyle(defStyle)
	width, height := g.Screen.Size()
	snakeStyle := tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorGreen)

	for {
		g.Screen.Clear()
		g.snakeBody.Update(width, height)
		g.Screen.SetContent(g.snakeBody.X, g.snakeBody.Y, ' ', nil, snakeStyle)
		time.Sleep(40 * time.Millisecond)
		g.Screen.Show()
	}

}

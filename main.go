package main

import "ttt-go/ttt"

func main() {
  board := ttt.NewBoard(9)
  ui := new(ttt.UI)
  pg := PlayerGenerator{}
  tttRules := ttt.NewRules(board)
  ttt.NewGame(ui, pg)
}

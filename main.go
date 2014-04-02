package main

import "ttt-go/ttt"

func main() {
  board := ttt.NewBoard(9)
  ui := new(ttt.UI)
  ttt.NewGame(ui)
  tttRules := ttt.NewRules(board)
  ui.PrintBoard(tttRules)
}

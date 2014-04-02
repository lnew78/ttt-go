package main

import "ttt-go/ttt"

func main() {
  board := ttt.NewBoard(9)
  ui := new(ttt.UI)
  tttRules := ttt.NewRules(board)
  ui.PrintBoard(tttRules)
}

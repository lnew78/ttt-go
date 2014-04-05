package main

import "ttt-go/ttt"

func main() {
  board := ttt.NewBoard(9)
  tttRules := ttt.NewRules(board)
  pg := ttt.PlayerGenerator{}
  ui := new(ttt.UI)
  minimax := ttt.NewMinimax(board, tttRules)
  ttt.NewGame(board, tttRules, pg, ui, minimax)
}

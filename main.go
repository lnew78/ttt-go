package main

import "ttt-go/ttt"

func main() {
  board := ttt.NewBoard(9)
  ui := new(ttt.UI)
  pg := ttt.PlayerGenerator{}
  tttRules := ttt.NewRules(board)
  minimax := ttt.NewMinimax(board, tttRules)
  ttt.NewGame(ui, pg, tttRules, board, minimax)
}

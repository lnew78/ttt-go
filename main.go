package main

import "ttt-go/ttt"

func main() {
  board := ttt.NewBoard(9)
  ui := new(ttt.UI)
  tttRules := ttt.TicTacToeRules{}
  ui.PrintBoard(tttRules, board)
}

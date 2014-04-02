package ttt

import "strings"

type TicTacToeRules struct {
  board *Board
}

const XMARK string  = "x"
const OMARK string  = "o"
const NOWINNER string = "none"

func NewRules(board *Board) TicTacToeRules {
  tttRules := TicTacToeRules{}
  tttRules.board = board
  return tttRules
}

func (tttRules TicTacToeRules) Rows() ([]string, []string, []string) {
  row1 := []string{tttRules.board.spaces[0], tttRules.board.spaces[1], tttRules.board.spaces[2]}
  row2 := []string{tttRules.board.spaces[3], tttRules.board.spaces[4], tttRules.board.spaces[5]}
  row3 := []string{tttRules.board.spaces[6], tttRules.board.spaces[7], tttRules.board.spaces[8]}
  return row1, row2, row3
}

func (tttRules TicTacToeRules) Columns(board *Board) ([]string, []string, []string) {
  col1 := []string{board.spaces[0], board.spaces[3], board.spaces[6]}
  col2 := []string{board.spaces[1], board.spaces[4], board.spaces[7]}
  col3 := []string{board.spaces[2], board.spaces[5], board.spaces[8]}
  return col1, col2, col3
}

func (tttRules TicTacToeRules) Diagonals(board *Board) ([]string, []string) {
  diag1 := []string{board.spaces[0], board.spaces[4], board.spaces[8]}
  diag2 := []string{board.spaces[2], board.spaces[4], board.spaces[6]}
  return diag1, diag2
}

func (tttRules TicTacToeRules) IsTie(board *Board) bool {
 return board.IsFull() && tttRules.Winner(board) == "none"
}

func (tttRules TicTacToeRules) Winner(board *Board) string {
  return tttRules.checkCombos(board)
}

func (tttRules TicTacToeRules) checkCombos(board *Board) string {
  row1, row2, row3 := tttRules.Rows()
  col1, col2, col3 := tttRules.Columns(board)
  diag1, diag2 := tttRules.Diagonals(board)
  winningCombos := [][]string{row1, row2, row3, col1, col2, col3, diag1, diag2}
  winner := tttRules.checkForWinner(winningCombos)
  return winner
}

func (tttRules TicTacToeRules) checkForWinner(winningCombos [][]string) string {
  winner := ""
  for i := range winningCombos {
    winner = tttRules.checkComboForWinner(winningCombos[i])
    if winner != NOWINNER {
      break
    }
  }
  return winner
}

func (tttRules TicTacToeRules) checkComboForWinner(winningCombo []string) string {
  xCount := 0
  oCount := 0
  for _,mark := range winningCombo {
    if mark == XMARK {
      xCount++
    } else if mark == OMARK {
      oCount++
    }

    if xCount == 3 {
      return XMARK
    } else if oCount == 3 {
      return OMARK
    }
  }
  return NOWINNER
}

func (tttRules TicTacToeRules) NextPlayerMark(board *Board) string {
  mark := ""
  if board.NumberOfAvailableSpaces() % 2 == 0 {
     mark = OMARK
  } else {
    mark = XMARK
  }
  return mark
}

func (tttRules TicTacToeRules) IsGameOver(board *Board) bool {
  return tttRules.Winner(board) != "none" || tttRules.IsTie(board)
}

func (tttRules TicTacToeRules) IsMarkValid(mark string) bool {
  return strings.ToLower(mark) == XMARK || strings.ToLower(mark) == OMARK
}

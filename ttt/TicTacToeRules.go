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

func (tttRules TicTacToeRules) Columns() ([]string, []string, []string) {
  col1 := []string{tttRules.board.spaces[0], tttRules.board.spaces[3], tttRules.board.spaces[6]}
  col2 := []string{tttRules.board.spaces[1], tttRules.board.spaces[4], tttRules.board.spaces[7]}
  col3 := []string{tttRules.board.spaces[2], tttRules.board.spaces[5], tttRules.board.spaces[8]}
  return col1, col2, col3
}

func (tttRules TicTacToeRules) Diagonals() ([]string, []string) {
  diag1 := []string{tttRules.board.spaces[0], tttRules.board.spaces[4], tttRules.board.spaces[8]}
  diag2 := []string{tttRules.board.spaces[2], tttRules.board.spaces[4], tttRules.board.spaces[6]}
  return diag1, diag2
}

func (tttRules TicTacToeRules) IsTie() bool {
 return tttRules.board.IsFull() && tttRules.Winner() == "none"
}

func (tttRules TicTacToeRules) Winner() string {
  return tttRules.checkCombos()
}

func (tttRules TicTacToeRules) checkCombos() string {
  row1, row2, row3 := tttRules.Rows()
  col1, col2, col3 := tttRules.Columns()
  diag1, diag2 := tttRules.Diagonals()
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

func (tttRules TicTacToeRules) Mark(key string) string {
  constants := make(map[string]string)
  constants["XMARK"] = XMARK
  constants["OMARK"] = OMARK

  return constants[key]
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

func (tttRules TicTacToeRules) NextPlayerMark() string {
  mark := ""
  if tttRules.board.NumberOfAvailableSpaces() % 2 == 0 {
     mark = OMARK
  } else {
    mark = XMARK
  }
  return mark
}

func (tttRules TicTacToeRules) IsGameOver() bool {
  return tttRules.Winner() != "none" || tttRules.IsTie()
}

func (tttRules TicTacToeRules) IsMarkValid(mark string) bool {
  return strings.ToLower(mark) == XMARK || strings.ToLower(mark) == OMARK
}

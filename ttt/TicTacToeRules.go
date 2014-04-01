package ttt

type TicTacToeRules struct {
}

func (tttRules *TicTacToeRules) Rows(board *Board) ([]string, []string, []string) {
  row1 := []string{board.spaces[0], board.spaces[1], board.spaces[2]}
  row2 := []string{board.spaces[3], board.spaces[4], board.spaces[5]}
  row3 := []string{board.spaces[6], board.spaces[7], board.spaces[8]}
  return row1, row2, row3
}

func (tttRules *TicTacToeRules) Columns(board *Board) ([]string, []string, []string) {
  col1 := []string{board.spaces[0], board.spaces[3], board.spaces[6]}
  col2 := []string{board.spaces[1], board.spaces[4], board.spaces[7]}
  col3 := []string{board.spaces[2], board.spaces[5], board.spaces[8]}
  return col1, col2, col3
}

func (tttRules *TicTacToeRules) Diagonals(board *Board) ([]string, []string) {
  diag1 := []string{board.spaces[0], board.spaces[4], board.spaces[8]}
  diag2 := []string{board.spaces[2], board.spaces[4], board.spaces[6]}
  return diag1, diag2
}

func (tttRules *TicTacToeRules) Winner(board *Board) string {
  return tttRules.checkCombos(board)
}

func (tttRules *TicTacToeRules) checkCombos(board *Board) string {
  row1, row2, row3 := tttRules.Rows(board)
  col1, col2, col3 := tttRules.Columns(board)
  diag1, diag2 := tttRules.Diagonals(board)
  winningCombos := [][]string{row1, row2, row3, col1, col2, col3, diag1, diag2}
  winner := tttRules.checkForWinner(winningCombos)
  return winner
}

func (tttRules *TicTacToeRules) checkForWinner(winningCombos [][]string) string {
  xCount := 0
  oCount := 0
  for i := range winningCombos {
    for _,mark := range winningCombos[i] {
      if mark == "x" {
        xCount++
      } else if mark == "o" {
        oCount++
      }

      if xCount == 3 {
        return "x"
      } else if oCount == 3 {
        return "o"
      }
    }
  }
  return "none"
}

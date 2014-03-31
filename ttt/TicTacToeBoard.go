package ttt

type TicTacToeBoard struct {
}

func (tttBoard *TicTacToeBoard) Rows(board *Board) ([]string, []string, []string) {
  row1 := []string{board.spaces[0], board.spaces[1], board.spaces[2]}
  row2 := []string{board.spaces[3], board.spaces[4], board.spaces[5]}
  row3 := []string{board.spaces[6], board.spaces[7], board.spaces[8]}
  return row1, row2, row3
}

func (tttBoard *TicTacToeBoard) Columns(board *Board) ([]string, []string, []string) {
  col1 := []string{board.spaces[0], board.spaces[3], board.spaces[6]}
  col2 := []string{board.spaces[1], board.spaces[4], board.spaces[7]}
  col3 := []string{board.spaces[2], board.spaces[5], board.spaces[8]}
  return col1, col2, col3
}

func (tttBoard *TicTacToeBoard) Diagonals(board *Board) ([]string, []string) {
  diag1 := []string{board.spaces[0], board.spaces[4], board.spaces[8]}
  diag2 := []string{board.spaces[2], board.spaces[4], board.spaces[6]}
  return diag1, diag2
}

package ttt_test

import (
	. "ttt-go/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TicTacToeBoard", func() {
  var board = NewBoard(9)
  var tttBoard = new (TicTacToeBoard)

  It("Get the board's rows", func() {
    testRow1 := []string{"1", "2", "3"}
    testRow2 := []string{"4", "5", "6"}
    testRow3 := []string{"7", "8", "9"}
    row1,row2, row3 := tttBoard.Rows(board)
    Expect(row1).To(Equal(testRow1))
    Expect(row2).To(Equal(testRow2))
    Expect(row3).To(Equal(testRow3))
  })

  It("Get the board's columns", func() {
    testCol1 := []string{"1", "4", "7"}
    testCol2 := []string{"2", "5", "8"}
    testCol3 := []string{"3", "6", "9"}
    col1, col2, col3 := tttBoard.Columns(board)
    Expect(col1).To(Equal(testCol1))
    Expect(col2).To(Equal(testCol2))
    Expect(col3).To(Equal(testCol3))
  })

  It("Get the board's diagonals", func() {
    testDiag1 := []string{"1", "5", "9"}
    testDiag2 := []string{"3", "5", "7"}
    diag1, diag2 := tttBoard.Diagonals(board)
    Expect(diag1).To(Equal(testDiag1))
    Expect(diag2).To(Equal(testDiag2))
  })
})

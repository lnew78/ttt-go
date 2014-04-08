package ttt_test

import (
	. "github.com/lnew78/ttt-go/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TTTRules", func() {
  var board = NewBoard(9)
  var tttRules = NewRules(board)

  It("Get the board's rows", func() {
    testRow1 := []string{"1", "2", "3"}
    testRow2 := []string{"4", "5", "6"}
    testRow3 := []string{"7", "8", "9"}
    row1,row2, row3 := tttRules.Rows()
    Expect(row1).To(Equal(testRow1))
    Expect(row2).To(Equal(testRow2))
    Expect(row3).To(Equal(testRow3))
  })

  It("Get the board's columns", func() {
    testCol1 := []string{"1", "4", "7"}
    testCol2 := []string{"2", "5", "8"}
    testCol3 := []string{"3", "6", "9"}
    col1, col2, col3 := tttRules.Columns()
    Expect(col1).To(Equal(testCol1))
    Expect(col2).To(Equal(testCol2))
    Expect(col3).To(Equal(testCol3))
  })

  It("Get the board's diagonals", func() {
    testDiag1 := []string{"1", "5", "9"}
    testDiag2 := []string{"3", "5", "7"}
    diag1, diag2 := tttRules.Diagonals()
    Expect(diag1).To(Equal(testDiag1))
    Expect(diag2).To(Equal(testDiag2))
  })

  Context("Determine a tie game", func() {
    It("Tie game if board is full and there's no winner", func() {
      board.FillSpaceAt(1, "x")
      board.FillSpaceAt(2, "o")
      board.FillSpaceAt(3, "x")
      board.FillSpaceAt(4, "x")
      board.FillSpaceAt(5, "0")
      board.FillSpaceAt(6, "x")
      board.FillSpaceAt(7, "o")
      board.FillSpaceAt(8, "x")
      board.FillSpaceAt(9, "o")

     Expect(tttRules.IsTie()).To(Equal(true))
    })

    It("Is not a tie game if the board isn't full and there is no winner", func() {
      board.InitializeSpaces()

      board.FillSpaceAt(8, "x")
      board.FillSpaceAt(9, "o")

      Expect(tttRules.IsTie()).To(Equal(false))
    })
  })

  Context("Determine a winner", func() {
    It("Check rows", func() {
      board.InitializeSpaces()

      board.FillSpaceAt(1, "x")
      board.FillSpaceAt(2, "x")
      board.FillSpaceAt(3, "x")

      Expect(tttRules.Winner()).To(Equal("x"))
    })

    It("Check columns", func() {
      board.InitializeSpaces()

      board.FillSpaceAt(2, "o")
      board.FillSpaceAt(5, "o")
      board.FillSpaceAt(8, "o")

      Expect(tttRules.Winner()).To(Equal("o"))
    })

    It("Check diagonals", func() {
      board.InitializeSpaces()

      board.FillSpaceAt(3, "x")
      board.FillSpaceAt(5, "x")
      board.FillSpaceAt(7, "x")

      Expect(tttRules.Winner()).To(Equal("x"))
    })
  })

  It("'o' makes move if available space count is even", func() {
    board.InitializeSpaces()
    board.FillSpaceAt(5, "x")

    Expect(tttRules.NextPlayerMark()).To(Equal("o"))
  })

  It("'x' makes moves if available space count is even", func() {
    board.InitializeSpaces()

    Expect(tttRules.NextPlayerMark()).To(Equal("x"))
  })

  It("Game is over if there's a winner", func() {
    board.InitializeSpaces()

    board.FillSpaceAt(1, "x")
    board.FillSpaceAt(2, "x")
    board.FillSpaceAt(3, "x")

    Expect(tttRules.IsGameOver()).To(Equal(true))
  })

  It("Game is over if there's a tie", func() {
    board.InitializeSpaces()
    board.FillSpaceAt(1, "x")
    board.FillSpaceAt(2, "o")
    board.FillSpaceAt(3, "x")
    board.FillSpaceAt(4, "x")
    board.FillSpaceAt(5, "o")
    board.FillSpaceAt(6, "x")
    board.FillSpaceAt(7, "o")
    board.FillSpaceAt(8, "x")
    board.FillSpaceAt(9, "o")

    Expect(tttRules.IsGameOver()).To(Equal(true))
  })

  It("Check for valid mark", func() {
    Expect(tttRules.IsMarkValid("x")).To(Equal(true))
    Expect(tttRules.IsMarkValid("X")).To(Equal(true))
    Expect(tttRules.IsMarkValid("o")).To(Equal(true))
    Expect(tttRules.IsMarkValid("O")).To(Equal(true))
    Expect(tttRules.IsMarkValid("b")).To(Equal(false))
  })

  It("Get constant values", func() {
    Expect(tttRules.Mark("XMARK")).To(Equal("x"))
    Expect(tttRules.Mark("OMARK")).To(Equal("o"))
  })
})

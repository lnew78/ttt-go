package ttt_test

import (
  . "github.com/lnew78/ttt-go/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board", func() {
  var board = NewBoard(9)

    It("Spaces start at 1 and end at board size", func() {
      testArray :=  []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

      Expect(board.Spaces()).To(Equal(testArray))
    })

    It("Spaces can be filled", func() {
      testArray :=  []string{"x", "2", "3", "4", "5", "6", "7", "8", "9"}
      board.FillSpaceAt(1, "x")

      Expect(board.Spaces()).To(Equal(testArray))
    })

    It("Spaces can be reset to their default value", func() {
      testArray :=  []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
      board.FillSpaceAt(1, "x")
      board.ResetSpaceAt(1)

      Expect(board.Spaces()).To(Equal(testArray))
    })

    It("Board is full when all spaces are filled", func() {
      board.FillSpaceAt(1, "x")
      board.FillSpaceAt(2, "o")
      board.FillSpaceAt(3, "x")
      board.FillSpaceAt(4, "x")
      board.FillSpaceAt(5, "o")
      board.FillSpaceAt(6, "x")
      board.FillSpaceAt(7, "o")
      board.FillSpaceAt(8, "x")
      board.FillSpaceAt(9, "o")

      Expect(board.IsFull()).To(Equal(true))
    })

    It("Board not full when all spaces are not filled", func() {
      board.InitializeSpaces()
      board.FillSpaceAt(5, "o")

      Expect(board.IsFull()).To(Equal(false))
    })

    It("Get count of available spaces", func() {
      board.InitializeSpaces()
      Expect(board.NumberOfAvailableSpaces()).To(Equal(9))
      board.FillSpaceAt(1, "x")
      Expect(board.NumberOfAvailableSpaces()).To(Equal(8))
    })

    It("Get list of available spaces", func() {
      board.InitializeSpaces()
      testArray :=  []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
      Expect(board.AvailableSpaces()).To(Equal(testArray))

      board.InitializeSpaces()
      board.FillSpaceAt(1, "x")
      board.FillSpaceAt(2, "o")
      board.FillSpaceAt(3, "x")
      board.FillSpaceAt(4, "x")
      board.FillSpaceAt(5, "o")
      board.FillSpaceAt(6, "x")

      testArray = []string{"7", "8", "9"}
      Expect(board.AvailableSpaces()).To(Equal(testArray))
    })
})

package ttt_test

import (
  . "ttt-go/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board", func() {
  var board = NewBoard(9)

    It("Spacs start at 1 and end at board size", func() {
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

    It("Full when all spaces are filled", func() {
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

    It("Not full when all spaces are not filled", func() {
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

    It("Get list of availlable spaces", func() {
      testArray :=  []string{"2", "3", "4", "5", "6", "7", "8", "9"}
      board.InitializeSpaces()
      board.FillSpaceAt(1, "x")
      Expect(board.AvailableSpaces()).To(Equal(testArray))
    })
})

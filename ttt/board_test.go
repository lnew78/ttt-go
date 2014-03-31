package ttt_test

import (
  . "ttt-go/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board", func() {
  var board = NewBoard(9)

    It("Should have numbered spaces starting at 1 and ending at board size", func() {
      testArray :=  []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

      Expect( board.Spaces()).To(Equal(testArray))
    })

    It("Space can be filled", func() {
      testArray :=  []string{"x", "2", "3", "4", "5", "6", "7", "8", "9"}
      board.FillSpaceAt(1, "x")

      Expect( board.Spaces()).To(Equal(testArray))
    })
})

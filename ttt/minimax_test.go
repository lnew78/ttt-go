package ttt_test

import (
	. "github.com/lnew78/ttt-go/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Minimax", func() {
  var board = NewBoard(9)
  var tttRules = NewRules(board)
  var minimax = NewMinimax(board, tttRules)
  var bestScore = make(map[int]int)

  It("Take the win when possible", func() {
    board.InitializeSpaces()
    board.FillSpaceAt(1, "x")
    board.FillSpaceAt(2, "o")
    board.FillSpaceAt(3, "x")
    board.FillSpaceAt(4, "o")
    board.FillSpaceAt(5, "x")
    board.FillSpaceAt(6, "o")

    Expect(minimax.Move(0, bestScore)).To(Equal(7))
  })

  It("Block the opponent from winning", func() {
    board.InitializeSpaces()
    board.FillSpaceAt(1, "x")
    board.FillSpaceAt(5, "o")
    board.FillSpaceAt(2, "x")

    Expect(minimax.Move(0, bestScore)).To(Equal(3))
  })

  It("Defend the edges trap", func() {
    board.InitializeSpaces()
    board.FillSpaceAt(2, "x")
    board.FillSpaceAt(5, "o")
    board.FillSpaceAt(4, "x")

    Expect(minimax.Move(0, bestScore)).To(Equal(7))
  })

  It("Defend the corners trap", func() {
    board.InitializeSpaces()
    board.FillSpaceAt(1, "x")
    board.FillSpaceAt(5, "o")
    board.FillSpaceAt(7, "x")
    bestScore := make(map[int]int)

    Expect(minimax.Move(0, bestScore)).To(Equal(4))
  })

  It("Get the max key/value of a map by value", func() {
    testMap := make(map[int]int)
    testMap[1] = 2
    testMap[2] = 5

    testKey := 2
    testValue := 5
    key, value := minimax.MaxKeyValue(testMap)

    Expect(key).To(Equal(testKey))
    Expect(value).To(Equal(testValue))
  })
})

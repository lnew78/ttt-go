package ttt

import (
  "strconv"
)

type Board struct {
  numberOfSpaces int
  spaces []string
}

func NewBoard(size int) *Board {
  board := new ( Board )
  board.numberOfSpaces = size
  board.InitializeSpaces()
  return board
}

func (board *Board) InitializeSpaces() {
  board.spaces = make( []string, board.numberOfSpaces)
  for i := range  board.spaces {
    board.spaces[i] = strconv.Itoa(i + 1)
  }
}

func (board *Board) Spaces() []string {
  return board.spaces
}

func (board *Board) FillSpaceAt(space int, mark string) {
  if board.IsSpaceAvailableAt(space) {
    board.spaces[space - 1] = mark
  }
}

func (board *Board) ResetSpaceAt(space int) {
  board.spaces[space - 1] = strconv.Itoa(space)
}

func (board *Board) IsSpaceAvailableAt(space int) bool {
  return board.spaces[space - 1] == strconv.Itoa(space)
}

func (board *Board) IsFull() bool {
    return board.spaceAvailableCount() == 0
}

func (board *Board) NumberOfAvailableSpaces() int {
  return board.spaceAvailableCount()
}

func (board *Board) spaceAvailableCount() int {
  spaceAvailableCount := 0
  for i := 1; i <=  len(board.spaces); i++ {
    if board.IsSpaceAvailableAt(i) {
      spaceAvailableCount++
    }
  }
  return spaceAvailableCount
}

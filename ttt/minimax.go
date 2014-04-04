package ttt

import (
  "strconv"
  "fmt"
)

type Minimax struct {
  board *Board
  tttRules TicTacToeRules
}

func NewMinimax(board *Board, tttRules TicTacToeRules) *Minimax {
  minimax := new(Minimax)
  minimax.board = board
  minimax.tttRules = tttRules
  return minimax
}

func (minimax Minimax) Move(depth int, bestScore map[int]int) int {
  if minimax.tttRules.IsTie(){
    return 0
  }
  if minimax.tttRules.IsGameOver(){
    return -1
  }

  for _,space := range minimax.board.AvailableSpaces() {
    newspace,err := strconv.Atoi(space)
    if err != nil {
      fmt.Println(err)
    }

    minimax.board.FillSpaceAt(newspace, minimax.tttRules.NextPlayerMark())
    bestScore[newspace] = -1 * minimax.Move(depth + 1, make(map[int]int))
    minimax.board.ResetSpaceAt(newspace)
  }

  bestMove, highestScore := minimax.MaxKeyValue(bestScore)
  if depth > 0 {
    return highestScore
  } else {
    return bestMove
  }
}

func (minimax Minimax) MaxKeyValue(hashMap map[int]int) (keyToSave int, valueToSave int) {
  valueToSave = -1000
  for key, value := range hashMap {
    if value > valueToSave {
      keyToSave = key
      valueToSave = value
    }
  }
  return keyToSave, valueToSave
}

package ttt

import(
  "strconv"
  "strings"
)

type Game struct {
  ui UI
  playerGenerator PlayerGenerator
}

func NewGame(ui *UI, playerGenerator PlayerGenerator) *Game {
  game := new(Game)
  game.StartNewGame()
  game.playerCount()
  return game
}

func (game Game) StartNewGame() {
  game.ui.PrintMsg("Welcome to Tic Tac Toe!")
}

func (game Game) playerCount() int {
  game.ui.PrintMsg("How many players are playing? -- Type '1' or '2' and press [Enter]")
  input := strings.TrimSuffix(game.ui.GetInput(), "\n")

  numberOfPlayers, err := strconv.Atoi(input)
  if err != nil {
    return game.playerCount()
  }
  if numberOfPlayers == 1 || numberOfPlayers == 2 {
    return numberOfPlayers
  } else {
    return game.playerCount()
  }
}

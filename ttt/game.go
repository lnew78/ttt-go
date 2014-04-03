package ttt

import(
  "strconv"
  "strings"
  //"fmt"
)

type Game struct {
  ui UI
  playerGenerator PlayerGenerator
  tttRules TicTacToeRules
  player1 Player
  player2 Player
}

func NewGame(ui *UI, playerGenerator PlayerGenerator, tttRules TicTacToeRules) *Game {
  game := new(Game)
  game.StartNewGame()
  players := game.initPlayers(game.playerCount())
  game.player1 = players[0]
  game.player2 = players[1]
  player1Mark := game.getMarkChoice()
  game.setPlayerMarks(players, player1Mark)
  return game
}

func (game Game) StartNewGame() {
  game.ui.PrintMsg("Welcome to Tic Tac Toe!")
}

func (game Game) playerCount() int {
  game.ui.PrintMsg("How many players are playing? -- Type '1' or '2' and press [Enter].")
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

func (game Game) initPlayers(playerCount int) []Player {
  return game.playerGenerator.GeneratePlayers(playerCount)
}

func (game Game) getMarkChoice() string {
  game.ui.PrintMsg("Player 1, choose your mark. -- Type 'x' or 'o' and press [Enter].")
  input := strings.TrimSuffix(game.ui.GetInput(), "\n")

  if game.tttRules.IsMarkValid(input) {
    return input
  } else {
    return game.getMarkChoice()
  }
}

func (game Game) setPlayerMarks(players []Player, firstPlayerMark string) {
  game.player1.SetMark(firstPlayerMark)
  if firstPlayerMark == "x" {
    game.player2.SetMark("o")
  } else {
    game.player2.SetMark("x")
  }
}

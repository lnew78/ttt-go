package ttt

import(
  "strconv"
  "strings"
)

type Game struct {
  ui UI
  playerGenerator PlayerGenerator
  tttRules TicTacToeRules
  board *Board
  player1 Player
  player2 Player
}

func NewGame(ui *UI, playerGenerator PlayerGenerator, tttRules TicTacToeRules, board *Board) *Game {
  game := new(Game)
  game.board = board
  game.ui.PrintMsg("Welcome to Tic Tac Toe!")
  players := game.initPlayers(game.playerCount())
  game.player1 = players[0]
  game.player2 = players[1]
  player1Mark := game.getMarkChoice()
  game.setPlayerMarks(players, player1Mark)
  game.tttRules = tttRules
  game.ui.PrintBoard(game.tttRules)
  game.RunGameSequence()
  return game
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

func (game *Game) setPlayerMarks(players []Player, firstPlayerMark string) {
  game.player1.SetMark(firstPlayerMark)
  if firstPlayerMark == "x" {
    game.player2.SetMark("o")
    game.ui.PrintMsg("Great! Player 2, you will be 'o'!")
  } else {
    game.player2.SetMark("x")
    game.ui.PrintMsg("Great! Player 2, you will be 'x'!")
  }
}

func (game Game) RunGameSequence() {
  if game.tttRules.IsGameOver() {
    game.endGame()
  } else {
    game.makeMove(game.firstPlayerToMakeMove())
    game.makeMove(game.secondPlayerToMakeMove())
    game.RunGameSequence()
  }
}

func (game Game) makeMove(player Player) {
  game.getPlayerMove(player)
  if game.tttRules.IsGameOver() {
    game.endGame()
  }
}

func (game Game) getPlayerMove(player Player) {
  game.ui.PrintMsg("Make your move..........")
  input := strings.TrimSuffix(game.ui.GetInput(), "\n")

  move, err := strconv.Atoi(input)
  if err != nil {
    game.getPlayerMove(player)
  }
  game.board.FillSpaceAt(move, player.Mark())
  game.ui.PrintBoard(game.tttRules)
}

func (game Game) firstPlayerToMakeMove() Player {
  if game.player1.Mark() == "x" {
    return game.player1
  } else {
    return game.player2
  }
}

func (game Game) secondPlayerToMakeMove() Player {
  if game.player1.Mark() == "o" {
    return game.player1
  } else {
    return game.player2
  }
}

func (game Game) endGame() {
  game.ui.PrintMsg("game blah blah")
}

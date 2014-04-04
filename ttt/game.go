package ttt

import(
  "strconv"
  "strings"
  "os"
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
  players := game.generatePlayers(game.playerCount())
  game.assignPlayers(players)
  player1Mark := game.getMarkChoice()
  game.setPlayerMarks(players, player1Mark)
  game.tttRules = tttRules
  game.ui.PrintBoard(game.tttRules)
  game.RunGameSequence()
  return game
}

func (game Game) playerCount() int {
  numberOfPlayers := game.getNumericInput("How many players are playing? -- Type '1' or '2' and press [Enter].")
  if numberOfPlayers == 1 || numberOfPlayers == 2 {
    return numberOfPlayers
  }
  return game.playerCount()
}

func (game Game) generatePlayers(playerCount int) []Player {
  return game.playerGenerator.GeneratePlayers(playerCount)
}

func (game Game) assignPlayers(humanPlayers []Player) {
  if len(humanPlayers) == 1 {
    game.player1 = humanPlayers[0]
  } else if len(humanPlayers) == 2 {
    game.player1 = humanPlayers[0]
    game.player2 = humanPlayers[1]
  }
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
  if firstPlayerMark == game.tttRules.Mark("XMARK") {
    game.player2.SetMark(game.tttRules.Mark("OMARK"))
    game.ui.PrintMsg("Great! Player 2, you will be 'o'!")
  } else {
    game.player2.SetMark(game.tttRules.Mark("OMARK"))
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
  move := game.getPlayerMove(player)
  game.board.FillSpaceAt(move, player.Mark())
  game.ui.PrintBoard(game.tttRules)
  if game.tttRules.IsGameOver() {
    game.endGame()
  }
}

func (game Game) getPlayerMove(player Player) int {
  move := game.getNumericInput("Make your move.....")
  if game.board.IsSpaceAvailableAt(move) {
    return move
  }
  return game.getPlayerMove(player)
}

func (game Game) getNumericInput(message string) int {
  game.ui.PrintMsg(message)
  input := strings.TrimSuffix(game.ui.GetInput(), "\n")

  number, err := strconv.Atoi(input)
  if err != nil {
    return game.getNumericInput(message)
  }
  return number
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
  game.ui.PrintMsg("Game Over. Blah blah blah")
  os.Exit(0)
}

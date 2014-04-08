package ttt

import(
  "strings"
  "os"
)

type Game struct {
  board *Board
  tttRules TTTRules
  playerGenerator PlayerGenerator
  ui UI
  minimax *Minimax
  player1 Player
  player2 Player
}

func NewGame(board *Board, tttRules TTTRules, playerGenerator PlayerGenerator, ui *UI, minimax *Minimax) *Game {
  game := new(Game)
  game.board = board
  game.minimax = minimax
  game.ui.PrintMsg("Welcome to Tic Tac Toe!")
  players := game.generatePlayers(game.playerCount())
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
  numberOfPlayers := game.ui.GetNumericInput("How many players are playing? -- Type '1' or '2' and press [Enter].")
  if numberOfPlayers == 1 || numberOfPlayers == 2 {
    return numberOfPlayers
  }
  return game.playerCount()
}

func (game Game) generatePlayers(playerCount int) []Player {
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
  if firstPlayerMark == game.tttRules.Mark("XMARK") {
    game.player2.SetMark(game.tttRules.Mark("OMARK"))
    game.ui.PrintMsgWithData("Great! Player 2, you will be ", game.tttRules.Mark("OMARK"))
  } else {
    game.player2.SetMark(game.tttRules.Mark("XMARK"))
    game.ui.PrintMsgWithData("Great! Player 2, you will be ", game.tttRules.Mark("XMARK"))
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
  if player.Type == "computer" {
    game.ui.PrintMsg("I'm thinking of a move.  One moment please...")
    return game.minimax.Move(0, make(map[int]int))
  } else {
    game.ui.PrintMsgWithData(player.Mark(), "it's your turn")
    move := game.ui.GetNumericInput("Make your move...")
    if game.board.IsSpaceAvailableAt(move) {
      return move
    }
    return game.getPlayerMove(player)
  }
}

func (game Game) firstPlayerToMakeMove() Player {
  if game.player1.Mark() == game.tttRules.Mark("XMARK") {
    return game.player1
  } else {
    return game.player2
  }
}

func (game Game) secondPlayerToMakeMove() Player {
  if game.player1.Mark() == game.tttRules.Mark("OMARK") {
    return game.player1
  } else {
    return game.player2
  }
}

func (game Game) endGame() {
  if game.tttRules.IsTie() {
    game.ui.PrintMsg("It's a tie! You got lucky.")
  } else {
    game.ui.PrintMsgWithData(game.tttRules.Winner(), "wins!")
  }
  os.Exit(0)
}

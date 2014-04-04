package ttt

type PlayerGenerator struct {
  player Player
  players []Player
}

func (pg PlayerGenerator) GeneratePlayers(numberOfHumanPlayers int) []Player {
  if numberOfHumanPlayers == 1 {
    player1 := pg.createPlayer("human")
    player2 := pg.createPlayer("computer")
    pg.players = []Player{player1, player2}
  } else if numberOfHumanPlayers == 2 {
    player1 := pg.createPlayer("human")
    player2 := pg.createPlayer("human")
    pg.players = []Player{player1, player2}
  }
  return pg.players
}

func (pg PlayerGenerator) createPlayer(playerType string)Player {
  player := Player{Type:playerType}
  return player
}

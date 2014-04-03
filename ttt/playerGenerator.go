package ttt

type PlayerGenerator struct {
  player Player
  players []Player
}

func (pg PlayerGenerator) GeneratePlayers(numberOfHumanPlayers int) []Player {
  if numberOfHumanPlayers == 2 {
    player1 := pg.createPlayer()
    player2 := pg.createPlayer()
    pg.players = []Player{player1, player2}
  }
  return pg.players
}

func (pg PlayerGenerator) createPlayer() Player {
  player := Player{}
  return player
}

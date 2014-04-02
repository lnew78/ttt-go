package ttt

type PlayerGenerator struct {
  numberOfHumanPlayers int
  player *Player
  players []*Player
}

func NewPlayerGenerator(humanPlayerCount int) *PlayerGenerator {
  pg := new(PlayerGenerator)
  pg.numberOfHumanPlayers = humanPlayerCount
  pg.generatePlayers()
  return pg
}

func (pg *PlayerGenerator) generatePlayers() {
  if pg.numberOfHumanPlayers == 2 {
    player1 := new(Player)
    player2 := pg.createPlayer()
    pg.players = []*Player{player1, player2}
  }
}

func (pg PlayerGenerator) createPlayer() *Player {
  player := new(Player)
  return player
}

func (pg *PlayerGenerator) Players() []*Player {
  return pg.players
}

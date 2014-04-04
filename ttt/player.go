package ttt

type Player struct {
  mark string
  move int
  Type string
}

func (player *Player) SetMark(mark string) {
  player.mark = mark
}

func (player Player) Mark() string {
  return player.mark
}

func (player *Player) SetMove(move int) {
  player.move = move
}

func (player Player) Move() int {
  return player.move
}

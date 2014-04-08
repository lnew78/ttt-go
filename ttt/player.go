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

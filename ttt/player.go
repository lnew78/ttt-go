package ttt

type Player struct {}

func (player Player) Mark(mark string) string {
  return mark
}

func (player Player) Move(move int) int {
  return move
}

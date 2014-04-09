package ttt

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
  var pg = PlayerGenerator{}
  var players = pg.GeneratePlayers(2)
  var player1 = players[0]
  var player2 = players[1]
  var game = Game{player1:player1, player2:player2}


  It("Set the player's marks", func() {

    game.setPlayerMarks(players, "x")
    Expect(game.player1.Mark()).To(Equal("x"))
    Expect(game.player2.Mark()).To(Equal("o"))

    game.setPlayerMarks(players, "o")
    Expect(game.player1.Mark()).To(Equal("o"))
    Expect(game.player2.Mark()).To(Equal("x"))
  })

  It("X should always go first", func() {
    game.setPlayerMarks(players, "x")
    Expect(game.firstPlayerToMakeMove()).To(Equal(game.player1))
    Expect(game.secondPlayerToMakeMove()).To(Equal(game.player2))

    game.setPlayerMarks(players, "o")
    Expect(game.firstPlayerToMakeMove()).To(Equal(game.player2))
    Expect(game.secondPlayerToMakeMove()).To(Equal(game.player1))
  })
})

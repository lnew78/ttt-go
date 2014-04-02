package ttt_test

import (
	. "ttt-go/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PlayerGenerator", func() {
  var pg = NewPlayerGenerator(2)
  var  testPlayer1 = new(Player)
  var  testPlayer2 = new(Player)

  It("Generate 2 human players", func() {
    testPlayerArray := []*Player{testPlayer1, testPlayer2}
    Expect(pg.Players()).To(Equal(testPlayerArray))
  })

})

package ttt_test

import (
	. "ttt-go/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PlayerGenerator", func() {
  var pg = PlayerGenerator{}
  var  testPlayer1 = Player{}
  var  testPlayer2 = Player{}

  It("Generate 2 human players", func() {
    testPlayerArray := []Player{testPlayer1, testPlayer2}
    Expect(pg.GeneratePlayers(2)).To(Equal(testPlayerArray))
  })

})

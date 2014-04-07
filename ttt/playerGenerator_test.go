package ttt_test

import (
	. "github.com/lnew78/ttt-go/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PlayerGenerator", func() {
  var pg = PlayerGenerator{}
  var  testPlayer1 = Player{Type:"human"}
  var  testPlayer2 = Player{Type:"human"}
  var  testCompPlayer = Player{Type:"computer"}

  It("Generate 1 human player", func() {
    testPlayerArray := []Player{testPlayer1, testCompPlayer}
    Expect(pg.GeneratePlayers(1)).To(Equal(testPlayerArray))
  })

  It("Generate 2 human players", func() {
    testPlayerArray := []Player{testPlayer1, testPlayer2}
    Expect(pg.GeneratePlayers(2)).To(Equal(testPlayerArray))
  })

})

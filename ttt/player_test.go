package ttt_test

import (
	. "github.com/lnew78/ttt-go/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Player", func() {
  var player = Player{Type:"human"}

  It("Player's mark can be set and retrieved", func() {
    player.SetMark("x")
    Expect(player.Mark()).To(Equal("x"))
  })

  It("Player's type can be retrieved", func() {
    Expect(player.Type).To(Equal("human"))
    player2 := Player{Type:"computer"}
    Expect(player2.Type).To(Equal("computer"))
  })
})

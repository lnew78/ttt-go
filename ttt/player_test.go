package ttt_test

import (
	. "ttt-go/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Player", func() {
  var player = Player{}

  It("Return the player's mark", func() {
    Expect(player.Mark("x")).To(Equal("x"))
  })

  It("Return the player's move", func() {
    Expect(player.Move(1)).To(Equal(1))
  })
})

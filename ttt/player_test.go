package ttt_test

import (
	. "ttt-go/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Player", func() {
  var player = Player{}

  It("Player's mark can be set and retrieved", func() {
    player.SetMark("x")
    Expect(player.Mark()).To(Equal("x"))
  })

  It("Player's move can be set and retrieved", func() {
    player.SetMove(1)
    Expect(player.Move()).To(Equal(1))
  })
})

package gate

import (
	"leaf-tutorial/game"
	"leaf-tutorial/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
}

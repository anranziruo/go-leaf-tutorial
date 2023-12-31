package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"leaf-tutorial/msg"
	"reflect"
)

func init() {
	handler(&msg.Hello{}, handleHello)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)

	log.Debug("hello %v", m.Name)

	a.WriteMsg(&msg.Hello{
		Name: "client",
	})
}

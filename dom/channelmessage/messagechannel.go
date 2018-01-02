package channelmessage

import (
	"github.com/matthewmueller/joy/macro"
)

func New() *MessageChannel {
	macro.Rewrite("new MessageChannel()")
	return &MessageChannel{}
}

type MessageChannel struct {
}

func (*MessageChannel) Port1() (port1 *MessagePort) {
	macro.Rewrite("$_.port1")
	return port1
}

func (*MessageChannel) Port2() (port2 *MessagePort) {
	macro.Rewrite("$_.port2")
	return port2
}

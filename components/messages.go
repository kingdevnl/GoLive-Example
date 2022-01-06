package components

import (
	"ChatApp/message"
	"github.com/kingdevnl/GoLive"
	"log"
)

type Messages struct {
	GoLive.Component
}

func (c *Messages) OnMount(state *GoLive.State, args []interface{}) {
	state.Set("messages", &message.Messages)

	c.On(state, "messageAdded", func(state *GoLive.State, data GoLive.Map) {
		log.Println("messageAdded")
		c.ReRender(state)
	})
}

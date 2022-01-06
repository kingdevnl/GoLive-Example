package components

import (
	"ChatApp/message"
	"github.com/kingdevnl/GoLive"
)

type Inputs struct {
	GoLive.Component
}

func (c *Inputs) OnMount(state *GoLive.State, args []interface{}) {
	state.Set("message", "")
	c.Register("AddMessage", c.AddMessage)
}

func (c *Inputs) AddMessage(state *GoLive.State, data GoLive.Map) {

	ws := GoLive.GetWebSocketByState(state)

	if ws != nil {
		// get the session using the id
		message.Messages = append(message.Messages, message.NewMessage(ws.Conn.Locals("user").(string), state.Get("message").(string)))

		GoLive.EmitEvent("messageAdded", nil)

		state.Set("message", "")
		c.ReRender(state)
	}

}

package alef

import (
	"github.com/DrmagicE/gmqtt/server"
)

func (a *Alef) HookWrapper() server.HookWrapper {
	return server.HookWrapper{
		OnMsgArrivedWrapper: a.OnMsgArrivedWrapper,
	}
}

func (a *Alef) OnMsgArrivedWrapper(pre server.OnMsgArrived) server.OnMsgArrived {
	// panic("impermanent me")
	return nil
}

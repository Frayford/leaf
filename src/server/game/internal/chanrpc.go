package internal

import (
	entity "server/entity"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

var agents = make(map[gate.Agent]struct{})

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	log.Debug("w")
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	log.Debug("Player Online")
	agents[a] = struct{}{}
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	if a.UserData() != nil {
		player := a.UserData().(*entity.PlayerInfo)
		//player.online = false
		player.IsOnLine = 0
		//关闭player
		if player.Agent != nil {
			player.Agent.Close()
			player.Agent = nil
		}
		log.Debug("Player Offline " + player.Name)
		//并没有去除的操作，看之后会不会有影响

	}

	delete(agents, a)
	_ = a
}

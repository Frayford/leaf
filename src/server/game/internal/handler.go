package internal

import (
	"reflect"
	entity "server/entity"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	// 向当前模块（game 模块）注册 Login 消息的消息处理函数 handleLogin
	handler(&msg.LoginMsg{}, handleLogin)
	handler(&msg.ChatMsg{}, handlerChatMsg)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLogin(args []interface{}) {
	// 收到的 Login 消息
	m := args[0].(*msg.LoginMsg)
	// 消息的发送者
	a := args[1].(gate.Agent)

	player := loginPlayer(*m)

	if player.Name != "" {
		player.Agent = a

		a.SetUserData(player)
	} else {
		a.Close()
	}
	// 输出收到的消息的内容
	log.Debug("hello %v", m.GetName())
	/*for a := range agents {
		// 给发送者回应一个 Hello 消息
		a.WriteMsg(&msg.LoginMsg{
			Name: proto.String("client" + m.GetName()),
		})
	}
	*/
}

func handlerChatMsg(args []interface{}) {
	//收到的 ChatMsg 消息
	m := args[0].(*msg.ChatMsg)
	//消息的发送者
	//a := args[1].(gate.Agent)

	Room := RoomMap.Get(m.GetKey()).(*entity.RoomInfo)

	for x := 0; x < len(Room.Players); x++ {
		// 给发送者回应一个 Hello 消息
		log.Debug(m.GetName())
		if Room.Players[x].IsOnLine == 1 {
			Room.Players[x].Agent.WriteMsg(&msg.ChatMsg{
				Name: m.Name, Data: m.Data,
			})
		}
	}

}

package internal

import (
	"server/base"
	entity "server/entity"
	msg "server/msg"

	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"github.com/name5566/leaf/util"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

//房间
var RoomMap util.Map

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}

func loginPlayer(m msg.LoginMsg) *entity.PlayerInfo {
	if RoomMap.Get(m.GetKey()) == nil {
		//新房间
		addRoom(m)
		var player entity.PlayerInfo
		player.Name = m.GetName()
		player.Key = m.GetKey()
		//player.online = true
		player.IsOnLine = 1
		//找到那个房间并把那个玩家存进去
		Room := RoomMap.Get(m.GetKey()).(*entity.RoomInfo)

		Room.Players = append(Room.Players, &player)

		return &player
	} else {
		//房间已经存在
		var player entity.PlayerInfo
		player.Name = m.GetName()
		player.Key = m.GetKey()
		player.IsOnLine = 1
		//player. = true

		//找到那个房间并把那个玩家存进去
		Room := RoomMap.Get(m.GetKey()).(*entity.RoomInfo)

		Room.Players = append(Room.Players, &player)

		return &player
	}
}

//建立一个房间
func addRoom(m msg.LoginMsg) {
	//建立一个新房间
	if RoomMap.Get(m.GetKey()) == nil {
		var Room entity.RoomInfo
		Room.Key = m.GetKey()
		Room.OnlineCount = 0
		Room.Players = make([]*entity.PlayerInfo, 0, 5)
		log.Debug("New Room-> %v ", m.GetKey())
		RoomMap.Set(Room.Key, &Room)
	} else {
		log.Debug("Old Room-> %v ", m.GetKey())
	}
}

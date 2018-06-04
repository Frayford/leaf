package entity

import (
	"sync"

	"github.com/name5566/leaf/gate"
)

type RoomInfo struct {
	//读锁，适用于读多写少的场景
	//1、可以随便读。多个goroutin同时读。
	//2、写的时候，啥都不能干。不能读，也不能写。
	sync.RWMutex
	Key         string
	OnlineCount int
	Players     []*PlayerInfo
}

type PlayerInfo struct {
	sync.RWMutex
	Name     string
	Key      string
	IsOnLine int

	Agent gate.Agent
}

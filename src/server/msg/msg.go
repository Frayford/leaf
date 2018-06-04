package msg

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()

func init() {
	var id = Processor.Register(&LoginMsg{})
	log.Debug("LoginMsg (id= %v)", id)
	id = Processor.Register(&ChatMsg{})
	log.Debug("ChatMsg (id= %v)", id)
	id = Processor.Register(&RoomData{})
	log.Debug("RoomData (id= %v)", id)
}

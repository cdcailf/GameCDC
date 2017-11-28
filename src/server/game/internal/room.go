package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
)

type Room struct {
	roomID      uint
	roomPlayers map[uint]gate.Agent
}

func newRoom(roomID uint) *Room {
	fmt.Println("newRoomInit-->", roomID)
	room := Room{roomID: roomID}

	return &room
}

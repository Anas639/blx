package udp

import (
	"encoding/binary"
	"net"

	"github.com/anas639/blx/internal/event"
)

type udpBroadcaster struct{}

func NewUDPBroadcaster() event.EventBroadcaster {
	return &udpBroadcaster{}
}

func (this *udpBroadcaster) SendEvent(payload event.EventPayload) error {
	addr, err := net.ResolveUDPAddr("udp", ":6969")
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	buffer := []byte{}
	buffer = append(buffer, payload.Type)
	buffer = binary.BigEndian.AppendUint64(buffer, uint64(payload.TaskId))

	_, err = conn.Write(buffer)
	if err != nil {
		return err
	}

	return nil
}

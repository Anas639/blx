package udp

import (
	"encoding/binary"
	"net"

	"github.com/anas639/blx/internal/event"
)

type udpListener struct {
	payload chan event.EventPayload
}

func NewUDPListener() event.EventListener {
	return &udpListener{
		payload: make(chan event.EventPayload),
	}
}

func (this *udpListener) Listen() (chan event.EventPayload, error) {
	addr, err := net.ResolveUDPAddr("udp", ":6969")
	if err != nil {
		return nil, err
	}
	go this.start(addr)
	return this.payload, nil
}

func (this *udpListener) Close() error {
	return nil
}

func (this *udpListener) start(addr *net.UDPAddr) {
	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		this.payload <- event.NewPayloadError(err)
	}

	for {
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			this.payload <- event.NewPayloadError(err)
		}

		if n == 9 {
			eventType := buffer[0]
			taskId := binary.BigEndian.Uint64(buffer[1:9])
			this.payload <- event.NewPayload(eventType, int64(taskId))
		}
	}
}

package udp

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"

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
	go this.start()
	return this.payload, nil
}

func (this *udpListener) Close() error {
	return nil
}

func (this *udpListener) start() {
	var conn *net.UDPConn
	for {
		c, err := this.connect(6969)
		if err != nil {
			this.payload <- event.NewPayloadError(err)
		} else {
			conn = c
			break
		}
		time.Sleep(1 * time.Second)
	}
	defer conn.Close()
	for {
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)

		if err != nil {
			this.payload <- event.NewPayloadError(err)
			continue
		}

		if n == 9 {
			eventType := buffer[0]
			taskId := binary.BigEndian.Uint64(buffer[1:9])
			this.payload <- event.NewPayload(eventType, int64(taskId))
		}

	}
}

func (this *udpListener) connect(port int) (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

package protocol

import (
	"github.com/bitini111/mvta/cluster"
)

type TriggerRequest struct {
	Event cluster.Event
	GID   string
	UID   int64
}

type TriggerReply struct {
	Code int
}

type DeliverRequest struct {
	GID     string
	NID     string
	CID     int64
	UID     int64
	Message *Message
}

type DeliverReply struct {
	Code int
}

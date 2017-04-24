package service

import (
	"time"
	"container/list"
)

type eventType int

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
)

var archiveSize = 20
// Event archives.
var archive = list.New()

type wsEvent struct {
	Type      eventType // JOIN, LEAVE, MESSAGE
	User     SocketUser
	Timestamp int // Unix timestamp (secs)
	Content   string
}

func NewWsEvent(ep eventType,user SocketUser,msg string) (wsEvent) {
	return wsEvent{Type:ep,User:user,Timestamp:int(time.Now().Unix()),Content:msg}
}



// NewArchive saves new event to archive list.
func NewArchive(event wsEvent) {
	if archive.Len() >= archiveSize {
		archive.Remove(archive.Front())
	}
	archive.PushBack(event)
}
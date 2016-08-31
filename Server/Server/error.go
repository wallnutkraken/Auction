package Server

import (
	"encoding/json"
)

type Error struct {
	What string
}

func (e *Error) JSON() ([]byte, error) {
	return json.Marshal(*e)
}

func ErrorResponse(what string) Reply {
	reply := Reply{}
	err := Error{What: what}
	reply.ReplyType = "error"
	valueJson, _ := err.JSON()
	reply.ValueJson = string(valueJson)
	return reply
}

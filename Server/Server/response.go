package Server

import (
	"encoding/json"
)

type Reply struct {
	ReplyType string
	ValueJson string
}

func (r *Reply) JSON() ([]byte, error) {
	return json.Marshal(*r)
}

package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wallnutkraken/Auction/Server/Auction"
	"github.com/wallnutkraken/Auction/Server/Server"
)

func TestReplyJSON(t *testing.T) {
	auct := Auction.NewAuction()
	auct.NewPimp(GeneratePimp())

	jsonBytes, err := auct.ActivePimpsJSON()
	assert.NoError(t, err)
	t.Log("Generated JSON for pimps:", string(jsonBytes))

	reply := Server.Reply{ReplyType: "list", ValueJson: string(jsonBytes)}
	listReply, err := json.Marshal(reply)
	assert.NoError(t, err)
	t.Log("Generated JSON for reply:", string(listReply))

	readReply := Server.Reply{}
	err = json.Unmarshal(listReply, readReply)
	assert.NoError(t, err)
	assert.Equal(t, reply, readReply)
}

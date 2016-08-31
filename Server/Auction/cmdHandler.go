package Auction

import (
	"github.com/wallnutkraken/Auction/Server/Server"
)

func bid(command Server.Cmd, client Server.Client, auct Auction) error {
	var pimpId int
	var bidValue int
	var err error
	if pimpId, err = command.ArgInt(0); err != nil {
		return client.Send(Server.ErrorResponse(err.Error()))
	}
	if bidValue, err = command.ArgInt(1); err != nil {
		return client.Send(Server.ErrorResponse(err.Error()))
	}
	err = auct.Bid(client, pimpId, bidValue)
}

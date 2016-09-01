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
	return auct.Bid(client, pimpId, bidValue)
}

func list(command Server.Cmd, client Server.Client, auct Auction) error {
	/* Command currently unused */
	pimps, err := auct.ActivePimpsJSON()
	if err != nil {
		return err
	}
	response := Server.Reply{ReplyType: command.Command, ValueJson: string(pimps)}
	return client.Send(response)
}

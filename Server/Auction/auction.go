package Auction

import (
	"errors"

	"github.com/wallnutkraken/Auction/Server/Server"
)

const (
	DurationOfPimp = 500
)

type auction struct {
	Pimps map[int]Pimp
	Users []Server.Client
}

func (a *auction) Bid(cl Server.Client, pimpId int, bid int) error {
	bidPimp := a.Pimps[pimpId]
	if bidPimp == nil {
		return a.ErrorNoPimp()
	}
	if bidPimp.GetTimeLeft() == 0 {
		return a.ErrorOutOfTime()
	}
	if bid > bidPimp.GetCurrentBid() {
		bidPimp.SetCurrentBid(bid)
		bidPimp.SetTopBidder(cl)
		bidPimp.AddBidder(cl)
		return nil
	}
	return a.ErrorBidTooLow()
}

func (a *auction) ErrorNoPimp() error {
	return errors.New("No such pimp")
}

func (a *auction) ErrorPimpExists() error {
	return errors.New("Pimp already exists")
}

func (a *auction) ErrorOutOfTime() error {
	return errors.New("Out of time")
}

func (a *auction) ErrorBidTooLow() error {
	return errors.New("Bid too low")
}

func (a *auction) AddUser(client Server.Client) {
	a.Users = append(a.Users, client)
}

func (a *auction) DeletePimp(pimpId int) {
	a.Pimps[pimpId] = nil
}

func (a *auction) FindPimp(pimpId int) (Pimp, error) {
	foundPimp := a.Pimps[pimpId]
	if foundPimp == nil {
		return nil, a.ErrorNoPimp()
	}
	return foundPimp, nil
}

func (a *auction) NewPimp(newPimp Pimp) error {
	if a.Pimps[newPimp.GetId()] == nil {
		a.Pimps[newPimp.GetId()] = newPimp
		return nil
	}
	return a.ErrorPimpExists()
}

func (a *auction) RemoveUser(cl Server.Client) {
	for x := 0; x < len(a.Users); x++ {
		if a.Users[x].GetConnection() == cl.GetConnection() {
			//a.Users = append(a.Users[0:x], a.Users[x:])
			panic("aaa")
		}
	}
}

func (a *auction) ExecCommand(command Server.Cmd, client Server.Client) {
	switch command.Command {
	case "bid":
		bid(command, client, a)
	}
}

type Auction interface {
	FindPimp(int) (Pimp, error)
	NewPimp(Pimp) error
	DeletePimp(int)
	AddUser(Server.Client)
	RemoveUser(Server.Client)
	Bid(Server.Client, int, int) error
	ExecCommand(Server.Cmd, Server.Client)
}

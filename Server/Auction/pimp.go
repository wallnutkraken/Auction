package Auction

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/wallnutkraken/Auction/Server/Server"
)

type pimp struct {
	Id          int
	StartPrice  int
	ExpDate     int64
	BidItem     Item
	CurrentBid  int
	prevBidders []Server.Client
	topBidder   Server.Client
}

func (p *pimp) Serialize() ([]byte, error) {
	return json.Marshal(*p)
}

func NewPimp(startingPrice int, item Item) Pimp {
	p := new(pimp)
	p.StartPrice = startingPrice
	p.CurrentBid = startingPrice
	p.BidItem = item
	p.ExpDate = time.Now().UTC().Unix() + DurationOfPimp
	p.prevBidders = make([]Server.Client, 0)
	return p
}

func (p *pimp) GetStartingPrice() int {
	return p.StartPrice
}

func (p *pimp) GetTimeLeft() int64 {
	if p.ExpDate <= time.Now().UTC().Unix() {
		return 0
	}
	return p.ExpDate - time.Now().UTC().Unix()
}

func (p *pimp) GetItem() Item {
	return p.BidItem
}

func (p *pimp) GetCurrentBid() int {
	return p.CurrentBid
}

func (p *pimp) SetCurrentBid(bid int) {
	p.CurrentBid = bid
}

func (p *pimp) GetPrevBidders() []Server.Client {
	return p.prevBidders
}

func (p *pimp) hasBidder(bidder Server.Client) bool {
	for _, currBidder := range p.prevBidders {
		if currBidder == bidder {
			return true
		}
	}
	return false
}

func (p *pimp) AddBidder(bidder Server.Client) error {
	if p.hasBidder(bidder) {
		return errors.New("Bidder already exists")
	}
	p.prevBidders = append(p.prevBidders, bidder)
	return nil
}

func (p *pimp) GetTopBidder() Server.Client {
	return p.topBidder
}

func (p *pimp) SetTopBidder(newBidder Server.Client) {
	p.topBidder = newBidder
}

func (p *pimp) GetId() int {
	return p.Id
}

type Pimp interface {
	GetStartingPrice() int
	GetTimeLeft() int64
	GetItem() Item
	GetCurrentBid() int
	SetCurrentBid(int)
	GetPrevBidders() []Server.Client
	AddBidder(Server.Client) error
	GetTopBidder() Server.Client
	SetTopBidder(Server.Client)
	GetId() int
	Serialize() ([]byte, error)
}

func PimpFromJSON(content []byte) (Pimp, error) {
	p := new(pimp)
	p.BidItem = NewItem("default")
	err := json.Unmarshal(content, p)
	return p, err
}

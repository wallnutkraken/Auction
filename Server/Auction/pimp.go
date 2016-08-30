package Auction

import (
	"encoding/json"
	"errors"
	"time"
)

type pimp struct {
	Price       int
	ExpDate     int64
	BidItem     Item
	CurrentBid  int
	PrevBidders []string
}

func (p *pimp) Serialize() ([]byte, error) {
	return json.Marshal(*p)
}

func NewPimp(startingPrice int, item Item) Pimp {
	p := new(pimp)
	p.Price = startingPrice
	p.CurrentBid = startingPrice
	p.BidItem = item
	p.ExpDate = time.Now().UTC().Unix() + DurationOfPimp
	p.PrevBidders = make([]string, 0)
	return p
}

func (p *pimp) GetStartingPrice() int {
	return p.Price
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

func (p *pimp) GetPrevBidders() []string {
	return p.PrevBidders
}

func (p *pimp) hasBidder(bidder string) bool {
	for _, currBidder := range p.PrevBidders {
		if currBidder == bidder {
			return true
		}
	}
	return false
}

func (p *pimp) AddBidder(bidder string) error {
	if p.hasBidder(bidder) {
		return errors.New("Bidder already exists")
	}
	p.PrevBidders = append(p.PrevBidders, bidder)
	return nil
}

type Pimp interface {
	GetStartingPrice() int
	GetTimeLeft() int64
	GetItem() Item
	GetCurrentBid() int
	SetCurrentBid(int)
	GetPrevBidders() []string
	AddBidder(string) error
	Serialize() ([]byte, error)
}

func PimpFromJSON(content []byte) (Pimp, error) {
	p := new(pimp)
	err := json.Unmarshal(content, p)
	return p, err
}

package Auction

type item struct {
	Name string
}

func (i *item) GetName() string {
	return i.Name
}

func NewItem(name string) Item {
	i := new(item)
	i.Name = name
	return i
}

type Item interface {
	GetName() string
}

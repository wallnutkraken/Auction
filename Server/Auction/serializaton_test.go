package Auction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializingPimp(t *testing.T) {
	p := NewPimp(100, NewItem("ExampleItem"))
	t.Log("Created pimp:", p)
	bytes, err := p.Serialize()
	assert.NoError(t, err)
	t.Log("Generated JSON from pimp:", string(bytes))

	readPimp, err := PimpFromJSON(bytes)
	assert.NoError(t, err)
	t.Log("Deserialized pimp:", readPimp)
	assert.Equal(t, p, readPimp)
}

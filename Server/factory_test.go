package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomWord(t *testing.T) {
	assert.NotEqual(t, getWord(), "")
}

func TestRandomPimp(t *testing.T) {
	pimp1 := GeneratePimp()
	pimp2 := GeneratePimp()
	pimp3 := GeneratePimp()

	/* No names are the same */
	assert.NotEqual(t, pimp1.GetItem().GetName(), pimp2.GetItem().GetName())
	assert.NotEqual(t, pimp2.GetItem().GetName(), pimp3.GetItem().GetName())
	assert.NotEqual(t, pimp1.GetItem().GetName(), pimp3.GetItem().GetName())

	/* No ids are the same */
	assert.NotEqual(t, pimp1.GetId(), pimp2.GetId())
	assert.NotEqual(t, pimp2.GetId(), pimp3.GetId())
	assert.NotEqual(t, pimp1.GetId(), pimp3.GetId())
}

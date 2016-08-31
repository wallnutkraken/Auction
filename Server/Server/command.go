package Server

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Cmd struct {
	Command string
	Args    []string
}

func (c *Cmd) ArgInt(argIndex int) (int, error) {
	if c.Args == nil || argIndex >= len(c.Args) {
		return 0, c.NoArgError()
	}
	return strconv.Atoi(c.Args[argIndex])
}

func (c *Cmd) NoArgError() error {
	return errors.New("Argument index out of range")
}

func CmdFromJSON(jsonContent []byte) (Cmd, error) {
	newCmd := &Cmd{}
	err := json.Unmarshal(jsonContent, newCmd)
	return *newCmd, err
}

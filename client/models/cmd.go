package models

import (
	"encoding/json"
)

type CommandMessage interface {
	GetCategory() string
	GetCommand() string
	Serialize() ([]byte, error)
}

type CMD3Message struct {
	Category string
	Command string
	FromWallet string
	ToWallet string
	Amount string 
}

func (c CMD3Message) GetCategory() string {
	return c.Category
}

func (c CMD3Message) GetCommand() string {
	return c.Command
}

func (c CMD3Message) Serialize() ([]byte, error) {
	return json.Marshal(c)
}

type CMD5Message struct {
	Category string
	Command string
	Sha256Content string
}

func (c CMD5Message) GetCategory() string {
	return c.Category
}

func (c CMD5Message) GetCommand() string {
	return c.Command
}

func (c CMD5Message) Serialize() ([]byte, error) {
	return json.Marshal(c)
}
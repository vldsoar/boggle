package models

import (
	"github.com/segmentio/ksuid"
	"net"
)

type Peer struct {
	ID         string       `json:"id,omitempty"`
	Username   string       `json:"username,omitempty"`
	Initial	   bool 		`json:"-"`
	Socket     net.Conn 	`json:"-"`
}

func NewPeer(username string) *Peer {
	return &Peer{
		ID: ksuid.New().String(),
		Username: username,
	}
}

func (p *Peer) GetInitial() bool {
	return p.Initial
}

func (p *Peer) SetInitial(initial bool)  {
	p.Initial = initial
}

func (p *Peer) GetUsername() string {
	return p.Username
}

func (p *Peer) SetUsername(username string)  {
	p.Username = username
}

func (p *Peer) GetID() string {
	return p.ID
}

func (p *Peer) SetID(id string)  {
	p.ID = id
}

func (p *Peer) SetSocket(s net.Conn) {
	p.Socket = s
}



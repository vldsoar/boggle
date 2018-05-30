package models

import "reflect"

// Representa a estrutura de uma Mensagem
// que será trocada entre cliente/cliente e cliente/servidor
type Message struct {
	ID 	 	uint 		`json:"id,omitempty"`   	// Id da mensagem
	From 	string 		`json:"from,omitempty"` 	// Remetente
	Type 	string 		`json:"type,omitempty"` 	// Tipo de mensagem (CONNECT, DISCONNECT, ADD...)
	Timeout uint 		`json:"timeout,omitempty"`	// Tempo de resposta
	Payload map[string]interface{} `json:"payload"` // Corpo da mensagem
}

func NewMessage() *Message {
	return &Message{
		Payload: make(map[string]interface{}),
	}
}

func (this *Message) GetID() uint {
	return this.ID
}

func (this *Message) SetID(id uint) *Message {
	this.ID = id
	return this
}

func (this *Message) GetFrom() string {
	return this.From
}

func (this *Message) SetFrom(from string) *Message {
	this.From = from
	return this
}

func (this *Message) GetType() string {
	return this.Type
}

func (this *Message) SetType(kind string) *Message {
	this.Type = kind
	return this
}

func (this *Message) IsEmpty() bool {
	return reflect.DeepEqual(this, Message{})
}
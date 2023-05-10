package examples

import (
	"reflect"

	"datastore/domain"
)

type Animal struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	OwnerID string `json:"owner_id"`
}

func NewAnimal(name, id, ownerID string) domain.IObject {
	return &Animal{
		Name:    name,
		ID:      id,
		Type:    "animal",
		OwnerID: ownerID,
	}
}

func (p *Animal) GetKind() string {
	return reflect.TypeOf(p).String()
}

func (p *Animal) GetID() string {
	return p.ID
}

func (p *Animal) GetName() string {
	return p.Name
}

func (p *Animal) SetID(s string) {
	p.ID = s
}

func (p *Animal) SetName(s string) {
	p.Name = s
}

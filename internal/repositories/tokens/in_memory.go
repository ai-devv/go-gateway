package tokens

import (
	"sync"
)

type InMemory struct {
	m *sync.Map
}

func (im *InMemory) Check(token Token) bool {
	_, ok := im.m.Load(token)

	return ok
}

func NewInMemory(tokens ...Token) *InMemory {
	m := &sync.Map{}

	for _, token := range tokens {
		m.Store(token, true)
	}

	return &InMemory{
		m,
	}
}

package state

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

type InMemory struct {
	m *sync.Map
}

func (im *InMemory) Save(_ context.Context, key string, state State) error {
	stateBytes, err := json.Marshal(state)

	if err != nil {
		return err
	}

	im.m.Store(key, stateBytes)

	return nil
}

func (im *InMemory) Pull(_ context.Context, key string) (State, error) {
	rawState, ok := im.m.Load(key)

	if !ok {
		return nil, fmt.Errorf("key not exists: \"%s\"", key)
	}

	var state State

	if err := json.Unmarshal(rawState.([]byte), &state); err != nil {
		return nil, err
	}

	im.m.Delete(key)

	return state, nil
}

func NewInMemory() *InMemory {
	return &InMemory{
		m: &sync.Map{},
	}
}

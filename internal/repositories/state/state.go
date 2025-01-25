package state

import "context"

type State map[string]any

type Repository interface {
	Save(ctx context.Context, key string, data State) error

	Pull(ctx context.Context, key string) (State, error)
}

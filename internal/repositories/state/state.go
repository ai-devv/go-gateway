package state

type State map[string]any

type Repository interface {
	Save(key string, data State) error

	Pull(key string) (State, error)
}

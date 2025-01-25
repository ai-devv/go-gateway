package state

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	rdb *redis.Client
}

func (r *Redis) Save(ctx context.Context, key string, state State) error {
	stateBytes, err := json.Marshal(state)

	if err != nil {
		return err
	}

	err = r.rdb.Set(ctx, key, stateBytes, 3600*time.Second).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) Pull(ctx context.Context, key string) (State, error) {
	rawState, err := r.rdb.Get(ctx, key).Bytes()

	if err != nil {
		return nil, err
	}

	err = r.rdb.Del(ctx, key).Err()

	if err != nil {
		return nil, err
	}

	var state State

	if err := json.Unmarshal(rawState, &state); err != nil {
		return nil, err
	}

	return state, nil
}

func NewRedis(rdb *redis.Client) *Redis {
	return &Redis{
		rdb,
	}
}

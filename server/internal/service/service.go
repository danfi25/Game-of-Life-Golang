package service

import (
	"game/server/pkg/life"
	"math/rand"
	"time"
)

// LifeService stores the game state
type LifeService struct {
	currentWorld *life.World
	nextWorld    *life.World
}

func New(height, width int) (*LifeService, error) {
	rand.NewSource(time.Now().UTC().UnixNano())

	currentWorld, err := life.NewWorld(height, width)
	if err != nil {
		return nil, err
	}
	// let's fill this chaotically for an example
	currentWorld.RandInit(40)

	newWorld, err := life.NewWorld(height, width)
	if err != nil {
		return nil, err
	}

	ls := LifeService{
		currentWorld: currentWorld,
		nextWorld:    newWorld,
	}

	return &ls, nil
}

// NewState receives game state
func (ls *LifeService) NewState() *life.World {
	life.NextState(ls.currentWorld, ls.nextWorld)

	ls.currentWorld = ls.nextWorld

	return ls.currentWorld
}

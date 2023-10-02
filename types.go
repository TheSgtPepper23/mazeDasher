package main

import "time"

type MapTensor struct {
	Width  uint8
	Height uint8
	Tensor [][]uint8
}

type Level struct {
	Name     string
	Origin   string
	BestTime time.Duration
}

type GameState struct {
	MaxLevel       int
	CurrentLevel   int
	ExistingLevels []*Level
}

type IStorage interface {
	SaveState() error
	GetState() error
}

type FileStorage struct {
	Filename     string
	CurrentState GameState
}

package main

import (
	"sync"
)

type randomItemGenerator struct {
	titles     []string
	descs      []string
	titleIndex int
	descIndex  int
	mtx        *sync.Mutex
	shuffle    *sync.Once
}

func (r *randomItemGenerator) reset() { _ = "STUB: not implemented"; return }

func (r *randomItemGenerator) next() item { _ = "STUB: not implemented"; return *new(item) }

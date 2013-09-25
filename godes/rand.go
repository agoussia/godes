// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package godes

import (
	"math/rand"
	"time"
)

type Distribution struct {
	generator *rand.Rand
}

type UniformDistr struct {
	Distribution
}

func NewUniformDistr() *UniformDistr {
	rd := Distribution{rand.New(rand.NewSource(time.Now().UnixNano()))}
	dist := UniformDistr{rd}
	return &dist
}

func (b *UniformDistr) Get(min float64, max float64) float64 {
	return b.generator.Float64()*(max-min) + min
}

func (b *UniformDistr) Clear() {
	b.generator = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package godes

import (
	"math/rand"
	"time"
)

type distribution struct {
	generator *rand.Rand
}

// Clear reinitiate the random generator
func (b *distribution) Clear() {
	b.generator = rand.New(rand.NewSource(time.Now().UnixNano()))
}

//UniformDistr represents the generator for the uniform distribution
type UniformDistr struct {
	distribution
}

//NewUniformDistr initiats the generator for the uniform distribution
func NewUniformDistr() *UniformDistr {
	dist := UniformDistr{distribution{rand.New(rand.NewSource(time.Now().UnixNano()))}}
	return &dist
}

// Get returns new radom value from the uniform distribution generator
func (b *UniformDistr) Get(min float64, max float64) float64 {
	return b.generator.Float64()*(max-min) + min
}

//NormalDistr represents the generator for the normal distribution
type NormalDistr struct {
	distribution
}

//NewNormalDistr initiats the generator for the normal distribution
func NewNormalDistr() *NormalDistr {
	dist := NormalDistr{distribution{rand.New(rand.NewSource(time.Now().UnixNano()))}}
	return &dist
}

// Get returns new radom value from the normal distribution generator
func (b *NormalDistr) Get(mean float64, sigma float64) float64 {
	return b.generator.NormFloat64()*sigma + mean
}

//ExpDistr represents the generator for the exponential distribution
type ExpDistr struct {
	distribution
}

//NewExpDistr initiats the generator for the exponential distribution
func NewExpDistr() *ExpDistr {

	dist := ExpDistr{distribution{rand.New(rand.NewSource(time.Now().UnixNano()))}}
	return &dist
}

// Get returns new radom value from the exponential distribution generator
func (b *ExpDistr) Get(lambda float64) float64 {
	return b.generator.ExpFloat64() / lambda
}

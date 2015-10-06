// Copyright 2015 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
//
// Godes  is the general-purpose simulation library
// which includes the  simulation engine  and building blocks
// for modeling a wide variety of systems at varying levels of details.
//
// Godes contains set of built-in functions for generating random numbers
// for commonly used probability distributions (see examples for the usage).
// Each of the distrubutions in Godes has one or more parameter values associated with it:
// 		Uniform: Min, Max
// 		Normal: Mean and Standard Deviation
// 		Exponential: Lambda
//		Triangular: Min, Mode, Max

package godes

import (
	"math"
	"math/rand"
	//"fmt"
)

var seedCount int64 = 100000

type distribution struct {
	generator *rand.Rand
}


//UniformDistr represents the generator for the uniform distribution
type UniformDistr struct {
	distribution
}

//NewUniformDistr initiats the generator for the uniform distribution
func NewUniformDistr(repetion bool) *UniformDistr {
	if repetion {
		seedCount++
		return &UniformDistr{distribution{rand.New(rand.NewSource(seedCount))}}
	} else {
		return &UniformDistr{distribution{rand.New(rand.NewSource(GetCurComputerTime()))}}
	}	
}

// Get returns new radom value from the uniform distribution generator
func (b *UniformDistr) Get(min float64, max float64) float64 {
	return b.generator.Float64()*(max-min) + min
}

// NormalDistr represents the generator for the normal distribution
type NormalDistr struct {
	distribution
}

// NewNormalDistr initiats the generator for the normal distribution
func NewNormalDistr(repetion bool) *NormalDistr {
	if repetion {
		seedCount++
		return &NormalDistr{distribution{rand.New(rand.NewSource(seedCount))}}
	} else {
		return &NormalDistr{distribution{rand.New(rand.NewSource(GetCurComputerTime()))}}
	}
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
// If repetition flag is true, the generator will generate the same sequences for every execution
func NewExpDistr(repetion bool) *ExpDistr {

	if repetion {
		seedCount++
		return &ExpDistr{distribution{rand.New(rand.NewSource(seedCount))}}
	} else {
		return &ExpDistr{distribution{rand.New(rand.NewSource(GetCurComputerTime()))}}
	}

}

// Get returns new radom value from the exponential distribution generator
func (b *ExpDistr) Get(lambda float64) float64 {
	return b.generator.ExpFloat64() / lambda
}

//TriangularDistr represents the generator for the triangular distribution
type TriangularDistr struct {
	distribution
}

//NewTriangularDistr initiats the generator for the triangular distribution
// If repetition flag is true, the generator will generate the same sequences for every execution
func NewTriangularDistr(repetion bool) *TriangularDistr {
	if repetion {
		seedCount++
		return &TriangularDistr{distribution{rand.New(rand.NewSource(seedCount))}}
	} else {
		return &TriangularDistr{distribution{rand.New(rand.NewSource(GetCurComputerTime()))}}
	}
}

// Get returns new radom value from the triangular distribution generator
func (bd *TriangularDistr) Get(a float64, b float64, c float64) float64 {
	u := bd.generator.Float64()
	f := (c - a) / (b - a)
	if u < f {
		return a + math.Sqrt(u*(b-a)*(c-a))
	} else {
		return b - math.Sqrt((1.-u)*(b-a)*(b-c))
	}
}

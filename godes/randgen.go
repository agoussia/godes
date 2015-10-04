// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package godes

import (
	"math"
	"math/rand"
	"time"
	//"fmt"
)

const (
	UPPER = 99999999999
	LOWER = -99999999999
)

var curTime int64

func getCurTime() int64 {

	ct := time.Now().UnixNano()
	if ct > curTime {
		curTime = ct
		return ct
	} else if ct == curTime {
		curTime = ct + 1
		return curTime
	} else {
		curTime++
		return curTime
	}

}

type AvgCounter struct {
	min   float64
	max   float64
	sum   float64
	count int64
}

// Clear reinitiate the counter

func (a *AvgCounter) Init() {
	a.min = UPPER
	a.max = LOWER
	a.sum = 0
	a.count = 0

}

// Clear reinitiate the counter
func (a *AvgCounter) Add(item float64) {
	a.sum = a.sum + item
	a.count++

	if a.min == UPPER {
		a.min = item
	}
	if a.max == LOWER {
		a.max = item
	}
	if item < a.min {
		a.min = item
	}
	if item > a.max {
		a.max = item
	}
	//fmt.Printf("=====%3.2f   %3.2f   %3.2f\n", item,a.min,a.max)

}

func (a *AvgCounter) GetAverage() float64 {
	return a.sum / float64(a.count)
}

func (a *AvgCounter) GetMin() float64 {
	return a.min
}

func (a *AvgCounter) GetMax() float64 {
	return a.max
}

func (a *AvgCounter) GetCount() int64 {
	return a.count
}

//New
func NewAvgCounter() *AvgCounter {
	avg := AvgCounter{}
	return &avg
}

type distribution struct {
	generator *rand.Rand
}

// Clear reinitiate the random generator
func (b *distribution) Clear() {
	b.generator = rand.New(rand.NewSource(getCurTime()))
}

//UniformDistr represents the generator for the uniform distribution
type UniformDistr struct {
	distribution
}

//NewUniformDistr initiats the generator for the uniform distribution
func NewUniformDistr() *UniformDistr {
	dist := UniformDistr{distribution{rand.New(rand.NewSource(getCurTime()))}}
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
	dist := NormalDistr{distribution{rand.New(rand.NewSource(getCurTime()))}}
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

	dist := ExpDistr{distribution{rand.New(rand.NewSource(getCurTime()))}}
	return &dist
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

func NewTriangularDistr() *TriangularDistr {

	dist := TriangularDistr{distribution{rand.New(rand.NewSource(getCurTime()))}}
	return &dist
}

// Get returns new radom value from the triangular distribution generator
func (bd *TriangularDistr) Get(a float64,b float64, c float64) float64 {
	
	u:=bd.generator.Float64()
	f:=(c-a)/(b-a)
	if u < f {
		return a+math.Sqrt(u*(b-a)*(c-a))	
	}else{
		return b-math.Sqrt((1.-u)*(b-a)*(b-c))
	}	
}




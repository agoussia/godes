// Copyright 2015 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
//
// Godes  is the general-purpose simulation library
// which includes the  simulation engine  and building blocks
// for modeling a wide variety of systems at varying levels of details.
//
// Godes Main Features:
//
//1.Active Objects:
//All active objects in Godes shall implement the RunnerInterface
//
//2.Random Generators:
//Godes contains set of built-in functions for generating random numbers for commonly used probability distributions.
//Each of the distrubutions in Godes has one or more parameter values associated with it:Uniform (Min, Max), Normal (Mean and Standard Deviation), Exponential (Lambda), Triangular(Min, Mode, Max)
//
//3.Queues:
//Godes implements operations with FIFO and LIFO queues
//
//4.BooleanControl :
//Godes uses BooleanControl variables as a locks for
//syncronizing execution of multiple runners
//
//5.StatCollector:
//The object calculates and prints statistical parameters for set of samples collected during the simulation.
//
//See examples for usage.
package godes


// BooleanControl is a boolean control variable
type BooleanControl struct {
	state bool
}

// NewBooleanControl constructs a BooleanControl
func NewBooleanControl() *BooleanControl {
	return &BooleanControl{state: false}
}

//Wait stops the runner  untill the BooleanControll bc is set to true
func (bc *BooleanControl) Wait(b bool) {
	if bc.state == b {
		//do nothing
	} else {
		modl.booleanControlWait(bc, b)
	}
}

//Wait stops the runner  untill the BooleanControll bc is set to true or timeout
func (bc *BooleanControl) WaitAndTimeout(b bool, timeOut float64) {
	if bc.state == b {
		//do nothing
	} else {
		modl.booleanControlWaitAndTimeout(bc, b, timeOut)
	}
}

// Set changes the value of bc
func (bc *BooleanControl) Set(b bool) {

	if bc.state == b {
		//do nothing
	} else {
		bc.state = b
	}
}

// getState returns value of bc
func (bc *BooleanControl) GetState() bool {
	return bc.state
}

// Clear sets the bc value to default false
func (bc *BooleanControl) Clear() {
	bc.state = false
}

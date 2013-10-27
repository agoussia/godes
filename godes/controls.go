// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package godes

// BooleanControl is a boolean control variable
type BooleanControl struct {
	state bool
}

// NewBooleanControl returns a BooleanControl
func NewBooleanControl() *BooleanControl {
	return &BooleanControl{state: false}
}

//Wait stops the runner  untill the BooleanControll bc is set to true
func (bc *BooleanControl) Wait(b bool) {
	if bc.state == b {
		//do nothing
	} else {
		model.booleanControlWait(bc, b)
	}

}

//Wait stops the runner  untill the BooleanControll bc is set to true or timeout
func (bc *BooleanControl) WaitAndTimeout(b bool, timeOut float64) {
	if bc.state == b {
		//do nothing
	} else {
		model.booleanControlWaitAndTimeout(bc, b, timeOut)
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

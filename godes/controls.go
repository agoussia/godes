// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package godes

type BooleanControl struct {
	state bool
}

func NewBooleanControl() *BooleanControl {
	return &BooleanControl{state: false}
}

func (bc *BooleanControl) Wait(b bool) {
	if bc.state == b {
		//do nothing
	} else {
		model.booleanControlWait(bc, b)
	}

}

func (bc *BooleanControl) Set(b bool) {

	if bc.state == b {
		//do nothing
	} else {
		bc.state = b
	}

}

func (bc *BooleanControl) getState() bool {
	return bc.state

}

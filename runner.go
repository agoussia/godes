// Copyright 2015 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
//
// Godes  is the general-purpose simulation library
// which includes the  simulation engine  and building blocks
// for modeling a wide variety of systems at varying levels of details.
//
// All active objects in Godes shall implement the RunnerInterface
// See examples for the usage.
//

package godes

import (
	"fmt"
	"time"
)

type RunnerInterface interface {
	Run()
	setState(i int)
	getState() int
	setChannel(c chan int)
	getChannel() chan int
	setInternalId(id int)
	getInternalId() int
	setMovingTime(m float64)
	getMovingTime() float64
	setMarkTime(m time.Time)
	getMarkTime() time.Time
	setPriority(p int)
	getPriority() int
	setWaitingForBool(p bool)
	getWaitingForBool() bool
	setWaitingForBoolControl(p *BooleanControl)
	getWaitingForBoolControl() *BooleanControl
	setWaitingForBoolControlTimeoutId(id int)
	getWaitingForBoolControlTimeoutId() int
	
}

type Runner struct {
	state                          int
	channel                        chan int
	internalId                     int
	movingTime                     float64
	markTime                       time.Time
	priority                       int
	waitingForBool                 bool
	waitingForBoolControl          *BooleanControl
	waitingForBoolControlTimeoutId int
	//schedulledTime					 float64
}

type TimeoutRunner struct {
	*Runner
	original      RunnerInterface
	timeoutPeriod float64
}

func (timeOut *TimeoutRunner) Run() {
	Advance(timeOut.timeoutPeriod)
	if timeOut.original.getWaitingForBoolControl != nil && timeOut.original.getWaitingForBoolControlTimeoutId() == timeOut.internalId {
		timeOut.original.setState(rUNNER_STATE_READY)
		timeOut.original.setWaitingForBoolControl(nil)
		modl.addToMovingList(timeOut.original)
		delete(modl.waitingConditionMap, timeOut.original.getInternalId())
	}

}

func newRunner() *Runner {
	return &Runner{}
}

func (b *Runner) Run() {
	fmt.Println("Run Run Run Run")
}

func (b *Runner) setState(i int) {
	b.state = i
}

func (b *Runner) getState() int {
	return b.state
}

func (b *Runner) setChannel(c chan int) {
	b.channel = c
}

func (b *Runner) getChannel() chan int {
	return b.channel
}


func (b *Runner) setInternalId(i int) {
	b.internalId = i

}
func (b *Runner) getInternalId() int {
	return b.internalId
}

func (b *Runner) setMovingTime(m float64) {
	b.movingTime = m

}
func (b *Runner) getMovingTime() float64 {
	return b.movingTime
}




func (b *Runner) setMarkTime(m time.Time) {
	b.markTime = m

}
func (b *Runner) getMarkTime() time.Time {
	return b.markTime
}

func (b *Runner) setPriority(p int) {
	b.priority = p
}
func (b *Runner) getPriority() int {
	return b.priority
}

func (b *Runner) setWaitingForBool(p bool) {
	b.waitingForBool = p

}

func (b *Runner) getWaitingForBool() bool {
	return b.waitingForBool

}

func (b *Runner) setWaitingForBoolControl(p *BooleanControl) {
	b.waitingForBoolControl = p

}

func (b *Runner) getWaitingForBoolControl() *BooleanControl {
	return b.waitingForBoolControl
}

func (b *Runner) setWaitingForBoolControlTimeoutId(p int) {
	b.waitingForBoolControlTimeoutId = p

}

func (b *Runner) getWaitingForBoolControlTimeoutId() int {
	return b.waitingForBoolControlTimeoutId
}

func (b *Runner) IsShedulled() bool{
	if b.state == rUNNER_STATE_SCHEDULED {
		return true
	}
	return false
}

func (b *Runner) GetMovingTime() float64 {
	if b.state == rUNNER_STATE_SCHEDULED {
		return b.movingTime
	}else{
		panic ("Runner is Not Shedulled ")
	}
}




func (b *Runner) String() string {

	var st = ""

	switch b.state {
	case rUNNER_STATE_READY:
		st = "READY"
	case rUNNER_STATE_ACTIVE:
		st = "ACTIVE"
	case rUNNER_STATE_WAITING_COND:
		st = "WAITING_COND"
	case rUNNER_STATE_SCHEDULED:
		st = "SCHEDULED"
	case rUNNER_STATE_INTERRUPTED:
		st = "INTERRUPTED"
	case rUNNER_STATE_TERMINATED:
		st = "TERMINATED"

	default:
		panic("Unknown state")
	}
	return fmt.Sprintf(" st=%v ch=%v id=%v mt=%v mk=%v pr=%v", st, b.channel, b.internalId, b.movingTime, b.markTime, b.priority)
}

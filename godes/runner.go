// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package godes

import (
	"fmt"
	"time"
)

type RunnerInterface interface {
	Run()

	setState(i int)
	GetState() int

	setChannel(c chan int)
	getChannel() chan int

	setId(id int)
	GetId() int

	setMovingTime(m float64)
	GetMovingTime() float64

	setMarkTime(m time.Time)
	GetMarkTime() time.Time

	setPriority(p int)
	GetPriority() int

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
	id                             int
	movingTime                     float64
	markTime                       time.Time
	priority                       int
	waitingForBool                 bool
	waitingForBoolControl          *BooleanControl
	waitingForBoolControlTimeoutId int
}

type TimeoutRunner struct {
	*Runner
	original      RunnerInterface
	timeoutPeriod float64
}

func (timeOut *TimeoutRunner) Run() {
	Advance(timeOut.timeoutPeriod)

	if timeOut.original.getWaitingForBoolControl != nil && timeOut.original.getWaitingForBoolControlTimeoutId() == timeOut.id {
		timeOut.original.setState(RUNNER_STATE_READY)
		timeOut.original.setWaitingForBoolControl(nil)
		model.addToMovingList(timeOut.original)
		delete(model.waitingConditionMap, timeOut.original.GetId())

	}

}

func NewRunner() *Runner {

	return &Runner{}
}

func (b *Runner) Run() {
	fmt.Println("Run Run Run Run")
}

func (b *Runner) setState(i int) {
	b.state = i
}

func (b *Runner) GetState() int {
	return b.state
}

func (b *Runner) setChannel(c chan int) {
	b.channel = c
}

func (b *Runner) getChannel() chan int {
	return b.channel
}

func (b *Runner) setId(i int) {
	b.id = i

}
func (b *Runner) GetId() int {
	return b.id
}

func (b *Runner) setMovingTime(m float64) {
	b.movingTime = m

}
func (b *Runner) GetMovingTime() float64 {
	return b.movingTime
}

func (b *Runner) setMarkTime(m time.Time) {
	b.markTime = m

}
func (b *Runner) GetMarkTime() time.Time {
	return b.markTime
}

func (b *Runner) setPriority(p int) {
	b.priority = p
}
func (b *Runner) GetPriority() int {
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

func (b *Runner) String() string {

	var st = ""

	switch b.state {
	case RUNNER_STATE_READY:
		st = "READY"
	case RUNNER_STATE_ACTIVE:
		st = "ACTIVE"
	case RUNNER_STATE_WAITING_COND:
		st = "WAITING_COND"
	case RUNNER_STATE_SCHEDULED:
		st = "SCHEDULED"
	case RUNNER_STATE_INTERRUPTED:
		st = "INTERRUPTED"
	case RUNNER_STATE_TERMINATED:
		st = "TERMINATED"

	default:
		panic("Unknown state")
	}

	return fmt.Sprintf(" st=%v ch=%v id=%v mt=%v mk=%v pr=%v", st, b.channel, b.id, b.movingTime, b.markTime, b.priority)
}

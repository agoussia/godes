// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package godes

const simulationSecondScale = 100
const TRANSITION_ACTIVATE = 0
const TRANSITION_IC_YIELD_TO = 1
const TRANSITION_YIELD = 2
const TRANSITION_TERMINATE = 3
const TRANSITION_WAIT = 4
const TRANSITION_WAIT_UNTIL = 5
const TRANSITION_ADVANCE = 6
const TRANSITION_INTERRUPT = 7
const TRANSITION_RESUME = 8
const TRANSITION_REACTIVATE = 9
const TRANSITION_IC = 10
const RUNNER_STATE_READY = 1
const RUNNER_STATE_ACTIVE = 2
const RUNNER_STATE_WAITING = 3
const RUNNER_STATE_WAITING_COND = 4
const RUNNER_STATE_SCHEDULED = 5
const RUNNER_STATE_INTERRUPTED = 6
const RUNNER_STATE_TERMINATED = 7

var model *Model

var Stime float64 = 0

// Main Functions

func startSimulationRun(verbose bool) {

	if model != nil {
		panic("model is already active")
	}
	Stime = 0
	model = NewModel(verbose)
	model.simulationActive = true
	model.control()
	//assuming that it comes from the main go routine

}

func WaitUntilDone() {
	if model == nil {
		panic(" not initilized")
	}
	model.waitUntillDone()
}

func ActivateRunner(runner RunnerInterface) {
	if runner == nil {
		panic("runner is nil")
	}
	if model == nil {
		startSimulationRun(false)
	}
	model.activate(runner)
}

func Advance(interval float64) {
	if model == nil {
		startSimulationRun(false)
	}
	model.advance(interval)
}

func Verbose(v bool) {
	if model == nil {
		startSimulationRun(v)
	}
	model.DEBUG = v
}

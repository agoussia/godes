// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package godes

const simulationSecondScale = 100

const RUNNER_STATE_READY = 1
const RUNNER_STATE_ACTIVE = 2
const RUNNER_STATE_WAITING_COND = 4
const RUNNER_STATE_SCHEDULED = 5
const RUNNER_STATE_INTERRUPTED = 6
const RUNNER_STATE_TERMINATED = 7

var model *Model

var stime float64 = 0

// Main Functions

func startSimulationRun(verbose bool) {

	if model != nil {
		panic("model is already active")
	}
	stime = 0
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

func Clear() {
	if model == nil {
		panic(" No model exist")
	} else {

		stime = 0
		model = NewModel(model.DEBUG)
		model.simulationActive = true
		model.control()
	}
}
func GetSystemTime() float64 {
	return stime
}

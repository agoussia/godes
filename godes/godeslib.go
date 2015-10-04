// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*
Package godes  is the general-purpose simulation library
which includes the  simulation engine  and building blocks
for modeling a wide variety
of systems at varying levels of detail.
*/
package godes

const simulationSecondScale = 100

const RUNNER_STATE_READY = 0
const RUNNER_STATE_ACTIVE = 1
const RUNNER_STATE_WAITING_COND = 2
const RUNNER_STATE_SCHEDULED = 3
const RUNNER_STATE_INTERRUPTED = 4
const RUNNER_STATE_TERMINATED = 5

var model *Model

var stime float64 = 0

// StartSimulationRun start the simulation in the verbose mode when verbose is true
func createModel(verbose bool) {

	if model != nil {
		panic("model is already active")
	}
	stime = 0
	model = newModel(verbose)
	//model.simulationActive = true
	//model.control()
	//assuming that it comes from the main go routine

}

// WaitUntilDone stops the main goroutine and waits until all the runners execute the Run()
func WaitUntilDone() {
	if model == nil {
		panic(" not initilized")
	}
	model.waitUntillDone()
}

//ActivateRunner the runner
//If the library is not activated its start the libary in the non verbouse mode
func AddRunner(runner RunnerInterface) {
	if runner == nil {
		panic("runner is nil")
	}
	if model == nil {
		createModel(false)
	}
	model.add(runner)
}

func Interrupt(runner RunnerInterface) {
	if runner == nil {
		panic("runner is nil")
	}
	if model == nil {
		panic("model is nil")
	}
	model.interrupt(runner)
}

func Resume(runner RunnerInterface, timeChange float64) {
	if runner == nil {
		panic("runner is nil")
	}
	if model == nil {
		panic("model is nil")
	}
	model.resume(runner, timeChange)
}

func Run() {
	if model == nil {
		createModel(false)
	}
	//assuming that it comes from the main go routine
	if model.activeRunner == nil {
		panic("runner is nil")
	}

	if model.activeRunner.GetId() != 0 {
		panic("it comes from not from the main go routine")
	}

	model.simulationActive = true
	model.control()

}

//Advance the simulation time
func Advance(interval float64) {
	if model == nil {
		createModel(false)
	}
	model.advance(interval)
}

// Verbose sets the library in the verbose mode
func Verbose(v bool) {
	if model == nil {
		createModel(v)
	}
	model.DEBUG = v
}

// Clear the library between the runs
func Clear() {
	if model == nil {
		panic(" No model exist")
	} else {

		stime = 0
		model = newModel(model.DEBUG)
		//model.simulationActive = true
		//model.control()
	}
}

// GetSystemTime retuns the current simulation time
func GetSystemTime() float64 {
	return stime
}

// GetSystemTime retuns the current simulation time
func Yield()  {
	Advance(0.01)
}
// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"godes"
)

const PT_MEAN = 10.0          //	Avg. processing time in minutes
const PT_SIGMA = 2.0          //	Sigma of processing time
const MTTF = 300.0            // 	Mean time to failure in minutes
const REPAIR_TIME = 30.0      //	Time it takes to repair a machine in minutes
const REPAIR_TIME_SIGMA = 1.0 //	Sigma of repair time

const NUM_MACHINES = 10
const SHUT_DOWN_TIME = 4 * 7 * 24 * 60

// random generator for the processing time - normal distribution
var processingGen *godes.NormalDistr = godes.NewNormalDistr()

// random generator for the  time   until the next failure for a machine - exponential distribution
var breaksGen *godes.ExpDistr = godes.NewExpDistr()

// true when repairman is available for carrying a repair
var repairManAvailableSwt *godes.BooleanControl = godes.NewBooleanControl()

type Machine struct {
	*godes.Runner
	partsCount int
	number     int
}

func (machine *Machine) Run() {
	for {
		godes.Advance(processingGen.Get(PT_MEAN, PT_SIGMA))
		machine.partsCount = machine.partsCount + 1
		if godes.GetSystemTime() > SHUT_DOWN_TIME {
			break
		}
	}
}

type MachineRepair struct {
	*godes.Runner
	machine *Machine
}

func (machineRepair *MachineRepair) Run() {
	machine := machineRepair.machine
	for {
		godes.Advance(breaksGen.Get(1 / MTTF))
		if machine.GetState() == godes.RUNNER_STATE_SCHEDULED {
			breakTime := godes.GetSystemTime()
			//interrupt machine
			godes.Interrupt(machine)
			repairManAvailableSwt.Wait(true)
			//repair
			repairManAvailableSwt.Set(false)
			godes.Advance(processingGen.Get(REPAIR_TIME, REPAIR_TIME_SIGMA))
			//release repairman
			repairManAvailableSwt.Set(true)
			//resume machine and change the scheduled time to compensate the idle time
			godes.Resume(machine, godes.GetSystemTime()-breakTime)
		}

		if godes.GetSystemTime() > SHUT_DOWN_TIME {
			break
		}
	}
}

func main() {
	var m *Machine
	x := make(map[int]*Machine)
	for i := 0; i < NUM_MACHINES; i++ {
		m = &Machine{&godes.Runner{}, 0, i}
		godes.AddRunner(m)
		godes.AddRunner(&MachineRepair{&godes.Runner{}, m})
		x[i] = m
	}
	repairManAvailableSwt.Set(true)
	godes.Run()
	godes.WaitUntilDone()
	//print results
	for i := 0; i < NUM_MACHINES; i++ {
		m = x[i]
		fmt.Printf(" Machine # %v %v \n", m.number, m.partsCount)
	}
}

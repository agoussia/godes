// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package main

import (
	"fmt"

	"github.com/agoussia/godes"
)

const PT_MEAN = 10.0          //	Avg. processing time in minutes
const PT_SIGMA = 2.0          //	Sigma of processing time
const MTTF = 300.0            // 	Mean time to failure in minutes
const REPAIR_TIME = 30.0      //	Time it takes to repair a machine in minutes
const REPAIR_TIME_SIGMA = 1.0 //	Sigma of repair time

const NUM_MACHINES = 10
const SHUT_DOWN_TIME = 4 * 7 * 24 * 60

// random generator for the processing time - normal distribution
var processingGen *godes.NormalDistr = godes.NewNormalDistr(true)

// random generator for the  time   until the next failure for a machine - exponential distribution
var breaksGen *godes.ExpDistr = godes.NewExpDistr(true)

// true when repairman is available for carrying a repair
var repairManAvailableSwt *godes.BooleanControl = godes.NewBooleanControl()

type Machine struct {
	*godes.Runner
	partsCount int
	number     int
	finished   bool
}

func (machine *Machine) Run() {
	for {
		godes.Advance(processingGen.Get(PT_MEAN, PT_SIGMA))
		machine.partsCount++
		if godes.GetSystemTime() > SHUT_DOWN_TIME {
			machine.finished = true
			break
		}

	}
	fmt.Printf(" Machine # %v %v \n", machine.number, machine.partsCount)
}

type MachineRepair struct {
	*godes.Runner
	machine *Machine
}

func (machineRepair *MachineRepair) Run() {
	machine := machineRepair.machine
	for {
		godes.Advance(breaksGen.Get(1 / MTTF))
		if machine.finished {
			break
		}

		interrupted := godes.GetSystemTime()
		godes.Interrupt(machine)
		repairManAvailableSwt.Wait(true)
		if machine.finished {
			break
		}
		repairManAvailableSwt.Set(false)
		godes.Advance(processingGen.Get(REPAIR_TIME, REPAIR_TIME_SIGMA))
		if machine.finished {
			break
		}
		//release repairman
		repairManAvailableSwt.Set(true)
		//resume machine and change the scheduled time to compensate delay
		godes.Resume(machine, godes.GetSystemTime()-interrupted)

	}

}

func main() {

	godes.Run()
	repairManAvailableSwt.Set(true)
	var m *Machine
	for i := 0; i < NUM_MACHINES; i++ {
		m = &Machine{&godes.Runner{}, 0, i, false}
		godes.AddRunner(m)
		godes.AddRunner(&MachineRepair{&godes.Runner{}, m})
	}
	godes.WaitUntilDone()
}

/* OUTPUT
Machine # 4 3599
Machine # 5 3625
Machine # 7 3580
Machine # 2 3724
Machine # 1 3554
Machine # 9 3581
Machine # 0 3585
Machine # 6 3604
Machine # 3 3639
Machine # 8 3625
*/

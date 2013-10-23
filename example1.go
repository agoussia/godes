// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"godes"
)

// the arrival and service are two random number generators for the uniform  distribution
var arrival *godes.UniformDistr = godes.NewUniformDistr()
var service *godes.UniformDistr = godes.NewUniformDistr()

// tableBusy is the boolean control variable than can be accessed and changed by number of Runners
var tableBusy *godes.BooleanControl = godes.NewBooleanControl()

// the Visitor is a Runner
// any type of the Runner should be defined as struct // with the *godes.Runner as anonimous field
type Visitor struct {
	*godes.Runner
	number int
}

var visitorsCount int = 0

func (vst Visitor) Run() { // Any runner should have the Run method
	fmt.Printf("Visitor %v arrives at time= %6.3f \n", vst.number, godes.GetSystemTime())
	tableBusy.Wait(false) // this will wait till the tableBusy control becomes false
	tableBusy.Set(true)   // sets the tableBusy control to true - the table is busy
	fmt.Printf("Visitor %v gets the table at time= %6.3f \n", vst.number, godes.GetSystemTime())
	godes.Advance(service.Get(10, 60)) //the function advance the simulation time by the value in the argument
	tableBusy.Set(false)               // sets the tableBusy control to false - the table is idle
	fmt.Printf("Visitor %v leaves at time= %6.3f \n", vst.number, godes.GetSystemTime())
}
func main() {
	var shutdown_time float64 = 8 * 60
	godes.Run()
	for {
		//godes.GetSystemTime() is the current simulation time
		if godes.GetSystemTime() < shutdown_time {
			//the function acivates the Runner
			godes.AddRunner(Visitor{&godes.Runner{}, visitorsCount})
			godes.Advance(arrival.Get(0, 70))
			visitorsCount++
		} else {
			break
		}
	}
	godes.WaitUntilDone() // waits for all the runners to finish the Run()
}

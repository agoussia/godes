// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package main

import (
	"fmt"

	"github.com/godes"
)

// the arrival and service are two random number generators for the uniform  distribution
var arrival *godes.UniformDistr = godes.NewUniformDistr(true)
var service *godes.UniformDistr = godes.NewUniformDistr(true)

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
	fmt.Printf("%-6.3f \t Visitor %v arrives \n", godes.GetSystemTime(), vst.number)
	tableBusy.Wait(false) // this will wait till the tableBusy control becomes false
	tableBusy.Set(true)   // sets the tableBusy control to true - the table is busy
	fmt.Printf("%-6.3f \t Visitor %v gets the table \n", godes.GetSystemTime(), vst.number)
	godes.Advance(service.Get(10, 60)) //the function advance the simulation time by the value in the argument
	tableBusy.Set(false)               // sets the tableBusy control to false - the table is idle
	fmt.Printf("%-6.3f \t Visitor %v leaves \n", godes.GetSystemTime(), vst.number)
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

/* OUTPUT
0.000  	 Visitor 0 arrives
0.000  	 Visitor 0 gets the table
13.374 	 Visitor 0 leaves
37.486 	 Visitor 1 arrives
37.486 	 Visitor 1 gets the table
60.558 	 Visitor 1 leaves
98.737 	 Visitor 2 arrives
98.737 	 Visitor 2 gets the table
107.468 	 Visitor 3 arrives
146.824 	 Visitor 2 leaves
146.824 	 Visitor 3 gets the table
149.471 	 Visitor 4 arrives
171.623 	 Visitor 3 leaves
171.623 	 Visitor 4 gets the table
187.234 	 Visitor 4 leaves
207.523 	 Visitor 5 arrives
207.523 	 Visitor 5 gets the table
230.922 	 Visitor 6 arrives
245.859 	 Visitor 5 leaves
245.859 	 Visitor 6 gets the table
261.770 	 Visitor 7 arrives
269.668 	 Visitor 8 arrives
272.368 	 Visitor 6 leaves
272.368 	 Visitor 7 gets the table
290.484 	 Visitor 7 leaves
290.484 	 Visitor 8 gets the table
310.261 	 Visitor 9 arrives
333.570 	 Visitor 8 leaves
333.570 	 Visitor 9 gets the table
338.323 	 Visitor 10 arrives
354.874 	 Visitor 9 leaves
354.874 	 Visitor 10 gets the table
393.826 	 Visitor 10 leaves
397.720 	 Visitor 11 arrives
397.720 	 Visitor 11 gets the table
409.123 	 Visitor 12 arrives
436.817 	 Visitor 13 arrives
447.731 	 Visitor 14 arrives
455.705 	 Visitor 11 leaves
455.705 	 Visitor 13 gets the table
482.955 	 Visitor 13 leaves
482.955 	 Visitor 12 gets the table
496.034 	 Visitor 12 leaves
496.034 	 Visitor 14 gets the table
555.822 	 Visitor 14 leaves
*/

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

// the Visitor is a Runner
// any type of the Runner should be defined as struct // with the *godes.Runner as anonimous field
type Visitor struct {
	*godes.Runner
	number int
}

var visitorsCount int = 0

func (vst Visitor) Run() { // Any runner should have the Run method
	fmt.Printf("Visitor %v arrives at time= %6.3f \n", vst.number, godes.Stime)

}
func main() {
	var shutdown_time float64 = 8 * 60
	for {
		//godes.Stime is the current simulation time
		if godes.Stime < shutdown_time {
			//the function acivates the Runner
			godes.ActivateRunner(Visitor{&godes.Runner{}, visitorsCount})
			godes.Advance(arrival.Get(0, 70))
			visitorsCount++
		} else {
			break
		}
	}
	godes.WaitUntilDone() // waits for all the runners to finish the Run()
}

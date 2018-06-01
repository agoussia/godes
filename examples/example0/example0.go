// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/agoussia/godes"
)

// the arrival and service are two random number generators for the uniform  distribution
var arrival *godes.UniformDistr = godes.NewUniformDistr(true)

// the Visitor is a Runner
// any type of the Runner should be defined as struct
// with the *godes.Runner as anonimous field
type Visitor struct {
	*godes.Runner
	number int
}

var visitorsCount int = 0

func (vst *Visitor) Run() { // Any runner should have the Run method
	fmt.Printf(" %-6.3f \t Visitor # %v arrives \n", godes.GetSystemTime(), vst.number)
}
func main() {
	var shutdown_time float64 = 8 * 60
	godes.Run()
	for {
		//godes.Stime is the current simulation time
		if godes.GetSystemTime() < shutdown_time {
			//the function acivates the Runner
			godes.AddRunner(&Visitor{&godes.Runner{}, visitorsCount})
			//this advance the system time
			godes.Advance(arrival.Get(0, 70))
			visitorsCount++
		} else {
			break
		}
	}
	// waits for all the runners to finish the Run()
	godes.WaitUntilDone()
}
/* 	OUTPUT:
 0.000  	 Visitor # 0 arrives 
 37.486 	 Visitor # 1 arrives 
 98.737 	 Visitor # 2 arrives 
 107.468 	 Visitor # 3 arrives 
 149.471 	 Visitor # 4 arrives 
 207.523 	 Visitor # 5 arrives 
 230.922 	 Visitor # 6 arrives 
 261.770 	 Visitor # 7 arrives 
 269.668 	 Visitor # 8 arrives 
 310.261 	 Visitor # 9 arrives 
 338.323 	 Visitor # 10 arrives 
 397.720 	 Visitor # 11 arrives 
 409.123 	 Visitor # 12 arrives 
 436.817 	 Visitor # 13 arrives 
 447.731 	 Visitor # 14 arrives 
*/
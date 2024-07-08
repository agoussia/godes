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
var service *godes.UniformDistr = godes.NewUniformDistr(true)

// true when waiter should act
var waitersSwt *godes.BooleanControl = godes.NewBooleanControl()

// FIFO Queue for the arrived
var visitorArrivalQueue *godes.FIFOQueue = godes.NewFIFOQueue("0")

// the Visitor is a Passive Object
type Visitor struct {
	id int
}

// the Waiter is a Runner
type Waiter struct {
	*godes.Runner
	id int
}

var visitorsCount int = 0
var shutdown_time float64 = 4 * 60

func (waiter *Waiter) Run() {
	for {
		waitersSwt.Wait(true)
		if visitorArrivalQueue.Len() > 0 {
			visitorArrivalQueue.Get()
			if visitorArrivalQueue.Len() == 0 {
				waitersSwt.Set(false)
			}
			godes.Advance(service.Get(10, 60)) //advance the simulation time by the visitor service time

		}
		if godes.GetSystemTime() > shutdown_time && visitorArrivalQueue.Len() == 0 {
			break
		}

	}
}

func main() {
	for runs := 0; runs < 5; runs++ {
		for i := 0; i < 2; i++ {
			godes.AddRunner(&Waiter{&godes.Runner{}, i})
		}
		godes.Run()
		for {
			visitorArrivalQueue.Place(Visitor{visitorsCount})
			waitersSwt.Set(true)
			godes.Advance(arrival.Get(0, 30))
			visitorsCount++
			if godes.GetSystemTime() > shutdown_time {
				break
			}
		}
		waitersSwt.Set(true)
		godes.WaitUntilDone() // waits for all the runners to finish the Run()
		fmt.Printf(" Run # %v \t Average Time in Queue=%6.3f \n", runs, visitorArrivalQueue.GetAverageTime())
		//clear after each run
		waitersSwt.Clear()
		visitorArrivalQueue.Clear()
		godes.Clear()

	}
}

/* OUTPUT

Run # 0 	 Average Time in Queue=15.016
 Run # 1 	 Average Time in Queue=17.741
 Run # 2 	 Average Time in Queue=49.046
 Run # 3 	 Average Time in Queue=30.696
 Run # 4 	 Average Time in Queue=14.777
*/

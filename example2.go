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

// true when waiter should act
var waitersSwt *godes.BooleanControl = godes.NewBooleanControl()

// waitingQueue co..
var visitorArrivalQueue *godes.FIFOQueue = godes.NewFIFOQueue()

// the Visitor is a Runner
// any type of the Runner should be defined as struct // with the *godes.Runner as anonimous field
type Visitor struct {
	*godes.Runner
	id int
}

type Waiter struct {
	*godes.Runner
	id int
}

var visitorsCount int = 0
var shutdown_time float64 = 8 * 60

func (waiter Waiter) Run() {

	for {
		waitersSwt.Wait(true)
		if visitorArrivalQueue.Len() > 0 {
			visitor := visitorArrivalQueue.Get()
			if visitorArrivalQueue.Len() == 0 {
				waitersSwt.Set(false)
			}
			fmt.Printf("Visitor %v is invited by waiter %v at time= %6.3f \n", visitor.(Visitor).id, waiter.id, godes.Stime)
			visitor.Run()
			fmt.Printf("Visitor %v leaves at time= %6.3f \n", visitor.(Visitor).id, godes.Stime)
		} else {
			if godes.Stime > shutdown_time {
				fmt.Printf("Waiter  %v ends the work at time= %6.3f \n", waiter.id, godes.Stime)
				break
			}
		}
	}
}

func (vst Visitor) Run() { // Any runner should have the Run method
	godes.Advance(service.Get(10, 60)) //the function advance the simulation time by the value in the argument
}

func main() {

	var visitor Visitor
	for i := 0; i < 3; i++ {
		godes.ActivateRunner(Waiter{&godes.Runner{}, i})
	}
	for {
		//godes.Stime is the current simulation time
		if godes.Stime < shutdown_time {
			//the function acivates the Runner
			godes.Advance(arrival.Get(0, 30))
			visitor = Visitor{&godes.Runner{}, visitorsCount}
			visitorArrivalQueue.Place(visitor)
			fmt.Printf("Visitor %v arrives at time= %6.3f \n", visitor.id, godes.Stime)
			waitersSwt.Set(true)

			visitorsCount++
		} else {
			break
		}
	}
	waitersSwt.Set(true)
	godes.WaitUntilDone() // waits for all the runners to finish the Run()
	fmt.Printf("Average Waiting Time %v  \n", visitorArrivalQueue.GetAverageTime())
}

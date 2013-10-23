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

// FIFO Queue for the arrived
var visitorArrivalQueue *godes.FIFOQueue = godes.NewFIFOQueue("arrivalQueue")

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
			visitor := visitorArrivalQueue.Get()
			if visitorArrivalQueue.Len() == 0 {
				waitersSwt.Set(false)
			}
			fmt.Printf("Visitor %v is invited by waiter %v at %6.3f \n", visitor.(Visitor).id, waiter.id, godes.GetSystemTime())
			godes.Advance(service.Get(10, 60)) //advance the simulation time by the visitor service time
			fmt.Printf("Visitor %v leaves at= %6.3f \n", visitor.(Visitor).id, godes.GetSystemTime())

		}
		if godes.GetSystemTime() > shutdown_time && visitorArrivalQueue.Len() == 0 {
			fmt.Printf("Waiter  %v ends the work at %6.3f \n", waiter.id, godes.GetSystemTime())
			break
		}

	}
}

func main() {

	var visitor Visitor
	for i := 0; i < 2; i++ {
		godes.AddRunner(&Waiter{&godes.Runner{}, i})
	}
	godes.Run()
	for {

		visitorArrivalQueue.Place(Visitor{visitorsCount})
		fmt.Printf("Visitor %v arrives at time= %6.3f \n", visitor.id, godes.GetSystemTime())
		waitersSwt.Set(true)
		godes.Advance(arrival.Get(0, 30))
		visitorsCount++
		if godes.GetSystemTime() > shutdown_time {
			break
		}
	}
	waitersSwt.Set(true)
	godes.WaitUntilDone() // waits for all the runners to finish the Run()
	fmt.Printf("Average Waiting Time %6.3f  \n", visitorArrivalQueue.GetAverageTime())
}

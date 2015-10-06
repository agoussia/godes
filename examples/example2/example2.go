// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"godes"
)

// the arrival and service are two random number generators for the uniform  distribution
var arrival *godes.UniformDistr = godes.NewUniformDistr(true)
var service *godes.UniformDistr = godes.NewUniformDistr(true)

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
			fmt.Printf("%-6.3f \t Visitor %v is invited by waiter %v  \n", godes.GetSystemTime(), visitor.(Visitor).id, waiter.id)
			godes.Advance(service.Get(10, 60)) //advance the simulation time by the visitor service time
			fmt.Printf("%-6.3f \t Visitor %v leaves \n", godes.GetSystemTime(), visitor.(Visitor).id)

		}
		if godes.GetSystemTime() > shutdown_time && visitorArrivalQueue.Len() == 0 {
			fmt.Printf("%-6.3f \t Waiter  %v ends the work \n", godes.GetSystemTime(), waiter.id)
			break
		}
	}
}

func main() {

	for i := 0; i < 2; i++ {
		godes.AddRunner(&Waiter{&godes.Runner{}, i})
	}
	godes.Run()
	for {

		visitorArrivalQueue.Place(Visitor{visitorsCount})
		fmt.Printf("%-6.3f \t Visitor %v arrives \n", godes.GetSystemTime(), visitorsCount)
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
/* OUTPUT
0.000  	 	Visitor 0 arrives 
0.000  	 	Visitor 0 is invited by waiter 0  
13.374 		 Visitor 0 leaves 
16.066 	 	Visitor 1 arrives 
16.066 		 Visitor 1 is invited by waiter 1  
39.137 	 	Visitor 1 leaves 
42.316 		 Visitor 2 arrives 
42.316 	 	Visitor 2 is invited by waiter 1  
46.058 	 	Visitor 3 arrives 
46.058 	 	Visitor 3 is invited by waiter 0  
64.059 	 	Visitor 4 arrives 
70.857 	 	Visitor 3 leaves 
70.857 	 	Visitor 4 is invited by waiter 0  
86.468 	 	Visitor 4 leaves 
88.938 	 	Visitor 5 arrives 
88.938 		Visitor 5 is invited by waiter 0  
90.403 	 	Visitor 2 leaves 
98.966 	 	Visitor 6 arrives 
98.966 	 	Visitor 6 is invited by waiter 1  
112.187 	 Visitor 7 arrives 
115.572 	 Visitor 8 arrives 
125.475 	 Visitor 6 leaves 
125.475 	 Visitor 7 is invited by waiter 1  
127.275 	 Visitor 5 leaves 
127.275 	 Visitor 8 is invited by waiter 0  
132.969 	 Visitor 9 arrives 
143.591 	 Visitor 7 leaves 
143.591 	 Visitor 9 is invited by waiter 1  
144.995 	 Visitor 10 arrives 
164.895 	 Visitor 9 leaves 
164.895 	 Visitor 10 is invited by waiter 1  
170.361 	 Visitor 8 leaves 
170.451 	 Visitor 11 arrives 
170.451 	 Visitor 11 is invited by waiter 0  
175.338 	 Visitor 12 arrives 
187.207 	 Visitor 13 arrives 
191.885 	 Visitor 14 arrives 
203.848 	 Visitor 10 leaves 
203.848 	 Visitor 12 is invited by waiter 1  
213.596 	 Visitor 15 arrives 
228.436 	 Visitor 11 leaves 
228.436 	 Visitor 13 is invited by waiter 0  
231.098 	 Visitor 12 leaves 
231.098 	 Visitor 14 is invited by waiter 1  
231.769 	 Visitor 16 arrives 
241.515 	 Visitor 13 leaves 
241.515 	 Visitor 15 is invited by waiter 0  
287.864 	 Visitor 15 leaves 
287.864 	 Visitor 16 is invited by waiter 0  
290.886 	 Visitor 14 leaves 
290.886 	 Waiter  1 ends the work 
330.903 	 Visitor 16 leaves 
330.903 	 Waiter  0 ends the work 
Average Waiting Time 15.016  
*/
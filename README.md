#Godes#

Libriary to Build Discrete Event Simulation Models in Go (http://golang.org/)

Copyright (c) 2013 Alex Goussiatiner agoussia@yahoo.com

###Features###
Godes is the general-purpose simulation library which includes the  simulation engine  and building blocks for modeling a wide variety of systems at varying levels of detail.

###Advantages###
* Godes is easy to learn for the people familiar with the Go and the elementary simulation concept
* Godes model executes fast  as Go compiles to machine code.
* Godes model is multiplatform as Go compiler targets the Linux, Mac OS X, FreeBSD, Microsoft Windows, etc
* Godes model can be embedded in various computer systems and over the network
* Speed of the Godes model compilation is high
* Variety of the IDE with debuggers are available for Go and Godes as well
* The Godes model can use all of the GO's features and libraries
* Code Security - the Godes includes the  source code for the library and Go is an open source project supported by Google
* Godes is free open source software under MIT license

###Installation###
* Download, install and test the Go at your machine. See instructions at http://golang.org/doc/install
* Optionally install one of the free Go IDE (i.e.LiteIDE X)
* Download Godes package
* Test one of the examples provided

###Examples###

####Simulation Case 0. Basic Features####
During the working day the visitors are entering the restaurant at random intervals and immideatly get the table.
The inter arrival interval is the random variable with uniform distribution from 0 to 70 minutes.
The last visitor gets admitted not later than 8 hours after the opening.
The simulation itself is terminated when the last visitors enters the restaurant.
```go
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
// any type of the Runner should be defined as struct
// with the *godes.Runner as anonimous field
type Visitor struct {
	*godes.Runner
	number int
}

var visitorsCount int = 0

func (vst Visitor) Run() { // Any runner should have the Run method
	fmt.Printf("Visitor %v arrives at time= %6.3f \n", vst.number, godes.GetSystemTime())

}
func main() {
	var shutdown_time float64 = 8 * 60
	for {
		//godes.Stime is the current simulation time
		if godes.GetSystemTime() < shutdown_time {
			//the function acivates the Runner
			godes.ActivateRunner(Visitor{&godes.Runner{}, visitorsCount})
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
Results
Visitor 0 arrives at time=  0.000 
Visitor 1 arrives at time= 62.761 
Visitor 2 arrives at time= 89.380 
Visitor 3 arrives at time= 133.868 
Visitor 4 arrives at time= 189.023 
Visitor 5 arrives at time= 229.752 
Visitor 6 arrives at time= 291.620 
Visitor 7 arrives at time= 334.445 
Visitor 8 arrives at time= 358.918 
Visitor 9 arrives at time= 381.318 
Visitor 10 arrives at time= 424.361 
Visitor 11 arrives at time= 446.308 
```
####Simulation Case 1.  Boolean Controls####
The restaurant has only one table to sit on. During the working day the visitors are entering the restaurant at random intervals
and wait for the table to be available. The inter arrival interval is the random variable with uniform distribution from 0 to 70 minutes.
The time spent in the restaurant is the random variable with uniform distribution from 10 to 60 minutes.
The last visitor gets admitted not later than 8 hours after the opening.
The simulation itself is terminated when the last visitors has left the restaurant.
```go
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
	for {
		//godes.GetSystemTime() is the current simulation time
		if godes.GetSystemTime() < shutdown_time {
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
Results
Visitor 0 arrives at time=  0.000 
Visitor 0 gets the table at time=  0.000 
Visitor 0 leaves at time= 59.984 
Visitor 1 arrives at time= 69.978 
Visitor 1 gets the table at time= 69.978 
Visitor 1 leaves at time= 129.261 
Visitor 2 arrives at time= 138.974 
Visitor 2 gets the table at time= 138.974 
Visitor 3 arrives at time= 158.395 
Visitor 2 leaves at time= 162.846 
Visitor 3 gets the table at time= 162.846 
Visitor 4 arrives at time= 172.806 
Visitor 5 arrives at time= 180.486 
Visitor 3 leaves at time= 183.140 
Visitor 4 gets the table at time= 183.140 
Visitor 4 leaves at time= 198.625 
Visitor 5 gets the table at time= 198.625 
Visitor 6 arrives at time= 228.787 
Visitor 5 leaves at time= 243.126 
Visitor 6 gets the table at time= 243.126 
Visitor 7 arrives at time= 275.738 
Visitor 6 leaves at time= 286.662 
Visitor 7 gets the table at time= 286.662 
Visitor 8 arrives at time= 316.774 
Visitor 7 leaves at time= 325.974 
Visitor 8 gets the table at time= 325.974 
Visitor 9 arrives at time= 366.170 
Visitor 8 leaves at time= 371.257 
Visitor 9 gets the table at time= 371.257 
Visitor 10 arrives at time= 377.635 
Visitor 11 arrives at time= 385.511 
Visitor 9 leaves at time= 389.446 
Visitor 10 gets the table at time= 389.446 
Visitor 10 leaves at time= 405.072 
Visitor 11 gets the table at time= 405.072 
Visitor 12 arrives at time= 430.294 
Visitor 11 leaves at time= 447.059 
Visitor 12 gets the table at time= 447.059 
Visitor 13 arrives at time= 447.493 
Visitor 12 leaves at time= 469.345 
Visitor 13 gets the table at time= 469.345 
Visitor 13 leaves at time= 519.771 
```
####Simulation Case 2.  Queues####
During the four working hours the visitors are entering the restaurant at random intervals and form the arrival queue. 
The inter arrival interval is the random variable with uniform distribution from 0 to 30 minutes. The restaurant employs two waiters who are servicing one visitor in a time. The service time  is the random variable with uniform distribution from 10 to 60 minutes. 
The simulation itself is terminated when 
* Simulation time passes the four hours 
* Both waiters have finished servicing  
* There is no visitors in the arrival queue. 

The model  calculates the average (arithmetic mean) of  the visitors waiting time
```go
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
var visitorArrivalQueue *godes.FIFOQueue = godes.NewFIFOQueue()

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

func (waiter Waiter) Run() {

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
		godes.ActivateRunner(Waiter{&godes.Runner{}, i})
	}
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

Results
Visitor 0 arrives at time=  0.000 
Visitor 0 is invited by waiter 0 at  0.000 
Visitor 0 arrives at time= 20.748 
Visitor 1 is invited by waiter 1 at 20.748 
Visitor 0 arrives at time= 43.495 
Visitor 0 leaves at= 44.579 
Visitor 2 is invited by waiter 0 at 44.579 
Visitor 0 arrives at time= 44.983 
Visitor 2 leaves at= 57.059 
Visitor 3 is invited by waiter 0 at 57.059 
Visitor 0 arrives at time= 60.895 
Visitor 1 leaves at= 68.661 
Visitor 4 is invited by waiter 1 at 68.661 
Visitor 0 arrives at time= 90.616 
Visitor 0 arrives at time= 91.911 
Visitor 3 leaves at= 93.579 
Visitor 5 is invited by waiter 0 at 93.579 
Visitor 5 leaves at= 105.738 
Visitor 6 is invited by waiter 0 at 105.738 
Visitor 0 arrives at time= 109.690 
Visitor 0 arrives at time= 125.536 
Visitor 4 leaves at= 128.195 
Visitor 7 is invited by waiter 1 at 128.195 
Visitor 6 leaves at= 145.370 
Visitor 8 is invited by waiter 0 at 145.370 
Visitor 0 arrives at time= 155.007 
Visitor 0 arrives at time= 163.835 
Visitor 7 leaves at= 164.604 
Visitor 9 is invited by waiter 1 at 164.604 
Visitor 0 arrives at time= 165.978 
Visitor 9 leaves at= 189.317 
Visitor 10 is invited by waiter 1 at 189.317 
Visitor 0 arrives at time= 194.411 
Visitor 10 leaves at= 202.889 
Visitor 11 is invited by waiter 1 at 202.889 
Visitor 8 leaves at= 204.489 
Visitor 12 is invited by waiter 0 at 204.489 
Visitor 0 arrives at time= 217.891 
Visitor 12 leaves at= 253.621 
Visitor 13 is invited by waiter 0 at 253.621 
Visitor 11 leaves at= 260.278 
Waiter  1 ends the work at 260.278 
Visitor 13 leaves at= 312.358 
Waiter  0 ends the work at 312.358 
Average Waiting Time 13.847 
```
####Simulation Case 3.  Multiple Runs####

```go
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
var visitorArrivalQueue *godes.FIFOQueue = godes.NewFIFOQueue()

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

func (waiter Waiter) Run() {
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
			godes.ActivateRunner(Waiter{&godes.Runner{}, i})
		}
		for {
			//godes.Stime is the current simulation time

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
		fmt.Printf(" Run # %v Average Waiting Time %6.3f  \n", runs, visitorArrivalQueue.GetAverageTime())
		//clear after each run
		arrival.Clear()
		service.Clear()
		waitersSwt.Clear()
		visitorArrivalQueue.Clear()
		godes.Clear()

	}
}

Results
 Run # 0 Average Waiting Time 17.461  
 Run # 1 Average Waiting Time 22.264  
 Run # 2 Average Waiting Time 14.501  
 Run # 3 Average Waiting Time 20.446  
 Run # 4 Average Waiting Time 11.195  
```

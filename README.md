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
* Code Security - the Godes includes the  source code for the model,  Go is an open source project supported by Google

###Installation###
*	Download, install and test the Go at your machine. See instructions at http://golang.org/doc/install
* Optionally install one of the free Go IDE (i.e.LiteIDE X)
* Download Godes package
* Test one of the examples provided

###Examples###

####Simulation Case 0.Basic Features####
During the working day the visitors are entering the restaurant at random intervals and immideatly get the table.
The inter arrival interval is the random variable with uniform distribution from 0 to 70 minutes.
The last visitor gets admitted not later than 8 hours after the opening.
The simulation itself is terminated when the last visitors enters the restaurant.

```go
package main

import (
	"fmt"
	"godes"
)

// the arrival is random number generators for the uniform  distribution
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

Results
Visitor 0 arrives at time=  0.000 
Visitor 1 arrives at time=  6.279 
Visitor 2 arrives at time= 32.193 
Visitor 3 arrives at time= 66.223 
Visitor 4 arrives at time= 122.069 
Visitor 5 arrives at time= 159.754 
Visitor 6 arrives at time= 166.704 
Visitor 7 arrives at time= 189.620 
Visitor 8 arrives at time= 193.194 
Visitor 9 arrives at time= 241.867 
Visitor 10 arrives at time= 252.249 
Visitor 11 arrives at time= 297.207 
Visitor 12 arrives at time= 332.281 
Visitor 13 arrives at time= 341.150 
Visitor 14 arrives at time= 354.308 
Visitor 15 arrives at time= 416.638 
Visitor 16 arrives at time= 449.534 
Visitor 17 arrives at time= 475.263 
```
####Simulation Case 1. Boolean Controls####
The restaurant has only one table to sit on. During the working day the visitors are entering the restaurant at random intervals
and wait for the table to be available. The inter arrival interval is the random variable with uniform distribution from 0 to 70 minutes.
The time spent in the restaurant is the random variable with uniform distribution from 10 to 60 minutes.
The last visitor gets admitted not later than 8 hours after the opening.
The simulation itself is terminated when the last visitors has left the restaurant.
```go
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
	fmt.Printf("Visitor %v arrives at time= %6.3f \n", vst.number, godes.Stime)
	tableBusy.Wait(false) // this will wait till the tableBusy control becomes false
	tableBusy.Set(true)   // sets the tableBusy control to true - the table is busy
	fmt.Printf("Visitor %v gets the table at time= %6.3f \n", vst.number, godes.Stime)
	godes.Advance(service.Get(10, 60)) //the function advance the simulation time by the value in the argument
	tableBusy.Set(false)               // sets the tableBusy control to false - the table is idle
	fmt.Printf("Visitor %v leaves at time= %6.3f \n", vst.number, godes.Stime)
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
Results
Visitor 0 arrives at time=  0.000 
Visitor 0 gets the table at time=  0.000 
Visitor 0 leaves at time= 36.055 
Visitor 1 arrives at time= 36.477 
Visitor 1 gets the table at time= 36.477 
Visitor 1 leaves at time= 92.411 
Visitor 2 arrives at time= 100.784 
Visitor 2 gets the table at time= 100.784 
Visitor 3 arrives at time= 103.355 
Visitor 2 leaves at time= 112.620 
Visitor 3 gets the table at time= 112.620 
Visitor 4 arrives at time= 168.231 
Visitor 3 leaves at time= 168.960 
Visitor 4 gets the table at time= 168.960 
Visitor 4 leaves at time= 218.210 
Visitor 5 arrives at time= 223.180 
Visitor 5 gets the table at time= 223.180 
Visitor 6 arrives at time= 229.387 
Visitor 5 leaves at time= 237.614 
Visitor 6 gets the table at time= 237.614 
Visitor 7 arrives at time= 242.663 
Visitor 6 leaves at time= 257.096 
Visitor 7 gets the table at time= 257.096 
Visitor 8 arrives at time= 287.359 
Visitor 9 arrives at time= 293.332 
Visitor 7 leaves at time= 299.023 
Visitor 8 gets the table at time= 299.023 
Visitor 8 leaves at time= 313.288 
Visitor 9 gets the table at time= 313.288 
Visitor 10 arrives at time= 317.668 
Visitor 9 leaves at time= 340.672 
Visitor 10 gets the table at time= 340.672 
Visitor 11 arrives at time= 368.852 
Visitor 10 leaves at time= 387.232 
Visitor 11 gets the table at time= 387.232 
Visitor 12 arrives at time= 391.438 
Visitor 11 leaves at time= 413.365 
Visitor 12 gets the table at time= 413.365 
Visitor 13 arrives at time= 459.622 
Visitor 12 leaves at time= 472.067 
Visitor 13 gets the table at time= 472.067 
Visitor 13 leaves at time= 497.555 
```
####Simulation Case 2. Queues####
==============================
Now the restorant has only four waiters and visitors are making a FIFO Queue at the entrance. We need to find an average waiting time.
```go
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

Results
Visitor 0 arrives at time= 16.402 
Visitor 0 is invited by waiter 0 at time= 16.402 
Visitor 1 arrives at time= 29.638 
Visitor 1 is invited by waiter 1 at time= 29.638 
Visitor 2 arrives at time= 43.687 
Visitor 2 is invited by waiter 2 at time= 43.687 
Visitor 0 leaves at time= 53.739 
Visitor 3 arrives at time= 60.254 
Visitor 3 is invited by waiter 0 at time= 60.254 
Visitor 1 leaves at time= 61.698 
Visitor 2 leaves at time= 77.103 
Visitor 4 arrives at time= 89.303 
Visitor 4 is invited by waiter 1 at time= 89.303 
Visitor 3 leaves at time= 97.865 
Visitor 5 arrives at time= 99.497 
Visitor 5 is invited by waiter 0 at time= 99.497 
Visitor 6 arrives at time= 113.335 
Visitor 6 is invited by waiter 2 at time= 113.335 
Visitor 5 leaves at time= 126.486 
Visitor 7 arrives at time= 129.948 
Visitor 7 is invited by waiter 0 at time= 129.948 
Visitor 8 arrives at time= 139.667 
Visitor 9 arrives at time= 143.825 
Visitor 6 leaves at time= 146.399 
Visitor 8 is invited by waiter 2 at time= 146.399 
Visitor 4 leaves at time= 147.719 
Visitor 9 is invited by waiter 1 at time= 147.719 
Visitor 10 arrives at time= 163.542 
Visitor 11 arrives at time= 164.230 
Visitor 9 leaves at time= 164.649 
Visitor 10 is invited by waiter 1 at time= 164.649 
Visitor 7 leaves at time= 167.635 
Visitor 11 is invited by waiter 0 at time= 167.635 
Visitor 8 leaves at time= 172.598 
Visitor 11 leaves at time= 178.781 
Visitor 12 arrives at time= 180.497 
Visitor 12 is invited by waiter 2 at time= 180.497 
Visitor 13 arrives at time= 201.593 
Visitor 13 is invited by waiter 0 at time= 201.593 
Visitor 10 leaves at time= 207.511 
Visitor 14 arrives at time= 212.479 
Visitor 14 is invited by waiter 1 at time= 212.479 
Visitor 12 leaves at time= 217.610 
Visitor 15 arrives at time= 231.223 
Visitor 15 is invited by waiter 2 at time= 231.223 
Visitor 14 leaves at time= 240.623 
Visitor 16 arrives at time= 245.584 
Visitor 16 is invited by waiter 1 at time= 245.584 
Visitor 13 leaves at time= 246.752 
Visitor 17 arrives at time= 254.449 
Visitor 17 is invited by waiter 0 at time= 254.449 
Visitor 15 leaves at time= 272.463 
Visitor 18 arrives at time= 272.882 
Visitor 18 is invited by waiter 2 at time= 272.882 
Visitor 19 arrives at time= 273.244 
Visitor 17 leaves at time= 279.225 
Visitor 19 is invited by waiter 0 at time= 279.225 
Visitor 16 leaves at time= 279.518 
Visitor 20 arrives at time= 283.950 
Visitor 20 is invited by waiter 1 at time= 283.950 
Visitor 19 leaves at time= 289.828 
Visitor 20 leaves at time= 311.793 
Visitor 18 leaves at time= 313.603 
Visitor 21 arrives at time= 313.738 
Visitor 21 is invited by waiter 0 at time= 313.738 
Visitor 22 arrives at time= 338.856 
Visitor 22 is invited by waiter 1 at time= 338.856 
Visitor 23 arrives at time= 349.553 
Visitor 23 is invited by waiter 2 at time= 349.553 
Visitor 24 arrives at time= 356.131 
Visitor 21 leaves at time= 373.384 
Visitor 24 is invited by waiter 0 at time= 373.384 
Visitor 23 leaves at time= 377.381 
Visitor 25 arrives at time= 379.688 
Visitor 25 is invited by waiter 2 at time= 379.688 
Visitor 22 leaves at time= 390.720 
Visitor 24 leaves at time= 394.348 
Visitor 26 arrives at time= 400.944 
Visitor 26 is invited by waiter 1 at time= 400.944 
Visitor 27 arrives at time= 409.036 
Visitor 27 is invited by waiter 0 at time= 409.036 
Visitor 28 arrives at time= 427.995 
Visitor 25 leaves at time= 428.949 
Visitor 28 is invited by waiter 2 at time= 428.949 
Visitor 29 arrives at time= 430.793 
Visitor 27 leaves at time= 432.523 
Visitor 29 is invited by waiter 0 at time= 432.523 
Visitor 26 leaves at time= 446.370 
Visitor 29 leaves at time= 447.186 
Visitor 30 arrives at time= 451.491 
Visitor 30 is invited by waiter 1 at time= 451.491 
Visitor 31 arrives at time= 468.019 
Visitor 31 is invited by waiter 0 at time= 468.019 
Visitor 28 leaves at time= 470.548 
Visitor 32 arrives at time= 476.519 
Visitor 32 is invited by waiter 2 at time= 476.519 
Visitor 30 leaves at time= 495.989 
Visitor 33 arrives at time= 499.624 
Visitor 33 is invited by waiter 1 at time= 499.624 
Visitor 32 leaves at time= 500.685 
Visitor 31 leaves at time= 505.565 
Visitor 33 leaves at time= 548.133 
Average Waiting Time 1.2075535094620022  
```

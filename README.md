#Godes#

Libriary to Build Discrete Event Simulation Models in Go (http://golang.org/)

Copyright (c) 2013, 2015 Alex Goussiatiner agoussia@yahoo.com

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

####Example 0. Covers: Basic Features####
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

Results
 0.000  	 Visitor # 0 arrives
 65.490 	 Visitor # 1 arrives
 77.122 	 Visitor # 2 arrives
 100.740 	 Visitor # 3 arrives
 156.033 	 Visitor # 4 arrives
 196.230 	 Visitor # 5 arrives
 238.047 	 Visitor # 6 arrives
 299.696 	 Visitor # 7 arrives
 321.409 	 Visitor # 8 arrives
 336.237 	 Visitor # 9 arrives
 379.230 	 Visitor # 10 arrives
 439.952 	 Visitor # 11 arrives
 467.317 	 Visitor # 12 arrives
```
***

####Example 1.  Covers:  Boolean Controls####
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

Results
0.000  	 	Visitor 0 arrives
0.000  	 	Visitor 0 gets the table
33.668 	 	Visitor 0 leaves
64.166 	 	Visitor 1 arrives
64.166 	 	Visitor 1 gets the table
71.445 	 	Visitor 2 arrives
103.401 	 Visitor 1 leaves
103.401 	 Visitor 2 gets the table
121.797 	 Visitor 3 arrives
144.534 	 Visitor 2 leaves
144.534 	 Visitor 3 gets the table
150.740 	 Visitor 4 arrives
184.333 	 Visitor 3 leaves
184.333 	 Visitor 4 gets the table
197.497 	 Visitor 4 leaves
212.729 	 Visitor 5 arrives
212.729 	 Visitor 5 gets the table
264.648 	 Visitor 6 arrives
272.160 	 Visitor 5 leaves
272.160 	 Visitor 6 gets the table
294.490 	 Visitor 7 arrives
321.236 	 Visitor 6 leaves
321.236 	 Visitor 7 gets the table
321.964 	 Visitor 8 arrives
353.377 	 Visitor 7 leaves
353.377 	 Visitor 8 gets the table
354.954 	 Visitor 9 arrives
360.047 	 Visitor 10 arrives
361.397 	 Visitor 11 arrives
405.146 	 Visitor 8 leaves
405.146 	 Visitor 9 gets the table
429.129 	 Visitor 12 arrives
453.826 	 Visitor 9 leaves
453.826 	 Visitor 12 gets the table
476.401 	 Visitor 12 leaves
476.401 	 Visitor 10 gets the table
488.002 	 Visitor 10 leaves
488.002 	 Visitor 11 gets the table
501.840 	 Visitor 11 leaves
```
***

####Example 2.  Covers:  Queues####
During the four working hours the visitors are entering the restaurant at random intervals and form the arrival queue.
The inter arrival interval is the random variable with uniform distribution from 0 to 30 minutes. The restaurant employs two waiters who are servicing one visitor in a time. The service time  is the random variable with uniform distribution from 10 to 60 minutes.
The simulation itself is terminated when
* Simulation time passes the four hours
* Both waiters have finished servicing  
* There is no visitors in the arrival queue.

The model  calculates the average (arithmetic mean) of  the visitors waiting time
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
Results
0.000  	 	Visitor 0 arrives
0.000  	 	Visitor 0 is invited by waiter 0  
5.333  	 	Visitor 1 arrives
5.333  	 	Visitor 1 is invited by waiter 1  
19.893 	 	Visitor 1 leaves
26.236 	 	Visitor 2 arrives
26.236 	 	Visitor 2 is invited by waiter 1  
46.823 	 	Visitor 0 leaves
52.189 	 	Visitor 3 arrives
52.189 	 	Visitor 3 is invited by waiter 0  
61.310 	 	Visitor 2 leaves
64.796 	 	Visitor 4 arrives
64.796 	 	Visitor 4 is invited by waiter 1  
73.180 	 	Visitor 5 arrives
76.037 	 	Visitor 6 arrives
89.228 	 	Visitor 7 arrives
92.261 	 	Visitor 3 leaves
92.261 	 	Visitor 5 is invited by waiter 0  
96.226 	 	Visitor 8 arrives
109.172 	 Visitor 4 leaves
109.172 	 Visitor 6 is invited by waiter 1  
118.855 	 Visitor 5 leaves
118.855 	 Visitor 7 is invited by waiter 0  
121.444 	 Visitor 9 arrives
141.457 	 Visitor 10 arrives
149.288 	 Visitor 7 leaves
149.288 	 Visitor 8 is invited by waiter 0  
151.889 	 Visitor 11 arrives
160.821 	 Visitor 6 leaves
160.821 	 Visitor 9 is invited by waiter 1  
169.579 	 Visitor 8 leaves
169.579 	 Visitor 10 is invited by waiter 0  
173.890 	 Visitor 12 arrives
187.193 	 Visitor 13 arrives
189.220 	 Visitor 9 leaves
189.220 	 Visitor 11 is invited by waiter 1  
198.068 	 Visitor 14 arrives
208.362 	 Visitor 10 leaves
208.362 	 Visitor 12 is invited by waiter 0  
211.847 	 Visitor 15 arrives
221.005 	 Visitor 16 arrives
229.585 	 Visitor 17 arrives
230.239 	 Visitor 18 arrives
232.081 	 Visitor 11 leaves
232.081 	 Visitor 13 is invited by waiter 1  
263.779 	 Visitor 12 leaves
263.779 	 Visitor 14 is invited by waiter 0  
264.924 	 Visitor 13 leaves
264.924 	 Visitor 15 is invited by waiter 1  
277.866 	 Visitor 15 leaves
277.866 	 Visitor 16 is invited by waiter 1  
303.579 	 Visitor 16 leaves
303.579 	 Visitor 17 is invited by waiter 1  
310.750 	 Visitor 14 leaves
310.750 	 Visitor 18 is invited by waiter 0  
325.254 	 Visitor 17 leaves
325.254 	 Waiter  1 ends the work
351.524 	 Visitor 18 leaves
351.524 	 Waiter  0 ends the work
Average Waiting Time 34.171  
```
***

####Example 3.  Covers: Multiple Runs####
The same simulation scenario as for the example 3 repeats five times
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
		fmt.Printf(" Run # %v  %v  \n", runs, visitorArrivalQueue)
		//clear after each run
		arrival.Clear()
		service.Clear()
		waitersSwt.Clear()
		visitorArrivalQueue.Clear()
		godes.Clear()

	}
}


Results
 Run # 0   Average Time=67.804   
 Run # 1   Average Time=58.378   
 Run # 2   Average Time=25.189   
 Run # 3   Average Time=13.909   
 Run # 4   Average Time=50.269   
```
***

####Example 4.  Machine Shop (Covers: Interrupt and Resume) ####
A workshop has *n* identical machines. A stream of jobs (enough to
keep the machines busy) arrives. Each machine breaks down
periodically. Repairs are carried out by one repairman.
The repairman continues them when he is done
with the machine repair. The workshop works continuously.

```go
package main

import (
	"fmt"
	"godes"
)

const PT_MEAN = 10.0          //	Avg. processing time in minutes
const PT_SIGMA = 2.0          //	Sigma of processing time
const MTTF = 300.0            // 	Mean time to failure in minutes
const REPAIR_TIME = 30.0      //	Time it takes to repair a machine in minutes
const REPAIR_TIME_SIGMA = 1.0 //	Sigma of repair time

const NUM_MACHINES = 10
const SHUT_DOWN_TIME = 4 * 7 * 24 * 60

// random generator for the processing time - normal distribution
var processingGen *godes.NormalDistr = godes.NewNormalDistr()

// random generator for the  time   until the next failure for a machine - exponential distribution
var breaksGen *godes.ExpDistr = godes.NewExpDistr()

// true when repairman is available for carrying a repair
var repairManAvailableSwt *godes.BooleanControl = godes.NewBooleanControl()

type Machine struct {
	*godes.Runner
	partsCount int
	number     int
}

func (machine *Machine) Run() {
	for {
		godes.Advance(processingGen.Get(PT_MEAN, PT_SIGMA))
		machine.partsCount = machine.partsCount + 1
		if godes.GetSystemTime() > SHUT_DOWN_TIME {
			break
		}
	}
}

type MachineRepair struct {
	*godes.Runner
	machine *Machine
}

func (machineRepair *MachineRepair) Run() {
	machine := machineRepair.machine
	for {
		godes.Advance(breaksGen.Get(1 / MTTF))
		if machine.GetState() == godes.RUNNER_STATE_SCHEDULED {
			breakTime := godes.GetSystemTime()
			//interrupt machine
			godes.Interrupt(machine)
			repairManAvailableSwt.Wait(true)
			//repair
			repairManAvailableSwt.Set(false)
			godes.Advance(processingGen.Get(REPAIR_TIME, REPAIR_TIME_SIGMA))
			//release repairman
			repairManAvailableSwt.Set(true)
			//resume machine and change the scheduled time to compensate the idle time
			godes.Resume(machine, godes.GetSystemTime()-breakTime)
		}

		if godes.GetSystemTime() > SHUT_DOWN_TIME {
			break
		}
	}
}

func main() {
	var m *Machine
	x := make(map[int]*Machine)
	for i := 0; i < NUM_MACHINES; i++ {
		m = &Machine{&godes.Runner{}, 0, i}
		godes.AddRunner(m)
		godes.AddRunner(&MachineRepair{&godes.Runner{}, m})
		x[i] = m
	}
	repairManAvailableSwt.Set(true)
	godes.Run()
	godes.WaitUntilDone()
	//print results
	for i := 0; i < NUM_MACHINES; i++ {
		m = x[i]
		fmt.Printf(" Machine # %v %v \n", m.number, m.partsCount)
	}
}

Results
Machine # 0 3690
 Machine # 1 3661
 Machine # 2 3653
 Machine # 3 3622
 Machine # 4 3658
 Machine # 5 3628
 Machine # 6 3670
 Machine # 7 3521
 Machine # 8 3610
 Machine # 9 3659
```
***

####Example 5.  Bank Renege (Covers: Wait with timeout) ####
This example models a bank counter and customers arriving at random times. Each customer has a certain patience. It waits to get to the counter until she’s at the end of her tether. If she gets to the counter, she uses it for a while before releasing it.

```go
package main

import (
	"fmt"
	"godes"
)

const NEW_CUSTOMERS = 5          // Total number of customers
const INTERVAL_CUSTOMERS = 12.00 // Generate new customers roughly every x minites
const SERVICE_TIME = 12.0
const MIN_PATIENCE = 1 // Min. customer patience
const MAX_PATIENCE = 3 // Max. customer patience

// random generator for the arrival interval - expovariate distribution
var arrivalGen *godes.ExpDistr = godes.NewExpDistr()

// random generator for the patience time time - uniform distribution
var patienceGen *godes.UniformDistr = godes.NewUniformDistr()

// random generator for the  service time - expovariate distribution
var serviceGen *godes.ExpDistr = godes.NewExpDistr()

// true when Counter
var counterAvailable *godes.BooleanControl = godes.NewBooleanControl()

type Customer struct {
	*godes.Runner
	name int
}

func (customer *Customer) Run() {

	arrivalTime := godes.GetSystemTime()
	patience := patienceGen.Get(MIN_PATIENCE, MAX_PATIENCE)
	fmt.Printf("  %6.3f  Customer %v : Here I am   My patience=%6.3f  \n", godes.GetSystemTime(), customer.name, patience)

	counterAvailable.WaitAndTimeout(true, patience)
	if !counterAvailable.GetState() {
		fmt.Printf("  %6.3f  Customer %v : RENEGED after  %6.3f \n", godes.GetSystemTime(), customer.name, godes.GetSystemTime()-arrivalTime)
	} else {
		counterAvailable.Set(false)

		fmt.Printf("  %6.3f  Customer %v : Waited %6.3f \n", godes.GetSystemTime(), customer.name, godes.GetSystemTime()-arrivalTime)
		godes.Advance(serviceGen.Get(1 / SERVICE_TIME))
		fmt.Printf("  %6.3f  Customer %v : Finished \n", godes.GetSystemTime(), customer.name)
		counterAvailable.Set(true)

	}

}

func main() {
	counterAvailable.Set(true)
	godes.Run()
	for i := 0; i < NEW_CUSTOMERS; i++ {
		godes.AddRunner(&Customer{&godes.Runner{}, i})
		godes.Advance(arrivalGen.Get(1 / INTERVAL_CUSTOMERS))
	}

	godes.WaitUntilDone()

}


Results
  0.000  	Customer 0 : Here I am   My patience= 1.208  
  0.000   	Customer 0 : Waited  0.000
  4.135  	Customer 0 : Finished
 13.514  	Customer 1 : Here I am   My patience= 2.473  
 13.514  	Customer 1 : Waited  0.000
 18.172  	Customer 2 : Here I am   My patience= 1.980  
 19.749  	Customer 1 : Finished
 19.749  	Customer 2 : Waited  1.576
 24.204  	Customer 3 : Here I am   My patience= 2.795  
 26.999  	Customer 3 : RENEGED after   2.795
 27.836  	Customer 2 : Finished
 40.964  	Customer 4 : Here I am   My patience= 2.090  
 40.964  	Customer 4 : Waited  0.000
 42.849 	Customer 4 : Finished
 ```
 ***
####Example 6.  Bank - Single Run (Covers: FIFO Queue, Parallel Resources, Collection and processing of statistics) ####
Procces Description: A bank employs three tellers and the customers form a queue for all three tellers. The doors of the bank close after eight hours.
The simulation is ended when the last customer has been served.

Task:Execute single simulation run, calculate Average, Standard Deviation,
confidence intervall lower and upper bounds,minimum	 and Maximum for the
following performance measures: total elapsed time, queue length, queueing time
service time.

Model Features:
*FIFO Queue*
The customer object is placed in the FIFO arrival queue as soon as the customer is created.

*Parallel Resources*
The application constructs Tellers object to model tellers as a set of resources.
The object 'provides' tellers to the customer located in the Queue head and "releases" the teller when customer is serviced.
Maximum 3 tellers can be provided simultaneously.
The interlocking between catching request is performed using godes BooleanControl object.

*Collection and processing of statistics*
While finishing a customer run  the application creates data arrays for each measure. At the end of simulation, the application creates StatCollection object and performs descriptive statistical analysis. The following statistical parameters are calculated for each measure array:
	#Observ - number of observations
	Average - average (mean) value
	Std Dev- standard deviation
	L-Bound-lower bound of the confidence interval  with 95% probability
	U-Bound-upper bound of the confidence interval  with 95% probability
	Minimum- minimum value
	Maximum- maximum value
	```go
	// Copyright 2015 Alex Goussiatiner. All rights reserved.
	// Use of this source code is governed by a MIT
	// license that can be found in the LICENSE file.
	package main
	/*
	Procces Description:
	====================
	A bank employs three tellers and the customers form a queue for all three tellers.
	The doors of the bank close after eight hours.
	The simulation is ended when the last customer has been served.
	*/

	import (
		"fmt"
		"godes"
	)
	//Input Parameters
	const (
		ARRIVAL_INTERVAL = 0.5
		SERVICE_TIME = 1.3
		SHUTDOWN_TIME = 8 * 60.
	)
	// the arrival and service are two random number generators for the exponential  distribution
	var arrival *godes.ExpDistr = godes.NewExpDistr()
	var service *godes.ExpDistr = godes.NewExpDistr()
	// true when any counter is available
	var counterSwt *godes.BooleanControl = godes.NewBooleanControl()
	// FIFO Queue for the arrived customers
	var customerArrivalQueue *godes.FIFOQueue = godes.NewFIFOQueue("0")


	var tellers *Tellers
	var measures [][]float64
	var titles = []string{
		"Elapsed Time",
		"Queue Length",
		"Queueing Time",
		"Service Time",
	}

	var availableTellers int = 0;
	// the Tellers is a Passive Object represebting resource
	type Tellers struct {
		max     int
	}

	func (tellers *Tellers) Catch(customer *Customer) {
		for {
			counterSwt.Wait(true)
			if customerArrivalQueue.GetHead().(*Customer).GetId() == customer.GetId() {
				break
			} else {
				godes.Yield()
			}
		}
		availableTellers++
		if availableTellers == tellers.max {
			counterSwt.Set(false)
		}
	}

	func (tellers *Tellers) Release() {
		availableTellers--
		counterSwt.Set(true)
	}

	// the Customer is a Runner
	type Customer struct {
		*godes.Runner
		id int
	}

	func (customer *Customer) Run() {
		a0 := godes.GetSystemTime()
		tellers.Catch(customer)
		a1 := godes.GetSystemTime()
		customerArrivalQueue.Get()
		qlength := float64(customerArrivalQueue.Len())
		godes.Advance(service.Get(1. / SERVICE_TIME))
		a2 := godes.GetSystemTime()
		tellers.Release()
		collectionArray := []float64{a2 - a0, qlength, a1 - a0, a2 - a1}
		measures = append(measures, collectionArray)
	}
	func main() {
		measures = [][]float64{}
		tellers = &Tellers{3}
		godes.Run()
		counterSwt.Set(true)
		count := 0
		for {
			customer := &Customer{&godes.Runner{}, count}
			customerArrivalQueue.Place(customer)
			godes.AddRunner(customer)
			godes.Advance(arrival.Get(1. / ARRIVAL_INTERVAL))
			if godes.GetSystemTime() > SHUTDOWN_TIME {
				break
			}
			count++
		}
		godes.WaitUntilDone() // waits for all the runners to finish the Run()
		collector := godes.NewStatCollector(titles, measures)
		collector.PrintStat()
		fmt.Printf("Finished \n")
	}
	Variable		#	Average	Std Dev	L-Bound	U-Bound	Minimum	Maximum
	Elapsed Time	999	 3.965	 3.700	 3.736	 4.195	 0.002	17.343
	Queue Length	999	 5.248	 7.573	 4.779	 5.718	 0.000	30.000
	Queueing Time	999	 2.647	 3.467	 2.432	 2.862	 0.000	14.122
	Service Time	999	 1.319	 1.295	 1.238	 1.399	 0.001	 8.617
	Finished
```
####Example 7.  Bank - Multiple Runs (Covers: FIFO Queue, Parallel Resources, Collection and processing of statistics) ####
Procces Description: A bank employs three tellers and the customers form a queue for all three tellers. The doors of the bank close after eight hours.
The simulation is ended when the last customer has been served.

	```go
	// Copyright 2015 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package main

/*
Procces Description:
===================
A bank employs three tellers and the customers form a queue for all three tellers.
The doors of the bank close after eight hours.
The simulation is ended when the last customer has been served.

Task
====
Execute multiple simulation runs, calculate Average, Standard Deviation, 
confidence intervall lower and upper bounds,minimum	 and Maximum for the
following performance measures: 
	total elapsed time, 
	queue length,
	queueing time
	service time.

Model Features:
===============
1. FIFO Queue
The customer object is placed in the FIFO arrival queue as soon as the customer is created.

2. Parallel Resources
The application constructs Tellers object to model tellers as a set of resources.
The object 'provides' tellers to the customer located in the Queue head and "releases" the teller when customer is serviced.
Maximum 3 tellers can be provided simultaneously.
The interlocking between catching request is performed using godes BooleanControl object.

3. Collection and processing of statistics
While finishing a customer run  the application creates data arrays for each measure. At the end of simulation, the application creates StatCollection object and performs descriptive statistical analysis. The following statistical parameters are calculated for each measure array:
	#Observ - number of observations
	Average - average (mean) value
	Std Dev- standard deviation
	L-Bound-lower bound of the confidence interval  with 95% probability
	U-Bound-upper bound of the confidence interval  with 95% probability
	Minimum- minimum value
	Maximum- maximum value
*/

import (
	"fmt"
	"godes"

)

//Input Parameters
const (
	ARRIVAL_INTERVAL = 0.5
	SERVICE_TIME     = 1.3
	SHUTDOWN_TIME    = 8 * 60.
	INDEPENDENT_RUNS = 100
)

// the arrival and service are two random number generators for the exponential  distribution
var arrival *godes.ExpDistr = godes.NewExpDistr()
var service *godes.ExpDistr = godes.NewExpDistr()

// true when any counter is available
var counterSwt *godes.BooleanControl = godes.NewBooleanControl()

// FIFO Queue for the arrived customers
var customerArrivalQueue *godes.FIFOQueue = godes.NewFIFOQueue("0")

var tellers *Tellers
var statistics [][]float64
var replicationStats [][]float64
var titles = []string{
	"Elapsed Time",
	"Queue Length",
	"Queueing Time",
	"Service Time",
}

var availableTellers int = 0

// the Tellers is a Passive Object represebting resource
type Tellers struct {
	max int
}

func (tellers *Tellers) Catch(customer *Customer) {
	for {
		counterSwt.Wait(true)
		if customerArrivalQueue.GetHead().(*Customer).GetId() == customer.GetId() {
			break
		} else {
			godes.Yield()
		}
	}
	availableTellers++
	if availableTellers == tellers.max {
		counterSwt.Set(false)
	}
}

func (tellers *Tellers) Release() {
	availableTellers--
	counterSwt.Set(true)
}

// the Customer is a Runner
type Customer struct {
	*godes.Runner
	id int
}

func (customer *Customer) Run() {
	a0 := godes.GetSystemTime()
	tellers.Catch(customer)
	a1 := godes.GetSystemTime()
	customerArrivalQueue.Get()
	qlength := float64(customerArrivalQueue.Len())
	godes.Advance(service.Get(1. / SERVICE_TIME))
	a2 := godes.GetSystemTime()
	tellers.Release()
	collectionArray := []float64{a2 - a0, qlength, a1 - a0, a2 - a1}
	replicationStats = append(replicationStats, collectionArray)
}
func main() {
	statistics = [][]float64{}

	tellers = &Tellers{3}
	for i := 0; i < INDEPENDENT_RUNS; i++ {
		replicationStats = [][]float64{}
		godes.Run()
		counterSwt.Set(true)
		customerArrivalQueue.Clear()
		count := 0
		for {
			customer := &Customer{&godes.Runner{}, count}
			customerArrivalQueue.Place(customer)
			godes.AddRunner(customer)
			godes.Advance(arrival.Get(1. / ARRIVAL_INTERVAL))
			if godes.GetSystemTime() > SHUTDOWN_TIME {
				break
			}
			count++
		}
		godes.WaitUntilDone() // waits for all the runners to finish the Run()
		godes.Clear()
		replicationCollector := godes.NewStatCollector(titles, replicationStats)

		collectionArray := []float64{
			replicationCollector.GetAverage(0),
			replicationCollector.GetAverage(1),
			replicationCollector.GetAverage(2),
			replicationCollector.GetAverage(3),
		}
		statistics = append(statistics, collectionArray)
	}

	collector := godes.NewStatCollector(titles, statistics)
	collector.PrintStat()
	fmt.Printf("Finished \n")
}
Variable		#	Average	Std Dev	L-Bound	U-Bound	Minimum	Maximum
Elapsed Time	100	 3.787	 1.699	 3.454	 4.120	 2.013	15.081
Queue Length	100	 4.983	 3.505	 4.296	 5.670	 1.463	27.886
Queueing Time	100	 2.489	 1.672	 2.161	 2.816	 0.795	13.655
Service Time	100	 1.299	 0.044	 1.290	 1.307	 1.202	 1.425
```
## Godes

Open Source Library to Build Discrete Event Simulation Models in Go/Golang (http://golang.org/)

Copyright (c) 2013-2015 Alex Goussiatiner agoussia@yahoo.com

### Features
Godes is the general-purpose simulation library which includes the  simulation engine  and building blocks for modeling a wide variety of systems at varying levels of details.

###### Active Objects
All active objects shall implement the RunnerInterface and have Run() method. For each active object Godes creates a goroutine - lightweight thread.

###### Random Generators
Godes contains set of built-in functions for generating random numbers for commonly used probability distributions.
Each of the distrubutions in Godes has one or more parameter values associated with it: Uniform (Min, Max), Normal (Mean and Standard Deviation), Exponential (Lambda), Triangular(Min, Mode, Max)

###### Queues
Godes implements operations with FIFO and LIFO queues

###### BooleanControl
Godes uses BooleanControl variable as a lock for
synchronizing execution of multiple runners

###### StatCollector
The Object calculates and prints statistical parameters for set of samples collected during the simulation.



### Library Docs
[![GoDoc](https://godoc.org/github.com/agoussia/godes?status.svg)](https://godoc.org/github.com/agoussia/godes)

### Advantages
* Godes is easy to learn for the people familiar with the Go and the elementary simulation concept.
* Godes model executes fast  as Go compiles to machine code. Its performace is similar to C++ in performance.
* Godes model is multiplatform as Go compiler targets the Linux, Mac OS X, FreeBSD, Microsoft Windows,etc.
* Godes model can be embedded in various computer systems and over the network.
* Speed of the Godes model compilation is high.
* Variety of the IDE with debuggers are available for Go and Godes as well.
* The Godes sumulation model can use all of the GO's features and libraries.
* Code Security - the Godes includes the  source code for the library and Go is an open source project supported by Google.
* Godes is free open source software under MIT license.

### Installation

```
$ go get github.com/agoussia/godes
```

### Examples

#### Example 0. Restaurant.Godes Basics

###### Proces Description
During the working day the visitors are entering the restaurant at random intervals and immediately get the table.
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
```
***

#### Example 1. Restaurant. Godes Boolean Controls
###### Proces Description
The restaurant has only one table to sit on. During the working day the visitors are entering the restaurant at random intervals
and wait for the table to be available. The inter arrival interval is the random variable with uniform distribution from 0 to 70 minutes.The time spent in the restaurant is the random variable with uniform distribution from 10 to 60 minutes.
The last visitor gets admitted not later than 8 hours after the opening.
The simulation itself is terminated when the last visitors has left the restaurant.
```go
package main

import (
	"fmt"
	"godes"
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
107.468  Visitor 3 arrives 
146.824  Visitor 2 leaves 
146.824  Visitor 3 gets the table 
149.471  Visitor 4 arrives 
171.623  Visitor 3 leaves 
171.623  Visitor 4 gets the table 
187.234  Visitor 4 leaves 
207.523  Visitor 5 arrives 
207.523  Visitor 5 gets the table 
230.922  Visitor 6 arrives 
245.859  Visitor 5 leaves 
245.859  Visitor 6 gets the table 
261.770  Visitor 7 arrives 
269.668  Visitor 8 arrives 
272.368  Visitor 6 leaves 
272.368  Visitor 7 gets the table 
290.484  Visitor 7 leaves 
290.484  Visitor 8 gets the table 
310.261  Visitor 9 arrives 
333.570  Visitor 8 leaves 
333.570  Visitor 9 gets the table 
338.323  Visitor 10 arrives 
354.874  Visitor 9 leaves 
354.874  Visitor 10 gets the table 
393.826  Visitor 10 leaves 
397.720  Visitor 11 arrives 
397.720  Visitor 11 gets the table 
409.123  Visitor 12 arrives 
436.817  Visitor 13 arrives 
447.731  Visitor 14 arrives 
455.705  Visitor 11 leaves 
455.705  Visitor 13 gets the table 
482.955  Visitor 13 leaves 
482.955  Visitor 12 gets the table 
496.034  Visitor 12 leaves 
496.034  Visitor 14 gets the table 
555.822  Visitor 14 leaves 
*/
```
***

#### Example 2.  Restaurant. Godes Queues
###### Proces Description
During the four working hours the visitors are entering the restaurant at random intervals and form the arrival queue. 
The inter arrival interval is the random variable with uniform distribution from 0 to 30 minutes. The restaurant employs two waiters who are servicing one visitor in a time. The service time  is the random variable with uniform distribution from 10 to 60 minutes. 
The simulation itself is terminated when 
* Simulation time passes the four hours 
* Both waiters have finished servicing  
* There are no visitors in the arrival queue. 

The model  calculates the average (arithmetic mean) of  the visitors waiting time
```go
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
13.374 		Visitor 0 leaves 
16.066 	 	Visitor 1 arrives 
16.066 		Visitor 1 is invited by waiter 1  
39.137 	 	Visitor 1 leaves 
42.316 		Visitor 2 arrives 
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
112.187 	Visitor 7 arrives 
115.572 	Visitor 8 arrives 
125.475 	Visitor 6 leaves 
125.475 	Visitor 7 is invited by waiter 1  
127.275 	Visitor 5 leaves 
127.275 	Visitor 8 is invited by waiter 0  
132.969 	Visitor 9 arrives 
143.591 	Visitor 7 leaves 
143.591 	Visitor 9 is invited by waiter 1  
144.995 	Visitor 10 arrives 
164.895 	Visitor 9 leaves 
164.895 	Visitor 10 is invited by waiter 1  
170.361 	Visitor 8 leaves 
170.451 	Visitor 11 arrives 
170.451 	Visitor 11 is invited by waiter 0  
175.338 	Visitor 12 arrives 
187.207 	Visitor 13 arrives 
191.885 	Visitor 14 arrives 
203.848 	Visitor 10 leaves 
203.848 	Visitor 12 is invited by waiter 1  
213.596 	Visitor 15 arrives 
228.436 	Visitor 11 leaves 
228.436 	Visitor 13 is invited by waiter 0  
231.098 	Visitor 12 leaves 
231.098 	Visitor 14 is invited by waiter 1  
231.769 	Visitor 16 arrives 
241.515 	Visitor 13 leaves 
241.515 	Visitor 15 is invited by waiter 0  
287.864 	Visitor 15 leaves 
287.864 	Visitor 16 is invited by waiter 0  
290.886 	Visitor 14 leaves 
290.886 	Waiter  1 ends the work 
330.903 	Visitor 16 leaves 
330.903 	Waiter  0 ends the work 
Average Waiting Time 15.016  
*/
```
***

#### Example 3. Restaurant. Multiple Runs
###### Proces Description
This is the same process as in Example 2. Simulation is repeated 5 times.
```go
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

```
***
#### Example 4.  Machine Shop. Godes Interrupt and Resume Feature.
###### Proces Description
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
var processingGen *godes.NormalDistr = godes.NewNormalDistr(true)

// random generator for the  time   until the next failure for a machine - exponential distribution
var breaksGen *godes.ExpDistr = godes.NewExpDistr(true)

// true when repairman is available for carrying a repair
var repairManAvailableSwt *godes.BooleanControl = godes.NewBooleanControl()

type Machine struct {
	*godes.Runner
	partsCount int
	number     int
	finished   bool
}

func (machine *Machine) Run() {
	for {
		godes.Advance(processingGen.Get(PT_MEAN, PT_SIGMA))
		machine.partsCount++
		if godes.GetSystemTime() > SHUT_DOWN_TIME {
			machine.finished = true
			break
		}

	}
	fmt.Printf(" Machine # %v %v \n", machine.number, machine.partsCount)
}

type MachineRepair struct {
	*godes.Runner
	machine *Machine
}

func (machineRepair *MachineRepair) Run() {
	machine := machineRepair.machine
	for {
		godes.Advance(breaksGen.Get(1 / MTTF))
		if machine.finished {
			break
		}

		interrupted := godes.GetSystemTime()
		godes.Interrupt(machine)
		repairManAvailableSwt.Wait(true)
		if machine.finished {
			break
		}
		repairManAvailableSwt.Set(false)
		godes.Advance(processingGen.Get(REPAIR_TIME, REPAIR_TIME_SIGMA))
		if machine.finished {
			break
		}
		//release repairman
		repairManAvailableSwt.Set(true)
		//resume machine and change the scheduled time to compensate delay
		godes.Resume(machine, godes.GetSystemTime()-interrupted)

	}

}

func main() {

	godes.Run()
	repairManAvailableSwt.Set(true)
	var m *Machine
	for i := 0; i < NUM_MACHINES; i++ {
		m = &Machine{&godes.Runner{}, 0, i, false}
		godes.AddRunner(m)
		godes.AddRunner(&MachineRepair{&godes.Runner{}, m})
	}
	godes.WaitUntilDone()
}
/* OUTPUT
 Machine # 1 3382 
 Machine # 7 3255 
 Machine # 4 3343 
 Machine # 5 3336 
 Machine # 6 3248 
 Machine # 0 3369 
 Machine # 2 3378 
 Machine # 9 3342 
 Machine # 8 3362 
 Machine # 3 3213 
*/
```
***

#### Example 5.  Bank Counter. Godes Wait with Timeout Feature.
###### Proces Description
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
var arrivalGen *godes.ExpDistr = godes.NewExpDistr(true)

// random generator for the patience time time - uniform distribution
var patienceGen *godes.UniformDistr = godes.NewUniformDistr(true)

// random generator for the  service time - expovariate distribution
var serviceGen *godes.ExpDistr = godes.NewExpDistr(true)

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
		fmt.Printf("  %6.3f  Customer %v : Reneged after  %6.3f \n", godes.GetSystemTime(), customer.name, godes.GetSystemTime()-arrivalTime)
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
/* OUTPUT
0.000  	Customer 0 : Here I am   My patience= 1.135  
0.000  	Customer 0 : Waited  0.000 
5.536  	Customer 0 : Finished 
11.141  Customer 1 : Here I am   My patience= 1.523  
11.141  Customer 1 : Waited  0.000 
34.482  Customer 2 : Here I am   My patience= 2.523  
36.123  Customer 1 : Finished 
36.123  Customer 2 : Waited  1.641 
38.869  Customer 2 : Finished 
39.943  Customer 3 : Here I am   My patience= 1.592  
39.943  Customer 3 : Waited  0.000 
49.113  Customer 4 : Here I am   My patience= 1.224  
50.337  Customer 4 : Reneged after   1.224 
59.948  Customer 3 : Finished 
*/
```
***
#### Example 6. Bank. Sngle Run, FIFO Queue, Parallel Resources, StatCollector
###### Proces Description
A bank employs three tellers and the customers form a queue for all three tellers. The doors of the bank close after eight hours. The simulation is ended when the last customer has been served.
###### Task
Execute single simulation run, calculate average, standard deviation,
confidence interval, lower and upper bounds, minimum and maximum values for the
following performance measures: total elapsed time, queue length, queueing time
service time.
###### Model Features
* **FIFO Queue.** The customer object is placed in the FIFO arrival queue as soon as the customer is created.
* **Parallel Resources.** The application constructs Tellers object to model tellers as a set of resources.
The object 'provides' tellers to the customer located in the Queue head and "releases" the teller when customer is serviced.
Maximum 3 tellers can be provided simultaneously. The interlocking between catching request is performed using godes BooleanControl object.
* **Collection and processing of statistics.** While finishing a customer run  the application creates data arrays for each measure. At the end of simulation, the application creates StatCollector object and performs descriptive statistical analysis. The following statistical parameters are calculated for each measure array:
	Observ - number of observations, Average - average (mean) value, Std Dev- standard deviation, L-Bound-lower bound of the confidence interval  with 95% probability, U-Bound-upper bound of the confidence interval  with 95% probability,
	Minimum value,Maximum value
```go
package main

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
var arrival *godes.ExpDistr = godes.NewExpDistr(true)
var service *godes.ExpDistr = godes.NewExpDistr(true)
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
		if customerArrivalQueue.GetHead().(*Customer).id == customer.id {
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
/* OUTPUT
Variable		#	Average	Std Dev	L-Bound	U-Bound	Minimum	Maximum
Elapsed Time	944	 2.591	 1.959	 2.466	 2.716	 0.005	11.189
Queue Length	944	 2.411	 3.069	 2.215	 2.607	 0.000	13.000
Queueing Time	944	 1.293	 1.533	 1.195	 1.391	 0.000	 6.994
Service Time	944	 1.298	 1.247	 1.219	 1.378	 0.003	 7.824
*/
```
#### Example 7.  Bank.  Multiple Runs, FIFO Queue, Parallel Resources, StatCollector
###### Procces Description
See example 6.
###### Task
Execute multiple simulation runs, calculate Average, Standard Deviation, 
confidence intervall lower and upper bounds,minimu and maximum for the
following performance measures: total elapsed time, queue length, queueing time, service time.
```go
package main
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
var arrival *godes.ExpDistr = godes.NewExpDistr(true)
var service *godes.ExpDistr = godes.NewExpDistr(true)

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

func (customer *Customer) GetId() int{
	return customer.id
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
/* OUTPUT
Variable	#	Average	Std Dev	L-Bound	U-Bound	Minimum	Maximum
Elapsed Time	100	 3.672	 1.217	 3.433	 3.910	 1.980	 8.722
Queue Length	100	 4.684	 2.484	 4.197	 5.171	 1.539	14.615
Queueing Time	100	 2.368	 1.194	 2.134	 2.602	 0.810	 7.350
Service Time	100	 1.304	 0.044	 1.295	 1.312	 1.170	 1.432
Finished 
*/
```

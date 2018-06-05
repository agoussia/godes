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

	"github.com/godes"
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

func (customer *Customer) GetId() int {
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
Variable		#	Average	Std Dev	L-Bound	U-Bound	Minimum	Maximum
Elapsed Time	100	 3.672	 1.217	 3.433	 3.910	 1.980	 8.722
Queue Length	100	 4.684	 2.484	 4.197	 5.171	 1.539	14.615
Queueing Time	100	 2.368	 1.194	 2.134	 2.602	 0.810	 7.350
Service Time	100	 1.304	 0.044	 1.295	 1.312	 1.170	 1.432
Finished
*/

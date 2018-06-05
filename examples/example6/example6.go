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
	"github.com/agoussia/godes"
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
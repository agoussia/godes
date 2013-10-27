// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
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

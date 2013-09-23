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

####Simulation Case 0.Using Basic Features####
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





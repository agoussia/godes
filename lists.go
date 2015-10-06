// Copyright 2015 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
//
// Godes  is the general-purpose simulation library
// which includes the  simulation engine  and building blocks
// for modeling a wide variety of systems at varying levels of details.
//
// Godes contains set of private functions for managing runners

package godes

import (
	"container/list"
	"fmt"
)

/*
MovingList
This list is sorted in descending order according to the value of the ball priority attribute.
Balls with identical priority values are sorted according to the FIFO principle.

			|
			|	|
			|	|
	     	|	|	|
	<-----	|	|	|	|
*/
func (mdl *model) addToMovingList(runner RunnerInterface) bool {

	if mdl.DEBUG {
		fmt.Printf("addToMovingList %v\n", runner)
	}

	if mdl.movingList == nil {
		mdl.movingList = list.New()
		mdl.movingList.PushFront(runner)
		return true
	}

	insertedSwt := false
	for element := mdl.movingList.Front(); element != nil; element = element.Next() {
		if runner.getPriority() > element.Value.(RunnerInterface).getPriority() {
			mdl.movingList.InsertBefore(runner, element)
			insertedSwt = true
			break
		}
	}
	if !insertedSwt {
		mdl.movingList.PushBack(runner)
	}
	return true
}

func (mdl *model) getFromMovingList() RunnerInterface {

	if mdl.movingList == nil {
		panic("MovingList was not initilized")
	}
	if mdl.DEBUG {
		runner := mdl.movingList.Front().Value.(RunnerInterface)
		fmt.Printf("getFromMovingList %v\n", runner)
	}
	runner := mdl.movingList.Front().Value.(RunnerInterface)
	mdl.movingList.Remove(mdl.movingList.Front())
	return runner

}

func (mdl *model) removeFromMovingList(runner RunnerInterface) {

	if mdl.movingList == nil {
		panic("MovingList was not initilized")
	}

	if mdl.DEBUG {
		fmt.Printf("removeFromMovingList %v\n", runner)
	}
	var found bool
	for e := mdl.movingList.Front(); e != nil; e = e.Next() {
		if e.Value == runner {
			mdl.movingList.Remove(e)
			found = true
			break
		}
	}

	if !found {
		//panic("not found in MovingList")
	}
}

/*
SchedulledList
This list is sorted in descending order according to the schedulled time
Priorites are not used

			|
			|	|
			|	|
			|	|	|
			|	|	|	|--->

*/
func (mdl *model) addToSchedulledList(runner RunnerInterface) bool {

	if mdl.scheduledList == nil {
		mdl.scheduledList = list.New()
		mdl.scheduledList.PushFront(runner)
		return true
	}
	insertedSwt := false
	for element := mdl.scheduledList.Back(); element != nil; element = element.Prev() {
		if runner.getMovingTime() < element.Value.(RunnerInterface).getMovingTime() {
			mdl.scheduledList.InsertAfter(runner, element)
			insertedSwt = true
			break
		}
	}
	if !insertedSwt {
		mdl.scheduledList.PushFront(runner)
	}

	if mdl.DEBUG {
		fmt.Println("===")
		fmt.Printf("addToSchedulledList %v\n", runner)
		for element := mdl.scheduledList.Front(); element != nil; element = element.Next() {
			fmt.Printf("elem %v\n", element.Value.(RunnerInterface))
		}
		fmt.Println("===")
	}
	return true
}

func (mdl *model) getFromSchedulledList() RunnerInterface {
	if mdl.scheduledList == nil {
		panic(" SchedulledList was not initilized")
	}
	if mdl.DEBUG {
		runner := mdl.scheduledList.Back().Value.(RunnerInterface)
		fmt.Printf("getFromSchedulledList %v\n", runner)
	}
	runner := mdl.scheduledList.Back().Value.(RunnerInterface)
	mdl.scheduledList.Remove(mdl.scheduledList.Back())
	return runner
}

func (mdl *model) removeFromSchedulledList(runner RunnerInterface) {
	if mdl.scheduledList == nil {
		panic("schedulledList was not initilized")
	}
	if modl.DEBUG {
		fmt.Printf("removeFrom schedulledListt %v\n", runner)
	}
	var found bool
	for e := mdl.scheduledList.Front(); e != nil; e = e.Next() {
		if e.Value == runner {
			mdl.scheduledList.Remove(e)
			found = true
			break
		}
	}
	if !found {
		panic("not found in scheduledList")
	}
	return
}

func (mdl *model) addToWaitingConditionMap(runner RunnerInterface) bool {

	if runner.getWaitingForBoolControl() == nil {
		panic(" addToWaitingConditionMap - no control ")
	}

	if mdl.DEBUG {
		fmt.Printf("addToWaitingConditionMap %v\n", runner)
	}

	if mdl.waitingConditionMap == nil {
		mdl.waitingConditionMap = make(map[int]RunnerInterface)

	}
	mdl.waitingConditionMap[runner.getInternalId()] = runner
	return true
}

func (mdl *model) addToInterruptedMap(runner RunnerInterface) bool {

	if mdl.DEBUG {
		fmt.Printf("addToInterruptedMap %v\n", runner)
	}
	if mdl.interruptedMap == nil {
		mdl.interruptedMap = make(map[int]RunnerInterface)
	}

	mdl.interruptedMap[runner.getInternalId()] = runner
	return true
}

func (mdl *model) removeFromInterruptedMap(runner RunnerInterface) bool {

	if mdl.DEBUG {
		fmt.Printf("removeFromInterruptedMap %v\n", runner)
	}

	_, ok := mdl.interruptedMap[runner.getInternalId()]

	if ok {
		delete(mdl.interruptedMap, runner.getInternalId())
	} else {
		panic("not found in interruptedMap")
	}
	return true
}

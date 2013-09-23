// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package godes

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type Model struct {
	mu                  sync.RWMutex
	activeRunner        RunnerInterface
	movingList          *list.List
	scheduledList       *list.List
	waitingList         *list.List
	waitingConditionMap map[int]RunnerInterface
	interruptedList     *list.List
	terminatedList      *list.List
	currentId           int
	controlChannel      chan int
	simulationActive    bool
	DEBUG               bool
}

//Initilization Of the Main Ball
func NewModel(verbose bool) *Model {

	var ball *Runner = NewRunner()
	ball.channel = make(chan int)
	ball.markTime = time.Now()
	ball.id = 0
	ball.state = RUNNER_STATE_ACTIVE //that is bypassing READY
	ball.priority = 100
	ball.setMarkTime(time.Now())
	var runner RunnerInterface = ball
	mdl := Model{activeRunner: runner, controlChannel: make(chan int), DEBUG: verbose}
	mdl.addToMovingList(runner)

	return &mdl
}

func (mdl *Model) advance(interval float64) bool {

	ch := mdl.activeRunner.getChannel()
	mdl.activeRunner.setMovingTime(Stime + interval)
	mdl.activeRunner.setState(RUNNER_STATE_SCHEDULED)
	mdl.removeFromMovingList(mdl.activeRunner)
	mdl.addToSchedulledList(mdl.activeRunner)
	//restart control channel and freez
	mdl.controlChannel <- TRANSITION_ADVANCE
	<-ch
	return true
}

func (mdl *Model) waitUntillDone() {

	if mdl.activeRunner.GetId() != 0 {
		panic("waitUntillDone initiated for not main ball")
	}

	mdl.removeFromMovingList(mdl.activeRunner)
	mdl.controlChannel <- 100
	for {

		if !model.simulationActive {
			break
		} else {
			if mdl.DEBUG {
				fmt.Println("waiting", mdl.movingList.Len())
			}
			time.Sleep(time.Millisecond * simulationSecondScale)
		}
	}
}

func (mdl *Model) activate(runner RunnerInterface) bool {

	mdl.currentId++
	runner.setChannel(make(chan int))
	runner.setMovingTime(Stime)
	runner.setId(mdl.currentId)
	runner.setState(RUNNER_STATE_READY)
	mdl.addToMovingList(runner)

	go func() {
		<-runner.getChannel()
		runner.setMarkTime(time.Now())
		runner.Run()
		if mdl.activeRunner == nil {
			panic("remove: activeRunner == nil")
		}
		mdl.removeFromMovingList(mdl.activeRunner)
		mdl.activeRunner.setState(RUNNER_STATE_TERMINATED)
		mdl.activeRunner = nil
		mdl.controlChannel <- 100
	}()
	return true

}

func (mdl *Model) booleanControlWait(b *BooleanControl, val bool) {

	ch := mdl.activeRunner.getChannel()
	if mdl.activeRunner == nil {
		panic("booleanControlWait - no runner")
	}

	mdl.removeFromMovingList(mdl.activeRunner)

	mdl.activeRunner.setState(RUNNER_STATE_WAITING_COND)
	mdl.activeRunner.setWaitingForBool(val)
	mdl.activeRunner.setWaitingForBoolControl(b)

	mdl.addToWaitingConditionMap(mdl.activeRunner)
	mdl.controlChannel <- 100
	<-ch

}

func (mdl *Model) booleanControlSet(b *BooleanControl) {
	ch := mdl.activeRunner.getChannel()
	if mdl.activeRunner == nil {
		panic("booleanControlSet - no runner")
	}
	mdl.controlChannel <- 100
	<-ch

}

func (mdl *Model) control() bool {

	if mdl.activeRunner == nil {
		panic("control: activeBall == nil")
	}

	go func() {
		var runner RunnerInterface
		for {
			<-mdl.controlChannel
			if mdl.waitingConditionMap != nil && len(mdl.waitingConditionMap) > 0 {
				for key, temp := range mdl.waitingConditionMap {
					if temp.getWaitingForBoolControl() == nil {
						panic("  no BoolControl")
					}
					if temp.getWaitingForBool() == temp.getWaitingForBoolControl().getState() {
						temp.setState(RUNNER_STATE_READY)
						mdl.addToMovingList(temp)
						delete(mdl.waitingConditionMap, key)
						break
					}
				}
			}

			//finding new runner
			runner = nil
			if mdl.movingList != nil && mdl.movingList.Len() > 0 {
				runner = mdl.getFromMovingList()
			}
			if runner == nil && mdl.scheduledList != nil && mdl.scheduledList.Len() > 0 {
				runner = mdl.getFromSchedulledList()
				if runner.GetMovingTime() < Stime {
					panic("control is seting simulation time in the past")
				} else {
					Stime = runner.GetMovingTime()
				}
				mdl.addToMovingList(runner)
			}
			if runner == nil {
				break
			}
			//restarting
			mdl.activeRunner = runner
			mdl.activeRunner.setState(RUNNER_STATE_ACTIVE)
			runner.setWaitingForBoolControl(nil)
			mdl.activeRunner.getChannel() <- -1

		}
		if mdl.DEBUG {
			fmt.Println("Finished")
		}
		mdl.simulationActive = false
	}()

	return true

}

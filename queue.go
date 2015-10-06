// Copyright 2015 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
//
// Godes  is the general-purpose simulation library
// which includes the  simulation engine  and building blocks
// for modeling a wide variety of systems at varying levels of details.
//



package godes

import (
	"container/list"
)

// Queue represents a FIFO or LIFO queue
type Queue struct {
	id      string
	fifo    bool
	sumTime float64
	count   int64
	qList   *list.List
	qTime   *list.List
	startTime float64
}

// FIFOQueue represents a FIFO queue
type FIFOQueue struct {
	Queue
}

// LIFOQueue represents a LIFO queue
type LIFOQueue struct {
	Queue
}

//GetAverageTime is average elapsed time for an object in the queue
func (q *Queue) GetAverageTime() float64 {
	return q.sumTime / float64(q.count)
}

//Len returns number of objects in the queue
func (q *Queue) Len() int {
	return q.qList.Len()
}

//GetAverageTime is average elapsed time for an object in the queue
func (q *Queue) GetAverageNumber() float64 {
	return q.sumTime / (stime-q.startTime)
}

//Place adds an object to the queue
func (q *Queue) Place(entity interface{}) {
	q.qList.PushFront(entity)
	q.qTime.PushFront(stime)
	if(q.startTime==0){
		q.startTime=stime
	}
}

// Get returns an object and removes it from the queue
func (q *Queue) Get() interface{} {

	var entity interface{}
	var timeIn float64
	if q.fifo {
		entity = q.qList.Back().Value
		timeIn = q.qTime.Back().Value.(float64)
		q.qList.Remove(q.qList.Back())
		q.qTime.Remove(q.qTime.Back())
	} else {
		entity = q.qList.Front().Value
		timeIn = q.qTime.Front().Value.(float64)
		q.qList.Remove(q.qList.Front())
		q.qTime.Remove(q.qTime.Front())
	}

	q.sumTime = q.sumTime + stime - timeIn
	q.count++

	return entity
}

// GetHead returns the head object (doesn't remove it from the queue)
func (q *Queue) GetHead() interface{} {
	var entity interface{}
	if q.fifo {
		entity = q.qList.Back().Value
	} else {
		entity = q.qList.Front().Value
	}
	return entity
}

// NewFIFOQueue itializes the FIFO queue
func NewFIFOQueue(mid string) *FIFOQueue {
	return &FIFOQueue{Queue{fifo: true, id: mid, qList: list.New(), qTime: list.New()}}
}

// NewLIFOQueue itializes the LIFO queue
func NewLIFOQueue(mid string) *LIFOQueue {
	return &LIFOQueue{Queue{fifo: false, id: mid, qList: list.New(), qTime: list.New()}}
}

//Clear reinitiates the queue
func (q *Queue) Clear() {
	q.sumTime = 0
	q.count = 0
	q.qList.Init()
	q.qTime.Init()
}

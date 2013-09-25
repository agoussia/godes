// Copyright 2013 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package godes

import (
	"container/list"
)

type Queue struct {
	fifo    bool
	sumTime float64
	count   int64
	qList   *list.List
	qTime   *list.List
}

type FIFOQueue struct {
	Queue
}

type LIFOQueue struct {
	Queue
}

func (q *Queue) GetAverageTime() float64 {
	return q.sumTime / float64(q.count)
}

func (q *Queue) Len() int {
	return q.qList.Len()
}

func (q *Queue) Place(entity interface{}) {
	q.qList.PushFront(entity)
	q.qTime.PushFront(stime)
}

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

func NewFIFOQueue() *FIFOQueue {
	return &FIFOQueue{Queue{fifo: true, qList: list.New(), qTime: list.New()}}
}

func NewLIFOQueue() *LIFOQueue {
	return &LIFOQueue{Queue{fifo: false, qList: list.New(), qTime: list.New()}}
}

func (q *Queue) Clear() {
	q.sumTime = 0
	q.count = 0
	q.qList.Init()
	q.qTime.Init()

}

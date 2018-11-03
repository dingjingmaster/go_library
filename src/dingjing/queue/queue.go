package queue

import (
	"library/go-logging"
	"sync"
)

type mlistNode struct {
	next  *mlistNode
	pre   *mlistNode
	value interface{}
}

type Queue struct {
	head, tail *mlistNode
	mutex      sync.Mutex
	size       int
	list       *mlistNode
	log        *logging.Logger
}

func New() *Queue {
	queue := new(Queue)
	queue.head = nil
	queue.tail = nil
	queue.mutex = sync.Mutex{}
	queue.size = 0
	queue.list = nil
	queue.log = logging.MustGetLogger("queue")

	return queue
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Push(e interface{}) {
	q.mutex.Lock()

	ele := new(mlistNode)
	ele.next = nil
	ele.pre = nil
	ele.value = e
	/* 第一次插入数据(头插) */
	if q.size <= 0 {
		q.head = ele
		q.tail = ele
	} else {
		ele.next = q.head
		q.head.pre = ele
		q.head = ele
	}
	q.size++
	q.log.Info("新数据加入队列!!!")

	q.mutex.Unlock()
}

func (q *Queue) Pop() interface{} {
	q.mutex.Lock()
	var ele interface{}

	if q.size > 0 {
		ele = q.tail.value
		q.tail = q.tail.pre
		q.size--
	} else {
		ele = nil
	}
	q.log.Info("数据弹出队列!!!")

	q.mutex.Unlock()

	return ele
}

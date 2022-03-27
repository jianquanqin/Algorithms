package BLC

import (
	"errors"
	"fmt"
)

//Test example

//func main() {
//	queue := &Queue{
//		Maxsize: 5,
//		Front:   -1,
//		Rear:    -1,
//	}
//	queue.In(1)
//	queue.In(2)
//	queue.In(3)
//	queue.In(4)
//	queue.In(5)
//	queue.In(6)
//
//	queue.Show()
//
//	val, _ := queue.Out()
//	fmt.Println("the out value is:", val)
//
//	queue.Show()
//
//}

//定义一个队列对象

type Queue struct {
	Maxsize int
	Array   [5]int
	Front   int
	Rear    int
}

//入队

func (queue *Queue) In(val int) (err error) {
	if queue.Rear == queue.Maxsize-1 {
		fmt.Println("queue is full")
		return errors.New("queue is full")
	} else {
		queue.Rear++
		queue.Array[queue.Rear] = val
	}
	return err
}

//显示队列

func (queue *Queue) Show() (val int) {
	for i := queue.Front + 1; i <= queue.Rear; i++ {
		fmt.Printf("array[%d]=%d\n", i, queue.Array[i])
	}
	return val
}

//出队

func (queue *Queue) Out() (val int, err error) {
	if queue.Rear == queue.Front {
		return -1, errors.New("queue empty")
	}
	queue.Front++
	val = queue.Array[queue.Front]
	return val, nil
}

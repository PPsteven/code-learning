package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func genBatch(nums []int, size int, d time.Duration) <-chan []int{
	in := make(chan int)
	out := make(chan []int)

	go func() {
		// 为什么使用 goroutine?
		// 因为 in channel 的管道容量有限，需要边消费边插入。不然管道会被堵住，造成 all goroutine asleep 的 deadlock
		defer close(in)
		for _, num := range nums {
			in <- num
		}
	}()

	go func() {
		defer close(out)
		batch := make([]int, 0)
		full := make(chan struct{}, 1) // buffer
		tick := time.Tick(d)

		for num := range in {
			batch = append(batch, num)
			if len(batch) == size {
				full <- struct{}{}
			}

			select {
			case <- full:
				out <- batch
				batch = nil
			case <- tick:
				out <- batch
				batch = nil
			default:
			}
		}
		if len(batch) > 0 {
			out <- batch
		}
	}()
	return out
}

func slowFuncs(threadNo int, in <-chan []int) error {
	time.Sleep(time.Second * 3)
	for nums := range in {
		fmt.Printf("#%d:%v\n", threadNo, nums)
	}
	return nil
}

func main() {
	nums := []int{0, 1, 8, 2, 4, 5, 4, 10, 3, 5, 7, 19, 20, 3, 1, 0, 22}
	out := genBatch(nums, 4, time.Second * 5)

	g, _ := errgroup.WithContext(context.TODO())
	nThread := 64
	for i := 0; i < nThread; i++ {
		g.Go(func() error {
			return slowFuncs(i, out)
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("g.Wait(): %v\n", err)
	}


	fmt.Printf("finish")
}
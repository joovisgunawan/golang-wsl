package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var num = 0
var numatomic int32 = 0
var mu sync.Mutex     //this is exactly like mutex c++ rtos
var wg sync.WaitGroup //this is like semaphore

func main() {
	fmt.Println("hello world")

	wg.Add(2) //add to key

	go func() {
		defer wg.Done() //take 1 key and once defer, release 1 key
		for i := 1; i < 5; i++ {
			mu.Lock()
			num++
			fmt.Println("Goroutine", num)
			mu.Unlock()
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done() //take 1 key and once defer, relese 1 key
		for i := 1; i < 5; i++ {
			mu.Lock()
			num--
			fmt.Println("Goroutine", num)
			mu.Unlock()
			time.Sleep(500 * time.Millisecond)
		}
	}()

	wg.Wait() //wait for all key to be release, if set to 2, means it should be to again
	fmt.Println("latest value ", num)
	time.Sleep(5 * time.Second)
	testAtomic()

}

func testAtomic() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		atomic.AddInt32(&numatomic, 1)
		fmt.Println("Atomic", atomic.LoadInt32(&numatomic))
		time.Sleep(500 * time.Millisecond)

	}()

	go func() {
		defer wg.Done()
		atomic.AddInt32(&numatomic, -1)
		fmt.Println("Atomic", atomic.LoadInt32(&numatomic))
		time.Sleep(500 * time.Millisecond)
	}()
	wg.Wait()
	fmt.Println("latest value atomic", numatomic)

}

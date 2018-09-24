package basic

import (
	"fmt"
	"sync"
)

func SyncTest() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go incrCountSync()
	}
	wg.Wait()
	fmt.Println(counter)
}

var counter = 0
var mutex sync.Mutex
var wg sync.WaitGroup

func incrCount() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		counter++
	}
}

func incrCountSync() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync/atomic"
	"time"
)

var id int32 = 0

func nextId() int32 {
	return atomic.AddInt32(&id, 1)
}

type MakerThread struct {
	random *rand.Rand
	buff   chan<- []string
	name   string
}

func (thread MakerThread) Start() {
	for {
		// Do any tasks
		time.Sleep(time.Duration(thread.random.Int31()))
		cake := fmt.Sprintf("[ Cake No. %d by %s ]", nextId(), thread.name)
		thread.buff <- []string{thread.name, cake}
		fmt.Printf("%s puts %s\n", thread.name, cake)
	}
}

type EaterThread struct {
	random *rand.Rand
	buff   <-chan []string
	name   string
}

func (thread EaterThread) Start() {
	for pair := range thread.buff {
		fmt.Printf("%s takes %s\n", pair[0], pair[1])
		// Do any tasks
		time.Sleep(time.Duration(thread.random.Int31()))
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	buffer := make(chan []string, 3)

	for i := 0; i < 3; i++ {
		go MakerThread{rand.New(rand.NewSource(time.Now().Unix())), buffer, fmt.Sprintf("MakerThread-%d", i)}.Start()
	}

	for i := 0; i < 3; i++ {
		go EaterThread{rand.New(rand.NewSource(time.Now().Unix())), buffer, fmt.Sprintf("EaterThread-%d", i)}.Start()
	}

	select {}
}

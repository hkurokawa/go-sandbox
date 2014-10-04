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
	buff   chan []string
	name   string
}

func (thread MakerThread) Start() {
	for {
		time.Sleep(time.Duration(thread.random.Int31()))
		cake := fmt.Sprintf("[ Cake No. %d by %s ]", nextId(), thread.name)
		thread.buff <- []string{thread.name, cake}
		fmt.Printf("%s puts %s\n", thread.name, cake)
	}
}

type EaterThread struct {
	random *rand.Rand
	buff   chan []string
	name   string
}

func (thread EaterThread) Start() {
	for {
		pair := <-thread.buff
		fmt.Printf("%s takes %s\n", pair[0], pair[1])
		time.Sleep(time.Duration(thread.random.Int31()))
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	buffer := make(chan []string, 3)

	go MakerThread{random: rand.New(rand.NewSource(time.Now().Unix())), buff: buffer, name: "MakerThread-1"}.Start()
	go MakerThread{random: rand.New(rand.NewSource(time.Now().Unix())), buff: buffer, name: "MakerThread-2"}.Start()
	go MakerThread{random: rand.New(rand.NewSource(time.Now().Unix())), buff: buffer, name: "MakerThread-3"}.Start()

	go EaterThread{random: rand.New(rand.NewSource(time.Now().Unix())), buff: buffer, name: "EaterThread-1"}.Start()
	go EaterThread{random: rand.New(rand.NewSource(time.Now().Unix())), buff: buffer, name: "EaterThread-2"}.Start()
	go EaterThread{random: rand.New(rand.NewSource(time.Now().Unix())), buff: buffer, name: "EaterThread-3"}.Start()

	select {}
}

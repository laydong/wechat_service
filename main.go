package main

import (
	"fmt"
	"sync"
	"time"
)

var path = "./conf/app.json"

//func main() {
//	err := conf.InitConfig(path)
//	if err != nil {
//		panic(err)
//	}
//
//}

var x int64
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
		fmt.Println(x)
		time.Sleep(10)
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

package golang

import (
	"fmt"
	"sync"
	"time"
)

//只执行一次的函数
var once sync.Once

func GoOnce() {

	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("count:", v, "---", i)
	}
	for i := 0; i < 10; i++ {

		go func() {
			once.Do(onced)
			fmt.Println("213")
		}()
	}
	time.Sleep(4000)
}
func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}

//sync.Pool的实现原理和适用场景
//这个类设计的目的是用来保存和复用临时对象，以减少内存分配，降低CG压力
//sync.Pool包在init的时候注册了一个poolCleanup函数，它会清除所有的pool里面所有缓存的对象，该函数注册进去之后会在每次gc之前都会调用，因此sync.Pool缓存的期限只是两次gc之间这段时间。
//正因为如此，我们是不可以使用sync.Pool去实现一个socket连接池的。
/*
func init() {
	runtime_registerPoolCleanup(poolCleanup)
}
*/

func GoPool() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	a := p.Get().(int)
	p.Put(1)
	b := p.Get().(int)

	println(a, b)
}

package golang

import (
	"fmt"
	"runtime"
)

//golang性能监控和资源占用情况的上报
func GoReport() {

	//NumGoroutine 可以返回当前程序的goroutine数目
	println("NumGoroutine = ", runtime.NumGoroutine())

	//GOMAXPROCS 用来设置或查询可以并发执行的goroutine数目，n大于1表示设置GOMAXPROCS值，否则表示查询当前的GOMAXPROCS值

	//设置当前的GOMAXPROCS的值为2
	runtime.GOMAXPROCS(2)

	//获取当前的GOMAXPROCS值
	println("GOMAXPROCS = ", runtime.GOMAXPROCS(0))
}

//使用无缓冲的通道来实现goroutine之间的同步等待
func GoChan() {
	c := make(chan struct{})

	go func(i chan struct{}) {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}

		println(sum)
		//写通道
		c <- struct{}{}
	}(c)

	//NumGoroutine 可以返回当前程序的goroutine数目
	println("NumGoroutine = ", runtime.NumGoroutine())

	//读通道c，通过通道进行同步等待
	<-c
}

//runtime.caller的用法
/*
函数的签名如下

func runtime.Caller(skip int) (pc uintptr, file string, line int, ok bool)

runtime.Caller返回当前的goroutine的栈上的函数调用信息，主要有当前的pc值和调用的文件和行号等信息，
其输入参数skip为要跳过的栈帧数，若为0则表示runtime.Caller的调用者
*/
func GoRunTimeCaller() {
	for skip := 0; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		fmt.Printf("skip = %v, pc = %v, file = %v, line = %v\n", skip, pc, file, line)
	}
}

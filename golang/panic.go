package golang

import (
	"fmt"
	"log"
	"runtime/debug"
)

//panic的处理,打印堆栈信息
func panic() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			fmt.Printf("panic recover! err: %v", err)
			debug.PrintStack()
		}
	}()
}

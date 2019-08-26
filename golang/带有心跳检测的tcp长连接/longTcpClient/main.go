package main

import (
	"fmt"
	"net"
)

var (
	Req_REGISTER byte = 1 // c register cid
	Res_REGISTER byte = 2 // s response

	Req_HEARTBEAT byte = 3 // s send heartbeat req
	Res_HEARTBEAT byte = 4 // c send heartbead res

	Req byte = 5 // cs send data
	Res byte = 6 // cs send ack
)

var Dch chan bool
var Readch chan []byte
var Writech chan []byte

func main() {
	Dch = make(chan bool)
	Readch = make(chan []byte)
	Writech = make(chan []byte)

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:6666")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("连接服务端失败:", err.Error())
		return
	}

	fmt.Println("已连接服务器")

	defer conn.Close()
	go Handler(conn)
	select {
	case <-Dch:
		fmt.Println("关闭连接")
	}
}

func Handler(conn *net.TCPConn) {
	//直到register ok
	data := make([]byte, 128)
	for {
		conn.Write([]byte{Req_REGISTER, '#', '2'})
		conn.Read(data)

		fmt.Println(string(data))
		if data[0] == Res_REGISTER {
			break
		}
	}

	fmt.Println("I 'm register")
	go RHandler(conn)
	go WHandler(conn)
	go Work()
}

func RHandler(conn *net.TCPConn) {
	for {
		//心跳包，回复ack
		data := make([]byte, 128)
		i, _ := conn.Read(data)
		if i == 0 {
			Dch <- true
			return
		}

		if data[0] == Req_HEARTBEAT {
			fmt.Println("recv ht pack")
			conn.Write([]byte{Res_HEARTBEAT, '#', 'h'})
			fmt.Println("send ht pack ack")
		} else if data[0] == Req {
			fmt.Println("recv data pack")
			fmt.Printf("%v\n", string(data[:2]))
			Readch <- data[:2]
			conn.Write([]byte{Res, '#'})
		}
	}
}

func WHandler(conn net.Conn) {
	for {
		select {
		case msg := <-Writech:
			fmt.Println(msg[0])
			fmt.Println("send data after:" + string(msg[1:]))
			conn.Write(msg)
		}
	}
}

func Work() {
	for {
		select {
		case msg := <-Readch:
			fmt.Println("work recv" + string(msg))
			Writech <- []byte{Req, '#', 'x', 'x', 'x', 'x', 'x'}
		}
	}
}

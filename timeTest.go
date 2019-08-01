package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}

// 客户端
func ping() {
	addr := "127.0.0.1:8848"
	conn, err := net.Dial("udp", addr)
	checkError(err)
	defer conn.Close()
	if err != nil {
		os.Exit(1)
	}
	for {
		_, _ = conn.Write([]byte("ping"))
	}
}

func recvPing(conn *net.UDPConn) {
	timeoutDuration := 5 * time.Millisecond
	data := make([]byte, 1024)
	for {
		err := conn.SetReadDeadline(time.Now().Add(time.Duration(timeoutDuration)))
		if err != nil {
			fmt.Printf("set err %s\n", err)
		}
		n, remoteAddr, err := conn.ReadFrom(data)
		if err != nil {
			if e, ok := err.(net.Error); !ok || !e.Timeout() {
				// 非超时错误处理
				fmt.Printf("no timeout\n")
			} else {
				// 收包超时，处理该情况
				fmt.Printf("timeout err\n")
			}
		}
		fmt.Printf("<%s> %s %s\n", remoteAddr, data[:n], time.Now())
		/*
			// 回包
			_, err = conn.WriteToUDP([]byte("world"), remoteAddr)
			if err != nil {
				fmt.Printf(err.Error())
			}
		*/
	}
}

// Server端
func pong() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":8848")
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	defer conn.Close()
	recvPing(conn)
}

func main() {
	// ping client 启动
	go ping()
	// ping server 启动
	go pong() //可以正常执行，使用go pong(),则会出现运行一会停止的情况，不退出也不打印日志

}

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool) //这里保存的是 客户端对应的消息通道
	for {
		select {
		case msg := <-messages: //④
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering: //在步骤③中，新的客户端接入的消息通道，会通过entering传送到这里   ⑤
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)   // 消息输出 ①
	go clientWriter(conn, ch) //消息输出
	who := conn.RemoteAddr().String()
	ch <- "You are " + who        // 消息输出到客户端
	messages <- who + "has login" // 广播消息 ②
	entering <- ch                //把当前client的消息通道，作为通道的元素，发送到登录队列 ③

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	leaving <- ch
	messages <- who + "has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
func main() {
	server, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}

}

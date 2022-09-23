package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "5000"
)

type Client struct {
	conn net.Conn
}

func main() {

	fmt.Println("Start Server... on port ", HOST, ":", PORT)

	listener, err := net.Listen("tcp", HOST+":"+PORT)

	checkErrr(err)

	defer listener.Close()

	for {

		conn, err := listener.Accept()

		checkErrr(err)

		defer conn.Close()

		client := &Client{
			conn: conn,
		}

		fmt.Printf("new connexion from - %s\n", conn.RemoteAddr().String())

		go client.handleIncommingRequest()
	}
}

func (client *Client) handleIncommingRequest() {

	reader := bufio.NewReader(client.conn)
	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
			return
		}

		if string(message) == "exit\n" {
			fmt.Printf("close connecxion from - %s\n", client.conn.RemoteAddr().String())
			client.conn.Close()
		}

		go ConvertVideo(message)
		go TakeScreenOfVideo(message)

		client.conn.Write([]byte("video recived..."))
		fmt.Printf("message received: %s", string(message))
	}
}

func checkErrr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func openServer() {
	// open server
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	defer server.Close()

	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	authenticated := false
	defer func() {
		connection.Close()
		fmt.Println("connection closed")
	}()

	defer func() {
		recover()
	}()

	for {
		data, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("client disconnected")
				break
			}
			fmt.Println("Error reading:", err.Error())
			write(&connection, ERROR+":invalid message")
			break
		}

		var msg Message
		err = json.Unmarshal([]byte(data), &msg)
		if err != nil {
			write(&connection, "Error: invalid message")
			break
		}

		if !authenticated && msg.Method != "auth" {
			write(&connection, ERROR+":not authenticated")
			break
		} else if msg.Method == "auth" {
			arr := strings.Split(msg.Data, ":")
			if len(arr) != 2 {
				write(&connection, ERROR+":invalid credentials")
				break
			}

			user, password := arr[0], arr[1]
			if _, ok := credentials[user]; !ok {
				write(&connection, ERROR+":invalid user")
				break
			}

			if credentials[user] != password {
				write(&connection, ERROR+":invalid password")
				break
			}

			authenticated = true
			write(&connection, OK)
		} else {
			manageDbInteraction(&connection, msg)
		}
	}
}

type Message struct {
	Method string
	Data   string
}

type Kv struct {
	Key   string
	Value string
}

func write(connection *net.Conn, data string) {
	_, err := (*connection).Write([]byte(data + "\n"))
	if err != nil {
		panic(err)
	}
}

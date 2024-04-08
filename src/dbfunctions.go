package main

import (
	"encoding/json"
	"net"
)

const (
	EMPTY = "0"
	OK    = "1"
	ERROR = "2"
)

func manageDbInteraction(connection *net.Conn, message Message) {
	switch message.Method {
	case "get":
		get(connection, message)
	case "set":
		set(connection, message)
	case "rm":
		remove(connection, message)
	default:
		write(connection, ERROR+":invalid method, use get, set or rm")
	}
}

func get(connection *net.Conn, message Message) {
	value, ok := db[message.Data]
	if !ok {
		write(connection, EMPTY)
		return
	}

	write(connection, OK+":"+value)
}

func set(connection *net.Conn, message Message) {
	var kv Kv
	err := json.Unmarshal([]byte(message.Data), &kv)
	if err != nil {
		write(connection, ERROR+":unable to parse data")
		return
	}
	db[kv.Key] = kv.Value
	changes++
	write(connection, OK)
}

func remove(connection *net.Conn, message Message) {
	if _, ok := db[message.Data]; !ok {
		write(connection, EMPTY)
		return
	}

	delete(db, message.Data)
	changes++
	write(connection, OK)
}

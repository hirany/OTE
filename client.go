package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gorilla/websocket"
)

const (
	perm = 755
)

type client struct {
	socket *websocket.Conn
	send   chan []byte
	file   *file
}

func (c *client) read() {
	for {
		_, msg, err := c.socket.ReadMessage()
		if err == nil {
			c.file.forward <- msg
			fmt.Printf("%s\n", msg)
			ioutil.WriteFile("fileData/filename.txt", msg, perm)
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	var err error
	for msg := range c.send {
		if string(msg) != "" {
			err = c.socket.WriteMessage(websocket.TextMessage, msg)
		}
		if err != nil {
			break
		}
	}
	c.socket.Close()
}

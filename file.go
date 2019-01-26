package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

type file struct {
	forward chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
}

func newFile() *file {
	return &file{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (file *file) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP_upgrade")
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		file:   file,
	}
	file.join <- client
	defer func() { file.leave <- client }()
	go client.write()
	client.read()
}

func (file *file) run() {
	for {
		select {
		case client := <-file.join:
			file.clients[client] = true
		case client := <-file.leave:
			delete(file.clients, client)
			close(client.send)
		case msg := <-file.forward:
			for client := range file.clients {
				select {
				case client.send <- msg:
				default:
					delete(file.clients, client)
					close(client.send)
				}
			}
		}
	}
}

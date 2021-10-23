package main

import (
	"chat-server-telnet/internal/server"
)

func main() {
	srv := server.NewServer(":10000")
	err := srv.RunChatServer()

	if err != nil {
		panic(err)
	}
}

package main

import (
	"chat-server-telnet/internal/server"
	"os"
)

func main() {
	srv := server.NewServer(os.Getenv("CHAT_PORT"))
	err := srv.RunChatServer()

	if err != nil {
		panic(err)
	}
}

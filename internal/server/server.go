package server

import (
	"fmt"
	"log"
	"net"
)

// NewServer initializes the chat server
func NewServer(port string) *Server {
	return &Server{
		Port: port,
	}
}

// Server contains port definition
type Server struct {
	Port string
}

type ChatServer interface {
	RunChatServer() error
}

type clients []net.Conn

var connectedClients clients

// RunChatServer starts the chat server
func (s *Server) RunChatServer() error {
	addresses, err := net.ResolveTCPAddr("tcp", s.Port)
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", addresses)
	if err != nil {
		return err
	}

	return s.initializeConnection(listener)
}

func (s *Server) initializeConnection(listener *net.TCPListener) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		connectedClients = append(connectedClients, conn)

		tcpAddr, err := net.ResolveTCPAddr("tcp", conn.LocalAddr().String())
		if err != nil {
			conn.Close()
			return err
		}

		conn.Write([]byte("Connected: " + tcpAddr.IP.String() + "\r\n"))

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	var buf [512]byte
	for {
		n, error := conn.Read(buf[0:])
		if error != nil {
			log.Fatalf(error.Error())
		}
		for _, v := range connectedClients {
			if v.RemoteAddr() != conn.RemoteAddr() {
				v.Write([]byte(buf[0:n]))
			}
		}
		fmt.Printf("%s", string(buf[0:n]))
	}
}

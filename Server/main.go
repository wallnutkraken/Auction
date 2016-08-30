package main

import (
	"Sockets/Server"
	"flag"
	"log"
	"net"
	"os"
	"sync"
)

var (
	logger *log.Logger
	wait   *sync.WaitGroup
)

func main() {
	ip := flag.String("ip", "[::1]", "Ip to listen on (IPv6 only)")
	flag.Parse()
	logfile, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}
	logger = log.New(os.Stdout, "[Sockets]", log.LUTC)
	defer logfile.Close()

	addr, err := net.ResolveTCPAddr("tcp6", *ip+":1632")
	if err != nil {
		logger.Fatalln(err)
	}

	logger.Println("Starting server on", addr.IP, "port", addr.Port)
	server, err := Server.NewServer(addr, "tcp6")
	if err != nil {
		logger.Fatalln(err)
	}

	listenChan := make(chan *net.TCPConn, 64)
	go Acceptance(server, listenChan)
	go Pass(listenChan)
	wait = &sync.WaitGroup{}
	wait.Add(1)
	wait.Wait()
}

func Acceptance(server Server.Server, callback chan *net.TCPConn) {
	for {
		conn, err := server.GetListner().AcceptTCP()
		if err == nil {
			callback <- conn
		} else {
			logger.Println("Refused connection:", err.Error(), "conn =", conn)
		}
	}
	wait.Done()
}

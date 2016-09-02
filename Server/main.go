package main

import (
	"Sockets/Server"
	"flag"
	"log"
	"net"
	"os"

	"github.com/wallnutkraken/Auction/Server/Auction"
)

var (
	logger  *log.Logger
	auction Auction.Auction
)

func main() {
	ip := flag.String("ip", "[::1]", "IP to listen on")
	port := flag.String("port", "1632", "Server running port")
	useIPv4 := flag.Bool("ipv4", false, "Uses IPv4. MUST be provided with a non-default IP arg.")
	flag.Parse()
	protocol := "tcp6"

	if *useIPv4 {
		protocol = "tcp"
	}

	logger = log.New(os.Stdout, "[Sockets]", log.LUTC)
	addr, err := net.ResolveTCPAddr(protocol, *ip+":"+*port)
	if err != nil {
		logger.Fatalln(err)
	}

	logger.Println("Starting server on", addr.IP, "port", addr.Port)
	server, err := Server.NewServer(addr, protocol)
	if err != nil {
		logger.Fatalln(err)
	}

	auction = Auction.NewAuction()
	auction.NewPimp(Auction.NewPimp(100, Auction.NewItem("Artifact")))

	listenChan := make(chan *net.TCPConn, 64)
	go Acceptance(server, listenChan)

	go AddRandomPimps()

	Pass(listenChan)
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
}

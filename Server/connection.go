package main

import (
	"net"

	S "github.com/wallnutkraken/Auction/Server/Server"
)

func Pass(sender chan *net.TCPConn) {
	for {
		connection := <-sender
		client, err := S.NewClient(connection)
		if err == nil {
			go ClientHandler(client)
		} else {
			logger.Println(err)
		}
	}
}

func ClientHandler(client S.Client) {
	buffer := make([]byte, 4096)
	for {
		if n, err := client.GetConnection().Read(buffer); err == nil {
			logger.Println("Recieved: n =", n, "success")
			content := buffer[:n]
			if MessageHandler(client, content) {
				if err := client.GetConnection().Close(); err != nil {
					logger.Println("Error closing connection with", client.GetConnection().RemoteAddr().String(),
						"because of", err)
				}
			}
		} else {
			logger.Println("Connection with", client.GetConnection().RemoteAddr().String(), "closed because", err)
			break
		}
	}
}

func MessageHandler(client S.Client, content []byte) bool {
	var response string
	var endConnection bool = false

	command, err := S.CmdFromJSON(content)
	if err != nil {
		logger.Println("Discarded bad command packet:", string(content), "because of", err)
	} else {
		err = auction.ExecCommand(command, client)
		if err != nil {
			logger.Println("Could not send response to", command.Command, "command because of", err)
		}
	}

	if response != "" {
		logger.Println("Responding to request [", content, "] from", client.GetConnection().RemoteAddr().String(), "with",
			response)
		if _, err := client.GetConnection().Write([]byte(response)); err != nil {
			logger.Println("Could not send response", response, "to", client.GetConnection().RemoteAddr().String,
				"because of", err)
		}
	}

	return endConnection
}

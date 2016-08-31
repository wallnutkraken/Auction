package main

import "net"

func Pass(sender chan *net.TCPConn) {
	for {
		connection := <-sender
		client, err := NewClient(connection)
		if err == nil {
			go ClientHandler(client)
		} else {
			logger.Println(err)
		}
	}
}

func ClientHandler(client Client) {
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

func MessageHandler(client Client, content []byte) bool {
	var response string
	var endConnection bool = false

	if string(content) == "hello" {
		response = "Hello to you too!"
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

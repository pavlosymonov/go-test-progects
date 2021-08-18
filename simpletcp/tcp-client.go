package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	for {
	 	reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Print(err.Error())
		}

		fmt.Fprintf(conn, text + "\n")

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Print(err.Error())
		}

		fmt.Print("Message from server:", message)
	}
}

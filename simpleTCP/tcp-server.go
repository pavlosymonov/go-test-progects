package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func writeFileMessage(message string) {
	file, err := os.OpenFile("messages.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	dataWriter := bufio.NewWriter(file)

	dataWriter.WriteString(message)

	dataWriter.Flush()
	file.Close()
}

func main() {

	fmt.Println("Launching server...")

	ln, err := net.Listen("tcp", ":8081")

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("Message Received: ", string(message))

		newMessage := strings.ToUpper(message)

		conn.Write([]byte(newMessage))

		writeFileString := fmt.Sprintf(
			"Received message: %s - Formatted message: %s;\n",
			strings.TrimSuffix(message, "\n"),
			strings.TrimSuffix(newMessage, "\n"))

		writeFileMessage(writeFileString)
	}
}
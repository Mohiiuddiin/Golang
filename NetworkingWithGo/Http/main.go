package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	//Listen Request
	netListener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer netListener.Close() //await
	log.Printf("Server started on localhost:8888")
	//Accept Request
	connection, err := netListener.Accept()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Remote Addrress : \n" + connection.RemoteAddr().String())

	byteString := make([]byte, 1024)

	//read bytestring
	byteReader, err := connection.Read(byteString)
	if err != nil {
		fmt.Println(err.Error())
	}

	requestString := string(byteString[:byteReader])

	fmt.Println("From Byte Reader : \n ", byteReader)
	fmt.Println("From Request String: \n " + requestString)

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta http-equiv="X-UA-Compatible" content="IE=edge"><meta name="viewport" content="width=device-width, initial-scale=1.0"><title>Document</title></head><body><h1>Name : Golum Mohammed Mohiuddin</h1><h2>I love to code with Golang</h2></body></html>`
	// fmt.Fprint(connection, "HTTP/1.1 200 OK\r\n")
	// fmt.Fprintf(connection, "Content-Length: %d\r\n", len(body))
	// fmt.Fprint(connection, "Content-Type: text/html\r\n")
	// fmt.Fprint(connection, "\r\n")
	// fmt.Fprint(connection, body)
	response := "HTTP/1.1 200 OK \r\n" +
		"Content-Length: %d \r\n" +
		"Content-Type: text/html \r\n" +
		"\r\n %s"
	message := fmt.Sprintf(response, len(body), body)
	fmt.Println(message)
	connection.Write([]byte(message))
}

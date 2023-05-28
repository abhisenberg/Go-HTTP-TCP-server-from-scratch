package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
)

func main() {

	//Step 1: Start listening for connections on the :8080 port
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()

	fmt.Println("Listening: ", listener)

	//Step 2: Keep listening for incoming connection requests (using infinite for loop)
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Connection: ", connection)

		defer connection.Close()

		//Step 3a: Write something on the connection
		io.WriteString(connection, "Hello from the server!")

		//Step 3b: Read something on the connection
		readdata, err := ioutil.ReadAll(connection)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(readdata))
	}

}

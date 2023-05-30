package main

import (
	"bufio"
	"fmt"
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

		//Step 3: Handle multiple connections
		go handle(connection)
	}

}

func handle(connection net.Conn) {
	defer connection.Close()
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		fmt.Println("###", scanner.Text())
	}
	fmt.Println("Sup bruhs!")
}

/*
The scanner.Scan() loop above will never stop on it's own, because
the scanner does not know when the request has ended and will keep
listening on it. This functionlity will have to be built manually.
*/

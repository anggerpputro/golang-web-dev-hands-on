package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

func main() {
	// create a new listener
	lt, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	// print to console
	fmt.Println("TCP Server Listening on port :8080")
	fmt.Println("...")

	// defer close the listener
	defer lt.Close()

	// loop forever until break
	for {
		// accepting new connection
		conn, err := lt.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		go handle(conn)
	}
}

// write function is used to write to incoming connection
func write(conn net.Conn) {
	// defer close connection
	defer conn.Close()

	// print to console
	fmt.Printf("[%v] Accepting new connection!\n", time.Now().Format(time.UnixDate))

	// send response to the new connection
	fmt.Fprintln(conn, "Hello from TCP Server!")
}

// handle function is used to handle a new incoming connection
func handle(conn net.Conn) {
	// request id
	rid := rand.Intn(100)

	// print to console
	fmt.Printf("[%d][%v] Accepting new connection!\n", rid, time.Now().Format(time.UnixDate))

	fmt.Printf("[%d]## Scanning Request:\n", rid)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()

		if ln == "exit" {
			fmt.Printf("[%d]## User Send Exit Request!\n", rid)
			fmt.Fprintf(conn, "\nBye NDES!\n\n")
			break
		}

		fmt.Printf("[%d]%s\n", rid, ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}

	fmt.Printf("[%d]## Closing this connection...\n", rid)

	// defer close the connection
	defer conn.Close()
}

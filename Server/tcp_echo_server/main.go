package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
)

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Printf("TYPE: %T\n", ln)
		ln = fmt.Sprint("FROM SERVER: " + ln)
		fmt.Fprintln(conn, ln)
	}
}

func main() {
	li, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		handle(conn)
	}-r
}


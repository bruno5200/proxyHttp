package main

import (
"log"
"fmt"
"net"
"os"
"io"
)

func main() {
	Connect()
}

func Connect() {
	log.Println("connecion starting")

	l, err := net.Listen("tcp","localhost:4545")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go proxy(conn)
	}
}

fun proxy() {
	defer conn.Close()

	remoteconnection, err := net.Dial("tcp", "site.net:443")

	if err != nil {
		log.Println(err)
		return
	}
	defer remoteconnection.Close()

	go io.Copy(remoteconnection, conn)
	io.Copy(conn, remoteconnection)
}

func respawnWhenError(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])

		if err != nil {
			log.Println(err)
			return
		}
		_ = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		
		os.Stderr.Write(buf[:n])
	}
}
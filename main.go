package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting the server ...")
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 8866})
	if err != nil {
		log.Fatalf("Udp Service listen report udp fail:%v", err)
	}
	defer conn.Close()
	var data = make([]byte, 1024*4)
	var raw []byte
	for {
		n, remoteAddr, err := conn.ReadFromUDP(data)
		if err == nil {
			raw = make([]byte, n)
			copy(raw, data[:n])
			fmt.Println("receive from ", remoteAddr, " data:", string(raw))
			go func() {
				conn.WriteToUDP(raw, remoteAddr)
			}()
		} else {
			log.Printf("ReadFromUDP error:%v", err)
		}
	}
}

package main

import (
	"log"

	"github.com/Lelisayohanes/fms-ingo/peer2peer"
)

func main() {
	// Your code here
	tr := peer2peer.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}

}

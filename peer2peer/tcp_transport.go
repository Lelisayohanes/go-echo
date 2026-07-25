package peer2peer

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represnent the remote node over the TCP established connection

type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	mu            sync.Mutex
	peers         map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddress,
		peers:         make(map[net.Addr]Peer),
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err := net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	//start accept loop
	go t.startAcceptLoop()

	return nil

}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {

	fmt.Printf(" New incoming connection %+v\n")

}

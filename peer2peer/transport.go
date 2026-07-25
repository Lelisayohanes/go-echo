package peer2peer

// peer is a node in the network, it can be a client or a server
// it represent remote node
type Peer interface {
}

// any thing that handles the communication between nodes in the network
// this can be of the form (TCP, UDP, HTTP, WebSocket, etc)
type Transport interface {
	listenAndAccept() error
}

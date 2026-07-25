package peer2peer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testTCPTransport(t *testing.T) {
	// Create a new TCP transport
	listenAddress := ":8080"
	transport := NewTCPTransport(listenAddress)

	//Server
	assert.Nil(t, transport.ListenAndAccept())

	select {}
}

package node

import (
	"context"
	"sync"

	"github.com/AnomalyFi/hypersdk/rpc"
	executionv1 "github.com/AnomalyFi/nodekit-sdk/structs"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ethereum/go-ethereum/log"
)

// NodeKitListenerHandler is the NodeKit listener handler.
type NodeKitListenerHandler struct {
	mu sync.Mutex

	websocketClient        *rpc.WebSocketClient
	executionServiceServer executionv1.ExecutionServiceServer
	host                   string
	chainId                string
}

// NewNodeKitListenerHandler creates a new NodeKit Listener.
// It registers the listener with the node so it can be stopped on shutdown.
func NewNodeKitListenerHandler(node *Node, execService executionv1.ExecutionServiceServer, cfg *Config) error {

	// Create a new WebSocketClient

	websocketClient, err := rpc.NewWebSocketClient(cfg.NodeKitWSHost)
	if err != nil {
		return err
	}

	serverHandler := &NodeKitListenerHandler{
		websocketClient:        websocketClient,
		executionServiceServer: execService,
		host:                   cfg.NodeKitWSHost,
		chainId:                cfg.NodeKitChainId,
	}

	node.RegisterNodeKitListener(serverHandler)
	return nil
}

// Start starts the NodeKit listener if it is enabled.
func (handler *NodeKitListenerHandler) Start() error {
	handler.mu.Lock()
	defer handler.mu.Unlock()

	chainID, err := ids.FromString(handler.chainId)

	if err != nil {
		return err
	}
	go handler.executionServiceServer.WSBlock(handler.host, chainID, context.Background(), handler.websocketClient)

	log.Info("NodeKit Listener started")
	return nil
}

// Stop stops the NodeKit listener.
func (handler *NodeKitListenerHandler) Stop() error {
	handler.mu.Lock()
	defer handler.mu.Unlock()

	handler.websocketClient.Close()
	// handler.server.Stop()
	log.Info("NodeKit listener stopped")
	return nil
}

package node

//TODO need to fix this to work with websocket client

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
	executionServiceServer *executionv1.ExecutionServiceServer
}

// NewNodeKitListenerHandler creates a new NodeKit Listener.
// It registers the listener with the node so it can be stopped on shutdown.
func NewNodeKitListenerHandler(node *Node, execService executionv1.ExecutionServiceServer, cfg *Config) error {

	// Create a new WebSocketClient
	//TODO need to fix this
	websocketClient, err := rpc.NewWebSocketClient("http://127.0.0.1:9650/ext/bc/2bLP6aabd9Hju4SNnn1dsE4Q8FNrAg3N1zeWmzYFky1yDzoFVr")
	if err != nil {
		return err
	}

	serverHandler := &NodeKitListenerHandler{
		websocketClient:        websocketClient,
		executionServiceServer: &execService,
	}

	// executionv1.RegisterExecutionServiceServer(server, execService)

	node.RegisterNodeKitListener(serverHandler)
	return nil
}

// Start starts the gRPC server if it is enabled.
func (handler *NodeKitListenerHandler) Start() error {
	handler.mu.Lock()
	defer handler.mu.Unlock()

	//TODO fix this
	JSONRPCEndpoint := "http://127.0.0.1:9650/ext/bc/2bLP6aabd9Hju4SNnn1dsE4Q8FNrAg3N1zeWmzYFky1yDzoFVr"

	chainID, err := ids.FromString("2bLP6aabd9Hju4SNnn1dsE4Q8FNrAg3N1zeWmzYFky1yDzoFVr")

	if err != nil {
		return err
	}
	go handler.executionServiceServer.WSBlock(JSONRPCEndpoint, chainID, context.Background(), handler.websocketClient)

	log.Info("NodeKit Listener started")
	return nil
}

// Stop stops the gRPC server.
func (handler *NodeKitListenerHandler) Stop() error {
	handler.mu.Lock()
	defer handler.mu.Unlock()

	handler.websocketClient.Close()
	// handler.server.Stop()
	log.Info("NodeKit listener stopped")
	return nil
}

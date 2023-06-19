package node

//TODO need to fix this to work with websocket client

import (
	"fmt"
	"sync"

	"github.com/AnomalyFi/hypersdk/rpc"
	"github.com/ethereum/go-ethereum/log"
	executionv1 "github.com/ethereum/go-ethereum/ws/execution"
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
	websocketClient, err := rpc.NewWebSocketClient("your_websocket_uri")
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

	// if handler.endpoint == "" {
	// 	return nil
	// }

	// // Start the gRPC server
	// lis, err := net.Listen("tcp", handler.endpoint)
	// if err != nil {
	// 	return err
	// }

	blockHash, err := handler.executionServiceServer.InitState()
	if err != nil {
		return err
	}
	fmt.Println(blockHash)

	//TODO create a method that listens to block and integrates with DoBlock
	// I can probably just create a method in the execution service itself to help with this
	//go handler.websocketClient.ListenBlock(context.Background(), nil) // Add this line to start listening to blocks

	log.Info("NodeKit Listener started")
	return nil
}

// Stop stops the gRPC server.
func (handler *NodeKitListenerHandler) Stop() error {
	handler.mu.Lock()
	defer handler.mu.Unlock()

	// handler.server.Stop()
	log.Info("gRPC server stopped", "endpoint", handler.endpoint)
	return nil
}

package ethapi

import (
	nodekittx "github.com/AnomalyFi/nodekit-sdk/tx"
	"github.com/ethereum/go-ethereum/log"
)

const secondaryChainID = "ethereum"

type NodeKitAPI struct {
	endpoint string
}

func NewNodeKitAPI(endpoint string) *NodeKitAPI {
	log.Info("NewNodeKitAPI", "endpoint", endpoint)
	return &NodeKitAPI{endpoint: endpoint}
}

func (api *NodeKitAPI) SubmitTransaction(tx []byte) error {
	return nodekittx.BuildAndSendTransaction("http://127.0.0.1:9650/ext/bc/2bLP6aabd9Hju4SNnn1dsE4Q8FNrAg3N1zeWmzYFky1yDzoFVr", "ethereum", tx)
}

package ethapi

import (
	nodekittx "github.com/AnomalyFi/nodekit-sdk/tx"
	"github.com/ethereum/go-ethereum/log"
)

const secondaryChainID = "ethereum"

type NodeKitAPI struct {
	endpoint string
	chainId  string
}

func NewNodeKitAPI(endpoint string, chainId string) *NodeKitAPI {
	log.Info("NewNodeKitAPI", "endpoint", endpoint)
	return &NodeKitAPI{endpoint: endpoint, chainId: chainId}
}

func (api *NodeKitAPI) SubmitTransaction(tx []byte) error {
	return nodekittx.BuildAndSendTransaction(api.endpoint, api.chainId, secondaryChainID, tx)
}

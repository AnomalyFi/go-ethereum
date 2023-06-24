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
	return nodekittx.BuildAndSendSecondaryTransaction(api.endpoint, secondaryChainID, tx)
}

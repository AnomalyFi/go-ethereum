package ethapi

//TODO this needs to be modified to work with our sdk

import (
	nodekittx "https://github.com/AnomalyFi/nodekit-sdk/tx"
	"github.com/ethereum/go-ethereum/log"
)

const secondaryChainID = "ethereum"

type NodeKitAPI struct {
	endpoint string
}

func NewNodeKitAPI(endpoint string) *MetroAPI {
	log.Info("NewNodeKitAPI", "endpoint", endpoint)
	return &NodeKitAPI{endpoint: endpoint}
}

func (api *NodeKitAPI) SubmitTransaction(tx []byte) error {
	return nodekittx.BuildAndSendSecondaryTransaction(api.endpoint, secondaryChainID, tx)
}

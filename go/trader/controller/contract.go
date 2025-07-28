package controller

import (

)

var Arbitrum = "https://sepolia-rollup.arbitrum.io/rpc"

type ContractHandler struct {
	cfg     *conf.Config
	ethRepo *commonrepository.EthRepository

	farmDB     *model.FarmDB

	eventMapLock       sync.RWMutex
	eventMap           map[string]map[string]*commonprotocol.ContractEventManage
	//...

	grpcClient pb.CoinmecaGrpcModuleClient
	grpcConn   *grpc.ClientConn

	start chan struct{}
}

func NewContractHandler(rep *model.Repositories, config *conf.Config, root *Controller) (IController, error) {
	r := &ContractHandler{

	}

	if err := rep.Get(..., &r.farmDB, ...); err != nil {
		return nil, err
	}


	r.newGrpcHandler(config)

	return r, nil
}

func convertToGrpcTxData(txData *commondatabase.TxData) *commondatabase.GrpcTxData {
	return &commondatabase.GrpcTxData{
		Id:          txData.Id.String(),
		BlockHash:   txData.BlockHash,
	
	}
}

func (c *ContractHandler) Start() error {

}

func (c *Controller) ProcessEvent(txData *commondatabase.GrpcTxData) {

}

func (c *ContractHandler) ParseEvent(contract *commonprotocol.Contract, log *types.Log, chainId, eventName string, event interface{}) (*int64, string, error) {

}


func (c *ContractHandler) processEvent(txData *commondatabase.GrpcTxData) {
	if txData == nil {
		commonlog.Logger.Error("processEvent",
			zap.String("txData type: ", fmt.Sprintf("%T", txData)),
		)
		return
	}
	
	//...

	chainId := fmt.Sprintf("%d", chainIdInt)

	//... 

	if contract, ok := c.GetContractAt(chainId, txData.To); ok {

		switch contract.Name {
			//...

		case commonprotocol.ContractFarm:
			c.processEventFarm(txData, receipt)

			//...
		default:
			commonlog.Logger.Warn("processEvent",
				zap.String("Event not set up yet", contract.Name),
			)
		}
	}
}



func (c *ContractHandler) processEventMarket(txData *commondatabase.GrpcTxData, receipt *types.Receipt) {

}

func (c *ContractHandler) processEventFarm(txData *commondatabase.GrpcTxData, receipt *types.Receipt) {
	
}


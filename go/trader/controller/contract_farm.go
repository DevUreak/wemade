package controller

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

// 비지니스 코드 
func (c *ContractHandler) HandleFarmEventByName(log *types.Log, chainId, eventName string, contract *commonprotocol.Contract) {
	if contract == nil {
		commonlog.Logger.Warn("HandleMarketEventByName",
			zap.String("chainId", chainId),
			zap.String("eventName", eventName),
		)
		return
	}

	switch eventName {
	case "Stake":
		c.Stake(log, contract, chainId, eventName)

	case "Unstake":
		c.Unstake(log, contract, chainId, eventName)


	default:
		commonlog.Logger.Warn("HandleFarmEventByName",
			zap.String(eventName, "invalid event name"),
		)
	}
}

func (c *ContractHandler) Stake(log *types.Log, contract *commonprotocol.Contract, chainId, eventName string) {

	event := farm.EventStake{}
	timestamp, date, err := c.ParseEvent(contract, log, chainId, eventName, &event)
	if err != nil {
		commonlog.Logger.Error(eventName,
			zap.String("UnmarshalEvent", err.Error()),
		)
		return
	}

	info, err := c.farmDB.GetFarm(chainId, contract.Address)
	if err != nil {
		commonlog.Logger.Error(eventName,
			zap.String("GetFarm", err.Error()),
		)
	}

	staking := commonutils.BigIntFromDecimal128(&info.Staked)

	amount, err := commonutils.Decimal128FromBigInt(event.Amount)
	if err != nil {
		commonlog.Logger.Error(eventName,
			zap.String("ConvertDecimalToBigInt", err.Error()),
		)
	}

	share, err := commonutils.Decimal128FromBigInt(new(big.Int).Div(new(big.Int).Mul(event.Amount, big.NewInt(100)), staking))
	if err != nil {
		commonlog.Logger.Error(eventName,
			zap.String("ConvertDecimalToBigInt", err.Error()),
		)
	}

	recent := farm.Recent{
		ChainId:  chainId,
		Time:     *timestamp,
		//...
	}

	c.UpdateFarmRecent(&recent)
	fmt.Println("recent : ", recent)
}


package controller

import (
	"context"
	"fmt"
	"log"

	pb "/grpc-module"
)

func (c *ContractHandler) newGrpcHandler(config *conf.Config) error {
	client, conn := pb.InitGrpcClientV2(config.Gclient.GrpcPort)
	c.grpcClient = client
	c.grpcConn = conn

	return nil
}

// 비지니스 코드
func (c *ContractHandler) sendFarmRecent(f *farm.Recent) {
	var recentType pb.FarmTradeType
	switch f.Type {
	case 0:
		recentType = pb.FarmTradeType_STAKING
	case 1:
		recentType = pb.FarmTradeType_UNSTAKE
	}

	req := &pb.GeneralRequest{
		Request: &pb.GeneralRequest_FarmRecent{
			FarmRecent: &pb.FarmRecentRequest{
				Data: &pb.FarmRecent{
					ChainId:  f.ChainId,
					Address:  f.Address,
					//...
				},
			},
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	response, err := c.grpcClient.Send(ctx, req)
	if err != nil {
		log.Fatalf("Error sending transaction to gRPC server: %v", err)
	}
	log.Printf("gRPC server response: %v", response.GetSuccess())

}


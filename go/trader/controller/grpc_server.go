package controller

import (
	"context"
	"fmt"

	pb "/grpc-module"
)

type TransactionService struct {
	controller *Controller
}

func NewTransactionService(controller *Controller) *TransactionService {
	return &TransactionService{
		controller: controller,
	}
}

func (s *TransactionService) Send(ctx context.Context, req *pb.GeneralRequest) (*pb.GeneralResponse, error) {
	// oneof를 사용하여 들어오는 요청의 타입을 확인합니다.
	switch x := req.Request.(type) {
	case *pb.GeneralRequest_Transaction:
		// Transaction 타입의 요청을 처리합니다.
		data := x.Transaction
		// 실제 트랜잭션 데이터 처리 로직을 여기에 추가합니다.
		txData := commondatabase.GrpcTxData{
			Id:          data.Data.Id,
			BlockHash:   data.Data.BlockHash,
			//...
			ChainId:     data.Data.ChainId,
			//...
			To:          data.Data.To,
			//...
		}
		s.controller.ProcessEvent(&txData)

	default:
		return nil, fmt.Errorf("unknown request type")
	}

	return &pb.GeneralResponse{Success: true}, nil
}

func StartGrpcServer(controller *Controller, port string) {
	transactionService := NewTransactionService(controller)

	// gRPC 서버 시작
	pb.StartGrpcServerV2(transactionService, port)
}

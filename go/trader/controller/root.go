package controller

import (

	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

type IController interface {
	Start() error
}

type ControlHandler func(repo *model.Repositories, config *conf.Config, root *Controller) (IController, error)
type Controller struct {
	cfg          *conf.Config
	serverNumber int

	ContractHandler *ContractHandler

	lock  sync.RWMutex
	eUnit map[reflect.Type]reflect.Value
}

func New(config *conf.Config, rep *model.Repositories) (*Controller, error) {
	contractHandler, err := NewContractHandler(rep, config, nil)
	if err != nil {
		return nil, err
	}

	r := &Controller{
		cfg:             config,
		eUnit:           make(map[reflect.Type]reflect.Value),
		ContractHandler: contractHandler.(*ContractHandler),
	}

	//...

	//++ grpc 서버 실행
	StartGrpcServer(r, config.Gserver.ServerAddr)

	return r, nil
}

func (c *Controller) Register(handler ControlHandler, rep *model.Repositories, cfg *conf.Config) error {

	if r, err := handler(rep, cfg, c); err != nil {
		return err
	} else if r != nil {

		//...
	}
	return nil
}

func (c *Controller) GetConfig() *conf.Config {
	return c.cfg
}

func (c *Controller) GetServerNum() int {
	return c.serverNumber
}

func (c *Controller) GetContractHandler() *ContractHandler {
	return c.ContractHandler
}

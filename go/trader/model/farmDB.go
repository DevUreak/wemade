package model

import (

	"context"
	"strings"


	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type FarmDB struct {
	config *conf.Config

	client     *mongo.Client
	colFarm    *mongo.Collection
	colHistory *mongo.Collection

	start chan struct{}
}

func NewFarmDB(config *conf.Config) (commondatabase.IRepository, error) {
	r := &FarmDB{
		config: config,
		start:  make(chan struct{}),
	}

	var err error
	credential := options.Credential{
		Username: r.config.Repositories["farmDB"]["username"].(string),
		Password: r.config.Repositories["farmDB"]["pass"].(string),
	}
	// ... 
	return r, nil
}

func (f *FarmDB) Start() error {
	return func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		close(f.start)
		return
	}()
}

// 비지니스 코드
func (f *FarmDB) SaveFarmRecent(recent *farm.Recent) error {
	filter, update := f.BsonForFarmsRecent(recent)
	option := options.Update().SetUpsert(true)

	_, err := f.colHistory.UpdateOne(
		context.Background(),
		filter,
		update,
		option,
	)
	if err != nil {
		commonlog.Logger.Debug("FarmDB",
			zap.String("SaveFarmRecent", err.Error()),
		)
		return err
	}

	return nil
}
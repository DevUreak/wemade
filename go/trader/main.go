package main

import (
	ctl "coinmeca-trader/controller"
	"coinmeca-trader/key"
	"coinmeca-trader/model"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"syscall"
	"golang.org/x/sync/errgroup"
)

var configFlag = flag.String("config", "./conf/config.toml", "toml file to use for configuration")
var httpFlag = flag.Int("http", 0, "router http port")

var (
	g errgroup.Group
)

func main() {
	flag.Parse()

	config := conf.NewConfig(*configFlag)

	if *httpFlag != 0 {
		config.Port.Http = *httpFlag
	}
	// ... 
	repositories, err := model.NewRepositories(config, keyManager)
	if err != nil {
		panic(err)
	}
	
	// grpc 서버 실행
	controller, err := ctl.New(config, repositories)
	if err != nil {
		panic(err)
	}

	// ... 
	// Setup the HTTP server
	mapi := &http.Server{
		Addr: config.Port.Server,
	}

	g.Go(func() error {
		return mapi.ListenAndServe()
	})
	// ... 

}

func ensureDirectory(dirName string) {
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		panic(fmt.Sprintf("Failed to create directory: %s", err))
	}
}

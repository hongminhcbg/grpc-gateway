package main

import (
	"github.com/hongminhcbg/grpc-gateway/conf"
	"github.com/hongminhcbg/grpc-gateway/gateway"
)

func main(){
	cfg, err := conf.GetConfig()
	if err != nil {
		panic(err)
	}

	gw := gateway.NewGateway(cfg, nil)
	gw.Run()

}

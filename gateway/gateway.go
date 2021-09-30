package gateway

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/hongminhcbg/grpc-gateway/api"
	"github.com/hongminhcbg/grpc-gateway/conf"
	"github.com/hongminhcbg/grpc-gateway/src/service"
)

func Run(config *conf.Config) error {
	gmux := runtime.NewServeMux(
		runtime.WithErrorHandler(CustomErrorHandler()),
		runtime.WithIncomingHeaderMatcher(CustomIncomingHeaderMatcher(config)),
	)
	mainService := service.NewService()
	err := api.RegisterUserServiceHandlerServer(context.Background(), gmux, mainService)
	if err != nil {
		panic(err)
	}

	gw := *&http.Server{
		Addr: fmt.Sprintf("0.0.0.0%s", config.HttpPort),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				gmux.ServeHTTP(w, r)
				return
			}

			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("content-type", "applicaton/json")
			w.Write([]byte(`{"code": 403, "message": "not allow this action"}`))
			return
		}),
	}

	go func() {
		log.Printf("starting http server at address %s\n", config.HttpPort)
		err = gw.ListenAndServe()
		if err != nil {
			panic("start http gateway error" + err.Error())
		}
	}()

	go func() {
		log.Printf("starting grpc server at address %s\n", config.GrpcPort)
		lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0%s", config.GrpcPort))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		api.RegisterUserServiceServer(s, mainService)
		err = s.Serve(lis)
		if err != nil {
			panic("start grpc gateway error" + err.Error())
		}
	}()

	return nil
}

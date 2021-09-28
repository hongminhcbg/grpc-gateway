package gateway

import (
	"context"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/hongminhcbg/grpc-gateway/api"
	"github.com/hongminhcbg/grpc-gateway/src/service"
)

var httpPort string = "9000"
var grpcPort string = "10000"
var grpcAddress string = "localhost:10000"
var httpAddress string = "localhost:9000"

func Run() error {
	gmux := runtime.NewServeMux(
		runtime.WithErrorHandler(CustomErrorHandler()),
	)
	mainService := service.NewService()
	err := api.RegisterUserServiceHandlerServer(context.Background(), gmux, mainService)
	if err != nil {
		panic(err)
	}

	gw := *&http.Server{
		Addr: httpAddress,
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
		log.Printf("starting http server at address %s\n", httpAddress)
		err = gw.ListenAndServe()
		if err != nil {
			panic("start http gateway error" + err.Error())
		}
	}()

	go func() {
		log.Printf("starting grpc server at address %s\n", grpcAddress)
		lis, err := net.Listen("tcp", grpcAddress)
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

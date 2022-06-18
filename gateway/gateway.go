package gateway

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/hongminhcbg/grpc-gateway/api"
	"github.com/hongminhcbg/grpc-gateway/conf"
	"github.com/hongminhcbg/grpc-gateway/src/service"
)

type Gateway struct {
	config      *conf.Config
	middlewares []func(http.Handler) http.Handler
}

func NewGateway(config *conf.Config, middlewares []func(http.Handler) http.Handler) *Gateway {
	return &Gateway{
		config:      config,
		middlewares: middlewares,
	}
}

func customMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("req = ", req, "method = ", info.FullMethod)
	return handler(ctx, req)
}

func (g *Gateway) Run() error {
	gmux := runtime.NewServeMux(
		runtime.WithErrorHandler(CustomErrorHandler()),
		runtime.WithIncomingHeaderMatcher(CustomIncomingHeaderMatcher(g.config)),
	)
	mainService := service.NewService()
	err := api.RegisterUserServiceHandlerServer(context.Background(), gmux, mainService)
	if err != nil {
		panic(err)
	}

	var hannler http.Handler
	hannler = gmux
	for _, middleware := range g.middlewares {
		hannler = middleware(hannler)
	}

	muxServer := http.NewServeMux()
	muxServer.Handle("/", hannler)

	gw := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0%s", g.config.HttpPort),
		Handler: muxServer,
	}

	go func() {
		log.Printf("starting http server at address %s\n", g.config.HttpPort)
		err = gw.ListenAndServe()
		if err != nil {
			panic("start http gateway error" + err.Error())
		}
	}()

	go func() {
		log.Printf("starting grpc server at address %s\n", g.config.GrpcPort)
		lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0%s", g.config.GrpcPort))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer(
			grpc.UnaryInterceptor(
				grpc_middleware.ChainUnaryServer(customMiddleware),
			),
		)
		api.RegisterUserServiceServer(s, mainService)
		err = s.Serve(lis)
		if err != nil {
			panic("start grpc gateway error" + err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c)
	for {
		select {
		case sig := <-c:
			if sig == syscall.SIGKILL || sig == syscall.SIGINT {
				log.Println("Received signal kill process")
				return nil
			}
		}
	}
}

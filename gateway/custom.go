package gateway

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func CustomErrorHandler() runtime.ErrorHandlerFunc {
	return func(ctx context.Context, _ *runtime.ServeMux, _ runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(fmt.Sprintf("{\"code\": 0, \"message\": \"%s\"}", err.Error())))
	}
}

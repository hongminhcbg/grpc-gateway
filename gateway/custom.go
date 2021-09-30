package gateway

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hongminhcbg/grpc-gateway/conf"
)

func CustomErrorHandler() runtime.ErrorHandlerFunc {
	return func(ctx context.Context, _ *runtime.ServeMux, _ runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(fmt.Sprintf("{\"code\": 0, \"message\": \"%s\"}", err.Error())))
	}
}

func CustomIncomingHeaderMatcher(config *conf.Config) runtime.HeaderMatcherFunc {
	return func(in string) (string, bool) {
		headers := strings.Split(config.Cors.AccessCOntrolAllowHeaders, ",")
		for _, h := range headers {
			if in == h {
				return in, true
			}
		}

		return "", false
	}

}

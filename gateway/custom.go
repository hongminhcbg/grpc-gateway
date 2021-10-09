package gateway

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hongminhcbg/grpc-gateway/conf"
	"github.com/hongminhcbg/grpc-gateway/src/erp"
)

const defaultError = `{"code": "500000", "message": "internal server error"}`

func CustomErrorHandler() runtime.ErrorHandlerFunc {
	return func(ctx context.Context, _ *runtime.ServeMux, _ runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
		w.Header().Set("content-type", "application/json")
		appErr, ok := err.(*erp.ApplicationError)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(defaultError))
			return
		}

		w.WriteHeader(appErr.Code / 1000)
		w.Write([]byte(fmt.Sprintf("{\"code\": %d, \"message\": \"%s\"}", appErr.Code, appErr.Message)))
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

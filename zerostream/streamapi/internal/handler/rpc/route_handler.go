package rpc

import (
	"net/http"

	"zerostream/streamapi/internal/logic/rpc"
	"zerostream/streamapi/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RouteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := rpc.NewRouteLogic(r.Context(), svcCtx)
		resp, err := l.Route()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

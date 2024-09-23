package rpc

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerostream/streamapi/internal/logic/rpc"
	"zerostream/streamapi/internal/svc"
	"zerostream/streamapi/internal/types"
)

func RecordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StreamRpcRecordRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := rpc.NewRecordLogic(r.Context(), svcCtx)
		resp, err := l.Record(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerostream/streamapi/internal/logic"
	"zerostream/streamapi/internal/svc"
	"zerostream/streamapi/internal/types"
)

func StreamapiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewStreamapiLogic(r.Context(), svcCtx)
		resp, err := l.Streamapi(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

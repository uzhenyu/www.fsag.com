package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zg6/z326/api/internal/logic"
	"zg6/z326/api/internal/svc"
)

func OrderNofitHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		l := logic.NewOrderNofitLogic(r.Context(), svcCtx)
		l.OrderNofit(r.Form)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			w.Write([]byte("成功"))
			httpx.Ok(w)
		}
	}
}

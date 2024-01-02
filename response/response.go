package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type (
	Response struct {
		Code    int64     `json:"code"`
		Message string    `json:"message"`
		Data    ReloadRes `json:"data"`
	}

	ReloadRes struct {
		Reload bool `json:"reload"`
	}
)

func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	res := &ReloadRes{
		Reload: true,
	}
	resRet := &Response{

		Code:    401,
		Message: "您的帐户异地登陆或令牌失效",
		Data:    *res,
	}
	httpx.WriteJson(w, http.StatusOK, resRet)
}

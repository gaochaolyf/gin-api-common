package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/jangozw/gin-api-common/consts"
	"net/http"
	"time"
)

// api result to client format here

type ApiResponse struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

type GinCtx struct {
	*gin.Context
}

//output data need use *gin.Context
// usage:
// utils.Ctx(c).Success("Ok")

func Ctx(c *gin.Context) *GinCtx {
	return &GinCtx{c}
}

func (c *GinCtx) Success(data interface{}) {
	result := ResponseSuccess(data)
	//logs.Logger().Infof("%s", ToJson(result))
	c.JSON(http.StatusOK, result)

}

func (c *GinCtx) SuccessSimple() {
	c.JSON(http.StatusOK, ResponseSuccessSimple())
}

func (c *GinCtx) Fail(errMsg interface{}) {
	c.FailWithCode(consts.ApiCodeError, errMsg)
}

func (c *GinCtx) FailWithCode(code int, errMsg interface{}) {
	c.JSON(http.StatusOK, ResponseFailWithCode(code, errMsg))
}

//当前登陆的用户id
func (c *GinCtx) GetLoginUid() int64 {
	return c.GetInt64(consts.CtxKeyLoginUser)
}

func ResponseSuccess(data interface{}) ApiResponse {
	msg, _ := consts.GetApiMsgByCode(consts.ApiCodeSuccess)
	return ApiResponse{
		consts.ApiCodeSuccess,
		msg,
		time.Now().Unix(),
		data,
	}
}

func ResponseSuccessSimple() ApiResponse {
	msg, _ := consts.GetApiMsgByCode(consts.ApiCodeSuccess)
	return ApiResponse{
		consts.ApiCodeSuccess,
		msg,
		time.Now().Unix(),
		struct{}{},
	}
}

func ResponseFailWithCode(code int, err interface{}) ApiResponse {
	msg, _ := consts.GetApiMsgByCode(code)
	return ApiResponse{
		code,
		msg + "! " + parseErrToMsg(err),
		time.Now().Unix(),
		struct{}{},
	}
}

func parseErrToMsg(err interface{}) (errMsg string) {
	if e, ok := err.(error); ok {
		errMsg = e.Error()
	} else if e, ok := err.(string); ok {
		errMsg = e
	} else {
		errMsg = ToJson(err)
	}
	return
}

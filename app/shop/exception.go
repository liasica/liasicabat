package shop

import (
    "battery/app/model"
    "battery/app/service"
    "battery/library/response"
    "github.com/gogf/gf/net/ghttp"
)

var ExceptionApi = exceptionApi{}

type exceptionApi struct {
}

// Report
// @Summary 门店-异常上报
// @Tags    门店
// @Accept  json
// @Produce json
// @Param   entity  body model.ExceptionReportReq true "请求数据"
// @Router  /sapi/exception [POST]
// @Success 200 {object} response.JsonResponse "返回结果"
func (*exceptionApi) Report(r *ghttp.Request) {
    var req model.ExceptionReportReq
    if err := r.Parse(&req); err != nil {
        response.Json(r, response.RespCodeArgs, err.Error())
    }
    manager := r.Context().Value(model.ContextShopManagerKey).(*model.ContextShopManager)
    // manager.ShopId
    req.ShopId = manager.ShopId
    if service.ExceptionService.Create(r.Context(), req) == nil {
        response.JsonOkExit(r)
    }
    response.JsonErrExit(r)
}

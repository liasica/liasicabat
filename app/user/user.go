package user

import (
    "fmt"
    "github.com/gogf/gf/net/ghttp"

    "battery/app/model"
    "battery/app/service"
    "battery/library/esign/sign"
    "battery/library/qr"
    "battery/library/response"
)

var RiderApi = riderApi{}

type riderApi struct {
}

// Login
// @Summary 骑手-用户登录
// @Tags    骑手
// @Accept  json
// @Produce json
// @Param   entity  body model.UserLoginReq true "登录数据"
// @Router  /rapi/login [POST]
// @Success 200 {object} response.JsonResponse{data=model.UserLoginRep}  "返回结果"
func (*riderApi) Login(r *ghttp.Request) {
    var req model.UserLoginReq
    if err := r.Parse(&req); err != nil {
        response.Json(r, response.RespCodeArgs, err.Error())
    }
    if data, err := service.UserService.Login(r.Context(), req); err != nil {
        response.Json(r, response.RespCodeArgs, err.Error())
    } else {
        response.JsonOkExit(r, data)
    }
}

// Auth
// @Summary 骑手-实名认证提交
// @Tags    骑手
// @Accept  json
// @Produce json
// @Param   entity  body model.UserRealNameAuthReq true "认证数据"
// @Router  /rapi/auth [POST]
// @Success 200 {object} response.JsonResponse{data=model.UserRealNameAuthRep}  "返回结果"
func (*riderApi) Auth(r *ghttp.Request) {
    var req model.UserRealNameAuthReq
    if err := r.Parse(&req); err != nil {
        response.Json(r, response.RespCodeArgs, err.Error())
    }
    if user, err := service.UserService.GetUserByIdCardNo(r.Context(), req.IdCardNo); err == nil && user.AuthState == model.AuthStateVerifySuccess {
        response.Json(r, response.RespCodeArgs, fmt.Sprintf("证件号码 %s 已认证超过，请检查证件号码", req.IdCardNo))
    }
    if res, err := service.UserService.RealNameAuthSubmit(r.Context(), req); err != nil {
        response.Json(r, response.RespCodeArgs, err.Error())
    } else {
        response.JsonOkExit(r, res)
    }
}

// AuthGet
// @Summary 骑手-获取实名认证状态 [每次切换页面都进行查询]
// @Tags    骑手
// @Accept  json
// @Produce json
// @Router  /rapi/auth [GET]
// @Success 200 {object} response.JsonResponse{data=int}  "返回结果"
func (*riderApi) AuthGet(r *ghttp.Request) {
    u := r.Context().Value(model.ContextRiderKey).(*model.ContextRider)
    response.JsonOkExit(r, u.AuthState)
}

// PushToken
// @Summary 骑手-上报推送token
// @Tags    骑手-消息
// @Accept  json
// @Produce json
// @Param   entity  body model.PushTokenReq true "登录数据"
// @Router  /rapi/device  [PUT]
// @Success 200 {object} response.JsonResponse  "返回结果"
func (*riderApi) PushToken(r *ghttp.Request) {
    var req model.PushTokenReq
    if err := r.Parse(&req); err != nil {
        response.Json(r, response.RespCodeArgs, err.Error())
    }
    err := service.UserService.PushToken(r.Context(), req)
    if err != nil {
        response.JsonErrExit(r, response.RespCodeSystemError)
    }
    response.JsonOkExit(r)
}

// Combo
// @Summary 骑手-获取骑手当前套餐详情
// @Tags    骑手
// @Accept  json
// @Produce json
// @Router  /rapi/combo  [GET]
// @Success 200 {object} response.JsonResponse{data=model.UserCurrentComboOrder}  "返回结果"
func (*riderApi) Combo(r *ghttp.Request) {
    rep, err := service.UserService.MyCombo(r.Context())
    if err != nil {
        response.Json(r, response.RespCodeArgs, err.Error())
    }
    profile := service.UserService.Profile(r.Context())
    rep.BatteryState = profile.User.BatteryState
    response.JsonOkExit(r, rep)
}

// ComboOrderQr
// @Summary 骑手-获取骑手当前套餐二维码
// @Tags    骑手
// @Accept  json
// @Produce json
// @Router  /rapi/combo_order/qr  [GET]
// @Success 200 {object} response.JsonResponse "返回结果, data字段为二维码图片数据，需要本地生成二维码"
func (*riderApi) ComboOrderQr(r *ghttp.Request) {
    u := r.Context().Value(model.ContextRiderKey).(*model.ContextRider)
    if u.GroupId > 0 {
        response.JsonOkExit(r, fmt.Sprintf("%d-%s-%d", u.GroupId, u.Qr, u.BatteryType))
    } else {
        if u.ComboOrderId == 0 {
            response.Json(r, response.RespCodeArgs, "还未购买套餐")
        }
        order, err := service.ComboOrderService.Detail(r.Context(), u.ComboOrderId)
        if err != nil {
            response.Json(r, response.RespCodeArgs, "为找到订单")
        }
        response.JsonOkExit(r, order.No)
    }
}

// Profile
// @Summary 骑手-首页
// @Tags    骑手
// @Accept  json
// @Produce json
// @Router  /rapi/home  [GET]
// @Success 200 {object} response.JsonResponse{data=model.UserProfileRep}  "返回结果"
func (*riderApi) Profile(r *ghttp.Request) {
    profile := service.UserService.Profile(r.Context())
    profile.Qr = qr.Code.AddPrefix(profile.Qr)
    response.JsonOkExit(r, profile)
}

// SignFile
// @Summary 骑手-签约文件地址
// @Tags    骑手
// @Accept  json
// @Produce json
// @Router  /rapi/sign_file  [GET]
// @Success 200 {object} response.JsonResponse{data=[]model.UserSignFileRep}  "返回结果"
func (*riderApi) SignFile(r *ghttp.Request) {
    u := r.Context().Value(model.ContextRiderKey).(*model.ContextRider)
    s, err := service.SignService.UserLatestDoneDetail(r.Context(), u.Id, u.ComboOrderId, u.GroupId)
    if err != nil {
        response.JsonErrExit(r, response.RespCodeNotFound)
    }
    res, err := sign.Service().SignFlowDocuments(s.FlowId)
    if err != nil || res.Code != 0 {
        response.JsonErrExit(r, response.RespCodeSystemError)
    }
    files := make([]*model.UserSignFileRepItem, len(res.Data.Docs))
    for i, f := range res.Data.Docs {
        files[i] = &model.UserSignFileRepItem{
            FileName: f.FileName,
            FileUrl:  f.FileUrl,
        }
    }
    response.JsonOkExit(r, files)
}

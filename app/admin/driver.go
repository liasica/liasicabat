// Copyright (C) liasica. 2021-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Created at 2021-10-16
// Based on apiv2 by liasica, magicrolan@qq.com.

package admin

import (
    "battery/app/model"
    "battery/app/service"
    "battery/library/request"
    "battery/library/response"
    "github.com/gogf/gf/net/ghttp"
)

var DriverApi = driverApi{}

type driverApi struct {
}

// Verify
// @Summary 认证列表
// @Tags    管理
// @Accept  json
// @Param   entity body model.UserVerifyReq true "请求参数"
// @Produce  json
// @Router  /admin/driver/verify [GET]
// @Success 200 {object} response.JsonResponse{data=model.ItemsWithTotal{items=[]model.UserVerifyListItem}}  "返回结果"
func (*driverApi) Verify(r *ghttp.Request) {
    req := new(model.UserVerifyReq)
    _ = request.ParseRequest(r, req)

    total, items := service.UserService.ListVerifyItems(r.Context(), req)
    response.ItemsWithTotal(r, total, items)
}

func (*driverApi) Personal(r *ghttp.Request) {
    req := new(model.UserPersonalReq)
    _ = request.ParseRequest(r, req)

    service.UserService.ListPersonalItems(r.Context(), req)
}

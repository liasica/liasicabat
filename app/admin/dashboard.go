// Copyright (C) liasica. 2021-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Created at 2021-10-24
// Based on apiv2 by liasica, magicrolan@qq.com.

package admin

import (
    "battery/app/dao"
    "battery/app/model"
    "battery/app/service"
    "battery/library/request"
    "battery/library/response"
    "github.com/gogf/gf/net/ghttp"
)

var DashboardApi = new(dashboardApi)

type dashboardApi struct {
}

// OpenCities
// @Summary 已开通城市
// @Tags    管理
// @Accept  json
// @Produce json
// @Router  /admin/dashboard/cities [GET]
// @Success 200 {object} response.JsonResponse{data=[]model.City} "返回结果"
func (d *dashboardApi) OpenCities(r *ghttp.Request) {
    response.JsonOkExit(r, service.DashboardService.OpenCities(r.Context()))
}

// Overview
// @Summary 统计概览
// @Tags    管理
// @Accept  json
// @Produce json
// @Param 	startDate query string false "开始日期"
// @Param 	endDate query string false "结束日期"
// @Router  /admin/dashboard/overview [GET]
// @Success 200 {object} response.JsonResponse{data=model.DashboardOverview} "返回结果"
func (*dashboardApi) Overview(r *ghttp.Request) {
    var req = new(model.DateBetween)
    _ = request.ParseRequest(r, req)
    ctx := r.Context()
    data := new(model.DashboardOverview)
    service.DashboardService.OverviewRiderCount(ctx, req, data)
    service.DashboardService.OverviewGroupCount(ctx, req, data)
    service.DashboardService.OverviewOrderTotal(ctx, req, data)
    response.JsonOkExit(r, data)
}

// Order
// @Summary 新增订单统计
// @Tags    管理
// @Accept  json
// @Produce json
// @Param 	cityId query int false "城市ID"
// @Param 	startDate query string true "开始日期"
// @Param 	endDate query string true "结束日期"
// @Router  /admin/dashboard/order [GET]
// @Success 200 {object} response.JsonResponse{data=[]model.DashboardOrder} "返回结果"
func (*dashboardApi) Order(r *ghttp.Request) {
    var req = new(model.DashboardOrderReq)
    _ = request.ParseRequest(r, req)
    if req.StartDate.AddDate(0, 0, 60).Before(req.EndDate) {
        response.JsonErrExit(r, response.RespCodeArgs, "时间范围太大")
    }
    ctx := r.Context()
    data := service.DashboardService.NewlyOrders(ctx, req)
    response.JsonOkExit(r, data)
}

// Business
// @Summary 业务统计
// @Tags    管理
// @Accept  json
// @Produce json
// @Param 	cityId query int false "城市ID"
// @Param 	startDate query string true "开始日期"
// @Param 	endDate query string true "结束日期"
// @Router  /admin/dashboard/business [GET]
// @Success 200 {object} response.JsonResponse{data=[]model.DashboardBusiness} "返回结果"
func (*dashboardApi) Business(r *ghttp.Request) {
    var req = new(model.DashboardBusinessReq)
    _ = request.ParseRequest(r, req)
    if req.StartDate.AddDate(0, 0, 60).Before(req.EndDate) {
        response.JsonErrExit(r, response.RespCodeArgs, "时间范围太大")
    }
    ctx := r.Context()
    data := service.DashboardService.Business(ctx, req)
    response.JsonOkExit(r, data)
}

// RankShop
// @Summary 门店订单数量排名
// @Tags    管理
// @Accept  json
// @Produce json
// @Router  /admin/dashboard/rankshop [GET]
// @Success 200 {object} response.JsonResponse{data=[]model.DashboardRankShop} "返回结果"
func (*dashboardApi) RankShop(r *ghttp.Request) {
    var items []model.DashboardRankShop
    c := dao.ComboOrder.Columns
    _ = dao.ComboOrder.Ctx(r.Context()).WithAll().
        Where(c.PayState, model.PayStateSuccess).
        WhereNot(c.Type, model.ComboOrderTypePenalty).
        Fields(`type, payState, cityId, shopId, COUNT(1) AS cnt`).
        OrderDesc("cnt").
        Group("shopId").
        Scan(&items)
    for k, item := range items {
        if item.Shop != nil {
            items[k].ShopName = item.Shop.Name
        }
        if item.City != nil {
            items[k].CityName = item.City.Name
        }
    }
    response.JsonOkExit(r, items)
}

// RankCity
// @Summary 门店订单数量排名[总订单分布]
// @Tags    管理
// @Accept  json
// @Produce json
// @Router  /admin/dashboard/rankcity [GET]
// @Success 200 {object} response.JsonResponse{data=[]model.DashboardRankShop} "返回结果"
func (*dashboardApi) RankCity(r *ghttp.Request) {
    var items []model.DashboardRankCity
    c := dao.ComboOrder.Columns
    _ = dao.ComboOrder.Ctx(r.Context()).WithAll().
        Where(c.PayState, model.PayStateSuccess).
        WhereNot(c.Type, model.ComboOrderTypePenalty).
        Fields(`type, payState, cityId, COUNT(1) AS cnt`).
        OrderDesc("cnt").
        Group("cityId").
        Scan(&items)
    for k, item := range items {
        if item.City != nil {
            items[k].CityName = item.City.Name
        }
    }
    response.JsonOkExit(r, items)
}

// Rider
// @Summary 骑手统计
// @Tags    管理
// @Accept  json
// @Produce json
// @Router  /admin/dashboard/rider [GET]
// @Success 200 {object} response.JsonResponse{data=[]model.DashboardRider} "返回结果"
func (*dashboardApi) Rider(r *ghttp.Request) {
    data := new(model.DashboardRider)
    _ = dao.User.Ctx(r.Context()).Fields("COUNT(1) AS total, SUM(IF(authState = 2, 1, 0)) AS verified, SUM(IF(DATE(batteryReturnAt) >= DATE(NOW()), 1, 0)) AS `using`").Scan(data)
    _ = dao.UserBiz.Ctx(r.Context()).Fields(`userId, COUNT(1) AS used`).Scan(data)
    response.JsonOkExit(r, data)
}

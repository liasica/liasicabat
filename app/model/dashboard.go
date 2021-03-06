// Copyright (C) liasica. 2021-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Created at 2021-10-24
// Based on apiv2 by liasica, magicrolan@qq.com.

package model

import "github.com/gogf/gf/util/gmeta"

// DashboardOverview 订单概览
type DashboardOverview struct {
    PersonalRiders int64   `json:"personalRiders"` // 个签用户数量
    GroupRiders    int64   `json:"groupRiders"`    // 团签用户数量
    GroupCnt       int64   `json:"groupCnt"`       // 团队数量
    TotalAmount    float64 `json:"totalAmount"`    // 订单金额总计
    Orders         int64   `json:"orders"`         // 总订单量
    PersonalAmount float64 `json:"personalAmount"` // 个签订单金额总计
    PersonalOrders int64   `json:"personalOrders"` // 个签订单数量
    Deposit        float64 `json:"deposit"`        // 个签押金
    GroupAmount    float64 `json:"groupAmount"`    // 团签订单金额总计
    GroupOrders    int64   `json:"groupOrders"`    // 团签订单数量
}

// DashboardOrderReq 新增统计请求体
type DashboardOrderReq struct {
    DateBetween
    CityId uint `json:"cityId"` // 城市ID
}

// DashboardOrder 新增订单统计
type DashboardOrder struct {
    Date    string  `json:"date"`    // 日期
    New     int64   `json:"new"`     // 新增订单
    Renewal int64   `json:"renewal"` // 续签订单
    Amount  float64 `json:"amount"`  // 新增订单金额
}

// DashboardBusinessReq 业务统计请求
type DashboardBusinessReq struct {
    DateBetween
    CityId uint `json:"cityId"` // 城市ID
}

// DashboardBusiness 业务统计
type DashboardBusiness struct {
    Date      string `json:"date"`      // 日期
    Renewal   int64  `json:"renewal"`   // 换电
    Pause     int64  `json:"pause"`     // 寄存
    Retrieval int64  `json:"retrieval"` // 恢复
}

// DashboardRankShop 门店订单排名
type DashboardRankShop struct {
    gmeta.Meta `json:"-" orm:"table:combo_order" swaggerignore:"true"`

    CityId   uint   `json:"-"`
    CityName string `json:"cityName"` // 城市名
    ShopId   uint   `json:"-"`
    ShopName string `json:"shopName"` // 门店名
    Cnt      int    `json:"cnt"`      // 订单数量

    City *Districts `json:"-" orm:"with:id=CityId"`
    Shop *Shop      `json:"-" orm:"with:id=ShopId"`
}

// DashboardRankCity 城市订单排名
type DashboardRankCity struct {
    gmeta.Meta `json:"-" orm:"table:combo_order" swaggerignore:"true"`

    CityId   uint   `json:"-"`
    CityName string `json:"cityName"` // 城市名
    Cnt      int    `json:"cnt"`      // 订单数量

    City *Districts `json:"-" orm:"with:id=CityId"`
}

// DashboardRider 用户统计
type DashboardRider struct {
    Total    uint64 `json:"total"`    // 注册骑手总数
    Verified uint64 `json:"verified"` // 已认证骑手
    Used     uint64 `json:"used"`     // 换过电的骑手
    Using    uint64 `json:"using"`    // 正在使用中的骑手
}

package service

import (
    "context"
    "github.com/gogf/gf/os/gtime"
    "time"

    "battery/app/dao"
    "battery/app/model"
)

var ShopBatteryRecordService = shopBatteryRecordService{}

type shopBatteryRecordService struct {
}

// User 骑手记录
func (*shopBatteryRecordService) User(ctx context.Context, recordType, bizType, shopId uint, bizId uint64, userName string, batteryType uint) error {
    _, err := dao.ShopBatteryRecord.Ctx(ctx).
        Fields(
            dao.ShopBatteryRecord.Columns.ShopId,
            dao.ShopBatteryRecord.Columns.BizId,
            dao.ShopBatteryRecord.Columns.BizType,
            dao.ShopBatteryRecord.Columns.UserName,
            dao.ShopBatteryRecord.Columns.BatteryType,
            dao.ShopBatteryRecord.Columns.Num,
            dao.ShopBatteryRecord.Columns.Type,
        ).
        Insert(model.ShopBatteryRecord{
            ShopId:      shopId,
            BizId:       bizId,
            BizType:     bizType,
            UserName:    userName,
            BatteryType: batteryType,
            Num:         1,
            Type:        recordType,
        })
    return err
}

// Platform 平台调拨
func (*shopBatteryRecordService) Platform(ctx context.Context, recordType, shopId, num, batteryType uint) error {
    _, err := dao.ShopBatteryRecord.Ctx(ctx).
        Fields(dao.ShopBatteryRecord.Columns.ShopId,
            dao.ShopBatteryRecord.Columns.Num,
            dao.ShopBatteryRecord.Columns.Type,
            dao.ShopBatteryRecord.Columns.BatteryType,
        ).Insert(model.ShopBatteryRecord{
        ShopId:      shopId,
        BatteryType: batteryType,
        Num:         num,
        Type:        recordType,
    })
    return err
}

// ShopList 门店获取电池记录
func (*shopBatteryRecordService) ShopList(ctx context.Context, shopId uint, recordType uint, st *gtime.Time, et *gtime.Time) (list []model.ShopBatteryRecord) {
    m := dao.ShopBatteryRecord.Ctx(ctx).
        Where(dao.ShopBatteryRecord.Columns.ShopId, shopId).
        Where(dao.ShopBatteryRecord.Columns.Type, recordType).
        OrderDesc(dao.ShopBatteryRecord.Columns.Id)
    if !st.IsZero() {
        m = m.WhereGTE(dao.ShopBatteryRecord.Columns.CreatedAt, st)
    }
    if !et.IsZero() {
        m = m.WhereLT(dao.ShopBatteryRecord.Columns.CreatedAt, et.Add(24*time.Hour))
    }
    _ = m.Scan(&list)
    return
}

// ShopDaysTotal 门店获取电池记录按天统计
func (*shopBatteryRecordService) ShopDaysTotal(ctx context.Context, days []int, recordType uint) (list []struct {
    Day int
    Cnt uint
}) {
    _ = dao.ShopBatteryRecord.Ctx(ctx).
        Fields(dao.ShopBatteryRecord.Columns.Day, "count(*) cnt").
        WhereIn(dao.ShopBatteryRecord.Columns.Day, days).
        Where(dao.ShopBatteryRecord.Columns.Type, recordType).
        Group(dao.ShopBatteryRecord.Columns.Day).
        Scan(&list)

    return
}

// ListAdmin 所有门店电池记录
func (*shopBatteryRecordService) ListAdmin(ctx context.Context, req *model.BatteryRecordListReq) (total int, items []model.BatteryRecordListItem) {
    c := dao.ShopBatteryRecord.Columns
    query := dao.ShopBatteryRecord.Ctx(ctx).
        OrderDesc(c.CreatedAt)
    if req.Type > 0 {
        query = query.Where(c.Type, req.Type)
    }
    if !req.StartTime.IsZero() {
        query = query.WhereGTE(c.CreatedAt, req.StartTime)
    }
    if !req.EndTime.IsZero() {
        query = query.WhereLT(c.CreatedAt, req.EndTime.Add(24*time.Hour))
    }
    if req.ShopId > 0 {
        query = query.Where(c.ShopId, req.ShopId)
    }
    _ = query.Page(req.PageIndex, req.PageLimit).Scan(&items)
    total, _ = query.Count()

    return
}

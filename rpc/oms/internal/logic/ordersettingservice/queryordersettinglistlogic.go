package ordersettingservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOrderSettingListLogic 查询订单设置列表
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
*/
type QueryOrderSettingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderSettingListLogic {
	return &QueryOrderSettingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderSettingList 查询订单设置列表
func (l *QueryOrderSettingListLogic) QueryOrderSettingList(in *omsclient.QueryOrderSettingListReq) (*omsclient.QueryOrderSettingListResp, error) {
	orderSetting := query.OmsOrderSetting
	q := orderSetting.WithContext(l.ctx)
	if in.Status != 2 {
		q = q.Where(orderSetting.Status.Eq(in.Status))
	}
	if in.IsDefault != 2 {
		q = q.Where(orderSetting.IsDefault.Eq(in.IsDefault))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询订单设置列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询订单设置列表失败")
	}

	var list []*omsclient.OrderSettingListData

	for _, item := range result {
		list = append(list, &omsclient.OrderSettingListData{
			Id:                  item.ID,                                          // 主键ID
			FlashOrderOvertime:  item.FlashOrderOvertime,                          // 秒杀订单超时关闭时间(分)
			NormalOrderOvertime: item.NormalOrderOvertime,                         // 正常订单超时时间(分)
			ConfirmOvertime:     item.ConfirmOvertime,                             // 发货后自动确认收货时间（天）
			FinishOvertime:      item.FinishOvertime,                              // 自动完成交易时间，不能申请售后（天）
			Status:              item.Status,                                      // 状态：0->禁用；1->启用
			IsDefault:           item.IsDefault,                                   // 是否默认：0->否；1->是
			CommentOvertime:     item.CommentOvertime,                             // 订单完成后自动好评时间（天）
			CreateBy:            item.CreateBy,                                    // 创建者
			CreateTime:          time_util.TimeToStr(item.CreateTime),             // 创建时间
			UpdateBy:            pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
			UpdateTime:          time_util.TimeToString(item.UpdateTime),          // 更新时间

		})
	}

	return &omsclient.QueryOrderSettingListResp{
		Total: count,
		List:  list,
	}, nil
}

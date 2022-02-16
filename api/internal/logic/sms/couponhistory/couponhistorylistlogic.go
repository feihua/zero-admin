package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) CouponHistoryListLogic {
	return CouponHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponHistoryListLogic) CouponHistoryList(req types.ListCouponHistoryReq) (*types.ListCouponHistoryResp, error) {
	resp, err := l.svcCtx.Sms.CouponHistoryList(l.ctx, &smsclient.CouponHistoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询优惠券使用记录列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询优惠券使用记录失败")
	}

	for _, data := range resp.List {

		fmt.Println(data)
	}
	var list []*types.ListtCouponHistoryData

	for _, item := range resp.List {
		list = append(list, &types.ListtCouponHistoryData{
			Id:             item.Id,
			CouponId:       item.CouponId,
			MemberId:       item.MemberId,
			CouponCode:     item.CouponCode,
			MemberNickname: item.MemberNickname,
			GetType:        item.GetType,
			CreateTime:     item.CreateTime,
			UseStatus:      item.UseStatus,
			UseTime:        item.UseTime,
			OrderId:        item.OrderId,
			OrderSn:        item.OrderSn,
		})
	}

	return &types.ListCouponHistoryResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询优惠券使用记录成功",
	}, nil
}

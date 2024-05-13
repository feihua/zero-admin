package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type CouponListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponListLogic(ctx context.Context, svcCtx *svc.ServiceContext) CouponListLogic {
	return CouponListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponListLogic) CouponList(req types.ListCouponReq) (*types.ListCouponResp, error) {
	resp, err := l.svcCtx.CouponService.CouponList(l.ctx, &smsclient.CouponListReq{
		Current:   req.Current,
		PageSize:  req.PageSize,
		Type:      req.Type,
		Name:      req.Name,
		Platform:  req.Platform,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		UseType:   req.UseType,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询优惠券列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询优惠券失败")
	}

	var list []*types.ListCouponData

	for _, item := range resp.List {
		list = append(list, &types.ListCouponData{
			Id:           item.Id,
			Type:         item.Type,
			Name:         item.Name,
			Platform:     item.Platform,
			Count:        item.Count,
			Amount:       item.Amount,
			PerLimit:     item.PerLimit,
			MinPoint:     item.MinPoint,
			StartTime:    item.StartTime,
			EndTime:      item.EndTime,
			UseType:      item.UseType,
			Note:         item.Note,
			PublishCount: item.PublishCount,
			UseCount:     item.UseCount,
			ReceiveCount: item.ReceiveCount,
			EnableTime:   item.EnableTime,
			Code:         item.Code,
			MemberLevel:  item.MemberLevel,
		})
	}

	return &types.ListCouponResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询优惠券成功",
	}, nil
}

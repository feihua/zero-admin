package coupon

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"strings"
)

// CouponListLogic 优惠券
/*
Author: LiuFeiHua
Date: 2024/5/15 9:21
*/
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

// CouponList 查询优惠券列表
func (l *CouponListLogic) CouponList(req types.ListCouponReq) (*types.ListCouponResp, error) {
	resp, err := l.svcCtx.CouponService.QueryCouponList(l.ctx, &smsclient.QueryCouponListReq{
		Current:   req.Current,
		PageSize:  req.PageSize,
		Type:      req.Type,
		Name:      strings.TrimSpace(req.Name),
		Platform:  req.Platform,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		UseType:   req.UseType,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询优惠券列表异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.ListCouponData

	for _, item := range resp.List {
		list = append(list, &types.ListCouponData{
			Id:           item.Id,           //
			Type:         item.Type,         // 优惠券类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券
			Name:         item.Name,         // 名称
			Platform:     item.Platform,     // 使用平台：0->全部；1->移动；2->PC
			Count:        item.Count,        // 数量
			Amount:       item.Amount,       // 金额
			PerLimit:     item.PerLimit,     // 每人限领张数
			MinPoint:     item.MinPoint,     // 使用门槛；0表示无门槛
			StartTime:    item.StartTime,    // 开始时间
			EndTime:      item.EndTime,      // 结束时间
			UseType:      item.UseType,      // 使用类型：0->全场通用；1->指定分类；2->指定商品
			Note:         item.Note,         // 备注
			PublishCount: item.PublishCount, // 发行数量
			UseCount:     item.UseCount,     // 已使用数量
			ReceiveCount: item.ReceiveCount, // 领取数量
			EnableTime:   item.EnableTime,   // 可以领取的日期
			Code:         item.Code,         // 优惠码
			MemberLevel:  item.MemberLevel,  // 可领取的会员类型：0->无限时
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

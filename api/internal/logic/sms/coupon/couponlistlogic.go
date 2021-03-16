package logic

import (
	"context"
	"fmt"
	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"
	"go-zero-admin/rpc/sms/smsclient"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
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
	resp, err := l.svcCtx.Sms.CouponList(l.ctx, &smsclient.CouponListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {

		fmt.Println(data)
	}
	var list []*types.ListtCouponData

	for _, item := range resp.List {
		StartTime, _ := time.Parse("2006-01-02 15:04:05", item.StartTime)
		EndTime, _ := time.Parse("2006-01-02 15:04:05", item.EndTime)
		list = append(list, &types.ListtCouponData{
			Id:           item.Id,
			Type:         item.Type,
			Name:         item.Name,
			Platform:     item.Platform,
			Count:        item.Count,
			Amount:       item.Amount,
			PerLimit:     item.PerLimit,
			MinPoint:     item.MinPoint,
			StartTime:    StartTime,
			EndTime:      EndTime,
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
		Message:  "",
	}, nil
}

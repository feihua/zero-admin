package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionListLogic 秒杀活动
/*
Author: LiuFeiHua
Date: 2024/5/14 10:51
*/
type FlashPromotionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionListLogic {
	return FlashPromotionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FlashPromotionList 查询秒杀活动
func (l *FlashPromotionListLogic) FlashPromotionList(req types.ListFlashPromotionReq) (*types.ListFlashPromotionResp, error) {
	resp, err := l.svcCtx.FlashPromotionService.FlashPromotionList(l.ctx, &smsclient.FlashPromotionListReq{
		Current:   req.Current,
		PageSize:  req.PageSize,
		Title:     req.Title,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Status:    req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询限时购列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询限时购表失败")
	}

	var list []*types.ListFlashPromotionData

	for _, item := range resp.List {
		list = append(list, &types.ListFlashPromotionData{
			Id:         item.Id,
			Title:      item.Title,
			StartDate:  item.StartDate,
			EndDate:    item.EndDate,
			Status:     item.Status,
			CreateTime: item.CreateTime,
		})
	}

	return &types.ListFlashPromotionResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询限时购表成功",
	}, nil
}

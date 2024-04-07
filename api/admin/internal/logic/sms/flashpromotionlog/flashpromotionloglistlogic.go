package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionLogListLogic {
	return FlashPromotionLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionLogListLogic) FlashPromotionLogList(req types.ListFlashPromotionLogReq) (*types.ListFlashPromotionLogResp, error) {
	resp, err := l.svcCtx.FlashPromotionLogService.FlashPromotionLogList(l.ctx, &smsclient.FlashPromotionLogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询限时购通知记录列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询限时购通知记录失败")
	}

	var list []*types.ListtFlashPromotionLogData

	for _, item := range resp.List {
		list = append(list, &types.ListtFlashPromotionLogData{
			Id:            item.Id,
			MemberId:      item.MemberId,
			ProductId:     item.ProductId,
			MemberPhone:   item.MemberPhone,
			ProductName:   item.ProductName,
			SubscribeTime: item.SubscribeTime,
			SendTime:      item.SendTime,
		})
	}

	return &types.ListFlashPromotionLogResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询限时购通知记录成功",
	}, nil
}

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

type FlashPromotionSessionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionSessionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionSessionListLogic {
	return FlashPromotionSessionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionSessionListLogic) FlashPromotionSessionList(req types.ListFlashPromotionSessionReq) (*types.ListFlashPromotionSessionResp, error) {
	resp, err := l.svcCtx.Sms.FlashPromotionSessionList(l.ctx, &smsclient.FlashPromotionSessionListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询限时购场次表列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询限时购场次表失败")
	}

	for _, data := range resp.List {

		fmt.Println(data)
	}
	var list []*types.ListtFlashPromotionSessionData

	for _, item := range resp.List {
		list = append(list, &types.ListtFlashPromotionSessionData{
			Id:         item.Id,
			Name:       item.Name,
			StartTime:  item.StartTime,
			EndTime:    item.EndTime,
			Status:     item.Status,
			CreateTime: item.CreateTime,
		})
	}

	return &types.ListFlashPromotionSessionResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询限时购场次表成功",
	}, nil
}

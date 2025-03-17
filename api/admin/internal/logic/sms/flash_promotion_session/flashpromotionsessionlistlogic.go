package flash_promotion_session

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"strconv"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionSessionListLogic 限时购场次
/*
Author: LiuFeiHua
Date: 2024/5/14 13:43
*/
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

// FlashPromotionSessionList 查询限时购场次
func (l *FlashPromotionSessionListLogic) FlashPromotionSessionList(req *types.ListFlashPromotionSessionReq) (*types.ListFlashPromotionSessionResp, error) {
	resp, err := l.svcCtx.FlashPromotionSessionService.QueryFlashPromotionSessionList(l.ctx, &smsclient.QueryFlashPromotionSessionListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询限时购场次表列表异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.ListFlashPromotionSessionData

	for _, item := range resp.List {
		list = append(list, &types.ListFlashPromotionSessionData{
			Id:         item.Id,         // 编号
			Name:       item.Name,       // 场次名称
			StartTime:  item.StartTime,  // 每日开始时间
			EndTime:    item.EndTime,    // 每日结束时间
			Status:     item.Status,     // 启用状态：0->不启用；1->启用
			CreateTime: item.CreateTime, // 创建时间
			Key:        strconv.FormatInt(item.Id, 10),
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

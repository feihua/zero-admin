package flash_promotion

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryFlashPromotionListLogic 查询限时购表列表
/*
Author: 刘飞华
Date: 2025/02/07 10:11:43
*/
type QueryFlashPromotionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryFlashPromotionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFlashPromotionListLogic {
	return &QueryFlashPromotionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryFlashPromotionList 查询限时购表列表
func (l *QueryFlashPromotionListLogic) QueryFlashPromotionList(req *types.QueryFlashPromotionListReq) (resp *types.QueryFlashPromotionListResp, err error) {
	result, err := l.svcCtx.FlashPromotionService.QueryFlashPromotionList(l.ctx, &smsclient.QueryFlashPromotionListReq{
		PageNum:   req.Current,
		PageSize:  req.PageSize,
		Title:     req.Title,     // 标题
		StartDate: req.StartDate, // 开始日期
		EndDate:   req.EndDate,   // 结束日期
		Status:    req.Status,    // 上下线状态

	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字限时购表列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryFlashPromotionListData

	for _, item := range result.List {
		list = append(list, &types.QueryFlashPromotionListData{
			Id:         item.Id,         // 编号
			Title:      item.Title,      // 标题
			StartDate:  item.StartDate,  // 开始日期
			EndDate:    item.EndDate,    // 结束日期
			Status:     item.Status,     // 上下线状态
			CreateTime: item.CreateTime, // 创建时间
		})
	}

	return &types.QueryFlashPromotionListResp{
		Code:     "000000",
		Message:  "查询限时购表列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}

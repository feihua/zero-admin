package flashpromotion

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

// QueryFlashPromotionDetailLogic 查询限时购表详情
/*
Author: 刘飞华
Date: 2025/02/07 10:11:43
*/
type QueryFlashPromotionDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryFlashPromotionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFlashPromotionDetailLogic {
	return &QueryFlashPromotionDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryFlashPromotionDetail 查询限时购表详情
func (l *QueryFlashPromotionDetailLogic) QueryFlashPromotionDetail(req *types.QueryFlashPromotionDetailReq) (resp *types.QueryFlashPromotionDetailResp, err error) {

	detail, err := l.svcCtx.FlashPromotionService.QueryFlashPromotionDetail(l.ctx, &smsclient.QueryFlashPromotionDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询限时购表详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryFlashPromotionDetailData{
		Id:         detail.Id,         // 编号
		Title:      detail.Title,      // 标题
		StartDate:  detail.StartDate,  // 开始日期
		EndDate:    detail.EndDate,    // 结束日期
		Status:     detail.Status,     // 上下线状态
		CreateTime: detail.CreateTime, // 创建时间
	}
	return &types.QueryFlashPromotionDetailResp{
		Code:    "000000",
		Message: "查询限时购表成功",
		Data:    data,
	}, nil
}

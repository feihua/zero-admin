package homerecommendsubject

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRecommendSubjectStatusLogic 人气推荐专题
/*
Author: LiuFeiHua
Date: 2024/5/14 9:43
*/
type UpdateRecommendSubjectStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRecommendSubjectStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRecommendSubjectStatusLogic {
	return &UpdateRecommendSubjectStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateRecommendSubjectStatus 批量修改推荐状态
func (l *UpdateRecommendSubjectStatusLogic) UpdateRecommendSubjectStatus(req *types.UpdateRecommendSubjectStatusReq) (resp *types.UpdateRecommendSubjectStatusResp, err error) {
	_, err = l.svcCtx.HomeRecommendSubjectService.UpdateRecommendSubjectStatus(l.ctx, &smsclient.UpdateRecommendSubjectStatusReq{
		Ids:             req.Ids,
		RecommendStatus: req.RecommendStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量修改推荐状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量修改推荐状态失败")
	}

	return &types.UpdateRecommendSubjectStatusResp{
		Code:    "000000",
		Message: "批量修改推荐状态成功",
	}, nil
}

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

// HomeRecommendSubjectDeleteLogic 人气推荐专题
/*
Author: LiuFeiHua
Date: 2024/5/14 9:43
*/
type HomeRecommendSubjectDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendSubjectDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendSubjectDeleteLogic {
	return HomeRecommendSubjectDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// HomeRecommendSubjectDelete 删除人气推荐专题
func (l *HomeRecommendSubjectDeleteLogic) HomeRecommendSubjectDelete(req types.DeleteHomeRecommendSubjectReq) (*types.DeleteHomeRecommendSubjectResp, error) {
	_, err := l.svcCtx.HomeRecommendSubjectService.HomeRecommendSubjectDelete(l.ctx, &smsclient.HomeRecommendSubjectDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除人气推荐专题异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除人气推荐专题失败")
	}
	return &types.DeleteHomeRecommendSubjectResp{
		Code:    "000000",
		Message: "",
	}, nil
}

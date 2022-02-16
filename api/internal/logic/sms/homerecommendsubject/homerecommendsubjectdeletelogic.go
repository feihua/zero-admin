package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *HomeRecommendSubjectDeleteLogic) HomeRecommendSubjectDelete(req types.DeleteHomeRecommendSubjectReq) (*types.DeleteHomeRecommendSubjectResp, error) {
	_, err := l.svcCtx.Sms.HomeRecommendSubjectDelete(l.ctx, &smsclient.HomeRecommendSubjectDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除人气推荐专题异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除人气推荐专题失败")
	}
	return &types.DeleteHomeRecommendSubjectResp{
		Code:    "000000",
		Message: "",
	}, nil
}

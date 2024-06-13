package homerecommendsubject

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
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
// 1.删除sms_home_recommend_subject的记录(sms-rpc)
// 2.修改cms_subject记录的状态为不推荐(cms-rpc)
func (l *HomeRecommendSubjectDeleteLogic) HomeRecommendSubjectDelete(req types.DeleteHomeRecommendSubjectReq) (*types.DeleteHomeRecommendSubjectResp, error) {

	// 1.删除sms_home_recommend_subject的记录(sms-rpc)
	_, err := l.svcCtx.HomeRecommendSubjectService.DeleteHomeRecommendSubject(l.ctx, &smsclient.DeleteHomeRecommendSubjectReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,删除人气推荐专题异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除人气推荐专题失败")
	}

	// 2.修改cms_subject记录的状态为不推荐(cms-rpc)
	_, err = l.svcCtx.SubjectService.UpdateSubjectRecommendStatus(l.ctx, &cmsclient.UpdateSubjectRecommendStatusReq{
		Ids:    req.SubjectIds,
		Status: 0,
	})
	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,更新人气推荐专题状态异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除人气推荐专题失败")
	}

	return &types.DeleteHomeRecommendSubjectResp{
		Code:    "000000",
		Message: "删除人气推荐专题成功",
	}, nil
}

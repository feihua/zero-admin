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

// HomeRecommendSubjectAddLogic 人气推荐专题
/*
Author: LiuFeiHua
Date: 2024/5/14 9:42
*/
type HomeRecommendSubjectAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendSubjectAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendSubjectAddLogic {
	return HomeRecommendSubjectAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// HomeRecommendSubjectAdd 添加人气推荐
// 1.根据subjectIds查询推商品专题(cms-rpc)
// 2.添加首页推荐记录(sms-rpc)
// 3.修改商品专题的推荐状态(cms-rpc)
func (l *HomeRecommendSubjectAddLogic) HomeRecommendSubjectAdd(req types.AddHomeRecommendSubjectReq) (*types.AddHomeRecommendSubjectResp, error) {
	// 1.根据subjectIds查询推商品专题(cms-rpc)
	listResp, _ := l.svcCtx.SubjectService.SubjectListByIds(l.ctx, &cmsclient.SubjectListByIdsReq{
		Ids: req.SubjectIds,
	})

	var list []*smsclient.HomeRecommendSubjectAddData

	for _, item := range listResp.List {
		list = append(list, &smsclient.HomeRecommendSubjectAddData{
			SubjectId:       item.Id,
			SubjectName:     item.Title,
			RecommendStatus: 1,
			Sort:            int32(item.Id),
		})
	}

	// 2.添加首页推荐记录(sms-rpc)
	_, err := l.svcCtx.HomeRecommendSubjectService.HomeRecommendSubjectAdd(l.ctx, &smsclient.HomeRecommendSubjectAddReq{
		RecommendSubjectAddData: list,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加人气推荐专题信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加人气推荐专题失败")
	}

	// 3.修改cms_subject记录的状态为推荐(cms-rpc)
	_, err = l.svcCtx.SubjectService.UpdateSubjectRecommendStatus(l.ctx, &cmsclient.UpdateSubjectRecommendStatusReq{
		Ids:    req.SubjectIds,
		Status: 1,
	})
	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,更新人气推荐专题状态异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加人气推荐专题失败")
	}

	return &types.AddHomeRecommendSubjectResp{
		Code:    "000000",
		Message: "添加人气推荐专题成功",
	}, nil
}

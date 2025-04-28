package subject

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateSubjectStatusLogic 更新专题表状态状态
/*
Author: 刘飞华
Date: 2025/02/04 15:04:17
*/
type UpdateSubjectStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSubjectStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubjectStatusLogic {
	return &UpdateSubjectStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateSubjectStatus 更新专题表状态
func (l *UpdateSubjectStatusLogic) UpdateSubjectStatus(req *types.UpdateSubjectStatusReq) (resp *types.BaseResp, err error) {
	updateBy := l.ctx.Value("userName").(string)
	_, err = l.svcCtx.SubjectService.UpdateSubjectStatus(l.ctx, &cmsclient.UpdateSubjectStatusReq{
		Ids:             req.Ids,             // 专题id
		RecommendStatus: req.RecommendStatus, // 推荐状态：0->不推荐；1->推荐
		ShowStatus:      req.ShowStatus,      // 显示状态：0->不显示；1->显示
		UpdateBy:        updateBy,            // 更新者
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新专题表状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}

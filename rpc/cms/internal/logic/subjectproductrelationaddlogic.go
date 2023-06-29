package logic

import (
	"context"
	"zero-admin/rpc/model/cmsmodel"

	"zero-admin/rpc/cms/cmsclient"
	"zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectProductRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectProductRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectProductRelationAddLogic {
	return &SubjectProductRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SubjectProductRelationAdd 专题关联
func (l *SubjectProductRelationAddLogic) SubjectProductRelationAdd(in *cmsclient.SubjectProductRelationAddReq) (*cmsclient.SubjectProductRelationAddResp, error) {
	_, err := l.svcCtx.CmsSubjectProductRelationModel.Insert(l.ctx, &cmsmodel.CmsSubjectProductRelation{
		SubjectId: in.SubjectId,
		ProductId: in.ProductId,
	})
	if err != nil {
		return nil, err
	}

	return &cmsclient.SubjectProductRelationAddResp{}, nil
}

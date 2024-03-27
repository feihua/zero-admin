package subjectproductrelationservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectProductRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectProductRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectProductRelationDeleteLogic {
	return &SubjectProductRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubjectProductRelationDeleteLogic) SubjectProductRelationDelete(in *cmsclient.SubjectProductRelationDeleteReq) (*cmsclient.SubjectProductRelationDeleteResp, error) {
	err := l.svcCtx.CmsSubjectProductRelationModel.DeleteByProductId(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}

	return &cmsclient.SubjectProductRelationDeleteResp{}, nil
}

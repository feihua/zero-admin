package subjectproductrelationservicelogic

import (
	"context"

	"zero-admin/rpc/cms/cmsclient"
	"zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectProductRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectProductRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectProductRelationUpdateLogic {
	return &SubjectProductRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubjectProductRelationUpdateLogic) SubjectProductRelationUpdate(in *cmsclient.SubjectProductRelationUpdateReq) (*cmsclient.SubjectProductRelationUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.SubjectProductRelationUpdateResp{}, nil
}

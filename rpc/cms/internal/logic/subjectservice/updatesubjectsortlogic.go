package subjectservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSubjectSortLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSubjectSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubjectSortLogic {
	return &UpdateSubjectSortLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新排序
func (l *UpdateSubjectSortLogic) UpdateSubjectSort(in *cmsclient.UpdateSubjectSortReq) (*cmsclient.UpdateSubjectResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdateSubjectResp{}, nil
}

package subjectservicelogic

import (
	"context"
	"errors"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"

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

// UpdateSubjectSort 更新排序
func (l *UpdateSubjectSortLogic) UpdateSubjectSort(in *cmsclient.UpdateSubjectSortReq) (*cmsclient.UpdateSubjectResp, error) {
	q := query.CmsSubject

	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).UpdateSimple(q.Sort.Value(in.Sort))

	if err != nil {
		logc.Errorf(l.ctx, "更新排序失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新排序失败")
	}

	return &cmsclient.UpdateSubjectResp{}, nil
}

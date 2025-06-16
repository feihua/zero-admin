package productspuservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNewStatusSortLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNewStatusSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNewStatusSortLogic {
	return &UpdateNewStatusSortLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateNewStatusSort 更新新品排序
func (l *UpdateNewStatusSortLogic) UpdateNewStatusSort(in *pmsclient.UpdateProductSortReq) (*pmsclient.UpdateProductSpuStatusResp, error) {
	q := query.PmsProductSpu

	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).UpdateSimple(q.NewStatusSort.Value(in.Sort), q.UpdateBy.Value(in.UpdateBy))

	if err != nil {
		logc.Errorf(l.ctx, "更新新品排序失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新新品排序失败")
	}
	return &pmsclient.UpdateProductSpuStatusResp{}, nil
}

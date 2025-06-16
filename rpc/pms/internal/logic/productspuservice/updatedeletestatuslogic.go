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

type UpdateDeleteStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeleteStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeleteStatusLogic {
	return &UpdateDeleteStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateDeleteStatus 修改删除状态
func (l *UpdateDeleteStatusLogic) UpdateDeleteStatus(in *pmsclient.UpdateProductSpuStatusReq) (*pmsclient.UpdateProductSpuStatusResp, error) {
	q := query.PmsProductSpu
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.IsDeleted, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "批量修改删除状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("批量修改删除状态失败")
	}

	return &pmsclient.UpdateProductSpuStatusResp{}, nil
}

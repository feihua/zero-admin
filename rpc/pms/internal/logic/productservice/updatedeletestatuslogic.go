package productservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateDeleteStatusLogic
/*
Author: LiuFeiHua
Date: 2024/5/15 16:07
*/
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

// UpdateDeleteStatus 批量修改删除状态
func (l *UpdateDeleteStatusLogic) UpdateDeleteStatus(in *pmsclient.UpdateProductStatusReq) (*pmsclient.UpdateProductStatusResp, error) {
	q := query.PmsProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.DeleteStatus, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "批量修改删除状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("批量修改删除状态失败")
	}

	return &pmsclient.UpdateProductStatusResp{}, nil
}

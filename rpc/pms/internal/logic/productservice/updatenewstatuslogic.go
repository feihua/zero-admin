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

// UpdateNewStatusLogic
/*
Author: LiuFeiHua
Date: 2024/5/15 16:08
*/
type UpdateNewStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNewStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNewStatusLogic {
	return &UpdateNewStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateNewStatus 批量设为新品
func (l *UpdateNewStatusLogic) UpdateNewStatus(in *pmsclient.UpdateProductStatusReq) (*pmsclient.UpdateProductStatusResp, error) {
	q := query.PmsProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.NewStatus, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "批量设为新品失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("批量设为新品失败")
	}

	return &pmsclient.UpdateProductStatusResp{}, nil
}

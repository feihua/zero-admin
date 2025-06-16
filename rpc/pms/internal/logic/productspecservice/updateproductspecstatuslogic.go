package productspecservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateProductSpecStatusLogic 更新商品规格
/*
Author: LiuFeiHua
Date: 2025/06/17 16:14:48
*/
type UpdateProductSpecStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductSpecStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductSpecStatusLogic {
	return &UpdateProductSpecStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductSpecStatus 更新商品规格状态
func (l *UpdateProductSpecStatusLogic) UpdateProductSpecStatus(in *pmsclient.UpdateProductSpecStatusReq) (*pmsclient.UpdateProductSpecStatusResp, error) {
	q := query.PmsProductSpec

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新商品规格状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新商品规格状态失败")
	}

	return &pmsclient.UpdateProductSpecStatusResp{}, nil
}

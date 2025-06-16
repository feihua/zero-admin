package productspecvalueservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateProductSpecValueStatusLogic 更新商品规格值
/*
Author: LiuFeiHua
Date: 2025/06/17 16:14:48
*/
type UpdateProductSpecValueStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductSpecValueStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductSpecValueStatusLogic {
	return &UpdateProductSpecValueStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductSpecValueStatus 更新商品规格值状态
func (l *UpdateProductSpecValueStatusLogic) UpdateProductSpecValueStatus(in *pmsclient.UpdateProductSpecValueStatusReq) (*pmsclient.UpdateProductSpecValueStatusResp, error) {
	q := query.PmsProductSpecValue

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新商品规格值状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新商品规格值状态失败")
	}

	return &pmsclient.UpdateProductSpecValueStatusResp{}, nil
}

package brandservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateBrandFactoryStatusLogic 商品品牌
/*
Author: LiuFeiHua
Date: 2024/5/13 17:03
*/
type UpdateBrandFactoryStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBrandFactoryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBrandFactoryStatusLogic {
	return &UpdateBrandFactoryStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateBrandFactoryStatus 批量更新厂家制造商状态
func (l *UpdateBrandFactoryStatusLogic) UpdateBrandFactoryStatus(in *pmsclient.UpdateBrandFactoryStatusReq) (*pmsclient.UpdateBrandStatusResp, error) {
	q := query.PmsBrand
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.FactoryStatus, in.FactoryStatus)

	if err != nil {
		logc.Errorf(l.ctx, "批量更新厂家制造商状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("批量更新厂家制造商状态失败")
	}

	return &pmsclient.UpdateBrandStatusResp{}, nil
}

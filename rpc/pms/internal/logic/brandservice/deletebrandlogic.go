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

// DeleteBrandLogic 删除品牌
/*
Author: LiuFeiHua
Date: 2024/6/12 16:26
*/
type DeleteBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBrandLogic {
	return &DeleteBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteBrand 删除品牌
func (l *DeleteBrandLogic) DeleteBrand(in *pmsclient.DeleteBrandReq) (*pmsclient.DeleteBrandResp, error) {
	q := query.PmsBrand
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除品牌失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除品牌失败")
	}

	return &pmsclient.DeleteBrandResp{}, nil
}

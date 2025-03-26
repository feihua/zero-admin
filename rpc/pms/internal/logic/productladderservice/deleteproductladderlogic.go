package productladderservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductLadderLogic 删除产品阶梯价格(只针对同商品)
/*
Author: LiuFeiHua
Date: 2024/6/12 17:07
*/
type DeleteProductLadderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLadderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLadderLogic {
	return &DeleteProductLadderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductLadder 删除产品阶梯价格(只针对同商品)
func (l *DeleteProductLadderLogic) DeleteProductLadder(in *pmsclient.DeleteProductLadderReq) (*pmsclient.DeleteProductLadderResp, error) {
	q := query.PmsProductLadder
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除产品阶梯价格失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除产品阶梯价格失败")
	}

	return &pmsclient.DeleteProductLadderResp{}, nil
}

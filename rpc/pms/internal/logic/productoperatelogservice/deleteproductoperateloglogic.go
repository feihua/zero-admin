package productoperatelogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductOperateLogLogic 删除
/*
Author: LiuFeiHua
Date: 2024/6/12 17:09
*/
type DeleteProductOperateLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductOperateLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductOperateLogLogic {
	return &DeleteProductOperateLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductOperateLog 删除
func (l *DeleteProductOperateLogLogic) DeleteProductOperateLog(in *pmsclient.DeleteProductOperateLogReq) (*pmsclient.DeleteProductOperateLogResp, error) {
	q := query.PmsProductOperateLog
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.DeleteProductOperateLogResp{}, nil
}

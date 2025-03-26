package productoperatelogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductOperateLogLogic 删除商品操作日志
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

// DeleteProductOperateLog 删除商品操作日志
func (l *DeleteProductOperateLogLogic) DeleteProductOperateLog(in *pmsclient.DeleteProductOperateLogReq) (*pmsclient.DeleteProductOperateLogResp, error) {
	q := query.PmsProductOperateLog
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除商品操作日志失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除商品操作日志失败")
	}

	return &pmsclient.DeleteProductOperateLogResp{}, nil
}
